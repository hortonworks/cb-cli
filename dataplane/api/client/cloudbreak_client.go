// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/dataplane/api/client/autoscale"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_audits"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_blueprints"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_clustertemplates"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_connectors"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_credentials"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_databases"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_environments"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_events"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_file_systems"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_flex_subscriptions"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_imagecatalogs"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_kerberos"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_kubernetes"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_ldaps"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_mpacks"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_proxies"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_recipes"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_smartsense_subscriptions"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_stacks"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4user_profiles"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4users"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4utils"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v4workspaces"
)

// Default cloudbreak HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new cloudbreak HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Cloudbreak {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new cloudbreak HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Cloudbreak {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new cloudbreak client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Cloudbreak {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Cloudbreak)
	cli.Transport = transport

	cli.Autoscale = autoscale.New(transport, formats)

	cli.V4WorkspaceID = v4_workspace_id.New(transport, formats)

	cli.V4WorkspaceIDAudits = v4_workspace_id_audits.New(transport, formats)

	cli.V4WorkspaceIDBlueprints = v4_workspace_id_blueprints.New(transport, formats)

	cli.V4WorkspaceIDClustertemplates = v4_workspace_id_clustertemplates.New(transport, formats)

	cli.V4WorkspaceIDConnectors = v4_workspace_id_connectors.New(transport, formats)

	cli.V4WorkspaceIDCredentials = v4_workspace_id_credentials.New(transport, formats)

	cli.V4WorkspaceIDDatabases = v4_workspace_id_databases.New(transport, formats)

	cli.V4WorkspaceIDEnvironments = v4_workspace_id_environments.New(transport, formats)

	cli.V4WorkspaceIDEvents = v4_workspace_id_events.New(transport, formats)

	cli.V4WorkspaceIDFileSystems = v4_workspace_id_file_systems.New(transport, formats)

	cli.V4WorkspaceIDFlexSubscriptions = v4_workspace_id_flex_subscriptions.New(transport, formats)

	cli.V4WorkspaceIDImagecatalogs = v4_workspace_id_imagecatalogs.New(transport, formats)

	cli.V4WorkspaceIDKerberos = v4_workspace_id_kerberos.New(transport, formats)

	cli.V4WorkspaceIDKubernetes = v4_workspace_id_kubernetes.New(transport, formats)

	cli.V4WorkspaceIDLdaps = v4_workspace_id_ldaps.New(transport, formats)

	cli.V4WorkspaceIDMpacks = v4_workspace_id_mpacks.New(transport, formats)

	cli.V4WorkspaceIDProxies = v4_workspace_id_proxies.New(transport, formats)

	cli.V4WorkspaceIDRecipes = v4_workspace_id_recipes.New(transport, formats)

	cli.V4WorkspaceIDSmartsenseSubscriptions = v4_workspace_id_smartsense_subscriptions.New(transport, formats)

	cli.V4WorkspaceIDStacks = v4_workspace_id_stacks.New(transport, formats)

	cli.V4userProfiles = v4user_profiles.New(transport, formats)

	cli.V4users = v4users.New(transport, formats)

	cli.V4utils = v4utils.New(transport, formats)

	cli.V4workspaces = v4workspaces.New(transport, formats)

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

// Cloudbreak is a client for cloudbreak
type Cloudbreak struct {
	Autoscale *autoscale.Client

	V4WorkspaceID *v4_workspace_id.Client

	V4WorkspaceIDAudits *v4_workspace_id_audits.Client

	V4WorkspaceIDBlueprints *v4_workspace_id_blueprints.Client

	V4WorkspaceIDClustertemplates *v4_workspace_id_clustertemplates.Client

	V4WorkspaceIDConnectors *v4_workspace_id_connectors.Client

	V4WorkspaceIDCredentials *v4_workspace_id_credentials.Client

	V4WorkspaceIDDatabases *v4_workspace_id_databases.Client

	V4WorkspaceIDEnvironments *v4_workspace_id_environments.Client

	V4WorkspaceIDEvents *v4_workspace_id_events.Client

	V4WorkspaceIDFileSystems *v4_workspace_id_file_systems.Client

	V4WorkspaceIDFlexSubscriptions *v4_workspace_id_flex_subscriptions.Client

	V4WorkspaceIDImagecatalogs *v4_workspace_id_imagecatalogs.Client

	V4WorkspaceIDKerberos *v4_workspace_id_kerberos.Client

	V4WorkspaceIDKubernetes *v4_workspace_id_kubernetes.Client

	V4WorkspaceIDLdaps *v4_workspace_id_ldaps.Client

	V4WorkspaceIDMpacks *v4_workspace_id_mpacks.Client

	V4WorkspaceIDProxies *v4_workspace_id_proxies.Client

	V4WorkspaceIDRecipes *v4_workspace_id_recipes.Client

	V4WorkspaceIDSmartsenseSubscriptions *v4_workspace_id_smartsense_subscriptions.Client

	V4WorkspaceIDStacks *v4_workspace_id_stacks.Client

	V4userProfiles *v4user_profiles.Client

	V4users *v4users.Client

	V4utils *v4utils.Client

	V4workspaces *v4workspaces.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Cloudbreak) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Autoscale.SetTransport(transport)

	c.V4WorkspaceID.SetTransport(transport)

	c.V4WorkspaceIDAudits.SetTransport(transport)

	c.V4WorkspaceIDBlueprints.SetTransport(transport)

	c.V4WorkspaceIDClustertemplates.SetTransport(transport)

	c.V4WorkspaceIDConnectors.SetTransport(transport)

	c.V4WorkspaceIDCredentials.SetTransport(transport)

	c.V4WorkspaceIDDatabases.SetTransport(transport)

	c.V4WorkspaceIDEnvironments.SetTransport(transport)

	c.V4WorkspaceIDEvents.SetTransport(transport)

	c.V4WorkspaceIDFileSystems.SetTransport(transport)

	c.V4WorkspaceIDFlexSubscriptions.SetTransport(transport)

	c.V4WorkspaceIDImagecatalogs.SetTransport(transport)

	c.V4WorkspaceIDKerberos.SetTransport(transport)

	c.V4WorkspaceIDKubernetes.SetTransport(transport)

	c.V4WorkspaceIDLdaps.SetTransport(transport)

	c.V4WorkspaceIDMpacks.SetTransport(transport)

	c.V4WorkspaceIDProxies.SetTransport(transport)

	c.V4WorkspaceIDRecipes.SetTransport(transport)

	c.V4WorkspaceIDSmartsenseSubscriptions.SetTransport(transport)

	c.V4WorkspaceIDStacks.SetTransport(transport)

	c.V4userProfiles.SetTransport(transport)

	c.V4users.SetTransport(transport)

	c.V4utils.SetTransport(transport)

	c.V4workspaces.SetTransport(transport)

}
