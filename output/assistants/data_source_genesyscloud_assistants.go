package assistants

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
   The data_source_genesyscloud_assistants.go contains the data source implementation
   for the resource.
*/

// dataSourceAssistantsRead retrieves by name the id in question
func dataSourceAssistantsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
