# Tembo Terraform Provider
The Terraform provider for [Tembo](https://tembo.io/).

## Requirements
- [Terraform](https://www.terraform.io/downloads.html) >= 1.0

# Example Usage

```hcl
terraform {
  required_providers {
    tembo = {
      source = "tembo-io/tembo"
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