terraform {
  required_providers {
    tembo = {
      source = "registry.terraform.io/tembo-io/tembo"
    }
  }
}

provider "tembo" {}

resource "tembo_cluster" "adarsh_db" {
  cluster_name    = "adarsh-db"
  organization_id = "test"
  cpu             = "4"
  stack           = "Standard"
  environment     = "dev"
  memory          = "8Gi"
  storage         = "10Gi"
}