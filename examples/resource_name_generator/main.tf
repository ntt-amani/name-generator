terraform {
  required_providers {
    hashicups = {
      source = "hashicorp.com/edu/hashicups"
    }
  }
  required_version = ">= 1.8.0"
}

provider "hashicups" {}


locals  {
     resource_name = split(" ",provider::hashicups::name_generator(3, "dev", "uksouth", "ken", "rg", "cache","general") )
}



output "resource_name" {
     value = local.resource_name[1]
}