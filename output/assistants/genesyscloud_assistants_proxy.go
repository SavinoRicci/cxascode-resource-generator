package assistants

import (
	"context"
	"github.com/mypurecloud/platform-client-sdk-go/v133/platformclientv2"
)

/*
The genesyscloud_assistants_proxy.go file contains the proxy structures and methods that interact
with the Genesys Cloud SDK. We use composition here for each function on the proxy so individual functions can be stubbed
out during testing.
*/

// internalProxy holds a proxy instance that can be used throughout the package
var internalProxy *assistantsProxy

// Type definitions for each func on our proxy so we can easily mock them out later
type createAssistantsFunc func(ctx context.Context, p *assistantsProxy, assistant *platformclientv2.Assistant) (*platformclientv2.Assistant, *platformclientv2.APIResponse, error)
type getAllAssistantsFunc func(ctx context.Context, p *assistantsProxy) (*[]platformclientv2.Assistant, *platformclientv2.APIResponse, error)
type getAssistantsIdByNameFunc func(ctx context.Context, p *assistantsProxy, name string) (string, *platformclientv2.APIResponse, bool, error)
type getAssistantsByIdFunc func(ctx context.Context, p *assistantsProxy, id string) (*platformclientv2.Assistant, *platformclientv2.APIResponse, error)
type deleteAssistantsFunc func(ctx context.Context, p *assistantsProxy, id string) (*platformclientv2.APIResponse, error)

// assistantsProxy contains all of the methods that call genesys cloud APIs.
type assistantsProxy struct {
	clientConfig              *platformclientv2.Configuration
	agentAssistantsApi        *platformclientv2.AgentAssistantsApi
	createAssistantsAttr      createAssistantsFunc
	getAllAssistantsAttr      getAllAssistantsFunc
	getAssistantsIdByNameAttr getAssistantsIdByNameFunc
	getAssistantsByIdAttr     getAssistantsByIdFunc
	deleteAssistantsAttr      deleteAssistantsFunc
}

// newAssistantsProxy initializes the assistants proxy with all of the data needed to communicate with Genesys Cloud
func newAssistantsProxy(clientConfig *platformclientv2.Configuration) *assistantsProxy {
	api := platformclientv2.NewAgentAssistantsApiWithConfig(clientConfig)
	return &assistantsProxy{
		clientConfig:              clientConfig,
		agentAssistantsApi:        api,
		createAssistantsAttr:      createAssistantsFn,
		getAllAssistantsAttr:      getAllAssistantsFn,
		getAssistantsIdByNameAttr: getAssistantsIdByNameFn,
		getAssistantsByIdAttr:     getAssistantsByIdFn,
		deleteAssistantsAttr:      deleteAssistantsFn,
	}
}

// getAssistantsProxy acts as a singleton to for the internalProxy.  It also ensures
// that we can still proxy our tests by directly setting internalProxy package variable
func getAssistantsProxy(clientConfig *platformclientv2.Configuration) *assistantsProxy {
	if internalProxy == nil {
		internalProxy = newAssistantsProxy(clientConfig)
	}

	return internalProxy
}

// createAssistants creates a Genesys Cloud assistants
func (p *assistantsProxy) createAssistants(ctx context.Context, assistants *platformclientv2.Assistant) (*platformclientv2.Assistant, *platformclientv2.APIResponse, error) {
	return p.createAssistantsAttr(ctx, p, assistants)
}

// getAssistants retrieves all Genesys Cloud assistants
func (p *assistantsProxy) getAllAssistants(ctx context.Context) (*[]platformclientv2.Assistant, *platformclientv2.APIResponse, error) {
	return p.getAllAssistantsAttr(ctx, p)
}

// getAssistantsIdByName returns a single Genesys Cloud assistants by a name
func (p *assistantsProxy) getAssistantsIdByName(ctx context.Context, name string) (string, *platformclientv2.APIResponse, bool, error) {
	return p.getAssistantsIdByNameAttr(ctx, p, name)
}

// getAssistantsById returns a single Genesys Cloud assistants by Id
func (p *assistantsProxy) getAssistantsById(ctx context.Context, id string) (*platformclientv2.Assistant, *platformclientv2.APIResponse, error) {
	return p.getAssistantsByIdAttr(ctx, p, id)
}

// deleteAssistants deletes a Genesys Cloud assistants by Id
func (p *assistantsProxy) deleteAssistants(ctx context.Context, id string) (*platformclientv2.APIResponse, error) {
	return p.deleteAssistantsAttr(ctx, p, id)
}

// createAssistantsFn is an implementation function for creating a Genesys Cloud assistants
func createAssistantsFn(ctx context.Context, p *assistantsProxy, assistants *platformclientv2.Assistant) (*platformclientv2.Assistant, *platformclientv2.APIResponse, error) {
	return nil, nil, nil
}

// getAllAssistantsFn is the implementation for retrieving all assistants in Genesys Cloud
func getAllAssistantsFn(ctx context.Context, p *assistantsProxy) (*[]platformclientv2.Assistant, *platformclientv2.APIResponse, error) {
	return nil, nil, nil
}

// getAssistantsIdByNameFn is an implementation of the function to get a Genesys Cloud assistants by name
func getAssistantsIdByNameFn(ctx context.Context, p *assistantsProxy, name string) (string, *platformclientv2.APIResponse, bool, error) {
	return "", nil, false, nil
}

// getAssistantsByIdFn is an implementation of the function to get a Genesys Cloud assistants by Id
func getAssistantsByIdFn(ctx context.Context, p *assistantsProxy, id string) (*platformclientv2.Assistant, *platformclientv2.APIResponse, error) {
	return nil, nil, nil
}

// deleteAssistantsFn is an implementation function for deleting a Genesys Cloud assistants
func deleteAssistantsFn(ctx context.Context, p *assistantsProxy, id string) (*platformclientv2.APIResponse, error) {
	return nil, nil
}
