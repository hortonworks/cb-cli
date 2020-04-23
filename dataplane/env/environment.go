package env

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hortonworks/cb-cli/dataplane/api-environment/client"
	"github.com/hortonworks/cb-cli/dataplane/api-environment/client/v1utils"
	sdxModel "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v1distrox"
	"github.com/hortonworks/cb-cli/dataplane/common"
	"github.com/hortonworks/cb-cli/dataplane/distrox"
	"github.com/hortonworks/cb-cli/dataplane/sdx"
	"strconv"
	"strings"
	"time"

	"github.com/hortonworks/cb-cli/dataplane/cloud"

	"github.com/hortonworks/cb-cli/dataplane/api-environment/model"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/oauth"
	"github.com/hortonworks/dp-cli-common/utils"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/hortonworks/cb-cli/dataplane/api-environment/client/v1env"
)

var EnvironmentHeader = []string{"Name", "Description", "CloudPlatform", "Status", "Credential", "Regions", "LocationName", "Crn"}

type environment struct {
	Name                  string   `json:"Name" yaml:"Name"`
	Description           string   `json:"Description" yaml:"Description"`
	CloudPlatform         string   `json:"CloudPlatform" yaml:"CloudPlatform"`
	Status                string   `json:"Status" yaml:"Status"`
	Credential            string   `json:"Credential" yaml:"Credential"`
	Regions               []string `json:"Regions" yaml:"Regions"`
	LocationName          string   `json:"LocationName" yaml:"LocationName"`
	Longitude             float64  `json:"Longitude" yaml:"Longitude"`
	Latitude              float64  `json:"Latitude" yaml:"Latitude"`
	Crn                   string   `json:"Crn" yaml:"Crn"`
	ParentEnvironmentName string   `json:"ParentEnvironmentName" yaml:"ParentEnvironmentName"`
}

type environmentOutTableDescribe struct {
	*environment
	ProxyConfigName string `json:"ProxyConfigName" yaml:"ProxyConfigName"`
	SdxStatus       string `json:"SdxStatus" yaml:"SdxStatus"`
	DataHubCount    int    `json:"DataHubCount" yaml:"DataHubCount"`
}

type environmentOutJsonDescribe struct {
	*environment
	ProxyConfig    model.ProxyResponse                       `json:"ProxyConfig" yaml:"ProxyConfig"`
	Network        model.EnvironmentNetworkV1Response        `json:"Network" yaml:"Network"`
	Tags           model.TagResponse                         `json:"Tags" yaml:"Tags"`
	Telemetry      model.TelemetryResponse                   `json:"Telemetry" yaml:"Telemetry"`
	Authentication model.EnvironmentAuthenticationV1Response `json:"Authentication" yaml:"Authentication"`
	FreeIpa        model.FreeIpaResponse                     `json:"FreeIpa" yaml:"FreeIpa"`
}

type environmentListJsonDescribe struct {
	*environment
	ProxyConfig model.ProxyViewResponse            `json:"ProxyConfig" yaml:"ProxyConfig"`
	Network     model.EnvironmentNetworkV1Response `json:"Network" yaml:"Network"`
	Tags        model.TagResponse                  `json:"Tags" yaml:"Tags"`
	Telemetry   model.TelemetryResponse            `json:"Telemetry" yaml:"Telemetry"`
	FreeIpa     model.FreeIpaResponse              `json:"FreeIpa" yaml:"FreeIpa"`
}

type features struct {
	ClusterLogsCollection *featureSetting `json:"ClusterLogsCollection,omitempty" yaml:"ClusterLogsCollection"`
	WorkloadAnalytics     *featureSetting `json:"WorkloadAnalytics,omitempty" yaml:"WorkloadAnalytics"`
}

type featureSetting struct {
	Enabled bool `json:"Enabled" yaml:"Enabled"`
}

type environmentClient interface {
	CreateEnvironmentV1(params *v1env.CreateEnvironmentV1Params) (*v1env.CreateEnvironmentV1OK, error)
	ListEnvironmentV1(params *v1env.ListEnvironmentV1Params) (*v1env.ListEnvironmentV1OK, error)
}

func (e *environment) DataAsStringArray() []string {
	return []string{e.Name, e.Description, e.CloudPlatform, e.Status, e.Credential, strings.Join(e.Regions, ","), e.LocationName, e.Crn}
}

func (e *environmentOutJsonDescribe) DataAsStringArray() []string {
	return append(e.environment.DataAsStringArray())
}

func (e *environmentOutTableDescribe) DataAsStringArray() []string {
	dhc := "UNKNOWN"
	if e.DataHubCount > -1 {
		dhc = strconv.Itoa(e.DataHubCount)
	}
	return append(e.environment.DataAsStringArray(), e.ProxyConfigName, e.SdxStatus, dhc)
}

func CreateEnvironmentFromTemplate(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "create environment from template")
	fileLocation := c.String(fl.FlEnvironmentTemplateFile.Name)
	log.Infof("[assembleStackTemplate] read environment template JSON from file: %s", fileLocation)
	content := utils.ReadFile(fileLocation)

	var req model.EnvironmentV1Request
	err := json.Unmarshal(content, &req)
	if err != nil {
		msg := fmt.Sprintf(`Invalid JSON format: %s. Please make sure that the json is valid (check for commas and double quotes).`, err.Error())
		utils.LogErrorMessageAndExit(msg)
	}

	if name := c.String(fl.FlNameOptional.Name); len(name) != 0 {
		req.Name = &name
	} else if req.Name == nil {
		utils.LogErrorMessageAndExit("Name of the environment must be set either in the template or with the --name command line option.")
	}
	createEnvironmentImpl(c, &req)
}

func EditEnvironmentFromTemplate(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "edit environment from template")
	fileLocation := c.String(fl.FlEnvironmentTemplateFile.Name)
	log.Infof("[EditEnvironmentFromTemplate] read environment edit template JSON from file: %s", fileLocation)
	content := utils.ReadFile(fileLocation)

	var req model.EnvironmentEditV1Request
	err := json.Unmarshal(content, &req)
	if err != nil {
		msg := fmt.Sprintf(`Invalid JSON format: %s. Please make sure that the json is valid (check for commas and double quotes).`, err.Error())
		utils.LogErrorMessageAndExit(msg)
	}
	var name string
	if name = c.String(fl.FlNameOptional.Name); len(name) == 0 {
		utils.LogErrorMessageAndExit("Name of the environment must be set with the --name command line option.")
	}

	envClient := oauth.NewEnvironmentClientFromContext(c)
	checkClientVersion(envClient.Environment, common.Version)
	resp, err := envClient.Environment.V1env.EditEnvironmentV1(v1env.NewEditEnvironmentV1Params().WithBody(&req).WithName(name))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	environment := resp.Payload

	log.Infof("[EditEnvironmentFromTemplate] environment has edited with name: %s, crn: %s", environment.Name, environment.Crn)
}

func EditEnvironmentTelemetryFeaturesFromTemplate(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "edit environment telemetry features from template")
	fileLocation := c.String(fl.FlEnvironmentTemplateFile.Name)
	log.Infof("[EditEnvironmentTelemetryFeaturesFromTemplate] read environment telemetry feature edit template JSON from file: %s", fileLocation)
	content := utils.ReadFile(fileLocation)

	var req model.FeaturesRequest
	err := json.Unmarshal(content, &req)
	if err != nil {
		msg := fmt.Sprintf(`Invalid JSON format: %s. Please make sure that the json is valid (check for commas and double quotes).`, err.Error())
		utils.LogErrorMessageAndExit(msg)
	}
	var name string
	if name = c.String(fl.FlNameOptional.Name); len(name) == 0 {
		utils.LogErrorMessageAndExit("Name of the environment must be set with the --name command line option.")
	}
	envClient := oauth.NewEnvironmentClientFromContext(c)
	checkClientVersion(envClient.Environment, common.Version)
	resp, err := envClient.Environment.V1env.ChangeTelemetryFeaturesInEnvironmentV1ByName(v1env.NewChangeTelemetryFeaturesInEnvironmentV1ByNameParams().WithBody(&req).WithName(name))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	environment := resp.Payload

	log.Infof("[EditEnvironmentTelemetryFeaturesFromTemplate] environment has edited (telemetry features) with name: %s, crn: %s", environment.Name, environment.Crn)
}

func createEnvironmentImpl(c *cli.Context, EnvironmentV1Request *model.EnvironmentV1Request) {
	log.Infof("[createEnvironmentImpl] create environment with name: %s", *EnvironmentV1Request.Name)
	envClient := oauth.NewEnvironmentClientFromContext(c)
	checkClientVersion(envClient.Environment, common.Version)
	resp, err := envClient.Environment.V1env.CreateEnvironmentV1(v1env.NewCreateEnvironmentV1Params().WithBody(EnvironmentV1Request))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	environment := resp.Payload

	log.Infof("[createEnvironmentImpl] environment created with name: %s, id: %s", *EnvironmentV1Request.Name, environment.Crn)
}

func GenerateAwsEnvironmentTemplate(c *cli.Context) error {
	cloud.SetProviderType(cloud.AWS)
	template := createEnvironmentTemplateWithNetwork(getNetworkMode(c), c)
	return printTemplate(template)
}

func GenerateAzureEnvironmentTemplate(c *cli.Context) error {
	cloud.SetProviderType(cloud.AZURE)
	template := createEnvironmentTemplateWithNetwork(getNetworkMode(c), c)
	return printTemplate(template)
}

func getNetworkMode(c *cli.Context) cloud.NetworkMode {
	switch c.Command.Names()[0] {
	case "create-new-network":
		return cloud.NEW_NETWORK_NEW_SUBNET
	case "use-existing-network":
		return cloud.EXISTING_NETWORK_EXISTING_SUBNET
	default:
		return cloud.NO_NETWORK
	}
}

func createEnvironmentTemplateWithNetwork(mode cloud.NetworkMode, c *cli.Context) model.EnvironmentV1Request {
	provider := cloud.GetProvider()
	template := model.EnvironmentV1Request{
		Name:           new(string),
		Description:    new(string),
		CredentialName: "____",
		Regions:        make([]string, 0),
		Location: &model.LocationV1Request{
			Name:      new(string),
			Longitude: 0,
			Latitude:  0,
		},
		Network: provider.GenerateDefaultNetworkWithParams(c.String, mode),
	}
	if credName := c.String(fl.FlEnvironmentCredentialOptional.Name); len(credName) != 0 {
		template.CredentialName = credName
	}
	if locationName := c.String(fl.FlEnvironmentLocationNameOptional.Name); len(locationName) != 0 {
		template.Location.Name = &locationName
	}
	if regions := utils.DelimitedStringToArray(c.String(fl.FlEnvironmentRegions.Name), ","); len(regions) != 0 {
		template.Regions = regions
	}
	return template
}

func printTemplate(template model.EnvironmentV1Request) error {
	resp, err := json.MarshalIndent(template, "", "\t")
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	fmt.Printf("%s\n", string(resp))
	return nil
}

func ListEnvironments(c *cli.Context) error {
	defer utils.TimeTrack(time.Now(), "list environments")

	envClient := oauth.NewEnvironmentClientFromContext(c)

	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	return listEnvironmentsImpl(envClient.Environment.V1env, output, c)
}

func listEnvironmentsImpl(envClient environmentClient, output utils.Output, c *cli.Context) error {
	resp, err := envClient.ListEnvironmentV1(v1env.NewListEnvironmentV1Params())
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	dixis := distrox.GetListOfDistroXs(c)
	sdxs := sdx.GetListOfSdx(c)

	var tableRows []utils.Row
	for _, e := range resp.Payload.Responses {
		row := &environment{
			Name:                  e.Name,
			Description:           e.Description,
			CloudPlatform:         e.CloudPlatform,
			Status:                e.EnvironmentStatus,
			Credential:            e.Credential.Name,
			Regions:               getRegionNames(e.Regions),
			LocationName:          e.Location.Name,
			Longitude:             e.Location.Longitude,
			Latitude:              e.Location.Latitude,
			Crn:                   e.Crn,
			ParentEnvironmentName: e.ParentEnvironmentName,
		}

		if output.Format != "table" && output.Format != "yaml" {
			envListJSON := environmentListJsonDescribe{
				environment: row,
			}
			if e.ProxyConfig != nil {
				envListJSON.ProxyConfig = *e.ProxyConfig
			}
			if e.Network != nil {
				envListJSON.Network = *e.Network
			}
			if e.Telemetry != nil {
				envListJSON.Telemetry = *e.Telemetry
			}
			if e.FreeIpa != nil {
				envListJSON.FreeIpa = *e.FreeIpa
			}
			tableRows = append(tableRows, &envListJSON)
		} else {
			envListTable := environmentOutTableDescribe{
				environment:  row,
				SdxStatus:    getSdxStatusByEnv(sdxs, row.Crn),
				DataHubCount: getDistroXCountByEnv(dixis, row.Crn),
			}
			if e.ProxyConfig != nil {
				envListTable.ProxyConfigName = *e.ProxyConfig.Name
			}
			tableRows = append(tableRows, &envListTable)
		}
	}

	if output.Format != "table" && output.Format != "yaml" {
		output.WriteList(append(EnvironmentHeader, "ProxyConfig", "Network", "Telemetry", "Tags"), tableRows)
	} else {
		output.WriteList(append(EnvironmentHeader, "ProxyConfig", "SdxStatus", "DataHubCount"), tableRows)
	}
	return nil
}

func getDistroXCountByEnv(dixlist *v1distrox.ListDistroXV1OK, envCrn string) int {
	quantity := 0
	if dixlist == nil {
		return 0
	}
	for _, response := range dixlist.Payload.Responses {
		if response.EnvironmentCrn == envCrn {
			quantity = quantity + 1
		}
	}
	return quantity
}

func getSdxStatusByEnv(sdxlist []*sdxModel.SdxClusterResponse, envCrn string) string {
	for _, response := range sdxlist {
		if response.EnvironmentCrn == envCrn {
			return response.Status
		}
	}
	return "-"
}

func DescribeEnvironment(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "describe an environment")

	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	envName := c.String(fl.FlName.Name)
	log.Infof("[DescribeEnvironment] describe environment by name: %s", envName)
	envClient := oauth.NewEnvironmentClientFromContext(c)

	resp, err := envClient.Environment.V1env.GetEnvironmentV1ByName(v1env.NewGetEnvironmentV1ByNameParams().WithName(envName))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	env := resp.Payload
	if output.Format != "table" && output.Format != "yaml" {
		output.Write(append(EnvironmentHeader, "ProxyConfig", "Network", "Telemetry", "Authentication", "FreeIpa"), convertResponseToJsonOutput(env))
	} else {
		dixis := distrox.GetListOfDistroXs(c)
		sdxs := sdx.GetListOfSdx(c)
		proxyConfig := ""
		if env.ProxyConfig != nil {
			proxyConfig = *env.ProxyConfig.Name
		}
		output.Write(append(EnvironmentHeader, "ProxyConfig", "SdxStatus", "DataHubCount"), convertResponseToTableOutput(env, proxyConfig, getSdxStatusByEnv(sdxs, env.Crn), getDistroXCountByEnv(dixis, env.Crn)))
	}
}

func StartEnvironment(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "start an environment")
	envName := c.String(fl.FlName.Name)
	log.Infof("[StartEnvironment] start environment by name: %s", envName)
	envClient := oauth.NewEnvironmentClientFromContext(c)

	err := envClient.Environment.V1env.StartEnvironmentByNameV1(v1env.NewStartEnvironmentByNameV1Params().WithName(envName))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
}

func StopEnvironment(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "stop an environment")
	envName := c.String(fl.FlName.Name)
	log.Infof("[StopEnvironment] stop environment by name: %s", envName)
	envClient := oauth.NewEnvironmentClientFromContext(c)

	err := envClient.Environment.V1env.StopEnvironmentByNameV1(v1env.NewStopEnvironmentByNameV1Params().WithName(envName))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
}

func DeleteEnvironment(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "delete environment")
	val := c.String(fl.FlNames.Name)
	envNames := strings.Split(val[1:len(val)-1], ",")
	forced := c.Bool(fl.FlForceOptional.Name)
	envClient := oauth.NewEnvironmentClientFromContext(c)
	log.Infof("[DeleteEnvironment] delete environment(s) by names: %s", envNames)
	_, err := envClient.Environment.V1env.DeleteEnvironmentsByName(v1env.NewDeleteEnvironmentsByNameParams().WithForced(&forced).WithBody(envNames))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
}

func ChangeCredential(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "change credential of environment")
	envName := c.String(fl.FlName.Name)
	credential := c.String(fl.FlCredential.Name)

	requestBody := &model.EnvironmentChangeCredentialV1Request{
		CredentialName: credential,
	}
	request := v1env.NewChangeCredentialInEnvironmentV1Params().WithName(envName).WithBody(requestBody)
	envClient := oauth.NewEnvironmentClientFromContext(c)
	checkClientVersion(envClient.Environment, common.Version)
	resp, err := envClient.Environment.V1env.ChangeCredentialInEnvironmentV1(request)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	environment := resp.Payload
	log.Infof("[ChangeCredential] credential of environment %s changed to: %s", environment.Name, *environment.Credential.Name)
}

func convertResponseToTableOutput(env *model.DetailedEnvironmentV1Response, proxyConfig string, sdxStatus string, datahubCount int) *environmentOutTableDescribe {
	return &environmentOutTableDescribe{
		environment: &environment{
			Name:          env.Name,
			Description:   env.Description,
			CloudPlatform: env.CloudPlatform,
			Status:        env.EnvironmentStatus,
			Credential:    *env.Credential.Name,
			Regions:       getRegionNames(env.Regions),
			LocationName:  env.Location.Name,
			Longitude:     env.Location.Longitude,
			Latitude:      env.Location.Latitude,
			Crn:           env.Crn,
		},
		ProxyConfigName: proxyConfig,
		SdxStatus:       sdxStatus,
		DataHubCount:    datahubCount,
	}
}

func convertResponseToJsonOutput(env *model.DetailedEnvironmentV1Response) *environmentOutJsonDescribe {
	result := &environmentOutJsonDescribe{
		environment: &environment{
			Name:                  env.Name,
			Description:           env.Description,
			CloudPlatform:         env.CloudPlatform,
			Status:                env.EnvironmentStatus,
			Credential:            *env.Credential.Name,
			Regions:               getRegionNames(env.Regions),
			LocationName:          env.Location.Name,
			Longitude:             env.Location.Longitude,
			Latitude:              env.Location.Latitude,
			Crn:                   env.Crn,
			ParentEnvironmentName: env.ParentEnvironmentName,
		},
	}
	if env.ProxyConfig != nil {
		result.ProxyConfig = *env.ProxyConfig
	}
	if env.Network != nil {
		result.Network = *env.Network
	}
	if env.Telemetry != nil {
		result.Telemetry = *env.Telemetry
	}
	if env.Authentication != nil {
		result.Authentication = *env.Authentication
	}
	if env.FreeIpa != nil {
		result.FreeIpa = *env.FreeIpa
	}
	if env.Tags != nil {
		result.Tags = *env.Tags
	}
	return result
}

func getRegionNames(region *model.CompactRegionV1Response) []string {
	var regions []string
	for k, v := range region.DisplayNames {
		regions = append(regions, fmt.Sprintf("%s (%s)", v, k))
	}
	return regions
}

func GetEnvCrnByName(envName string, c *cli.Context) string {
	defer utils.TimeTrack(time.Now(), "get env crn by name")

	log.Debugf("[getEnvCrnByName] get env crn by name: %s", envName)
	envClient := oauth.NewEnvironmentClientFromContext(c)
	envResp, err := envClient.Environment.V1env.GetEnvironmentV1ByName(v1env.NewGetEnvironmentV1ByNameParams().WithName(envName))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	return envResp.Payload.Crn
}

func checkClientVersion(client *client.Environment, version string) {
	versionCheckRequest := v1utils.NewCheckClientVersionOfEnvironmentV1Params().WithVersion(&version)
	resp, err := client.V1utils.CheckClientVersionOfEnvironmentV1(versionCheckRequest)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	valid := resp.Payload.VersionCheckOk
	message := resp.Payload.Message
	if !valid {
		utils.LogErrorAndExit(errors.New(message))
	}
}
