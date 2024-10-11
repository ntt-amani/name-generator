variable "DeployVar" {
  type = any
  #   object({
  #     Env          = string, 
  #     Region       = string,
  #     Project_code = string,
  #     Billing_code = string,
  #     App_name     = string,
  #     Project_name = string,
  #     Owner        = string
  # })
  default = {
    Env          = "dev"
    Region       = "uksouth"
    Project_code = "prj001"
    Billing_code = "bill001"
    App_name     = "webapp"
    Project_name = "st"
    Owner        = "built by terraform"
  }
}

variable "rg_ResourceVar" {
  type = map(list(any))
  #    object({
  #     Resource_category = string
  #     Resource_count    = number
  #     Resource_type     = string   
  #     Subscription_type = string
  #     Business_unit     = string
  #  } )
  default = {
    Resource_type     = "vnet"
    Resource_category = "general"
    Resource_count    = 1
    Business_unit     = "shared"
    Subscription_type = "prod"
  }
  default = {
    
  }

}
