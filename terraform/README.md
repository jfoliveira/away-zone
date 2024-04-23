# AwayZone infrastructure environments

Terraform project to provision all the infra needed in each `AwayZone` region.

## Directory layout

For each new cloud provider region to be supported a folder must be created, keeping the following directoy layout:

```md
- away-zone-environments
  - <environment>
    - <csp-region>
      backend.tf
      providers.tf
      terraform.tfvars
      .
      ,
      ... other terraform files
```

## Dependencies

### Terraform

It's recommended to use [tfswitch](https://tfswitch.warrensbox.com/Install/) to manage your `terraform` versions, as modules in this repo have different `terraform` version requirements.
