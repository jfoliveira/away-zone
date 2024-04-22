# AwayZone - Pulumi

A simple Pulumi program to deploy `away-zone-health-checker` helm chart to a Kubernetes cluster.

## Dependencies

- [Go](https://go.dev/doc/install)
- [pulumi](https://www.pulumi.com/docs/install/)

## Prepare local environment

Before running pulumi to deploy new changes, make sure the program dependencies are installed locally:

```shell
go mod tidy
```

## Usage

This `Go` program is automatically compiled and initialized by Pulumi.
Just run:

```shell
pulumi up
```

This will display a prompt that asks for an access token.
Follow the instructions to loging or generate a new access token:

```
Manage your Pulumi stacks by logging in.
Run `pulumi login --help` for alternative login options.
Enter your access token from https://app.pulumi.com/account/tokens
    or hit <ENTER> to log in using your browser:
```

Alternatively, to not be prompted for a pulumi stack on every execution, run:

```shell
pulumi up --stack <stack name>
```
