terraform {
  backend "s3" {
    region = "us-east-1"
    bucket = "away-zone-terraform-config"
    key    = "away-zone-environments/dev/azure-eastus2.tfstate"
    # TO DO: run terraform plan/apply in AwayZone CI/CD, reading AWS provider
    # credentials from secrets
    # shared_credentials_file = "/var/secrets/aws-dev-credentials"
  }
}
