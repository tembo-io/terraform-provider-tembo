# Tembo Terraform Provider
The Terraform provider for [Tembo](https://tembo.io/).

## Requirements
- [Terraform](https://www.terraform.io/downloads.html) >= 1.0

## Authentication

Tembo Terraform Provider needs an `access_token` to authenticate with the API. Generate a long-lived API token by following steps [here](https://tembo.io/docs/tembo-cloud/api#create-a-long-lived-api-token).

Note: To set `access_token` either create a local .tfvars file with the variable or set an environment variable with the access_token as shown below.

```bash
export TEMBO_ACCESS_TOKEN="TOKEN_GOES_HERE"
```

## Example Usage

```hcl
terraform {
  required_providers {
    tembo = {
      source = "tembo-io/tembo"
      version = ">= 0.1.0"
    }
  }
}

provider "tembo" {
  access_token = var.access_token
}

variable "access_token" {
  type = string
}
```
