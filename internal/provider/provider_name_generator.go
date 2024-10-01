/* trunk-ignore-all(golangci-lint/typecheck) */
package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure the implementation satisfies the desired interfaces.

var _ function.Function = ResourceNameGenerator{}

type ResourceNameGenerator struct{}

func NewResourceNameGenerator() function.Function {
    return ResourceNameGenerator{}
}

func (f ResourceNameGenerator) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
    resp.Name = "name_generator"
}

func (f ResourceNameGenerator) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
    resp.Definition = function.Definition{
        Summary:     "Generate name for a given resource",
        Description: "Given resource type,env,count,region parameters generate a resource name(can be list[n] of names)",
        Parameters: []function.Parameter{
            function.Int32Parameter{
                Name:        "noOfResources",
                Description: "Number of resources. Default to 1",
            },
            function.StringParameter{
                Name:        "env",
                Description: "Environment Name, eg dev, uat, qa, staging, prod. Default to dev",
            },
			function.StringParameter{
				Name:  		"region",
				Description: "Region, default to euwest2",
			},
			function.StringParameter{
				Name:  		"cloudType",
				Description: "Cloud Provider, eg. aws, az, gcp, default to az",
			},
			function.StringParameter{
				Name:  		"resource_type",
				Description: "resourceType, rg, vnet, snet, vm",
			},
			function.StringParameter{
				Name:  		"resource_category",
				Description: "category: general, networking, compute, storage, ai, integration, analytics",
			},
			function.StringParameter{
				Name:  		"app_name",
				Description: "application name, eg odie",
			},
        },
        Return: function.StringReturn{},
    }
}

func (f ResourceNameGenerator) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
    var noOfResources int64 = 1
    var env string = "dev"
    var region string = "uks"
	//var cloudType string = "az"
	var project_name string = "ken"
	var resource_type string 
	var resource_category string
	var newResourceName[99] string
	var app_name string
	var business_unit string = "fin"
	var subscription_type = "shared"
	//var rCnt string = 
	 
    // Read Terraform argument data into the variables
    resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &env))

   switch resource_category {
   //"mg-<business unit>[-<environment>]"
   //<business unit>-<subscription purpose>-<###>
   //rg-<app or service name>-<subscription purpose>-<###>
   //apim-<app or service name>   
   //id-<app or service name>-<environment>-<region name>-<###> - managed identity
   case "general":

    if resource_type == "mg" {
		newResourceName[noOfResources] =   resource_type + "-" + business_unit + "-" +  env 
	}	
	if resource_type == "sub" {
		newResourceName[noOfResources] =    business_unit + "-" +  subscription_type + "-" + strconv.FormatInt(noOfResources,4) 
	}
	if resource_type == "rg" {
		newResourceName[noOfResources] =    resource_type + "-" +  app_name + "-" + project_name + "-" +  env + "-" + strconv.FormatInt(noOfResources,4) 
	}
	if resource_type == "apim" {
		newResourceName[noOfResources] =    resource_type + "-" +  app_name + "-" + project_name + "-" +  env + "-" + strconv.FormatInt(noOfResources,4) 
	}
	if resource_type == "id" {
		newResourceName[noOfResources] =    resource_type + "-" +  app_name + "-" + project_name + "-" +  env + "-" + region + "-" + strconv.FormatInt(noOfResources,4) 
	}
	//st<project, app or service><###>
	//ssimp<project, app or service><environment>
	//cr<project, app or service><environment><###>
	// resource_type values = st, ssimp, cr	
   case "storage": {

	newResourceName[noOfResources] =   resource_type  + app_name  + project_name +   env + strconv.FormatInt(noOfResources,4)

   }
	case "databases": {

	//sqldb-<project, app or service>-<environment>
	//cosmos-<project, app or service>-<environment>
	//redis-<project, app or service>-<environment>
	// resource_type values = sqldb,cosmos,redis

	newResourceName[noOfResources] =   resource_type + "-" + app_name + "-" + project_name + "-" +  env + "-" + strconv.FormatInt(noOfResources,4)
	}
   case "compute": {
	// vm-<vm role>-<environment>-<###>
	// resource_type values - vm, app, func

	newResourceName[noOfResources] =   resource_type + "-" + app_name + "-" + project_name + "-" +  env + "-" + strconv.FormatInt(noOfResources,4)

   }
   case "networking":
	{	// eg. vnet-<subscription purpose>-<region>-<###>
		// resource_type values - vnet, snet, nic, pip, lbe, nsg, lgw, vgw, vcn, rt, dns

	newResourceName[noOfResources] =   resource_type + "-" + project_name + "-" + region + "-" +  env + "-" + strconv.FormatInt(noOfResources,4)
	if resource_type == "dns" {
		//<DNS A record for VM>.<region>.cloudapp.azure.com
	   newResourceName[noOfResources] =    app_name + "." + region + "." +  env + "." + "cloudapp.azure.com"
	}
	}
	case "ai":
		//resource type: ai search, openai, machine learning workspace
		// resource type values: srch, aai, mlw
		{
					newResourceName[noOfResources] =   resource_type + "-" + project_name + "-" + region + "-" +  env + "-" + strconv.FormatInt(noOfResources,4)

		}
	case "analytics":
		{}
	case "integration":
		{}

	default: {
			newResourceName[noOfResources] =   resource_type + "-" + project_name + "-" + region + "-" +  env + "-" + strconv.FormatInt(noOfResources,4)
	}
   }

   if resp.Error != nil {
		return
}

    // Set the result
    resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, newResourceName[noOfResources]))
}
