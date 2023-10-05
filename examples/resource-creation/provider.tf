terraform {
  required_providers {
    tembo = {
      source  = "tembo-io/tembo"
      version = ">= 0.1.0"
    }
  }
}

provider "tembo" {
}
