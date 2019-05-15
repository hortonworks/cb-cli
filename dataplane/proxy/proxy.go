package proxy

import (
	"strconv"
	"time"

	"github.com/hortonworks/cb-cli/dataplane/oauth"

	"strings"

	log "github.com/Sirupsen/logrus"
	v1Proxy "github.com/hortonworks/cb-cli/dataplane/api-environment/client/v1proxies"
	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/dp-cli-common/utils"
	"github.com/urfave/cli"
)

type proxyClient interface {
	CreateProxyConfigV1(params *v1Proxy.CreateProxyConfigV1Params) (*v1Proxy.CreateProxyConfigV1OK, error)
	ListProxyConfigsV1(params *v1Proxy.ListProxyConfigsV1Params) (*v1Proxy.ListProxyConfigsV1OK, error)
}

var Header = []string{"Name", "Host", "Port", "Protocol", "Environments"}

type proxy struct {
	Name         string `json:"Name" yaml:"Name"`
	Host         string `json:"Host" yaml:"Host"`
	Port         string `json:"Port" yaml:"Port"`
	Protocol     string `json:"Protocol" yaml:"Protocol"`
	Environments []string
}

func (p *proxy) DataAsStringArray() []string {
	return []string{p.Name, p.Host, p.Port, p.Protocol, strings.Join(p.Environments, ",")}
}

func CreateProxy(c *cli.Context) error {
	defer utils.TimeTrack(time.Now(), "create proxy")

	name := c.String(fl.FlName.Name)
	host := c.String(fl.FlProxyHost.Name)
	port := c.String(fl.FlProxyPort.Name)
	protocol := c.String(fl.FlProxyProtocol.Name)
	user := c.String(fl.FlProxyUser.Name)
	password := c.String(fl.FlProxyPassword.Name)
	environments := utils.DelimitedStringToArray(c.String(fl.FlEnvironmentsOptional.Name), ",")

	if protocol != "http" && protocol != "https" {
		utils.LogErrorMessageAndExit("Proxy protocol must be either http or https")
	}
	serverPort, _ := strconv.Atoi(port)

	envClient := oauth.NewEnvironmentClientFromContext(c)

	return createProxy(envClient.Environment.V1proxies, name, host, int32(serverPort), protocol, user, password, environments)
}

func createProxy(proxyClient proxyClient, name, host string, port int32, protocol, user, password string, environments []string) error {
	proxyRequest := &model.ProxyV1Request{
		Name:         &name,
		Host:         &host,
		Port:         &port,
		Protocol:     &protocol,
		UserName:     user,
		Password:     password,
		Environments: environments,
	}

	log.Infof("[createProxy] create proxy with name: %s", name)
	var proxy *model.ProxyV1Response
	resp, err := proxyClient.CreateProxyConfigV1(v1Proxy.NewCreateProxyConfigV1Params().WithBody(proxyRequest))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	proxy = resp.Payload

	log.Infof("[createProxy] proxy created with name: %s, id: %d", name, proxy.ID)
	return nil
}

func AttachProxyToEnvs(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "attach proxy to environments")

	proxyName := c.String(fl.FlName.Name)
	environments := utils.DelimitedStringToArray(c.String(fl.FlEnvironments.Name), ",")
	log.Infof("[AttachProxyToEnvs] attach proxy config '%s' to environments: %s", proxyName, environments)

	envClient := oauth.NewEnvironmentClientFromContext(c)
	attachRequest := v1Proxy.NewAttachProxyResourceToEnvironmentsParams().WithName(proxyName).WithBody(&model.EnvironmentNames{EnvironmentNames: environments})
	response, err := envClient.Environment.V1proxies.AttachProxyResourceToEnvironments(attachRequest)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	proxy := response.Payload
	log.Infof("[AttachProxyToEnvs] proxy config '%s' is now attached to the following environments: %s", *proxy.Name, proxy.Environments)
}

func DetachProxyFromEnvs(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "detach proxy from environments")

	proxyName := c.String(fl.FlName.Name)
	environments := utils.DelimitedStringToArray(c.String(fl.FlEnvironments.Name), ",")
	log.Infof("[DetachProxyFromEnvs] detach proxy config '%s' from environments: %s", proxyName, environments)

	envClient := oauth.NewEnvironmentClientFromContext(c)
	detachRequest := v1Proxy.NewDetachProxyResourceFromEnvironmentsParams().WithName(proxyName).WithBody(&model.EnvironmentNames{EnvironmentNames: environments})
	response, err := envClient.Environment.V1proxies.DetachProxyResourceFromEnvironments(detachRequest)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	proxy := response.Payload
	log.Infof("[DetachProxyFromEnvs] proxy config '%s' is now attached to the following environments: %s", *proxy.Name, proxy.Environments)
}

func ListProxies(c *cli.Context) error {
	defer utils.TimeTrack(time.Now(), "list proxies")

	envClient := oauth.NewEnvironmentClientFromContext(c)

	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	return listProxiesImpl(envClient.Environment.V1proxies, output.WriteList)
}

func listProxiesImpl(proxyClient proxyClient, writer func([]string, []utils.Row)) error {
	resp, err := proxyClient.ListProxyConfigsV1(v1Proxy.NewListProxyConfigsV1Params())
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	var tableRows []utils.Row
	for _, p := range resp.Payload.Responses {
		row := &proxy{
			Name:         *p.Name,
			Host:         *p.Host,
			Port:         strconv.Itoa(int(*p.Port)),
			Protocol:     *p.Protocol,
			Environments: p.Environments,
		}
		tableRows = append(tableRows, row)
	}

	writer(Header, tableRows)
	return nil
}

func DeleteProxy(c *cli.Context) error {
	defer utils.TimeTrack(time.Now(), "delete a proxy")

	proxyName := c.String(fl.FlName.Name)
	log.Infof("[DeleteProxy] delete proxy config by name: %s", proxyName)

	envClient := oauth.NewEnvironmentClientFromContext(c)

	if _, err := envClient.Environment.V1proxies.DeleteProxyConfigV1(v1Proxy.NewDeleteProxyConfigV1Params().WithName(proxyName)); err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[DeleteProxy] proxy config deleted: %s", proxyName)
	return nil
}
