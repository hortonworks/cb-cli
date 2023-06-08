package sdx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/hortonworks/cb-cli/dataplane/api-sdx/client/diagnostics"
	"github.com/hortonworks/cb-cli/dataplane/api-sdx/client/operation"
	"github.com/hortonworks/cb-cli/dataplane/api-sdx/model"

	"github.com/hortonworks/cb-cli/dataplane/api-sdx/client"
	"github.com/hortonworks/cb-cli/dataplane/api-sdx/client/sdxutils"
	"github.com/hortonworks/cb-cli/dataplane/common"

	"github.com/hortonworks/cb-cli/dataplane/api-sdx/client/internalsdx"
	"github.com/hortonworks/cb-cli/dataplane/api-sdx/client/sdx"
	sdxModel "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/oauth"
	"github.com/hortonworks/cb-cli/dataplane/types"
	"github.com/hortonworks/cb-cli/dataplane/utils"
	commonutils "github.com/hortonworks/dp-cli-common/utils"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var sdxClusterHeader = []string{"Crn", "Name", "EnvironmentName", "EnvironmentCrn", "StackCrn", "DatabaseServerCrn", "CludStorageLocation", "CloudStorageFileSystemType", "Status", "StatusReason", "RangerRazEnabled", "EnableMultiAz"}

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
	RangerRazEnabled           bool   `json:"RangerRazEnabled" yaml:"RangerRazEnabled"`
	EnableMultiAz              bool   `json:"EnableMultiAz" yaml:"EnableMultiAz"`
}

var sdxAdvertisedRuntimesHeader = []string{"RuntimeVersion", "DefaultRuntimeVersion"}

type sdxAdvertisedRuntimesOutput struct {
	RuntimeVersion        string `json:"RuntimeVersion" yaml:"RuntimeVersion"`
	DefaultRuntimeVersion bool   `json:"DefaultRuntimeVersion" yaml:"DefaultRuntimeVersion"`
}

func (r *sdxAdvertisedRuntimesOutput) DataAsStringArray() []string {
	return []string{r.RuntimeVersion, strconv.FormatBool(r.DefaultRuntimeVersion)}
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

type clientRuntimeSdx interface {
	Advertisedruntimes(params *sdx.AdvertisedruntimesParams) (*sdx.AdvertisedruntimesOK, error)
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
	serviceAccount := c.String(fl.FlCloudStorageServiceAccountOptional.Name)
	managedIdentity := c.String(fl.FlCloudStorageManagedIdentityOptional.Name)
	runtime := c.String(fl.FlRuntimeOptional.Name)
	withExternalDatabase := c.Bool(fl.FlWithExternalDatabaseOptional.Name)
	withoutExternalDatabase := c.Bool(fl.FlWithoutExternalDatabaseOptional.Name)
	withNonHaExternalDatabase := c.Bool(fl.FlWithNonHaExternalDatabaseOptional.Name)
	databaseEngineVersion := c.String(fl.FlRdsDatabaseEngineOptional.Name)
	spotPercentageString := c.String(fl.FlSpotPercentage.Name)
	spotPercentage := utils.ConvertToInt32Ptr(spotPercentageString)
	spotMaxPriceString := c.String(fl.FlSpotMaxPrice.Name)
	spotMaxPrice := utils.ConvertToFloat64Ptr(spotMaxPriceString)
	withRangerRazEnabled := c.Bool(fl.FlRangerRazEnabled.Name)
	withEnableMultiAz := c.Bool(fl.FlEnableMultiAz.Name)
	azureFlexibleServer := c.Bool(fl.FlAzureDatabaseFlexibleOptional.Name)
	imageCatalog := c.String(fl.FlImageCatalogOptional.Name)
	imageId := c.String(fl.FlImageIdOptional.Name)

	inputJson := assembleStackRequest(c)
	log.Infof("[CreateSdx] Runtime: %s", runtime)
	if inputJson != nil {
		createInternalSdx(envName, inputJson, c,
			name, baseLocation, instanceProfile, managedIdentity, serviceAccount, runtime,
			withoutExternalDatabase, withExternalDatabase, withNonHaExternalDatabase, databaseEngineVersion,
			spotPercentage, spotMaxPrice, withRangerRazEnabled, withEnableMultiAz, azureFlexibleServer)
	} else {
		createSdx(clusterShape, envName, c,
			name, baseLocation, instanceProfile, managedIdentity, serviceAccount, runtime,
			withoutExternalDatabase, withExternalDatabase, withNonHaExternalDatabase, databaseEngineVersion,
			spotPercentage, spotMaxPrice, withRangerRazEnabled, withEnableMultiAz, azureFlexibleServer,
			imageCatalog, imageId)
	}
}

func createSdx(clusterShape, envName string, c *cli.Context,
	name, cloudStorageBaseLocation, instanceProfile, managedIdentity, serviceAccount, runtime string,
	withoutExternalDatabase, withExternalDatabase, withNonHaExternalDatabase bool, databaseEngineVersion string,
	spotPercentage *int32, spotMaxPrice *float64, withRangerRazEnabled, withEnableMultiAz bool, azureFlexibleServer bool,
	imageCatalog, imageId string) {

	sdxRequest := createSdxRequest(clusterShape, envName, cloudStorageBaseLocation, instanceProfile, managedIdentity, serviceAccount, runtime,
		withExternalDatabase, withoutExternalDatabase, withNonHaExternalDatabase, databaseEngineVersion,
		spotPercentage, spotMaxPrice, withRangerRazEnabled, withEnableMultiAz, azureFlexibleServer, imageCatalog, imageId)

	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	checkClientVersion(sdxClient, common.Version)
	resp, err := sdxClient.Sdx.CreateCustomSdx(sdx.NewCreateCustomSdxParams().WithName(name).WithBody(sdxRequest))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	sdxCluster := resp.Payload
	log.Infof("[CreateSdx] SDX cluster created in environment: %s, with name: %s", envName, sdxCluster.Name)
}

func createSdxRequest(clusterShape, envName, cloudStorageBaseLocation, instanceProfile, managedIdentity, serviceAccount, runtime string,
	withExternalDatabase, withoutExternalDatabase, withNonHaExternalDatabase bool, databaseEngineVersion string,
	spotPercentage *int32, spotMaxPrice *float64, withRangerRazEnabled, withEnableMultiAz bool, azureFlexibleServer bool,
	imageCatalog, imageId string) *sdxModel.SdxCustomClusterRequest {

	var imageSettings *sdxModel.ImageSettingsV4Request = nil
	if len(imageCatalog) > 0 || len(imageId) > 0 {
		imageSettings = &sdxModel.ImageSettingsV4Request{}
		if len(imageCatalog) > 0 {
			imageSettings.Catalog = imageCatalog
		}
		if len(imageId) > 0 {
			imageSettings.ID = imageId
		}
	}

	sdxRequest := &sdxModel.SdxCustomClusterRequest{
		ClusterShape:           &clusterShape,
		Environment:            &envName,
		Runtime:                runtime,
		Tags:                   nil,
		CloudStorage:           nil,
		ExternalDatabase:       nil,
		Aws:                    nil,
		EnableRangerRaz:        withRangerRazEnabled,
		EnableMultiAz:          withEnableMultiAz,
		ImageSettingsV4Request: imageSettings,
	}

	setupCloudStorageIfNeeded(cloudStorageBaseLocation, instanceProfile, managedIdentity, serviceAccount, func(storage *sdxModel.SdxCloudStorageRequest) { sdxRequest.CloudStorage = storage })
	setupExternalDbIfNeeded(withExternalDatabase, withoutExternalDatabase, withNonHaExternalDatabase, databaseEngineVersion, azureFlexibleServer, &sdxRequest.ExternalDatabase)
	setupAwsSpotParameters(spotPercentage, spotMaxPrice, &sdxRequest.Aws)
	return sdxRequest
}

func setupExternalDbIfNeeded(withExternalDatabase, withoutExternalDatabase, withNonHaExternalDatabase bool, databaseEngineVersion string, azureFlexibleServer bool, sdxDatabase **sdxModel.SdxDatabaseRequest) {
	if withoutExternalDatabase && withExternalDatabase && withNonHaExternalDatabase {
		commonutils.LogErrorAndExit(errors.New("both withExternalDatabase and withoutExternalDatabase were set"))
	}
	if withNonHaExternalDatabase {
		externalDatabase := &sdxModel.SdxDatabaseRequest{
			AvailabilityType: "NON_HA",
		}
		*sdxDatabase = externalDatabase
	}
	if withExternalDatabase {
		externalDatabase := &sdxModel.SdxDatabaseRequest{
			AvailabilityType: "HA",
		}
		*sdxDatabase = externalDatabase
	}
	if withoutExternalDatabase {
		externalDatabase := &sdxModel.SdxDatabaseRequest{
			AvailabilityType: "NONE",
		}
		*sdxDatabase = externalDatabase
	}
	if len(databaseEngineVersion) != 0 {
		(*sdxDatabase).DatabaseEngineVersion = databaseEngineVersion
	}
	if azureFlexibleServer {
		(*sdxDatabase).SdxDatabaseAzureRequest = &model.SdxDatabaseAzureRequest{
			AzureDatabaseType: "FLEXIBLE_SERVER",
		}
	}
}

type CloudStorageSetter func(*sdxModel.SdxCloudStorageRequest)

func (s CloudStorageSetter) SetCloudStorage(req *sdxModel.SdxCloudStorageRequest) {
	s(req)
}

func setupCloudStorageIfNeeded(cloudStorageBaseLocation string, instanceProfile string, managedIdentity string, serviceAccount string, setter CloudStorageSetter) {
	if len(instanceProfile) > 0 && len(managedIdentity) > 0 {
		commonutils.LogErrorMessageAndExit("Please use managed identity or instance profile, but not both")
	}
	if len(instanceProfile) > 0 && len(cloudStorageBaseLocation) == 0 {
		commonutils.LogErrorMessageAndExit("Please use cloud storage base location if you define instance profile")
	}
	if len(managedIdentity) > 0 && len(cloudStorageBaseLocation) == 0 {
		commonutils.LogErrorMessageAndExit("Please use cloud storage base location if you define managed identity")
	}
	if len(serviceAccount) > 0 && len(cloudStorageBaseLocation) == 0 {
		commonutils.LogErrorMessageAndExit("Please use cloud storage base location if you define service account")
	}
	if len(cloudStorageBaseLocation) > 0 && len(instanceProfile) == 0 && len(managedIdentity) == 0 && len(serviceAccount) == 0 {
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
	if len(cloudStorageBaseLocation) > 0 && len(serviceAccount) > 0 {
		gcsStorage := &sdxModel.GcsCloudStorageV1Parameters{
			ServiceAccountEmail: &serviceAccount,
		}
		cloudStorage := &sdxModel.SdxCloudStorageRequest{
			Adls:           nil,
			AdlsGen2:       nil,
			BaseLocation:   cloudStorageBaseLocation,
			FileSystemType: "GCS",
			Gcs:            gcsStorage,
			S3:             nil,
			Wasb:           nil,
		}
		setter.SetCloudStorage(cloudStorage)
	}
}

func setupAwsSpotParameters(spotPercentage *int32, spotMaxPrice *float64, sdxAws **sdxModel.SdxAwsRequest) {
	if spotPercentage != nil || spotMaxPrice != nil {
		aws := &sdxModel.SdxAwsRequest{
			Spot: &sdxModel.SdxAwsSpotParameters{},
		}
		if spotPercentage != nil {
			aws.Spot.Percentage = spotPercentage
		}
		if spotMaxPrice != nil {
			aws.Spot.MaxPrice = *spotMaxPrice
		}
		*sdxAws = aws
	}
}

func createInternalSdx(envName string, inputJson *sdxModel.StackV4Request, c *cli.Context,
	name, cloudStorageBaseLocation, instanceProfile, managedIdentity, serviceAccount, runtime string,
	withoutExternalDatabase, withExternalDatabase, withNonHaExternalDatabase bool, databaseEngineVersion string,
	spotPercentage *int32, spotMaxPrice *float64, withRangerRazEnabled bool, withEnableMultiAz bool, azureFlexibleServer bool) {
	sdxInternalRequest := &sdxModel.SdxInternalClusterRequest{
		ClusterShape:     &(&types.S{S: sdxModel.SdxClusterRequestClusterShapeCUSTOM}).S,
		Environment:      &envName,
		Runtime:          runtime,
		StackV4Request:   inputJson,
		Tags:             nil,
		ExternalDatabase: nil,
		Aws:              nil,
		EnableRangerRaz:  withRangerRazEnabled,
		EnableMultiAz:    withEnableMultiAz,
	}

	setupCloudStorageIfNeeded(cloudStorageBaseLocation, instanceProfile, managedIdentity, serviceAccount, CloudStorageSetter(func(storage *sdxModel.SdxCloudStorageRequest) { sdxInternalRequest.CloudStorage = storage }))
	setupExternalDbIfNeeded(withExternalDatabase, withoutExternalDatabase, withNonHaExternalDatabase, databaseEngineVersion, azureFlexibleServer, &sdxInternalRequest.ExternalDatabase)
	setupAwsSpotParameters(spotPercentage, spotMaxPrice, &sdxInternalRequest.Aws)

	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	checkClientVersion(sdxClient, common.Version)
	resp, err := sdxClient.Internalsdx.CreateInternalSdx(internalsdx.NewCreateInternalSdxParams().WithName(name).WithBody(sdxInternalRequest))
	if err != nil {
		commonutils.LogErrorAndExit(err)
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
		commonutils.LogErrorAndExit(err)
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
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[RepairSdx] SDX cluster repair is started for: %s", name)
}

func formulateRequest(hostGroupToRepair string, hostGroupsToRepair []string) *sdxModel.SdxRepairRequest {
	if (len(hostGroupToRepair) == 0) == (len(hostGroupsToRepair[0]) == 0) {
		commonutils.LogErrorMessageAndExit(fmt.Sprintf("Please only specify either %s or %s", fl.FlHostGroupOptional.Name, fl.FlHostGroupsOptional.Name))
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
	defer commonutils.TimeTrack(time.Now(), "List sdx clusters in environment")
	envName := c.String(fl.FlEnvironmentNameOptional.Name)
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c))
	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	writer := output.WriteList
	listSdxClusterImpl(sdxClient.Sdx.Sdx, envName, writer)
	log.Infof("[ListSdx] SDX cluster list in environment: %s", envName)
}

func GetListOfSdx(c *cli.Context) []*sdxModel.SdxClusterResponse {
	defer commonutils.TimeTrack(time.Now(), "List sdx clusters by environment crn")
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c))
	resp, err := sdxClient.Sdx.Sdx.ListSdx(sdx.NewListSdxParams())
	if err != nil {
		return []*sdxModel.SdxClusterResponse{}
	}
	return resp.Payload
}

func listSdxClusterImpl(client clientSdx, envName string, writer func([]string, []commonutils.Row)) {
	resp, err := client.ListSdx(sdx.NewListSdxParams().WithEnvName(&envName))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}

	var tableRows []commonutils.Row
	for _, sdxCluster := range resp.Payload {
		tableRows = append(tableRows, &sdxClusterOutput{sdxCluster.Crn, sdxCluster.Name,
			sdxCluster.EnvironmentName,
			sdxCluster.EnvironmentCrn,
			sdxCluster.StackCrn,
			sdxCluster.DatabaseServerCrn,
			sdxCluster.CloudStorageBaseLocation,
			sdxCluster.CloudStorageFileSystemType,
			sdxCluster.Status,
			sdxCluster.StatusReason,
			sdxCluster.RangerRazEnabled,
			sdxCluster.EnableMultiAz})
	}

	writer(sdxClusterHeader, tableRows)
}

func SyncSdx(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Sync sdx cluster in environment")
	name := c.String(fl.FlName.Name)
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	checkClientVersion(sdxClient, common.Version)
	err := sdxClient.Sdx.SyncSdx(sdx.NewSyncSdxParams().WithName(name))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[SyncSdx] SDX cluster sync started for: %s", name)
}

func ListSdxAdvertisedRuntimesSdx(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "List sdx adertised runtimes")
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c))
	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	writer := output.WriteList
	listSdxAdvertisedRuntimesImpl(sdxClient.Sdx.Sdx, writer)
	log.Infof("[AdvertisedRuntimesSdx] SDX advertised runtimes")
}

func listSdxAdvertisedRuntimesImpl(client clientRuntimeSdx, writer func([]string, []commonutils.Row)) {
	resp, err := client.Advertisedruntimes(sdx.NewAdvertisedruntimesParams())
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}

	var tableRows []commonutils.Row
	for _, advertiseRuntime := range resp.Payload {
		tableRows = append(tableRows, &sdxAdvertisedRuntimesOutput{advertiseRuntime.RuntimeVersion, advertiseRuntime.DefaultRuntimeVersion})
	}

	writer(sdxAdvertisedRuntimesHeader, tableRows)
}

func RetrySdx(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Retry sdx cluster")
	name := c.String(fl.FlName.Name)
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	checkClientVersion(sdxClient, common.Version)
	_, err := sdxClient.Sdx.RetrySdx(sdx.NewRetrySdxParams().WithName(name))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[RetrySdx] SDX cluster retry started for: %s", name)
}

func createCloudStorageRequestForSdx() {

}

func DescribeSdx(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "describe SDX cluster")
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
		commonutils.LogErrorAndExit(err)
	}
	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
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
		sdxCluster.StatusReason,
		sdxCluster.RangerRazEnabled,
		sdxCluster.EnableMultiAz})
	log.Infof("[DescribeSdx] Describe a particular SDX cluster")
}

func describeSdxDetailed(sdxClient *sdx.Client, name string, c *cli.Context) {
	resp, err := sdxClient.GetSdxDetail(sdx.NewGetSdxDetailParams().WithName(name))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
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
			sdxCluster.StatusReason,
			sdxCluster.RangerRazEnabled,
			sdxCluster.EnableMultiAz},
		StackV4Response: *sdxCluster.StackV4Response,
	})
	log.Infof("[DescribeSdxDetailed] Describe a particular SDX cluster")
}

func StartSdx(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Start sdx cluster by name")
	name := c.String(fl.FlName.Name)
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	_, err := sdxClient.Sdx.StartSdxByName(sdx.NewStartSdxByNameParams().WithName(name))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[StartSdx] SDX cluster start executed for: %s", name)
}

func StopSdx(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Stop sdx cluster by name")
	name := c.String(fl.FlName.Name)
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	_, err := sdxClient.Sdx.StopSdxByName(sdx.NewStopSdxByNameParams().WithName(name))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[StopSdx] SDX cluster stop executed for: %s", name)
}

func printResponse(template *sdx.UpgradeDatalakeClusterOK, dryRun bool) error {
	var errorMessage error
	var response []byte

	if dryRun {
		resp, err := json.MarshalIndent(template.Payload, "", "\t")
		response = resp
		errorMessage = err
	} else if template.Payload.UpgradeCandidates == nil || len(template.Payload.UpgradeCandidates) == 0 {
		log.Infof("[UpgradeSDX] Reason field received: %s", template.Payload.Reason)
		resp, err := json.MarshalIndent(template.Payload.Reason, "", "\t")
		response = resp
		errorMessage = err
	} else {
		resp, err := json.MarshalIndent(template.Payload, "", "\t")
		response = resp
		errorMessage = err
	}

	if errorMessage != nil {
		commonutils.LogErrorAndExit(errorMessage)
	}
	commonutils.Println(string(response))
	return nil
}

func printResponseForPrepare(template *sdx.PrepareDatalakeClusterUpgradeOK, dryRun bool) error {
	var errorMessage error
	var response []byte

	if dryRun {
		resp, err := json.MarshalIndent(template.Payload, "", "\t")
		response = resp
		errorMessage = err
	} else if template.Payload.UpgradeCandidates == nil || len(template.Payload.UpgradeCandidates) == 0 {
		log.Infof("[UpgradeSDX] Reason field received: %s", template.Payload.Reason)
		resp, err := json.MarshalIndent(template.Payload.Reason, "", "\t")
		response = resp
		errorMessage = err
	} else {
		resp, err := json.MarshalIndent(template.Payload, "", "\t")
		response = resp
		errorMessage = err
	}

	if errorMessage != nil {
		commonutils.LogErrorAndExit(errorMessage)
	}
	commonutils.Println(string(response))
	return nil
}

func SdxClusterUpgrade(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Start SDX upgrade")
	dryRun := c.Bool(fl.FlDryRunOptional.Name)
	name := c.String(fl.FlName.Name)
	image := c.String(fl.FlImageIdOptional.Name)
	runtime := c.String(fl.FlRuntimeOptional.Name)
	lock := c.Bool(fl.FlLockComponentsOptional.Name)
	replaceVms := c.String(fl.FlReplaceVms.Name)
	showImages := c.Bool(fl.FlShowImagesOptional.Name)
	showLatestImages := c.Bool(fl.FlShowLatestImagesOptional.Name)
	backup := c.Bool(fl.FlUpgradeBackup.Name)
	rollingUpgrade := c.Bool(fl.FlRollingUpgradeOptional.Name)

	sdxRequest := createSdxUpgradeRequest(image, runtime, lock, dryRun, rollingUpgrade, replaceVms, showImages, showLatestImages, backup)
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	checkClientVersion(sdxClient, common.Version)

	resp, err := sdxClient.Sdx.UpgradeDatalakeCluster(sdx.NewUpgradeDatalakeClusterParams().WithName(name).WithBody(sdxRequest))
	if err != nil {
		commonutils.LogErrorAndExit(err)
		fmt.Printf("%s\n", err)
	}
	printResponse(resp, dryRun || showImages || showLatestImages)
}

func SdxClusterUpgradePrepare(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Start SDX upgrade")
	dryRun := c.Bool(fl.FlDryRunOptional.Name)
	name := c.String(fl.FlName.Name)
	image := c.String(fl.FlImageIdOptional.Name)
	runtime := c.String(fl.FlRuntimeOptional.Name)
	showImages := c.Bool(fl.FlShowImagesOptional.Name)
	showLatestImages := c.Bool(fl.FlShowLatestImagesOptional.Name)

	sdxRequest := createSdxUpgradeRequest(image, runtime, false, dryRun, false, "", showImages, showLatestImages, false)
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	checkClientVersion(sdxClient, common.Version)

	resp, err := sdxClient.Sdx.PrepareDatalakeClusterUpgrade(sdx.NewPrepareDatalakeClusterUpgradeParams().WithName(name).WithBody(sdxRequest))
	if err != nil {
		commonutils.LogErrorAndExit(err)
		fmt.Printf("%s\n", err)
	}
	printResponseForPrepare(resp, dryRun || showImages || showLatestImages)
}

func VerticalScaleSdx(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "vertical scale SDX")

	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	instanceType := c.String(fl.FlInstanceType.Name)

	req := &sdxModel.StackVerticalScaleV4Request{
		Group: &(&types.S{S: c.String(fl.FlGroupName.Name)}).S,
		Template: &sdxModel.InstanceTemplateV4Request{
			InstanceType: instanceType,
		},
	}
	name := c.String(fl.FlName.Name)
	log.Infof("[ScaleSDX] scaling SDX, name: %s", name)
	sdxClient.Sdx.VerticalScalingByName(sdx.NewVerticalScalingByNameParams().WithName(name).WithBody(req))

	log.Infof("[VerticalScaleSdx] SDX vertical scaled, name: %s", name)
}

func createSdxUpgradeRequest(imageId, runtime string, lockComponents, dryRun, rollingUpgrade bool, replaceVms string,
	showImages, showLatestImages, backup bool) *sdxModel.SdxUpgradeRequest {

	var showImagesString string
	if showLatestImages {
		showImagesString = "LATEST_ONLY"
	} else if showImages {
		showImagesString = "SHOW"
	} else {
		showImagesString = "DO_NOT_SHOW"
	}
	sdxRequest := &sdxModel.SdxUpgradeRequest{
		ImageID:               imageId,
		Runtime:               runtime,
		LockComponents:        lockComponents,
		DryRun:                dryRun,
		ReplaceVms:            replaceVms,
		ShowAvailableImages:   showImagesString,
		SkipBackup:            !backup,
		RollingUpgradeEnabled: rollingUpgrade,
	}
	return sdxRequest
}

func checkClientVersion(client *client.Datalake, version string) {
	versionCheckRequest := sdxutils.NewCheckClientVersionOfSdxParams().WithVersion(&version)
	resp, err := client.Sdxutils.CheckClientVersionOfSdx(versionCheckRequest)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	valid := resp.Payload.VersionCheckOk
	message := resp.Payload.Message
	if !valid {
		commonutils.LogErrorAndExit(errors.New(message))
	}
}

func GetVmLogs(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "get default monitored vm logs for DistroX CM clusters")
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	result, err := sdxClient.Diagnostics.GetSdxCmVMLogs(diagnostics.NewGetSdxCmVMLogsParams())
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	logs := result.Payload.Logs
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(logs); err != nil {
		commonutils.LogErrorAndExit(err)
	}
	fmt.Println(buf.String())
}

func CollectDiagnostics(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "get user synchronization state for an environment")
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	collectionRequest := assembleCollectionRequest(c)
	_, err := sdxClient.Diagnostics.CollectSdxCmDiagnostics(diagnostics.NewCollectSdxCmDiagnosticsParams().WithBody(collectionRequest))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	fmt.Println("Collection started")
}

func ListDiagnosticsCollections(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "get latest SDX diagnostics collection flows")
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	resp, err := sdxClient.Diagnostics.ListSdxDiagnosticsCollections(diagnostics.NewListSdxDiagnosticsCollectionsParams().WithCrn(c.String(fl.FlCrn.Name)))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	collections := resp.Payload.Collections
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(collections); err != nil {
		commonutils.LogErrorAndExit(err)
	}
	fmt.Println(buf.String())
}

func CancelDiagnosticsCollections(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "cancel running SDX diagnostics collection flows")
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	err := sdxClient.Diagnostics.CancelSdxDiagnosticsCollections(diagnostics.NewCancelSdxDiagnosticsCollectionsParams().WithCrn(c.String(fl.FlCrn.Name)))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	fmt.Println("Cancel running SDX diagnostics collections...")
}

func assembleCollectionRequest(c *cli.Context) *sdxModel.DiagnosticsCollectionRequest {
	stackCrn := c.String(fl.FlCrn.Name)
	collectionOnly := c.Bool(fl.FlCollectionOnly.Name)
	updatePackage := c.Bool(fl.FlUpdatePackage.Name)
	includeSaltLogs := c.Bool(fl.FlIncludeSaltLogs.Name)
	destinationOption := c.String(fl.FlCollectionDestination.Name)
	destination := "CLOUD_STORAGE"
	labelsOption := c.String(fl.FlCollectionLabels.Name)
	labels := make([]string, 0)
	if len(labelsOption) > 0 {
		labels = strings.Split(labelsOption, ",")
	}
	description := c.String(fl.FlDescriptionOptional.Name)
	issue := c.String(fl.FlCollectionIssue.Name)
	now := time.Now()
	startTime := strfmt.DateTime(now.AddDate(-10, 0, 0))
	endTime := strfmt.DateTime(now.AddDate(10, 0, 0))
	additionalLogsFileOption := c.String(fl.FlCollectionExtraLogsFile.Name)
	var additionalLogs []*model.VMLog
	if len(additionalLogsFileOption) > 0 {
		content := commonutils.ReadFile(additionalLogsFileOption)
		err := json.Unmarshal(content, &additionalLogs)
		if err != nil {
			commonutils.LogErrorAndExit(err)
		}
	}
	hostsOption := c.String(fl.FlCollectionHosts.Name)
	var hosts []string
	if len(hostsOption) > 0 {
		hosts = strings.Split(hostsOption, ",")
	}
	hostGroupsOption := c.String(fl.FlCollectionHostGroups.Name)
	var hostGroups []string
	if len(hostGroupsOption) > 0 {
		hostGroups = strings.Split(hostGroupsOption, ",")
	}

	if collectionOnly {
		destination = "LOCAL"
	} else if len(destinationOption) > 0 {
		if destinationOption == "ENG" {
			destination = "ENG"
		} else if destinationOption == "SUPPORT" {
			destination = "SUPPORT"
		}
	}
	request := sdxModel.DiagnosticsCollectionRequest{Destination: &destination, StackCrn: &stackCrn, Labels: labels,
		Issue: issue, Description: description, StartTime: startTime, EndTime: endTime, AdditionalLogs: additionalLogs,
		IncludeSaltLogs: includeSaltLogs, UpdatePackage: updatePackage, Hosts: hosts, HostGroups: hostGroups}
	return &request
}

func GetCMRoles(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "get CM roles for sdx cluster")
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	stackCrn := c.String(fl.FlCrn.Name)
	result, err := sdxClient.Diagnostics.GetSdxCmRoles(diagnostics.NewGetSdxCmRolesParams().WithStackCrn(stackCrn))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(result.Payload); err != nil {
		commonutils.LogErrorAndExit(err)
	}
	fmt.Println(buf.String())
}

func CollectCmDiagnostics(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "collect CM based distrox diagnostics")
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	cmCollectionRequest := assembleCMCollectionRequest(c)
	_, err := sdxClient.Diagnostics.CollectSdxCmBasedDiagnostics(diagnostics.NewCollectSdxCmBasedDiagnosticsParams().WithBody(cmCollectionRequest))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	fmt.Println("CM based collection started")
}

func assembleCMCollectionRequest(c *cli.Context) *sdxModel.CmDiagnosticsCollectionRequest {
	stackCrn := c.String(fl.FlCrn.Name)
	collectionOnly := c.Bool(fl.FlCollectionOnly.Name)
	updatePackage := c.Bool(fl.FlUpdatePackage.Name)
	monitorMetricsCollection := c.Bool(fl.FlMonitorMetricsCollection.Name)
	destinationOption := c.String(fl.FlCollectionDestination.Name)
	destination := "CLOUD_STORAGE"
	if collectionOnly {
		destination = "LOCAL"
	} else if len(destinationOption) > 0 {
		if destinationOption == "ENG" {
			destination = "ENG"
		} else if destinationOption == "SUPPORT" {
			destination = "SUPPORT"
		}
	}
	rolesOption := c.String(fl.FlCollectionRoles.Name)
	roles := make([]string, 0)
	if len(rolesOption) > 0 {
		roles = strings.Split(rolesOption, ",")
	}
	description := c.String(fl.FlDescriptionOptional.Name)
	issue := c.String(fl.FlCollectionIssue.Name)
	now := time.Now()
	startTime := strfmt.DateTime(now.AddDate(-10, 0, 0))
	endTime := strfmt.DateTime(now.AddDate(10, 0, 0))
	request := sdxModel.CmDiagnosticsCollectionRequest{Destination: &destination, StackCrn: &stackCrn, UpdatePackage: updatePackage,
		Roles: roles, StartTime: startTime, EndTime: endTime, Ticket: issue, Comments: description, EnableMonitorMetricsCollection: monitorMetricsCollection}
	return &request
}

func OperationProgress(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "check sdx provisioning flow")
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	stackCrn := c.String(fl.FlCrn.Name)
	result, err := sdxClient.Operation.GetOperationProgressByResourceCrn(operation.NewGetOperationProgressByResourceCrnParams().WithResourceCrn(stackCrn))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(result.Payload); err != nil {
		commonutils.LogErrorAndExit(err)
	}
	fmt.Println(buf.String())
}

func RotateCertificates(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Rotate AutoTLS certificates for sdx cluster")
	name := c.String(fl.FlName.Name)
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	body := model.CertificatesRotationV4Request{CertificateRotationType: model.CertificatesRotationV4RequestCertificateRotationTypeHOSTCERTS}
	_, err := sdxClient.Sdx.RotateAutoTLSCertificatesByName(sdx.NewRotateAutoTLSCertificatesByNameParams().WithName(name).WithBody(&body))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[RotateCerts] SDX cluster certificate rotation started for: %s", name)
}

func SdxClusterRecover(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Start SDX recovery")
	FlRecoverWithoutData := c.Bool(fl.FlRecoverWithoutDataOptional.Name)
	var recoverWithoutData string
	if FlRecoverWithoutData {
		recoverWithoutData = "RECOVER_WITHOUT_DATA"
	} else {
		recoverWithoutData = "RECOVER_WITH_DATA"
	}
	name := c.String(fl.FlName.Name)

	recoverSdxRequest := &sdxModel.SdxRecoveryRequest{
		Type: recoverWithoutData,
	}
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	checkClientVersion(sdxClient, common.Version)

	_, err := sdxClient.Sdx.RecoverDatalakeCluster(sdx.NewRecoverDatalakeClusterParams().WithName(name).WithBody(recoverSdxRequest))
	if err != nil {
		commonutils.LogErrorAndExit(err)
		fmt.Printf("%s\n", err)
	}
	log.Infof("[SdxClusterRecover] SDX cluster recover started for: %s", name)
}

func UpgradeDatabase(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Rotate AutoTLS certificates for sdx cluster")
	name := c.String(fl.FlName.Name)
	targetVersion := c.String(fl.FlTargetVersionOptional.Name)
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	body := model.SdxUpgradeDatabaseServerRequest{TargetMajorVersion: targetVersion}
	resp, err := sdxClient.Sdx.UpgradeDatalakeDatabaseByName(sdx.NewUpgradeDatalakeDatabaseByNameParams().WithName(name).WithBody(&body))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("SDX cluster database upgrade started for: %s", name)
	response, err := json.MarshalIndent(resp.Payload, "", "\t")
	if err != nil {
		commonutils.LogErrorAndExit(err)
	} else {
		commonutils.Println(string(response))
	}
}

func SyncComponentVersions(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Sync sdx cluster component versions from CM")
	name := c.String(fl.FlName.Name)
	sdxClient := ClientSdx(*oauth.NewSDXClientFromContext(c)).Sdx
	checkClientVersion(sdxClient, common.Version)
	_, err := sdxClient.Sdx.SyncCmOnDatalakeCluster(sdx.NewSyncCmOnDatalakeClusterParams().WithName(name))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[SyncVersionsSdx] SDX cluster component versions sync started for: %s", name)
}
