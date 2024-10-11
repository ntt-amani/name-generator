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
  type = any
  //type = map(list(any))
  #    object({
  #     Resource_category = string
  #     Resource_count    = number
  #     Resource_type     = string   
  #     Subscription_type = string
  #     Business_unit     = string
  #  } )
  # default = {
  #   Resource_type     = "vnet"
  #   Resource_category = "general"
  #   Resource_count    = 1
  #   Business_unit     = "shared"
  #   Subscription_type = "prod"
  # }
  default = {}
}

//subtype values = prod,shared, client
//business unit - mktg, corp, fin

//allowed values for resource_type, resource_category
//Management group = mg, general
//Subscription =     sub, general
//Resource group = rg, general
//api management service = apim, general
//managed identity = id, general

//Networking
//virtual network = vnet, networking
// subnet       = snet, networking
// network inc  = nic, networking
// public ip    = pip, networking
// load balancer external = lbe, networking
// NSG              = nsg, networking
// local network gatewat= lgw, networking
// virtual network gateway= vgw, networking
// vpn conn,   = vcn , networking
// dns label,    = dns, networking

//Compute and Web
// virtual machine = vm, compute
// webapp         = app, compute
// function app   = func, compute

// Databases
// azure sql server,  sqldb, db
// cosmos,            cosmos,db
// redis cache,       redis, db

// Storage
// storage account, sa, storage
// storsimple,      ssimp, storage
// acr,             cr, storage

// AI and ML
// AI Search,       srch, ai
// openai,        oai, ai
// ml workspace,  mlw, ai

// analytics and iot
// analysis services,   as, ait
// adf              ,     adf, ait
// synpase analytic workspace, synw, ait
// data lake sa,               dls, ait
// iot hub,                     iot, ait

// Integration
// Service Bus namespace,        sbns, int
// Service Bus,                 sbq, int
// Service Bus topic,           sbt, int


//allowed values for resource_category
