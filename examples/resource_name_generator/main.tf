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
     rg_resource_name = split(" ",provider::hashicups::name_generator(5, "dev", "uksouth", "ken", "rg", "cache","general") )
     storage_resource_name = split(" ",provider::hashicups::name_generator(5, "prod", "uksouth", "ken", "sa", "cache","storage") )
}


output "rg_resource_name" {
     value = local.rg_resource_name[1]
}

output "storage_resource_name" {
     value = local.storage_resource_name[4]
}