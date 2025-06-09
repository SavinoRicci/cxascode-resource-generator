package assistants

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-genesyscloud/genesyscloud/provider"
	resourceExporter "terraform-provider-genesyscloud/genesyscloud/resource_exporter"
	registrar "terraform-provider-genesyscloud/genesyscloud/resource_register"
)

/*
resource_genesycloud_assistants_schema.go holds four functions within it:

1.  The registration code that registers the Datasource, Resource and Exporter for the package.
2.  The resource schema definitions for the assistants resource.
3.  The datasource schema definitions for the assistants datasource.
4.  The resource exporter configuration for the assistants exporter.
*/
const resourceName = "genesyscloud_assistants"

// SetRegistrar registers all of the resources, datasources and exporters in the package
func SetRegistrar(regInstance registrar.Registrar) {
	regInstance.RegisterResource(resourceName, ResourceAssistants())
	regInstance.RegisterDataSource(resourceName, DataSourceAssistants())
	regInstance.RegisterExporter(resourceName, AssistantsExporter())
}

// ResourceAssistants registers the genesyscloud_assistants resource with Terraform
func ResourceAssistants() *schema.Resource {
	return &schema.Resource{
		Description: `Genesys Cloud assistants`,

		CreateContext: provider.CreateWithPooledClient(createAssistants),
		ReadContext:   provider.ReadWithPooledClient(readAssistants),
		UpdateContext: provider.UpdateWithPooledClient(updateAssistants),
		DeleteContext: provider.DeleteWithPooledClient(deleteAssistants),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		SchemaVersion: 1,
		Schema:        map[string]*schema.Schema{},
	}
}

// AssistantsExporter returns the resourceExporter object used to hold the genesyscloud_assistants exporter's config
func AssistantsExporter() *resourceExporter.ResourceExporter {
	return &resourceExporter.ResourceExporter{
		GetResourcesFunc: provider.GetAllWithPooledClient(getAllAuthAssistantss),
		RefAttrs:         map[string]*resourceExporter.RefAttrSettings{
			// TODO: Add any reference attributes here
		},
	}
}

// DataSourceAssistants registers the genesyscloud_assistants data source
func DataSourceAssistants() *schema.Resource {
	return &schema.Resource{
		Description: `Genesys Cloud assistants data source. Select an assistants by name`,
		ReadContext: provider.ReadWithPooledClient(dataSourceAssistantsRead),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Description: `assistants name`,
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}
