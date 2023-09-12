terraform {
  required_providers {
    tembo = {
      source = "registry.terraform.io/tembo-io/tembo"
    }
  }
}

provider "tembo" {}

resource "tembo_instance" "adarsh_db" {
  instance_name = "adarsh-db"
  org_id        = "org_30testfkw8WUVpJc6Zuhc"
  cpu           = "4"
  stack_type    = "Standard"
  environment   = "dev"
  memory        = "8Gi"
  storage       = "10Gi"
}

output "instance" {
  value = tembo_instance.adarsh_db
}
