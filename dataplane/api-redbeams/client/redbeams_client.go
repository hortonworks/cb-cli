// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/dataplane/api-redbeams/client/authorization"
	"github.com/hortonworks/cb-cli/dataplane/api-redbeams/client/database_servers"
	"github.com/hortonworks/cb-cli/dataplane/api-redbeams/client/databases"
	"github.com/hortonworks/cb-cli/dataplane/api-redbeams/client/flow"
	"github.com/hortonworks/cb-cli/dataplane/api-redbeams/client/flow_public"
	"github.com/hortonworks/cb-cli/dataplane/api-redbeams/client/v4operation"
	"github.com/hortonworks/cb-cli/dataplane/api-redbeams/client/v4progress"
	"github.com/hortonworks/cb-cli/dataplane/api-redbeams/client/v4utils"
)

// Default redbeams HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/redbeams/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new redbeams HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Redbeams {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new redbeams HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Redbeams {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new redbeams client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Redbeams {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Redbeams)
	cli.Transport = transport

	cli.Authorization = authorization.New(transport, formats)

	cli.DatabaseServers = database_servers.New(transport, formats)

	cli.Databases = databases.New(transport, formats)

	cli.Flow = flow.New(transport, formats)

	cli.FlowPublic = flow_public.New(transport, formats)

	cli.V4operation = v4operation.New(transport, formats)

	cli.V4progress = v4progress.New(transport, formats)

	cli.V4utils = v4utils.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Redbeams is a client for redbeams
type Redbeams struct {
	Authorization *authorization.Client

	DatabaseServers *database_servers.Client

	Databases *databases.Client

	Flow *flow.Client

	FlowPublic *flow_public.Client

	V4operation *v4operation.Client

	V4progress *v4progress.Client

	V4utils *v4utils.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Redbeams) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Authorization.SetTransport(transport)

	c.DatabaseServers.SetTransport(transport)

	c.Databases.SetTransport(transport)

	c.Flow.SetTransport(transport)

	c.FlowPublic.SetTransport(transport)

	c.V4operation.SetTransport(transport)

	c.V4progress.SetTransport(transport)

	c.V4utils.SetTransport(transport)

}
