terraform {
  required_providers {
    tembo = {
      source = "registry.terraform.io/tembo-io/tembo"
    }
  }
}

provider "tembo" {}

resource "tembo_instance" "adarsh_db" {
  instance_name = "adarsh-tf-provid9"
  org_id        = "org_2UdhszNbCVhLAXkZm30nz8pL778"
  cpu           = "1"
  stack_type    = "Standard"
  environment   = "dev"
  memory        = "4Gi"
  storage       = "10Gi"
}

output "instance" {
  value = tembo_instance.adarsh_db
}
