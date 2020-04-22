package sdx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hortonworks/cb-cli/dataplane/api-sdx/client"
	"github.com/hortonworks/cb-cli/dataplane/api-sdx/client/sdxutils"
	"github.com/hortonworks/cb-cli/dataplane/common"
	"os"
	"strings"
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

var sdxClusterHeader = []string{"Crn", "Name", "EnvironmentName", "EnvironmentCrn", "StackCrn", "DatabaseServerCrn", "CludStorageLocation", "CloudStorageFileSystemType", "Status", "StatusReason"}

type sdxClusterOutput struct {
	Crn                        string `json:"Crn" yaml:"Crn"`
	Name                       string `json:"Name" yaml:"Name"`
	Environment                string `json:"EnvironmentName" yaml:"EnvironmentName"`
	EnvironmentCrn             string `json:"EnvironmentCrn" yaml:"EnvironmentCrn"`
	StackCrn                   string `json:"StackCrn" yaml:"StackCrn"`
	DatabaseServerCrn          string `json:"DatabaseServerCrn" yaml:"DatabaseServerCrn"`
	CloudStorageBaseLocation   string `json:"CloudStorageBaseLocation" yaml:"CloudStorageBaseLocation"`
	CloudStorageFileSystemType string `json:"CloudStorageFileSystemType" yaml:"CloudStorageFileSystemType"`
	Status                     string `json:"Status" yaml:"Status"`
	StatusReason               string `json:"StatusReason" yaml:"StatusReason"`
}

type sdxClusterDetailedOutput struct {
	*sdxClusterOutput
	StackV4Response sdxModel.StackV4Response `json:"StackV4Response" yaml:"StackV4Response"`
}

func (r *sdxClusterOutput) DataAsStringArray() []string {
	return []string{r.Crn, r.Name, r.Environment, r.EnvironmentCrn, r.StackCrn, r.DatabaseServerCrn, r.CloudStorageBaseLocation, r.CloudStorageFileSystemType, r.Status, r.StatusReason}
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
	defer commonutils.TimeTrack(time.Now(), "Create SDX cluster")

	name := c.String(fl.FlName.Name)
	envName := c.String(fl.FlEnvironmentName.Name)
	clusterShape := c.String(fl.FlClusterShape.Name)
	baseLocation := c.String(fl.FlCloudStorageBaseLocationOptional.Name)
	instanceProfile := c.String(fl.FlCloudStorageInstanceProfileOptional.Name)
	managedIdentity := c.String(fl.FlCloudStorageManagedIdentityOptional.Name)
	runtime := c.String(fl.FlRuntimeOptional.Name)
	withExternalDatabase := c.Bool(fl.FlWithExternalDatabaseOptional.Name)
	withoutExternalDatabase := c.Bool(fl.FlWithoutExternalDatabaseOptional.Name)

	inputJson := assembleStackRequest(c)
	log.Infof("[CreateSdx] Runtime: %s", runtime)
	if inputJson != nil {
		createInternalSdx(envName, inputJson, c, name, baseLocation, instanceProfile, managedIdentity, runtime, withoutExternalDatabase, withExternalDatabase)
	} else {
		createSdx(clusterShape, envName, c, name, baseLocation, instanceProfile, managedIdentity, runtime, withoutExternalDatabase, withExternalDatabase)
	}
}

func createSdx(clusterShape string, envName string, c *cli.Context, name string, cloudStorageBaseLocation string, instanceProfile string, managedIdentity string, runtime string, withoutExternalDatabase bool, withExternalDatabase bool) {
	sdxRequest := createSdxRequest(clusterShape, envName, cloudStorageBaseLocation, instanceProfile, managedIdentity, runtime, withExternalDatabase, withoutExternalDatabase)

	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	checkClientVersion(sdxClient, common.Version)
	resp, err := sdxClient.Sdx.CreateSdx(sdx.NewCreateSdxParams().WithName(name).WithBody(sdxRequest))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	sdxCluster := resp.Payload
	log.Infof("[CreateSdx] SDX cluster created in environment: %s, with name: %s", envName, sdxCluster.Name)
}

func createSdxRequest(clusterShape string, envName string, cloudStorageBaseLocation string, instanceProfile string, managedIdentity string, runtime string, withExternalDatabase bool, withoutExternalDatabase bool) *sdxModel.SdxClusterRequest {
	sdxRequest := &sdxModel.SdxClusterRequest{
		ClusterShape:     &clusterShape,
		Environment:      &envName,
		Runtime:          runtime,
		Tags:             nil,
		CloudStorage:     nil,
		ExternalDatabase: nil,
	}
	setupCloudStorageIfNeeded(cloudStorageBaseLocation, instanceProfile, managedIdentity, CloudStorageSetter(func(storage *sdxModel.SdxCloudStorageRequest) { sdxRequest.CloudStorage = storage }))
	setupExternalDbIfNeeded(withExternalDatabase, withoutExternalDatabase, &sdxRequest.ExternalDatabase)
	return sdxRequest
}

func setupExternalDbIfNeeded(withExternalDatabase bool, withoutExternalDatabase bool, sdxDatabase **sdxModel.SdxDatabaseRequest) {
	if withoutExternalDatabase && withExternalDatabase {
		utils.LogErrorAndExit(errors.New("both withExternalDatabase and withoutExternalDatabase were set"))
	}
	if withExternalDatabase {
		externalDatabase := &sdxModel.SdxDatabaseRequest{
			Create: &withExternalDatabase,
		}
		*sdxDatabase = externalDatabase
	}
	if withoutExternalDatabase {
		externalDatabase := &sdxModel.SdxDatabaseRequest{
			Create: &(&types.B{B: false}).B,
		}
		*sdxDatabase = externalDatabase
	}
}

type CloudStorageSetter func(*sdxModel.SdxCloudStorageRequest)

func (s CloudStorageSetter) SetCloudStorage(req *sdxModel.SdxCloudStorageRequest) {
	s(req)
}

func setupCloudStorageIfNeeded(cloudStorageBaseLocation string, instanceProfile string, managedIdentity string, setter CloudStorageSetter) {
	if len(instanceProfile) > 0 && len(managedIdentity) > 0 {
		commonutils.LogErrorMessageAndExit("Please use managed identity or instance profile, but not both")
	}
	if len(instanceProfile) > 0 && len(cloudStorageBaseLocation) == 0 {
		commonutils.LogErrorMessageAndExit("Please use cloud storage base location if you define instance profile")
	}
	if len(managedIdentity) > 0 && len(cloudStorageBaseLocation) == 0 {
		commonutils.LogErrorMessageAndExit("Please use cloud storage base location if you define managed identity")
	}
	if len(cloudStorageBaseLocation) > 0 && len(instanceProfile) == 0 && len(managedIdentity) == 0 {
		commonutils.LogErrorMessageAndExit("Please use managed identity or instance profile if you define cloud storage")
	}
	if len(cloudStorageBaseLocation) > 0 && len(instanceProfile) > 0 {
		if len(instanceProfile) > 0 {
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
			setter.SetCloudStorage(cloudStorage)
		}
	}
	if len(cloudStorageBaseLocation) > 0 && len(managedIdentity) > 0 {
		adlsGen2Storage := &sdxModel.AdlsGen2CloudStorageV1Parameters{
			ManagedIdentity: managedIdentity,
		}
		cloudStorage := &sdxModel.SdxCloudStorageRequest{
			Adls:           nil,
			AdlsGen2:       adlsGen2Storage,
			BaseLocation:   cloudStorageBaseLocation,
			FileSystemType: "ADLS_GEN_2",
			Gcs:            nil,
			S3:             nil,
			Wasb:           nil,
		}
		setter.SetCloudStorage(cloudStorage)
	}
}

func createInternalSdx(envName string, inputJson *sdxModel.StackV4Request, c *cli.Context, name string, cloudStorageBaseLocation string, instanceProfile string, managedIdentity string, runtime string, withoutExternalDatabase bool, withExternalDatabase bool) {
	sdxInternalRequest := &sdxModel.SdxInternalClusterRequest{
		ClusterShape:     &(&types.S{S: sdxModel.SdxClusterRequestClusterShapeCUSTOM}).S,
		Environment:      &envName,
		Runtime:          runtime,
		StackV4Request:   inputJson,
		Tags:             nil,
		ExternalDatabase: nil,
	}

	setupCloudStorageIfNeeded(cloudStorageBaseLocation, instanceProfile, managedIdentity, CloudStorageSetter(func(storage *sdxModel.SdxCloudStorageRequest) { sdxInternalRequest.CloudStorage = storage }))
	setupExternalDbIfNeeded(withExternalDatabase, withoutExternalDatabase, &sdxInternalRequest.ExternalDatabase)

	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	checkClientVersion(sdxClient, common.Version)
	resp, err := sdxClient.Internalsdx.CreateInternalSdx(internalsdx.NewCreateInternalSdxParams().WithName(name).WithBody(sdxInternalRequest))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	sdxCluster := resp.Payload
	log.Infof("[CreateInternalSdx] SDX cluster created in environment: %s, with name: %s", envName, sdxCluster.Name)
}

func DeleteSdx(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "delete SDX cluster")
	name := c.String(fl.FlName.Name)
	forced := c.Bool(fl.FlForceOptional.Name)

	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	_, err := sdxClient.Sdx.DeleteSdx(sdx.NewDeleteSdxParams().WithName(name).WithForced(&forced))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[DeleteSdx] SDX cluster deleted in environment: %s", name)
}

func RepairSdx(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Sdx cluster repair")
	name := c.String(fl.FlName.Name)
	hostGroupToRepair := c.String(fl.FlHostGroupOptional.Name)
	hostGroupsToRepair := strings.Split(c.String(fl.FlHostGroupsOptional.Name), ",")

	hostGroupToRepairRequest := formulateRequest(hostGroupToRepair, hostGroupsToRepair)

	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	checkClientVersion(sdxClient, common.Version)
	_, err := sdxClient.Sdx.RepairSdxNode(sdx.NewRepairSdxNodeParams().WithName(name).WithBody(hostGroupToRepairRequest))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[RepairSdx] SDX cluster repair is started for: %s", name)
}

func formulateRequest(hostGroupToRepair string, hostGroupsToRepair []string) *sdxModel.SdxRepairRequest {
	if (len(hostGroupToRepair) == 0) == (len(hostGroupsToRepair[0]) == 0) {
		utils.LogErrorMessageAndExit(fmt.Sprintf("Please only specify either %s or %s", fl.FlHostGroupOptional.Name, fl.FlHostGroupsOptional.Name))
	}
	hostGroupToRepairRequest := &sdxModel.SdxRepairRequest{}

	if len(hostGroupToRepair) > 0 {
		hostGroupToRepairRequest.HostGroupName = hostGroupToRepair
	}

	if len(hostGroupsToRepair[0]) > 0 {
		hostGroupToRepairRequest.HostGroupNames = hostGroupsToRepair
	}

	return hostGroupToRepairRequest
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

func GetListOfSdx(c *cli.Context) []*sdxModel.SdxClusterResponse {
	defer utils.TimeTrack(time.Now(), "List sdx clusters by environment crn")
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c))
	resp, err := sdxClient.Sdx.Sdx.ListSdx(sdx.NewListSdxParams())
	if err != nil {
		return []*sdxModel.SdxClusterResponse{}
	}
	return resp.Payload
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
			sdxCluster.DatabaseServerCrn,
			sdxCluster.CloudStorageBaseLocation,
			sdxCluster.CloudStorageFileSystemType,
			sdxCluster.Status,
			sdxCluster.StatusReason})
	}

	writer(sdxClusterHeader, tableRows)
}

func SyncSdx(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "Sync sdx cluster in environment")
	name := c.String(fl.FlName.Name)
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	checkClientVersion(sdxClient, common.Version)
	err := sdxClient.Sdx.SyncSdx(sdx.NewSyncSdxParams().WithName(name))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[SyncSdx] SDX cluster sync started for: %s", name)
}

func RetrySdx(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "Retry sdx cluster")
	name := c.String(fl.FlName.Name)
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	checkClientVersion(sdxClient, common.Version)
	_, err := sdxClient.Sdx.RetrySdx(sdx.NewRetrySdxParams().WithName(name))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[RetrySdx] SDX cluster retry started for: %s", name)
}

func createCloudStorageRequestForSdx() {

}

func DescribeSdx(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "describe SDX cluster")
	name := c.String(fl.FlName.Name)
	detailed := c.Bool(fl.FlDetailedOptional.Name)
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx.Sdx

	if detailed {
		describeSdxDetailed(sdxClient, name, c)
	} else {
		describeSdx(sdxClient, name, c)
	}

}

func describeSdx(sdxClient *sdx.Client, name string, c *cli.Context) {
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
		sdxCluster.DatabaseServerCrn,
		sdxCluster.CloudStorageBaseLocation,
		sdxCluster.CloudStorageFileSystemType,
		sdxCluster.Status,
		sdxCluster.StatusReason})
	log.Infof("[DescribeSdx] Describe a particular SDX cluster")
}

func describeSdxDetailed(sdxClient *sdx.Client, name string, c *cli.Context) {
	resp, err := sdxClient.GetSdxDetail(sdx.NewGetSdxDetailParams().WithName(name))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	sdxCluster := resp.Payload

	output.Write(sdxClusterHeader, &sdxClusterDetailedOutput{
		sdxClusterOutput: &sdxClusterOutput{sdxCluster.Crn,
			sdxCluster.Name,
			sdxCluster.EnvironmentName,
			sdxCluster.EnvironmentCrn,
			sdxCluster.StackCrn,
			sdxCluster.DatabaseServerCrn,
			sdxCluster.CloudStorageBaseLocation,
			sdxCluster.CloudStorageFileSystemType,
			sdxCluster.Status,
			sdxCluster.StatusReason},
		StackV4Response: *sdxCluster.StackV4Response,
	})
	log.Infof("[DescribeSdxDetailed] Describe a particular SDX cluster")
}

func StartSdx(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "Start sdx cluster by name")
	name := c.String(fl.FlName.Name)
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	_, err := sdxClient.Sdx.StartSdxByName(sdx.NewStartSdxByNameParams().WithName(name))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[StartSdx] SDX cluster start executed for: %s", name)
}

func StopSdx(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "Stop sdx cluster by name")
	name := c.String(fl.FlName.Name)
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	_, err := sdxClient.Sdx.StopSdxByName(sdx.NewStopSdxByNameParams().WithName(name))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[StopSdx] SDX cluster stop executed for: %s", name)
}

func UpgradeSdx(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "Upgrade sdx cluster by name")
	name := c.String(fl.FlName.Name)
	dryRun := c.Bool(fl.FlDryRunOptional.Name)
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	checkClientVersion(sdxClient, common.Version)
	if dryRun {
		resp, err := sdxClient.Sdx.CheckForUpgrade(sdx.NewCheckForUpgradeParams().WithName(name))
		if err != nil {
			utils.LogErrorAndExit(err)
		}
		upgrade := resp.Payload.Upgrade
		if upgrade != nil && len(upgrade.ImageID) > 0 {
			fmt.Printf("There's new image for upgrade: %s\n", string(upgrade.ImageID))
		} else {
			log.Errorf("There's no new image for: %s", name)
		}
	} else {
		_, err := sdxClient.Sdx.UpgradeDatalakeCluster(sdx.NewUpgradeDatalakeClusterParams().WithName(name))
		if err != nil {
			utils.LogErrorAndExit(err)
		}
	}
	log.Infof("[UpgradeSdx] SDX cluster upgrade is in progress for: %s", name)
}

func checkClientVersion(client *client.Datalake, version string) {
	versionCheckRequest := sdxutils.NewCheckClientVersionOfSdxParams().WithVersion(&version)
	resp, err := client.Sdxutils.CheckClientVersionOfSdx(versionCheckRequest)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	valid := resp.Payload.VersionCheckOk
	message := resp.Payload.Message
	if !valid {
		utils.LogErrorAndExit(errors.New(message))
	}
}
