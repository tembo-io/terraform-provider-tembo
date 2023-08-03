terraform {
  required_providers {
    tembo = {
      source = "registry.terraform.io/tembo-io/tembo"
    }
  }
}

provider "tembo" {}

data "tembo_example" "example" {}