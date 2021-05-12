// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new azure API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for azure API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ApplyTKGConfigForAzure applies the changes to t k g configuration file for azure
*/
func (a *Client) ApplyTKGConfigForAzure(params *ApplyTKGConfigForAzureParams) (*ApplyTKGConfigForAzureOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplyTKGConfigForAzureParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "applyTKGConfigForAzure",
		Method:             "POST",
		PathPattern:        "/api/providers/azure/tkgconfig",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ApplyTKGConfigForAzureReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplyTKGConfigForAzureOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for applyTKGConfigForAzure: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateAzureRegionalCluster creates azure regional cluster
*/
func (a *Client) CreateAzureRegionalCluster(params *CreateAzureRegionalClusterParams) (*CreateAzureRegionalClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAzureRegionalClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAzureRegionalCluster",
		Method:             "POST",
		PathPattern:        "/api/providers/azure/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateAzureRegionalClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateAzureRegionalClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAzureRegionalCluster: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateAzureResourceGroup creates a new azure resource group
*/
func (a *Client) CreateAzureResourceGroup(params *CreateAzureResourceGroupParams) (*CreateAzureResourceGroupCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAzureResourceGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAzureResourceGroup",
		Method:             "POST",
		PathPattern:        "/api/providers/azure/resourcegroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateAzureResourceGroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateAzureResourceGroupCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAzureResourceGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateAzureVirtualNetwork creates a new azure virtual network
*/
func (a *Client) CreateAzureVirtualNetwork(params *CreateAzureVirtualNetworkParams) (*CreateAzureVirtualNetworkCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAzureVirtualNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAzureVirtualNetwork",
		Method:             "POST",
		PathPattern:        "/api/providers/azure/resourcegroups/{resourceGroupName}/vnets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateAzureVirtualNetworkReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateAzureVirtualNetworkCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAzureVirtualNetwork: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GenerateTKGConfigForAzure generates t k g configuration file for azure
*/
func (a *Client) GenerateTKGConfigForAzure(params *GenerateTKGConfigForAzureParams) (*GenerateTKGConfigForAzureOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGenerateTKGConfigForAzureParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "generateTKGConfigForAzure",
		Method:             "POST",
		PathPattern:        "/api/providers/azure/generate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GenerateTKGConfigForAzureReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GenerateTKGConfigForAzureOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for generateTKGConfigForAzure: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAzureEndpoint retrieves azure account params from environment variables
*/
func (a *Client) GetAzureEndpoint(params *GetAzureEndpointParams) (*GetAzureEndpointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAzureEndpointParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAzureEndpoint",
		Method:             "GET",
		PathPattern:        "/api/providers/azure",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAzureEndpointReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAzureEndpointOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAzureEndpoint: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAzureInstanceTypes retrieves list of supported azure instance types for a region
*/
func (a *Client) GetAzureInstanceTypes(params *GetAzureInstanceTypesParams) (*GetAzureInstanceTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAzureInstanceTypesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAzureInstanceTypes",
		Method:             "GET",
		PathPattern:        "/api/providers/azure/regions/{location}/instanceTypes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAzureInstanceTypesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAzureInstanceTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAzureInstanceTypes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAzureOSImages retrieves azure supported os images
*/
func (a *Client) GetAzureOSImages(params *GetAzureOSImagesParams) (*GetAzureOSImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAzureOSImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAzureOSImages",
		Method:             "GET",
		PathPattern:        "/api/providers/azure/osimages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAzureOSImagesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAzureOSImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAzureOSImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAzureRegions retrieves list of supported azure regions
*/
func (a *Client) GetAzureRegions(params *GetAzureRegionsParams) (*GetAzureRegionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAzureRegionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAzureRegions",
		Method:             "GET",
		PathPattern:        "/api/providers/azure/regions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAzureRegionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAzureRegionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAzureRegions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAzureResourceGroups retrieves list of azure resource groups for a subscription
*/
func (a *Client) GetAzureResourceGroups(params *GetAzureResourceGroupsParams) (*GetAzureResourceGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAzureResourceGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAzureResourceGroups",
		Method:             "GET",
		PathPattern:        "/api/providers/azure/resourcegroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAzureResourceGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAzureResourceGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAzureResourceGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAzureVnets retrieves list of azure virtual networks in a resource group
*/
func (a *Client) GetAzureVnets(params *GetAzureVnetsParams) (*GetAzureVnetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAzureVnetsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAzureVnets",
		Method:             "GET",
		PathPattern:        "/api/providers/azure/resourcegroups/{resourceGroupName}/vnets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAzureVnetsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAzureVnetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAzureVnets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SetAzureEndpoint validates and set azure credentials
*/
func (a *Client) SetAzureEndpoint(params *SetAzureEndpointParams) (*SetAzureEndpointCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetAzureEndpointParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setAzureEndpoint",
		Method:             "POST",
		PathPattern:        "/api/providers/azure",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SetAzureEndpointReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetAzureEndpointCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setAzureEndpoint: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
