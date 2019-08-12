package sdx

import (
	"encoding/json"
	"fmt"
	"github.com/hortonworks/cb-cli/dataplane/env"
	"os"
	"time"

	"github.com/hortonworks/cb-cli/dataplane/api-sdx/client/internalsdx"
	"github.com/hortonworks/cb-cli/dataplane/api-sdx/client/sdx"
	sdxModel "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/oauth"
	"github.com/hortonworks/cb-cli/dataplane/types"
	"github.com/hortonworks/dp-cli-common/utils"
	commonutils "github.com/hortonworks/dp-cli-common/utils"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var sdxClusterHeader = []string{"Crn", "Name", "EnvironmentName", "EnvironmentCrn", "StackCrn", "Status", "StatusReason"}

type sdxClusterOutput struct {
	Crn            string `json:"Crn" yaml:"Crn"`
	Name           string `json:"Name" yaml:"Name"`
	Environment    string `json:"environmentName" yaml:"environmentName"`
	EnvironmentCrn string `json:"environmentCrn" yaml:"environmentCrn"`
	StackCrn       string `json:"stackCrn" yaml:"stackCrn"`
	Status         string `json:"Status" yaml:"Status"`
	StatusReason   string `json:"StatusReason" yaml:"StatusReason"`
}

func (r *sdxClusterOutput) DataAsStringArray() []string {
	return []string{r.Crn, r.Name, r.Environment, r.EnvironmentCrn, r.StackCrn, r.Status, r.StatusReason}
}

type ClientSdx oauth.Sdx

type clientSdx interface {
	ListSdx(params *sdx.ListSdxParams) (*sdx.ListSdxOK, error)
}

func assembleStackRequest(c *cli.Context) *sdxModel.StackV4Request {
	path := c.String(fl.FlInputJson.Name)

	if path == "" {
		return nil
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		commonutils.LogErrorAndExit(err)
	}

	log.Infof("[assembleStackTemplateForSdx] read cluster create json from file: %s", path)
	content := commonutils.ReadFile(path)

	var req sdxModel.StackV4Request
	err := json.Unmarshal(content, &req)
	if err != nil {
		msg := fmt.Sprintf(`Invalid json format: %s. Please make sure that the json is valid (check for commas and double quotes).`, err.Error())
		commonutils.LogErrorMessageAndExit(msg)
	}

	return &req
}

func CreateSdx(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "create SDX cluster")

	name := c.String(fl.FlName.Name)
	envName := c.String(fl.FlEnvironmentName.Name)
	clusterShape := c.String(fl.FlClusterShape.Name)
	baseLocation := c.String(fl.FlCloudStorageBaseLocationOptional.Name)
	instanceProfile := c.String(fl.FlCloudStorageInstanceProfileOptional.Name)
	withExternalDatabase := c.Bool(fl.FlWithExternalDatabaseOptional.Name)

	inputJson := assembleStackRequest(c)

	if inputJson != nil {
		createInternalSdx(envName, inputJson, c, name, withExternalDatabase)
	} else {
		createSdx(clusterShape, envName, c, name, baseLocation, instanceProfile, withExternalDatabase)
	}
}

func createSdx(clusterShape string, envName string, c *cli.Context, name string, cloudStorageBaseLocation string, instanceProfile string, withExternalDatabase bool) {
	s3CloudStorage := &sdxModel.S3CloudStorageV1Parameters{
		InstanceProfile: &instanceProfile,
	}

	cloudStorage := &sdxModel.SdxCloudStorageRequest{
		Adls:           nil,
		AdlsGen2:       nil,
		BaseLocation:   cloudStorageBaseLocation,
		FileSystemType: "S3",
		Gcs:            nil,
		S3:             s3CloudStorage,
		Wasb:           nil,
	}

	externalDatabase := &sdxModel.SdxDatabaseRequest{
		Create: &withExternalDatabase,
	}

	sdxRequest := &sdxModel.SdxClusterRequest{
		ClusterShape:     &clusterShape,
		Environment:      &envName,
		Tags:             nil,
		CloudStorage:     cloudStorage,
		ExternalDatabase: externalDatabase,
	}
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	resp, err := sdxClient.Sdx.CreateSdx(sdx.NewCreateSdxParams().WithName(name).WithBody(sdxRequest))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	sdxCluster := resp.Payload
	log.Infof("[createSdx] SDX cluster created in environment: %s, with name: %s", envName, sdxCluster.Name)
}

func createInternalSdx(envName string, inputJson *sdxModel.StackV4Request, c *cli.Context, name string, withExternalDatabase bool) {
	externalDatabase := &sdxModel.SdxDatabaseRequest{
		Create: &withExternalDatabase,
	}

	if inputJson.EnvironmentCrn == nil || len(*inputJson.EnvironmentCrn) == 0 {
		envCrn := env.GetEnvCrnByName(envName, c)
		log.Debugf("[createInternalSdx] env crn is empty in stack request, update with: %s", envCrn)
		inputJson.EnvironmentCrn = &envCrn
	}

	sdxInternalRequest := &sdxModel.SdxInternalClusterRequest{
		ClusterShape:     &(&types.S{S: sdxModel.SdxClusterRequestClusterShapeCUSTOM}).S,
		Environment:      &envName,
		StackV4Request:   inputJson,
		Tags:             nil,
		ExternalDatabase: externalDatabase,
	}
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	resp, err := sdxClient.Internalsdx.CreateInternalSdx(internalsdx.NewCreateInternalSdxParams().WithName(name).WithBody(sdxInternalRequest))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	sdxCluster := resp.Payload
	log.Infof("[createInternalSdx] SDX cluster created in environment: %s, with name: %s", envName, sdxCluster.Name)
}

func DeleteSdx(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "delete SDX cluster")
	name := c.String(fl.FlName.Name)

	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	err := sdxClient.Sdx.DeleteSdx(sdx.NewDeleteSdxParams().WithName(name))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[deleteSdx] SDX cluster deleted in environment: %s", name)
}

func RepairSdx(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "delete SDX cluster")
	name := c.String(fl.FlName.Name)
	hostGroupToRepair := c.String(fl.FlHostGroup.Name)

	hostGroupToRepairRequest := &sdxModel.SdxRepairRequest{
		HostGroupName: &hostGroupToRepair,
	}

	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	err := sdxClient.Sdx.RepairSdxNode(sdx.NewRepairSdxNodeParams().WithName(name).WithBody(hostGroupToRepairRequest))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[repairSdx] SDX cluster repair is started for: %s", name)
}

func ListSdx(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "List sdx clusters in environment")
	envName := c.String(fl.FlEnvironmentNameOptional.Name)
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c))
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	writer := output.WriteList
	listSdxClusterImpl(sdxClient.Sdx.Sdx, envName, writer)
	log.Infof("[ListSdx] SDX cluster list in environment: %s", envName)
}

func listSdxClusterImpl(client clientSdx, envName string, writer func([]string, []utils.Row)) {
	resp, err := client.ListSdx(sdx.NewListSdxParams().WithEnvName(&envName))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	var tableRows []utils.Row
	for _, sdxCluster := range resp.Payload {
		tableRows = append(tableRows, &sdxClusterOutput{sdxCluster.Crn, sdxCluster.Name,
			sdxCluster.EnvironmentName,
			sdxCluster.EnvironmentCrn,
			sdxCluster.StackCrn,
			sdxCluster.Status,
			sdxCluster.StatusReason})
	}

	writer(sdxClusterHeader, tableRows)
}

func DescribeSdx(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "describe sdx cluster")
	name := c.String(fl.FlName.Name)
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx.Sdx
	resp, err := sdxClient.GetSdx(sdx.NewGetSdxParams().WithName(name))

	if err != nil {
		utils.LogErrorAndExit(err)
	}

	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	sdxCluster := resp.Payload
	output.Write(sdxClusterHeader, &sdxClusterOutput{sdxCluster.Crn,
		sdxCluster.Name,
		sdxCluster.EnvironmentName,
		sdxCluster.EnvironmentCrn,
		sdxCluster.StackCrn,
		sdxCluster.Status,
		sdxCluster.StatusReason})
	log.Infof("[DescribeSdx] Describe a particular SDX cluster")
}
