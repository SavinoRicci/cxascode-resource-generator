package assistants

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mypurecloud/platform-client-sdk-go/v133/platformclientv2"
	resourceExporter "terraform-provider-genesyscloud/genesyscloud/resource_exporter"
)

/*
The resource_genesyscloud_assistants.go contains all of the methods that perform the core logic for a resource.
*/

// getAllAuthAssistants retrieves all of the assistants via Terraform in the Genesys Cloud and is used for the exporter
func getAllAuthAssistantss(ctx context.Context, clientConfig *platformclientv2.Configuration) (resourceExporter.ResourceIDMetaMap, diag.Diagnostics) {
	return nil, nil
}

// createAssistants is used by the assistants resource to create Genesys cloud assistants
func createAssistants(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

// readAssistants is used by the assistants resource to read an assistants from genesys cloud
func readAssistants(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

// updateAssistants is used by the assistants resource to update an assistants in Genesys Cloud
func updateAssistants(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

// deleteAssistants is used by the assistants resource to delete an assistants from Genesys cloud
func deleteAssistants(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
