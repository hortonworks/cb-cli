package freeipa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hortonworks/cb-cli/dataplane/api-freeipa/client/v1dns"

	"os"

	"github.com/go-openapi/strfmt"
	"github.com/hortonworks/cb-cli/dataplane/api-freeipa/client/v1diagnostics"
	"github.com/hortonworks/cb-cli/dataplane/api-freeipa/client/v1freeipa"
	"github.com/hortonworks/cb-cli/dataplane/api-freeipa/client/v1freeipauser"
	"github.com/hortonworks/cb-cli/dataplane/api-freeipa/client/v1operation"
	"github.com/hortonworks/cb-cli/dataplane/api-freeipa/client/v1utils"
	"github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
	freeIpaModel "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
	"github.com/hortonworks/cb-cli/dataplane/env"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/oauth"
	"github.com/hortonworks/dp-cli-common/utils"
	commonutils "github.com/hortonworks/dp-cli-common/utils"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type ClientFreeIpa oauth.FreeIpa

type freeIpaOutDescibe struct {
	FreeIpa *freeIpaModel.DescribeFreeIpaV1Response `json:"freeIpa" yaml:"freeIpa"`
}

var header = []string{"CRN", "Name", "Status", "Status reason"}

func CreateFreeIpa(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "create FreeIpa cluster")
	FreeIpaRequest := assembleFreeIpaRequest(c)
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	resp, err := freeIpaClient.V1freeipa.CreateFreeIpaV1(v1freeipa.NewCreateFreeIpaV1Params().WithBody(FreeIpaRequest), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	freeIpaCluster := resp.Payload
	if freeIpaCluster.Name != nil {
		log.Infof("[createFreeIpa] FreeIpa cluster created with name: %s", *freeIpaCluster.Name)
	}
}

func DeleteFreeIpa(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "delete FreeIpa cluster")
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	err := freeIpaClient.V1freeipa.DeleteFreeIpaByEnvironmentV1(v1freeipa.NewDeleteFreeIpaByEnvironmentV1Params().WithEnvironment(&envCrn), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[deleteFreeIpa] FreeIpa cluster delete requested in environment %s", envName)
}

func DescribeFreeIpa(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "describe FreeIpa cluster")
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	resp, err := freeIpaClient.V1freeipa.GetFreeIpaByEnvironmentV1(v1freeipa.NewGetFreeIpaByEnvironmentV1Params().WithEnvironment(&envCrn), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[describeFreeIpa] FreeIpa cluster describe requested in enviornment %s", envName)
	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	iparesp := resp.Payload
	freeIpaOut := freeIpaOutDescibe{
		iparesp,
	}

	output.Write(header, &freeIpaOut)
}

func StartFreeIpa(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "start FreeIpa cluster")
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	err := freeIpaClient.V1freeipa.StartV1(v1freeipa.NewStartV1Params().WithEnvironment(&envCrn), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[startFreeIpa] FreeIpa cluster start requested in enviornment %s", envName)
}

func StopFreeIpa(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "stop FreeIpa cluster")
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	err := freeIpaClient.V1freeipa.StopV1(v1freeipa.NewStopV1Params().WithEnvironment(&envCrn), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[stopFreeIpa] FreeIpa cluster stop requested in enviornment %s", envName)
}

func RebootFreeIpa(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "reboot FreeIpa cluster")
	envName := c.String(fl.FlEnvironmentName.Name)
	RebootInstancesV1Request := assembleRebootRequest(c)
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	resp, err := freeIpaClient.V1freeipa.RebootV1(v1freeipa.NewRebootV1Params().WithBody(RebootInstancesV1Request), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[rebootFreeIpa] FreeIpa cluster reboot requested in environment %s", envName)
	operationStatus := resp.Payload

	if c.Bool(fl.FlWaitOptional.Name) && operationStatus.Status == "RUNNING" {
		getOperationStatus(c, *operationStatus.OperationID, "reboot")
	} else {
		writeOperationStatus(c, operationStatus)
	}
}

func assembleRebootRequest(c *cli.Context) *freeIpaModel.RebootInstancesV1Request {
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)
	force := c.Bool(fl.FlForceOptional.Name)
	nodes := c.String(fl.FlNodesOptional.Name)
	return &freeIpaModel.RebootInstancesV1Request{
		EnvironmentCrn: &envCrn,
		ForceReboot:    force,
		InstanceIds:    strings.Split(nodes, ","),
	}
}

func RepairFreeIpa(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "repair FreeIpa cluster")
	envName := c.String(fl.FlEnvironmentName.Name)
	RepairInstancesV1Request := assembleRepairRequest(c)
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	resp, err := freeIpaClient.V1freeipa.RepairV1(v1freeipa.NewRepairV1Params().WithBody(RepairInstancesV1Request), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[repairFreeIpa] FreeIpa cluster repair requested in environment %s", envName)
	operationStatus := resp.Payload

	if c.Bool(fl.FlWaitOptional.Name) && operationStatus.Status == "RUNNING" {
		getOperationStatus(c, *operationStatus.OperationID, "repair")
	} else {
		writeOperationStatus(c, operationStatus)
	}
}

func getOperationStatus(c *cli.Context, operationId string, command string) {
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	resp, err := freeIpaClient.V1operation.GetOperationStatusV1(v1operation.NewGetOperationStatusV1Params().WithOperationID(operationId))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	operationStatus := resp.Payload
	for c.Bool(fl.FlWaitOptional.Name) && operationStatus.Status == "RUNNING" {
		log.Infof("[%s] Status is RUNNING. Sleeping", command)
		time.Sleep(5 * time.Second)
		resp, err := freeIpaClient.V1operation.GetOperationStatusV1(v1operation.NewGetOperationStatusV1Params().WithOperationID(operationId))
		if err != nil {
			commonutils.LogErrorAndExit(err)
		}
		operationStatus = resp.Payload
	}
	log.Infof("[%s] Operation completed: %s", command, operationStatus.Status)
	writeOperationStatus(c, operationStatus)
}

func writeOperationStatus(c *cli.Context, operationStatus *freeIpaModel.OperationV1Status) {
	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	operationStatusOut := &freeIpaOutOperation{
		ID:        *operationStatus.OperationID,
		Status:    operationStatus.Status,
		Success:   convertSuccessDetailsModel(operationStatus.Success),
		Failure:   convertFailureDetailsModel(operationStatus.Failure),
		Error:     operationStatus.Error,
		StartTime: strconv.FormatInt(operationStatus.StartTime, 10),
		EndTime:   strconv.FormatInt(operationStatus.EndTime, 10),
	}
	output.Write(operationStatusHeader, operationStatusOut)
}

func assembleRepairRequest(c *cli.Context) *freeIpaModel.RepairInstancesV1Request {
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)
	force := c.Bool(fl.FlForceOptional.Name)
	nodes := c.String(fl.FlNodesOptional.Name)
	return &freeIpaModel.RepairInstancesV1Request{
		EnvironmentCrn: &envCrn,
		ForceRepair:    force,
		InstanceIds:    strings.Split(nodes, ","),
	}
}

func ListFreeIpa(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "list FreeIpa clusters")
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	resp, err := freeIpaClient.V1freeipa.ListFreeIpaClustersByAccountV1(v1freeipa.NewListFreeIpaClustersByAccountV1Params(), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[listFreeIpa] FreeIpa clusters list requested.")
	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}

	var tableRows []commonutils.Row

	for _, response := range resp.Payload {
		row := &freeIpaDetails{
			Name:           *response.Name,
			CRN:            *response.Crn,
			EnvironmentCrn: *response.EnvironmentCrn,
			Status:         response.Status,
			Domain:         response.Domain,
		}
		tableRows = append(tableRows, row)
	}
	output.WriteList(listHeader, tableRows)
}

func RemoteExecFreeIpa(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "remote exec on FreeIpa cluster")
	envName := c.String(fl.FlEnvironmentName.Name)
	remoteCommand := c.String(fl.FlRemoteCommand.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)
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
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	req := freeIpaModel.RemoteCommandsExecutionRequest{Command: remoteCommand, Hosts: hosts, HostGroups: hostGroups}
	resp, err := freeIpaClient.V1utils.RemoteExecV1(v1utils.NewRemoteExecV1Params().WithEnvironmentCrn(envCrn).WithBody(&req))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	commandOutputs := resp.Payload.Results
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(commandOutputs); err != nil {
		utils.LogErrorAndExit(err)
	}
	fmt.Println(buf.String())
}

func (f *freeIpaOutDescibe) DataAsStringArray() []string {
	return []string{
		*f.FreeIpa.Crn,
		*f.FreeIpa.Name,
		f.FreeIpa.Status,
		f.FreeIpa.StatusReason,
	}
}

func assembleFreeIpaRequest(c *cli.Context) *freeIpaModel.CreateFreeIpaV1Request {
	path := c.String(fl.FlInputJson.Name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		commonutils.LogErrorAndExit(err)
	}

	log.Infof("[assembleStackTemplate] read cluster create json from file: %s", path)
	content := commonutils.ReadFile(path)

	var req freeIpaModel.CreateFreeIpaV1Request
	err := json.Unmarshal(content, &req)
	if err != nil {
		msg := fmt.Sprintf(`Invalid json format: %s. Please make sure that the json is valid (check for commas and double quotes).`, err.Error())
		commonutils.LogErrorMessageAndExit(msg)
	}

	name := c.String(fl.FlName.Name)
	if len(name) != 0 {
		req.Name = &name
	}
	if req.Name == nil || len(*req.Name) == 0 {
		commonutils.LogErrorMessageAndExit("Name of the cluster must be set either in the template or with the --name command line option.")
	}
	if req.EnvironmentCrn == nil || len(*req.EnvironmentCrn) == 0 {
		environmentName := c.String(fl.FlEnvironmentNameOptional.Name)
		if len(environmentName) == 0 {
			commonutils.LogErrorMessageAndExit("Name of the environment must be set either in the template or with the --env-name command line option.")
		}
		crn := env.GetEnvirontmentCrnByName(c, environmentName)
		req.EnvironmentCrn = &crn
	}
	return &req
}

func SynchronizeAllUsers(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "sync all users to FreeIpa")

	users := c.StringSlice(fl.FlIpaUserCrnsSlice.Name)
	machineUsers := c.StringSlice(fl.FlIpaMachineUserCrnsSlice.Name)
	environments := c.StringSlice(fl.FlIpaEnvironmentCrnsSlice.Name)
	SynchronizeAllUsersV1Request := freeIpaModel.SynchronizeAllUsersV1Request{
		Users:        users,
		MachineUsers: machineUsers,
		Environments: environments,
	}

	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	resp, err := freeIpaClient.V1freeipauser.SynchronizeAllUsersV1(v1freeipauser.NewSynchronizeAllUsersV1Params().WithBody(&SynchronizeAllUsersV1Request))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	syncOperationStatus := resp.Payload

	if c.Bool(fl.FlWaitOptional.Name) && syncOperationStatus.Status == "RUNNING" {
		getSyncOperationStatus(c, *syncOperationStatus.OperationID, "synchronizeAllUsers")
	} else {
		writeSyncOperationStatus(c, syncOperationStatus)
	}
}

func SynchronizeCurrentUser(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "sync current user to FreeIpa")

	SynchronizeUserV1Request := freeIpaModel.SynchronizeUserV1Request{
		Environments: constructEnvCrnList(c),
	}

	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	resp, err := freeIpaClient.V1freeipauser.SynchronizeUserV1(v1freeipauser.NewSynchronizeUserV1Params().WithBody(&SynchronizeUserV1Request))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	syncOperationStatus := resp.Payload

	if c.Bool(fl.FlWaitOptional.Name) && syncOperationStatus.Status == "RUNNING" {
		getSyncOperationStatus(c, *syncOperationStatus.OperationID, "synchronizeUser")
	} else {
		writeSyncOperationStatus(c, syncOperationStatus)
	}
}

func SetPassword(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "set user password in FreeIpa")

	password := c.String(fl.FlIpaUserPassword.Name)
	SetPasswordV1Request := freeIpaModel.SetPasswordV1Request{
		Environments: constructEnvCrnList(c),
		Password:     password,
	}

	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	resp, err := freeIpaClient.V1freeipauser.SetPasswordV1(v1freeipauser.NewSetPasswordV1Params().WithBody(&SetPasswordV1Request))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	syncOperationStatus := resp.Payload

	if c.Bool(fl.FlWaitOptional.Name) && syncOperationStatus.Status == "RUNNING" {
		getSyncOperationStatus(c, *syncOperationStatus.OperationID, "setPassword")
	} else {
		writeSyncOperationStatus(c, syncOperationStatus)
	}
}

func constructEnvCrnList(c *cli.Context) []string {
	environmentCrns := c.StringSlice(fl.FlIpaEnvironmentCrnsSlice.Name)
	environmentNames := c.StringSlice(fl.FlIpaEnvironmentNamesOptionalSlice.Name)

	envCrnSet := make(map[string]bool, 0)
	for _, envName := range environmentNames {
		envCrn := env.GetEnvCrnByName(envName, c)
		envCrnSet[envCrn] = true
	}

	for _, envCrn := range environmentCrns {
		envCrnSet[envCrn] = true
	}

	crns := make([]string, 0, len(envCrnSet))
	for crn := range envCrnSet {
		crns = append(crns, crn)
	}

	return crns
}

func GetSyncOperationStatus(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "get status of a sync operation")

	operationId := c.String(fl.FlIpaSyncOperationId.Name)
	getSyncOperationStatus(c, operationId, "getSyncOperationStatus")
}

func GetVmLogs(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "get default monitored vm logs for freeipa clusters")
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	resp, err := freeIpaClient.V1diagnostics.GetFreeIpaVMLogsV1(v1diagnostics.NewGetFreeIpaVMLogsV1Params())
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	logs := resp.Payload.Logs
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(logs); err != nil {
		utils.LogErrorAndExit(err)
	}
	fmt.Println(buf.String())
}

func CollectDiagnostics(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "get user synchronization state for an environment")
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	collectionRequest := assembleCollectionRequest(c)
	_, err := freeIpaClient.V1diagnostics.CollectFreeIpaDiagnosticsV1(v1diagnostics.NewCollectFreeIpaDiagnosticsV1Params().WithBody(collectionRequest))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
}

func ListDiagnosticsCollections(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "get latest freeipa diagnostics collection flows")
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)
	resp, err := freeIpaClient.V1diagnostics.ListDiagnosticsCollectionsV1(v1diagnostics.NewListDiagnosticsCollectionsV1Params().WithEnvironmentCrn(envCrn))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	collections := resp.Payload.Collections
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(collections); err != nil {
		utils.LogErrorAndExit(err)
	}
	fmt.Println(buf.String())
}

func CancelDiagnosticsCollections(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "cancel running freeipa diagnostics collection flows")
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)
	err := freeIpaClient.V1diagnostics.CancelDiagnosticsCollectionsV1(v1diagnostics.NewCancelDiagnosticsCollectionsV1Params().WithEnvironmentCrn(envCrn))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	fmt.Println("Cancel running freeipa diagnostics collections...")
}

func assembleCollectionRequest(c *cli.Context) *freeIpaModel.DiagnosticsCollectionV1Request {
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)
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
		content := utils.ReadFile(additionalLogsFileOption)
		err := json.Unmarshal(content, &additionalLogs)
		if err != nil {
			utils.LogErrorAndExit(err)
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
	request := freeIpaModel.DiagnosticsCollectionV1Request{Destination: &destination, EnvironmentCrn: &envCrn, Labels: labels,
		Issue: issue, Description: description, StartTime: startTime, EndTime: endTime, AdditionalLogs: additionalLogs,
		IncludeSaltLogs: includeSaltLogs, UpdatePackage: updatePackage, Hosts: hosts, HostGroups: hostGroups}
	return &request
}

func getSyncOperationStatus(c *cli.Context, operationId string, command string) {
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	resp, err := freeIpaClient.V1freeipauser.GetSyncOperationStatusV1(v1freeipauser.NewGetSyncOperationStatusV1Params().WithOperationID(operationId))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	syncOperationStatus := resp.Payload
	for c.Bool(fl.FlWaitOptional.Name) && syncOperationStatus.Status == "RUNNING" {
		log.Infof("[%s] Status is RUNNING. Sleeping", command)
		time.Sleep(5 * time.Second)
		resp, err := freeIpaClient.V1freeipauser.GetSyncOperationStatusV1(v1freeipauser.NewGetSyncOperationStatusV1Params().WithOperationID(operationId))
		if err != nil {
			commonutils.LogErrorAndExit(err)
		}
		syncOperationStatus = resp.Payload
	}
	log.Infof("[%s] Operation completed: %s", command, syncOperationStatus.Status)
	writeSyncOperationStatus(c, syncOperationStatus)
}

func writeSyncOperationStatus(c *cli.Context, syncOperationStatus *freeIpaModel.SyncOperationV1Status) {
	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	syncStatus := &freeIpaOutSyncOperation{
		ID:        *syncOperationStatus.OperationID,
		Status:    syncOperationStatus.Status,
		SyncType:  *syncOperationStatus.SyncOperationType,
		Success:   convertSuccessDetailsModel(syncOperationStatus.Success),
		Failure:   convertFailureDetailsModel(syncOperationStatus.Failure),
		Error:     syncOperationStatus.Error,
		StartTime: strconv.FormatInt(syncOperationStatus.StartTime, 10),
		EndTime:   strconv.FormatInt(syncOperationStatus.EndTime, 10),
	}
	output.Write(syncStatusHeader, syncStatus)
}

func convertSuccessDetailsModel(sd []*freeIpaModel.SuccessDetailsV1) []successDetail {
	var successDetails = []successDetail{}
	for _, success := range sd {
		successDetails = append(successDetails, successDetail{
			Environment: success.Environment,
		})
	}
	return successDetails
}

func convertFailureDetailsModel(fd []*freeIpaModel.FailureDetailsV1) []failureDetail {
	var failureDetails = []failureDetail{}
	for _, failure := range fd {
		failureDetails = append(failureDetails, failureDetail{
			Environment: failure.Environment,
			Details:     failure.Message,
		})
	}
	return failureDetails
}

func GetEnvironmentSyncState(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "get user synchronization state for an environment")
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	resp, err := freeIpaClient.V1freeipauser.GetEnvironmentUserSyncStateV1(v1freeipauser.NewGetEnvironmentUserSyncStateV1Params().WithEnvironmentCrn(&envCrn))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[getEnvironmentSyncState] FreeIpa cluster start requested in environment %s", envName)
	writeEnvironmentUserSyncState(c, resp.Payload)
}

func writeEnvironmentUserSyncState(c *cli.Context, environmentUserSyncState *freeIpaModel.EnvironmentUserSyncV1State) {
	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	syncState := &freeIpaOutEnvironmentUserSyncState{
		State:                   *environmentUserSyncState.State,
		LastUserSyncOperationID: environmentUserSyncState.LastUserSyncOperationID,
	}
	output.Write(environmentUserSyncStateHeader, syncState)
}

func AddDnsARecord(c *cli.Context) {
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)
	hostName := c.String(fl.FlDnsHostname.Name)
	zone := c.String(fl.FlDnsZone.Name)
	ip := c.String(fl.FlDnsIp.Name)
	addDNSARecordV1Request := freeIpaModel.AddDNSARecordV1Request{
		EnvironmentCrn: &envCrn,
		DNSZone:        zone,
		Hostname:       &hostName,
		IP:             &ip,
		CreateReverse:  c.Bool(fl.FlDnsCreateReverse.Name),
	}
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	err := freeIpaClient.V1dns.AddDNSARecordV1(v1dns.NewAddDNSARecordV1Params().WithBody(&addDNSARecordV1Request))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
}

func DeleteDnsARecord(c *cli.Context) {
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)
	hostName := c.String(fl.FlDnsHostname.Name)
	zone := c.String(fl.FlDnsZone.Name)
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	err := freeIpaClient.V1dns.DeleteDNSARecordV1(v1dns.NewDeleteDNSARecordV1Params().WithHostname(&hostName).WithDNSZone(&zone).WithEnvironment(&envCrn))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
}

func AddDnsCnameRecord(c *cli.Context) {
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)
	cname := c.String(fl.FlDnsCname.Name)
	zone := c.String(fl.FlDnsZone.Name)
	target := c.String(fl.FlDnsTargetFqdn.Name)
	addDNSCnameRecordV1Request := freeIpaModel.AddDNSCnameRecordV1Request{
		EnvironmentCrn: &envCrn,
		DNSZone:        zone,
		Cname:          &cname,
		TargetFqdn:     &target,
	}
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	err := freeIpaClient.V1dns.AddDNSCnameRecordV1(v1dns.NewAddDNSCnameRecordV1Params().WithBody(&addDNSCnameRecordV1Request))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
}

func DeleteDnsCnameRecord(c *cli.Context) {
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)
	cname := c.String(fl.FlDnsCname.Name)
	zone := c.String(fl.FlDnsZone.Name)
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	err := freeIpaClient.V1dns.DeleteDNSCnameRecordV1(v1dns.NewDeleteDNSCnameRecordV1Params().WithCname(&cname).WithDNSZone(&zone).WithEnvironment(&envCrn))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
}

func ChangeImage(c *cli.Context) {
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)
	imageId := c.String(fl.FlImageId.Name)
	imageCatalogName := c.String(fl.FlImageCatalog.Name)
	osType := ""

	imageSettingsV1Request := model.ImageSettingsV1Request{Catalog: imageCatalogName, ID: imageId, Os: osType}
	request := model.ImageChangeV1Request{EnvironmentCrn: &envCrn, ImageSettings: &imageSettingsV1Request}

	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	_, err := freeIpaClient.V1freeipa.ChangeImageV1(v1freeipa.NewChangeImageV1Params().WithBody(&request), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
}

func UpdateSalt(c *cli.Context) {
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)

	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	_, err := freeIpaClient.V1freeipa.UpdateSaltV1(v1freeipa.NewUpdateSaltV1Params().WithEnvironment(&envCrn), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
}

func OperationProgress(c *cli.Context) {
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	result, err := freeIpaClient.V1operation.GetOperationProgressByEnvironmentCrn(v1operation.NewGetOperationProgressByEnvironmentCrnParams().WithEnvironmentCrn(envCrn))
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

func Upgrade(c *cli.Context) {
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)

	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	imageId := c.String(fl.FlImageIdOptional.Name)
	imageCatalog := c.String(fl.FlImageCatalogOptional.Name)
	imageSetting := model.ImageSettingsV1Request{Catalog: imageCatalog, ID: imageId}
	request := model.FreeIpaUpgradeV1Request{EnvironmentCrn: &envCrn, Image: &imageSetting}
	resp, err := freeIpaClient.V1freeipa.UpgradeFreeIpaV1(v1freeipa.NewUpgradeFreeIpaV1Params().WithBody(&request))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	printFormattedJsonResponse(resp.Payload)
}

func UpgradeOptions(c *cli.Context) {
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)

	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	imageCatalog := c.String(fl.FlImageCatalogOptional.Name)
	resp, err := freeIpaClient.V1freeipa.GetFreeIpaUpgradeOptionsV1(v1freeipa.NewGetFreeIpaUpgradeOptionsV1Params().WithEnvironmentCrn(&envCrn).WithCatalog(&imageCatalog))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	printFormattedJsonResponse(resp.Payload)
}

func Retry(c *cli.Context) {
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)

	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	resp, err := freeIpaClient.V1freeipa.RetryV1(v1freeipa.NewRetryV1Params().WithEnvironment(&envCrn), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	printFormattedJsonResponse(resp.Payload)
}

func ListRetryableFlows(c *cli.Context) {
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)

	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	resp, err := freeIpaClient.V1freeipa.ListRetryableFlowsV1(v1freeipa.NewListRetryableFlowsV1Params().WithEnvironment(&envCrn), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	printFormattedJsonResponse(resp.Payload)
}

func printFormattedJsonResponse(v interface{}) {
	response, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	commonutils.Println(string(response))
}

func Rebuild(c *cli.Context) {
	envName := c.String(fl.FlEnvironmentName.Name)
	envCrn := env.GetEnvirontmentCrnByName(c, envName)
	sourceFreeIpaCrn := c.String(fl.FlCrn.Name)

	request := model.RebuildV1Request{EnvironmentCrn: &envCrn, SourceCrn: &sourceFreeIpaCrn}
	freeIpaClient := ClientFreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	_, err := freeIpaClient.V1freeipa.RebuildV1(v1freeipa.NewRebuildV1Params().WithBody(&request), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
}
