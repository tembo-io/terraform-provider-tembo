terraform {
  required_providers {
    tembo = {
      source = "registry.terraform.io/tembo-io/tembo"
    }
  }
}

provider "tembo" {}

resource "tembo_instance" "adarsh_db" {
  instance_name = "test-tf-provider-16"
  org_id        = "org_2UdhszNbCVhLAXkZm30nz8pL778"
  cpu           = "1"
  stack_type    = "Standard"
  environment   = "dev"
  memory        = "4Gi"
  storage       = "10Gi"
  replicas      = 1
  # extra_domains_rw = ["sample-invalid-domain.test.tembo-development.com"]
  postgres_configs = [
    {
      name  = "max_connections"
      value = "200"
    },
    {
      name  = "wal_buffers"
      value = "10"
    }
  ]
}

output "instance" {
  value = tembo_instance.adarsh_db
}
