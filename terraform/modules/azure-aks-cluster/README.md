# Azure AKS Cluster

Terraform module to provision a cluster and related reosources on Azure Kubernetes Service (AKS).

## Example usage

```go
module "kubernetes_cluster" {
  # TO DO: consume versioned module from modules "repository"
  # source = "s3::https://s3.amazonaws.com/away-zone-tf-modules/azure-aks-cluster-0.1.0.zip"
  source             = "../../../modules/azure-aks-cluster"
  name               = "test-cluster-123"
  location           = "eastus2"
  environment        = "dev"
  kubernetes_version = "1.29"
}
```
