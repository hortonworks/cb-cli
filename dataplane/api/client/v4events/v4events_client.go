// Code generated by go-swagger; DO NOT EDIT.

package v4events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v4events API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v4events API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCloudbreakEventsListByCrn retrieves events by crn

Events are used to track stack creation initiated by cloudbreak users. Events are generated by the backend when resources requested by the user become available or unavailable
*/
func (a *Client) GetCloudbreakEventsListByCrn(params *GetCloudbreakEventsListByCrnParams) (*GetCloudbreakEventsListByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCloudbreakEventsListByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCloudbreakEventsListByCrn",
		Method:             "GET",
		PathPattern:        "/v4/events/crn/{crn}/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCloudbreakEventsListByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCloudbreakEventsListByCrnOK), nil

}

/*
GetCloudbreakEventsListByStack retrieves events by name

Events are used to track stack creation initiated by cloudbreak users. Events are generated by the backend when resources requested by the user become available or unavailable
*/
func (a *Client) GetCloudbreakEventsListByStack(params *GetCloudbreakEventsListByStackParams) (*GetCloudbreakEventsListByStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCloudbreakEventsListByStackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCloudbreakEventsListByStack",
		Method:             "GET",
		PathPattern:        "/v4/events/{name}/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCloudbreakEventsListByStackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCloudbreakEventsListByStackOK), nil

}

/*
GetEventsByStackNameInWorkspace retrieves events by name

Events are used to track stack creation initiated by cloudbreak users. Events are generated by the backend when resources requested by the user become available or unavailable
*/
func (a *Client) GetEventsByStackNameInWorkspace(params *GetEventsByStackNameInWorkspaceParams) (*GetEventsByStackNameInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEventsByStackNameInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEventsByStackNameInWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/events/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEventsByStackNameInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEventsByStackNameInWorkspaceOK), nil

}

/*
GetEventsInWorkspace retrieves events by timestamp long

Events are used to track stack creation initiated by cloudbreak users. Events are generated by the backend when resources requested by the user become available or unavailable
*/
func (a *Client) GetEventsInWorkspace(params *GetEventsInWorkspaceParams) (*GetEventsInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEventsInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEventsInWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEventsInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEventsInWorkspaceOK), nil

}

/*
GetStructuredEventsInWorkspace retrieves events by name

Events are used to track stack creation initiated by cloudbreak users. Events are generated by the backend when resources requested by the user become available or unavailable
*/
func (a *Client) GetStructuredEventsInWorkspace(params *GetStructuredEventsInWorkspaceParams) (*GetStructuredEventsInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStructuredEventsInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStructuredEventsInWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/events/{name}/structured",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStructuredEventsInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStructuredEventsInWorkspaceOK), nil

}

/*
GetStructuredEventsInWorkspaceByCrn retrieves events by crn

Events are used to track stack creation initiated by cloudbreak users. Events are generated by the backend when resources requested by the user become available or unavailable
*/
func (a *Client) GetStructuredEventsInWorkspaceByCrn(params *GetStructuredEventsInWorkspaceByCrnParams) (*GetStructuredEventsInWorkspaceByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStructuredEventsInWorkspaceByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStructuredEventsInWorkspaceByCrn",
		Method:             "GET",
		PathPattern:        "/v4/events/crn/{crn}/structured",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStructuredEventsInWorkspaceByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStructuredEventsInWorkspaceByCrnOK), nil

}

/*
GetStructuredEventsZipInWorkspace retrieves events in zip by name

Events are used to track stack creation initiated by cloudbreak users. Events are generated by the backend when resources requested by the user become available or unavailable
*/
func (a *Client) GetStructuredEventsZipInWorkspace(params *GetStructuredEventsZipInWorkspaceParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStructuredEventsZipInWorkspaceParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStructuredEventsZipInWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/events/{name}/zip",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStructuredEventsZipInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
