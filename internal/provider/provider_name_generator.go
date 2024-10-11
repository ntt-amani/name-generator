/* trunk-ignore-all(golangci-lint/typecheck) */
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the desired interfaces.
const (
	defaultNoOfResources = 1
	defaultEnv           = "dev"
	defaultRegion        = "uks"
	defaultProjectName   = "ken"
	defaultResourceType  = "rg"
	defaultBusinessUnit  = "fin"
	defaultSubType       = "shared"
)

var resourceCategoryHandlers = map[string]func(string, string, string, string, string, string, string, int) string{
	"general":    handleGeneralResource,
	"storage":    handleStorageResource,
	"databases":  handleDatabaseResource,
	"compute":    handleComputeResource,
	"networking": handleNetworkingResource,
	"ai":         handleAIResource,
	"analytics":  handleAnalyticsResource,
}

var _ function.Function = ResourceNameGenerator{}

type ResourceNameGenerator struct{}

func NewResourceNameGenerator() function.Function {
	return &ResourceNameGenerator{}
}

func (f ResourceNameGenerator) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "get_resource_name"
}

func (f ResourceNameGenerator) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Generate name for a given resource",
		Description: "Given resource type,env,count,region parameters generate a resource name(can be list[n] of names)",
		Parameters: []function.Parameter{
			function.Int64Parameter{
				Name:        "noOfResources",
				Description: "Number of resources. Default to 1",
			},
			function.StringParameter{
				Name:        "env",
				Description: "Environment Name, eg dev, uat, qa, staging, prod. Default to dev",
			},
			function.StringParameter{
				Name:        "region",
				Description: "Region, default to euwest2",
			},
			function.StringParameter{
				Name:        "resource_type",
				Description: "resourceType, rg, vnet, snet, vm",
			},
			function.StringParameter{
				Name:        "resource_category",
				Description: "category: general, networking, compute, storage, ai, integration, analytics",
			},
			function.StringParameter{
				Name:        "app_name",
				Description: "application name, eg data",
			},
			function.StringParameter{
				Name:        "project_name",
				Description: "project name, eg odie",
			},
			function.StringParameter{
				Name:        "business_unit",
				Description: "business unit, eg fin",
			},
			function.StringParameter{
				Name:        "sub_type",
				Description: "Sub type, eg shared",
			},
		},
		Return: function.ListReturn{
			ElementType: types.StringType,
		},
	}
}

func (f ResourceNameGenerator) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var noOfResources int = defaultNoOfResources
	var env, region, projectName, resourceType, appName, resourceCategory, businessUnit, subType string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &noOfResources, &env, &region, &resourceType, &resourceCategory, &appName, &projectName, &businessUnit, &subType))
	if resp.Error != nil {
		return
	}

	if env == "" {
		env = defaultEnv
	}
	if region == "" {
		region = defaultRegion
	}
	if projectName == "" {
		projectName = defaultProjectName
	}
	if resourceType == "" {
		resourceType = defaultResourceType
	}

	var newResourceNames []string
	handler, ok := resourceCategoryHandlers[resourceCategory]
	if !ok {
		handler = handleGeneralResource
	}

	for i := 1; i <= noOfResources; i++ {
		newResourceNames = append(newResourceNames, handler(resourceType, region, appName, projectName, businessUnit, subType, env, i))
	}

	listValue, diags := types.ListValueFrom(ctx, types.StringType, newResourceNames)

	if diags.HasError() {
		resp.Error = function.NewFuncError(diags.Errors()[0].Summary())
		return
	}
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, listValue))
}

func handleGeneralResource(resourceType, region, appName, projectName, businessUnit, subType, env string, i int) string {
	switch resourceType {
	case "mg":
		return fmt.Sprintf("%s-%s-%s", resourceType, businessUnit, env)
	case "sub":
		return fmt.Sprintf("%s-%s-%04d", businessUnit, subType, i)
	case "rg", "apim":
		return fmt.Sprintf("%s-%s-%s-%s-%04d", resourceType, appName, projectName, env, i)
	case "id":
		return fmt.Sprintf("%s-%s-%s-%s-%s-%04d", resourceType, appName, projectName, env, region, i)
	default:
		return fmt.Sprintf("%s-%s-%s-%s-%04d", resourceType, appName, projectName, env, i)
	}
}

// Implement similar functions for other resource categories

func handleStorageResource(resourceType, region, appName, projectName, businessUnit, subType, env string, i int) string {
	//st<project, app or service><###>
	//ssimp<project, app or service><environment>
	//cr<project, app or service><environment><###>
	switch resourceType {
	case "st", "ssimp", "cr":
		return fmt.Sprintf("%s%s%s%s%04d", resourceType, appName, projectName, env, i)
	default:
		return fmt.Sprintf("%s%s%s%s%04d", resourceType, appName, projectName, env, i)
	}

}

func handleDatabaseResource(resourceType, region, appName, projectName, businessUnit, subType, env string, i int) string {

	// sqldb-<project, app or service>-<environment>
	// cosmos-<project, app or service>-<environment>
	// redis-<project, app or service>-<environment>
	// resource_type values = sqldb,cosmos,redis
	switch resourceType {
	case "sqldb":
		return fmt.Sprintf("%s-%s-%s-%s-%04d", resourceType, appName, projectName, env, i)
	case "cosmos":
		return fmt.Sprintf("%s-%s-%s-%s-%04d", resourceType, appName, projectName, env, i)
	case "redis":
		return fmt.Sprintf("%s-%s-%s-%s-%04d", resourceType, appName, projectName, env, i)
	default:
		return fmt.Sprintf("%s-%s-%s-%s-%04d", resourceType, appName, projectName, env, i)
	}

}

func handleComputeResource(resourceType, region, appName, projectName, businessUnit, subType, env string, i int) string {

	// vm-<vm role>-<environment>-<###> ->
	// resource_type values - vm, app, func
	switch resourceType {
	case "vm", "app", "func":
		return fmt.Sprintf("%s-%s-%s-%s-%04d", resourceType, appName, projectName, env, i)
	default:
		return fmt.Sprintf("%s-%s-%s-%s-%04d", resourceType, appName, projectName, env, i)
	}
}

func handleNetworkingResource(resourceType, region, appName, projectName, businessUnit, subType, env string, i int) string {

	// eg. vnet-<subscription purpose>-<region>-<###>
	// resource_type values - vnet, snet, nic, pip, lbe, nsg, lgw, vgw, vcn, rt, dns

	switch resourceType {
	case "dns":
		return fmt.Sprintf("%s-%s-%s-%s-%04d", resourceType, appName, projectName, env, i)
	default:
		return fmt.Sprintf("%s-%s-%s.cloudapp.azure.com", appName, region, env)

	}
}

func handleAIResource(resourceType, region, appName, projectName, businessUnit, subType, env string, i int) string {

	//resource type: ai search, openai, machine learning workspace
	//srch-<project, app or service>-<environment>
	// resource type values: srch, aai, mlw

	switch resourceType {
	case "srch", "aai", "mlw":
		return fmt.Sprintf("%s-%s-%s-%s-%04d", resourceType, appName, projectName, env, i)
	default:
		return fmt.Sprintf("%s-%s-%s-%s-%04d", resourceType, appName, projectName, env, i)

	}
}

func handleAnalyticsResource(resourceType, region, appName, projectName, businessUnit, subType, env string, i int) string {

	//iot-<project, app or service>-<environment>
	//dls<project, app or service><environment>
	switch resourceType {
	case "iot", "dls":
		return fmt.Sprintf("%s-%s-%s-%s-%04d", resourceType, appName, projectName, env, i)
	default:
		return fmt.Sprintf("%s-%s-%s-%s-%04d", resourceType, appName, projectName, env, i)

	}

}

func handleIntegrationResource(resourceType, region, appName, projectName, businessUnit, subType, env string, i int) string {

	// service bus namespace - sbns-<project, app or service>-<environment>.servicebus.windows.net
	// service bus queue - sbq-<project, app or service>
	// service bus topic - sbt-<project, app or service>
	switch resourceType {
	case "sbns", "sbq", "sbt":
		return fmt.Sprintf("%s-%s-%s-%s-%04d", resourceType, appName, projectName, env, i)
	default:
		return fmt.Sprintf("%s-%s-%s-%s-%04d", resourceType, appName, projectName, env, i)

	}

}
