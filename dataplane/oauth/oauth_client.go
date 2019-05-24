package oauth

import (
	"github.com/go-openapi/strfmt"
	freeipaclient "github.com/hortonworks/cb-cli/dataplane/api-freeipa/client"
	sdxclient "github.com/hortonworks/cb-cli/dataplane/api-sdx/client"
	apiclient "github.com/hortonworks/cb-cli/dataplane/api/client"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	authapiclient "github.com/hortonworks/cb-cli/dataplane/oauthapi/client"
	u "github.com/hortonworks/cb-cli/dataplane/utils"
	"github.com/hortonworks/dp-cli-common/apikeyauth"
	"github.com/hortonworks/dp-cli-common/utils"
	"github.com/urfave/cli"
)

var PREFIX_TRIM = []string{"http://", "https://"}

type Cloudbreak struct {
	Cloudbreak *apiclient.Cloudbreak
}
type Dataplane struct {
	Dataplane *authapiclient.Dataplane
}
type Sdx struct {
	Sdx *sdxclient.Datalake
}
type FreeIpa struct {
	FreeIpa *freeipaclient.FreeIPA
}

func NewCloudbreakHTTPClientFromContext(c *cli.Context) *Cloudbreak {
	return NewCloudbreakHTTPClient(c.String(fl.FlServerOptional.Name), c.String(fl.FlApiKeyIDOptional.Name), c.String(fl.FlPrivateKeyOptional.Name))
}

func NewCloudbreakHTTPClient(address string, apiKeyID, privateKey string) *Cloudbreak {
	u.CheckServerAddress(address)
	var transport *utils.Transport
	baseAPIPath := "/cb/api"
	transport = apikeyauth.GetAPIKeyAuthTransport(address, baseAPIPath, apiKeyID, privateKey)
	return &Cloudbreak{Cloudbreak: apiclient.New(transport, strfmt.Default)}
}

func NewCloudbreakActorCrnHTTPClient(address string, actorCrn string) *Cloudbreak {
	u.CheckServerAddress(address)
	var transport *utils.Transport
	baseAPIPath := "/cb/api"
	transport = apikeyauth.GetActorCrnAuthTransport(address, baseAPIPath, actorCrn)
	return &Cloudbreak{Cloudbreak: apiclient.New(transport, strfmt.Default)}
}

// NewDataplaneHTTPClientFromContext : Initialize Dataplane client.
func NewDataplaneHTTPClientFromContext(c *cli.Context) *Dataplane {
	return NewDataplaneHTTPClient(c.String(fl.FlServerOptional.Name), c.String(fl.FlApiKeyIDOptional.Name), c.String(fl.FlPrivateKeyOptional.Name))
}

// NewDataplaneHTTPClient : Creates new Dataplane client
func NewDataplaneHTTPClient(address string, apiKeyID, privateKey string) *Dataplane {
	u.CheckServerAddress(address)
	var transport *utils.Transport
	const baseAPIPath string = "/"
	transport = apikeyauth.GetAPIKeyAuthTransport(address, baseAPIPath, apiKeyID, privateKey)
	return &Dataplane{Dataplane: authapiclient.New(transport, strfmt.Default)}
}

// NewDataplaneHTTPClientFromContext : Initialize Sdx client.
func NewSDXClientFromContext(c *cli.Context) *Sdx {
	return NewSDXClient(c.String(fl.FlServerOptional.Name), c.String(fl.FlApiKeyIDOptional.Name), c.String(fl.FlPrivateKeyOptional.Name))
}

func NewSDXClient(address string, apiKeyID, privateKey string) *Sdx {
	u.CheckServerAddress(address)
	var transport *utils.Transport
	const baseAPIPath string = "/dl/api"
	transport = apikeyauth.GetAPIKeyAuthTransport(address, baseAPIPath, apiKeyID, privateKey)
	return &Sdx{Sdx: sdxclient.New(transport, strfmt.Default)}
}

// NewDataplaneHTTPClientFromContext : Initialize FreeIpa client.
func NewFreeIpaClientFromContext(c *cli.Context) *FreeIpa {
	return NewFreeIpaClient(c.String(fl.FlServerOptional.Name), c.String(fl.FlApiKeyIDOptional.Name), c.String(fl.FlPrivateKeyOptional.Name))
}

func NewFreeIpaClient(address string, apiKeyID, privateKey string) *FreeIpa {
	u.CheckServerAddress(address)
	var transport *utils.Transport
	const baseAPIPath string = "/freeipa/api"
	transport = apikeyauth.GetAPIKeyAuthTransport(address, baseAPIPath, apiKeyID, privateKey)
	return &FreeIpa{FreeIpa: freeipaclient.New(transport, strfmt.Default)}
}
