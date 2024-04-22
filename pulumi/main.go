package main

import (
	"fmt"
	"log"
	"os"

	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apps/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dir)
		release, err := helm.NewRelease(ctx, "health-checker", &helm.ReleaseArgs{
			Chart: pulumi.String("../kubernetes/charts/away-zone-health-checker"),
			// TO DO: to be production ready we should deploy all away-zone charts to a remote helm repository
			// and get environment specific charts specifiying the chart version used in that environment.
			// Chart could be loaded either from repository files
			// Chart:     pulumi.String("../kubernetes/environments/dev/Chart.yaml"),
			//... or from remote repo:
			// Version: pulumi.String("0.2.0"),
			// RepositoryOpts: &helm.RepositoryOptsArgs{
			// 	Repo: pulumi.String("https://charts.helm.sh/stable"),
			// },

			Namespace: pulumi.String("away-zone"),
			Timeout:   pulumi.Int(10),
			ValueYamlFiles: pulumi.AssetOrArchiveArray{
				pulumi.NewFileAsset("../kubernetes/environments/dev/images.yaml"),
				pulumi.NewFileAsset("../kubernetes/environments/dev/values.yaml"),
			},
		})
		if err != nil {
			return err
		}

		// Await on the Status of helm release and use output to retrieve deployment details
		replicas := pulumi.All(release.Status.Namespace(), release.Status.Name()).
			ApplyT(func(r any) (any, error) {
				arr := r.([]any)
				namespace := arr[0].(*string)
				name := arr[1].(*string)

				deployment, err := appsv1.GetDeployment(ctx, "deployment", pulumi.ID(fmt.Sprintf("%s/%s-away-zone-health-checker", *namespace, *name)), nil)
				if err != nil {
					return nil, err
				}
				return deployment.Spec.Replicas(), nil
			})

		ctx.Export("chartName", release.Chart)
		ctx.Export("chartNamespace", release.Namespace)
		ctx.Export("releaseName", release.Name)
		ctx.Export("deploymentReplicas", replicas)

		return nil
	})
}
