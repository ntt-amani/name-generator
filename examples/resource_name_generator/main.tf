terraform {
  required_providers {
    hashicups = {
      source  = "hashicorp.com/edu/hashicups"
      version = "0.3.1"
    }
  }
  required_version = ">= 1.8.0"
}

provider "hashicups" {
}

locals {
  rg_resource_name      = provider::hashicups::get_resource_name(5, "dev", "uksouth", "rg", "general", "data", "ken", "corp", "client")[0]
  storage_resource_name = provider::hashicups::get_resource_name(5, "prod", "uksouth", "rg", "shared", "db", "storage", "mktg", "shared")[1]
}

output "rg_resource_name" {
  value = local.rg_resource_name
}

output "storage_resource_name" {
  value = local.storage_resource_name
}
