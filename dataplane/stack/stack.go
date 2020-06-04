package stack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hortonworks/cb-cli/dataplane/env"

	"github.com/hortonworks/cb-cli/dataplane/utils"
	commonutils "github.com/hortonworks/dp-cli-common/utils"

	"github.com/hortonworks/cb-cli/dataplane/common"
	"github.com/hortonworks/cb-cli/dataplane/oauth"

	v4stack "github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_stacks"
	"github.com/hortonworks/cb-cli/dataplane/api/model"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/types"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var stackHeader = []string{"Name", "CloudPlatform", "EnvironmentCrn", "StackStatus", "ClusterStatus"}
var instanceHeader = []string{"FQDN", "InstanceGroup", "PublicIP", "PrivateIP", "HostStatus", "StatusReason"}

type stackOut struct {
	common.CloudResourceOut
	EnvironmentCrn string `json:"EnvironmentCrn" yaml:"EnvironmentCrn"`
	StackStatus    string `json:"StackStatus" yaml:"StackStatus"`
	ClusterStatus  string `json:"ClusterStatus" yaml:"ClusterStatus"`
}

type stackOutDescribe struct {
	*model.StackV4Response
}

type stackOutRequest struct {
	*model.StackV4Request
}

type instanceOutDescribe struct {
	FQDN          string `json:"FQDN" yaml:"FQDN"`
	InstanceGroup string `json:"InstanceGroup" yaml:"InstanceGroup"`
	PublicIP      string `json:"PublicIP" yaml:"PublicIP"`
	PrivateIP     string `json:"PrivateIP" yaml:"PrivateIP"`
	HostStatus    string `json:"HostStatus" yaml:"HostStatus"`
	StatusReason  string `json:"StatusReason" yaml:"StatusReason"`
}

func (s *instanceOutDescribe) DataAsStringArray() []string {
	return []string{s.FQDN, s.InstanceGroup, s.PublicIP, s.PrivateIP, s.HostStatus, s.StatusReason}
}

func (s *stackOut) DataAsStringArray() []string {
	arr := []string{s.Name, s.CloudPlatform}
	arr = append(arr, s.EnvironmentCrn)
	arr = append(arr, s.StackStatus)
	arr = append(arr, s.ClusterStatus)
	return arr
}

func (s *stackOutDescribe) DataAsStringArray() []string {
	stack := convertResponseToStack(s.StackV4Response)
	return append(stack.DataAsStringArray(), s.Crn)
}

func (s *stackOutRequest) DataAsStringArray() []string {
	return []string{}
}

func CreateStack(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "create cluster")

	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	req := assembleStackRequest(c)
	cbClient := CloudbreakStack(*oauth.NewCloudbreakHTTPClientFromContext(c))
	utils.CheckClientVersion(cbClient.Cloudbreak.V4utils, common.Version)
	cbClient.createStack(workspaceID, req)
	if c.Bool(fl.FlWaitOptional.Name) {
		cbClient.waitForOperationToFinish(workspaceID, *req.Name, AVAILABLE, AVAILABLE)
	}
}

func ChangeImage(c *cli.Context) {
	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)

	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	imageId := c.String(fl.FlImageId.Name)
	imageCatalogName := c.String(fl.FlImageCatalog.Name)
	clusterName := c.String(fl.FlName.Name)
	log.Infof("[ChangeImage] changing image for stack stack, name: %s, imageid: %s, imageCatalog: %s", clusterName, imageId, imageCatalogName)
	req := model.StackImageChangeV4Request{ImageCatalogName: imageCatalogName, ImageID: &imageId}
	err, _ := cbClient.Cloudbreak.V4WorkspaceIDStacks.ChangeImageStackInWorkspaceV4(v4stack.NewChangeImageStackInWorkspaceV4Params().WithWorkspaceID(workspaceID).WithName(clusterName).WithBody(&req))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
}

func assembleStackRequest(c *cli.Context) *model.StackV4Request {
	path := c.String(fl.FlInputJson.Name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		commonutils.LogErrorAndExit(err)
	}

	log.Infof("[assembleStackTemplate] read cluster create json from file: %s", path)
	content := commonutils.ReadFile(path)

	var req model.StackV4Request
	err := json.Unmarshal(content, &req)
	if err != nil {
		msg := fmt.Sprintf(`Invalid json format: %s. Please make sure that the json is valid (check for commas and double quotes).`, err.Error())
		commonutils.LogErrorMessageAndExit(msg)
	}

	name := c.String(fl.FlName.Name)
	if len(name) != 0 {
		req.Name = &name
		if req.InstanceGroups != nil {
			for _, ig := range req.InstanceGroups {
				if ig.Azure != nil && ig.Azure.AvailabilitySet != nil {
					as := ig.Azure.AvailabilitySet
					as.Name = *req.Name + "-" + *ig.Name + "-as"
				}
			}
		}
	}
	if req.Name == nil || len(*req.Name) == 0 {
		commonutils.LogErrorMessageAndExit("Name of the cluster must be set either in the template or with the --name command line option.")
	}

	envName := c.String(fl.FlEnvironmentNameOptional.Name)
	if (req.EnvironmentCrn == nil || len(*req.EnvironmentCrn) == 0) && len(envName) > 0 {
		envCrn := env.GetEnvCrnByName(envName, c)
		log.Debugf("[assembleStackTemplate] env crn is empty in stack request, update with: %s", envCrn)
		req.EnvironmentCrn = &envCrn
	}

	ambariUser := c.String(fl.FlCMUserOptional.Name)
	ambariPassword := c.String(fl.FlCMPasswordOptional.Name)
	if len(ambariUser) != 0 || len(ambariPassword) != 0 {
		if req.Cluster != nil {
			if len(ambariUser) != 0 {
				req.Cluster.UserName = ambariUser
			}
			if len(ambariPassword) != 0 {
				req.Cluster.Password = ambariPassword
			}
		} else {
			commonutils.LogErrorMessageAndExit("Missing cluster node in JSON")
		}
	}
	return &req
}

func convertViewResponseToStack(s *model.StackViewV4Response) *stackOut {
	return &stackOut{
		CloudResourceOut: common.CloudResourceOut{
			Name:          *s.Name,
			CloudPlatform: s.CloudPlatform,
		},
		EnvironmentCrn: s.EnvironmentCrn,
		StackStatus:    s.Status,
		ClusterStatus:  utils.SafeClusterViewStatusConvert(s),
	}
}

func convertResponseToStack(s *model.StackV4Response) *stackOut {
	return &stackOut{
		CloudResourceOut: common.CloudResourceOut{
			Name:          *s.Name,
			Description:   utils.SafeClusterDescriptionConvert(s),
			CloudPlatform: s.CloudPlatform,
		},
		EnvironmentCrn: s.EnvironmentCrn,
		StackStatus:    s.Status,
		ClusterStatus:  utils.SafeClusterStatusConvert(s),
	}
}

func GetStackRequest(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "get request for stack")

	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	output := commonutils.Output{Format: "json"}
	resp, err := cbClient.Cloudbreak.V4WorkspaceIDStacks.GetStackRequestFromNameV4(v4stack.NewGetStackRequestFromNameV4Params().WithWorkspaceID(workspaceID).WithName(c.String(fl.FlName.Name)))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	s := resp.Payload
	output.Write([]string{}, &stackOutRequest{s})
}

func DescribeStack(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "describe stack")

	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	name := c.String(fl.FlName.Name)
	cbClient := CloudbreakStack(*oauth.NewCloudbreakHTTPClientFromContext(c))
	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	resp := cbClient.getStackByName(workspaceID, name)
	output.Write(append(stackHeader, "ID"), &stackOutDescribe{resp})
}

func DescribeStackInstances(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "describe instances")

	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	cbClient := CloudbreakStack(*oauth.NewCloudbreakHTTPClientFromContext(c))
	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	name := c.String(fl.FlName.Name)

	stack := cbClient.getStackByName(workspaceID, name)
	var tableRows []commonutils.Row
	for _, ig := range stack.InstanceGroups {
		for _, meta := range ig.Metadata {
			tableRows = append(tableRows, &instanceOutDescribe{
				FQDN:          meta.DiscoveryFQDN,
				InstanceGroup: meta.InstanceGroup,
				PublicIP:      meta.PublicIP,
				PrivateIP:     meta.PrivateIP,
				HostStatus:    meta.State,
				StatusReason:  meta.StatusReason,
			})
		}
	}
	output.WriteList(instanceHeader, tableRows)
}

type getStackInWorkspace interface {
	GetStackInWorkspaceV4(client *v4stack.GetStackInWorkspaceV4Params) (*v4stack.GetStackInWorkspaceV4OK, error)
}

func fetchStack(workspaceID int64, name string, client getStackInWorkspace) *model.StackV4Response {
	resp, err := client.GetStackInWorkspaceV4(v4stack.NewGetStackInWorkspaceV4Params().WithWorkspaceID(workspaceID).WithName(name))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	return resp.Payload
}

func ScaleStack(c *cli.Context) {
	desiredCount, err := strconv.Atoi(c.String(fl.FlDesiredNodeCount.Name))
	if err != nil {
		commonutils.LogErrorMessageAndExit("Unable to parse as number: " + c.String(fl.FlDesiredNodeCount.Name))
	}
	defer commonutils.TimeTrack(time.Now(), "scale stack")

	cbClient := CloudbreakStack(*oauth.NewCloudbreakHTTPClientFromContext(c))
	req := &model.StackScaleV4Request{
		DesiredCount: &(&types.I32{I: int32(desiredCount)}).I,
		Group:        &(&types.S{S: c.String(fl.FlGroupName.Name)}).S,
	}
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	name := c.String(fl.FlName.Name)
	log.Infof("[ScaleStack] scaling stack, name: %s", name)
	_, err = cbClient.Cloudbreak.V4WorkspaceIDStacks.PutScalingStackInWorkspaceV4(v4stack.NewPutScalingStackInWorkspaceV4Params().WithWorkspaceID(workspaceID).WithName(name).WithBody(req))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[ScaleStack] stack scaled, name: %s", name)

	if c.Bool(fl.FlWaitOptional.Name) {
		cbClient.waitForOperationToFinish(workspaceID, name, AVAILABLE, AVAILABLE)
	}
}

func StartStack(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "start stack")

	cbClient := CloudbreakStack(*oauth.NewCloudbreakHTTPClientFromContext(c))
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	name := c.String(fl.FlName.Name)
	log.Infof("[StartStack] starting stack, name: %s", name)
	_, err := cbClient.Cloudbreak.V4WorkspaceIDStacks.StartStackInWorkspaceV4(v4stack.NewStartStackInWorkspaceV4Params().WithWorkspaceID(workspaceID).WithName(name))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[StartStack] stack started, name: %s", name)

	if c.Bool(fl.FlWaitOptional.Name) {
		cbClient.waitForOperationToFinish(workspaceID, name, AVAILABLE, AVAILABLE)
	}
}

func StopStack(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "stop stack")

	cbClient := CloudbreakStack(*oauth.NewCloudbreakHTTPClientFromContext(c))
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	name := c.String(fl.FlName.Name)
	log.Infof("[StopStack] stopping stack, name: %s", name)
	_, err := cbClient.Cloudbreak.V4WorkspaceIDStacks.StopStackInWorkspaceV4(v4stack.NewStopStackInWorkspaceV4Params().WithWorkspaceID(workspaceID).WithName(name))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[StopStack] stack stopted, name: %s", name)

	if c.Bool(fl.FlWaitOptional.Name) {
		cbClient.waitForOperationToFinish(workspaceID, name, STOPPED, STOPPED)
	}
}

func SyncStack(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "sync stack")

	cbClient := CloudbreakStack(*oauth.NewCloudbreakHTTPClientFromContext(c))
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	name := c.String(fl.FlName.Name)
	log.Infof("[SyncStack] syncing stack, name: %s", name)
	_, err := cbClient.Cloudbreak.V4WorkspaceIDStacks.SyncStackInWorkspaceV4(v4stack.NewSyncStackInWorkspaceV4Params().WithWorkspaceID(workspaceID).WithName(name))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[SyncStack] stack synced, name: %s", name)
}

func UpdateSalt(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "update salt")

	cbClient := CloudbreakStack(*oauth.NewCloudbreakHTTPClientFromContext(c))
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	name := c.String(fl.FlName.Name)
	log.Infof("[UpdateSalt] update salt, name: %s", name)
	_, err := cbClient.Cloudbreak.V4WorkspaceIDStacks.UpdateSaltByName(v4stack.NewUpdateSaltByNameParams().WithWorkspaceID(workspaceID).WithName(name))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[UpdateSalt] salt updated, name: %s", name)
}

func RepairStackHostGroups(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "repair stack")

	var request model.ClusterRepairV4Request
	hostGroups := strings.Split(c.String(fl.FlHostGroups.Name), ",")
	request.HostGroups = hostGroups
	log.Infof("[RepairStack] repairing stack hostgroups: %s", hostGroups)
	repairStackCommon(c, request)
}

func RepairStackNodes(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "repair stack")

	var request model.ClusterRepairV4Request
	deleteVolumes := c.Bool(fl.FlDeleteVolumes.Name)
	nodes := strings.Split(c.String(fl.FlNodes.Name), ",")
	request.Nodes = &model.ClusterRepairNodesV4Request{DeleteVolumes: deleteVolumes, Ids: nodes}

	log.Infof("[RepairStackNodes] repairing stack nodes, deleteVolumes: %t, ids: %s", deleteVolumes, nodes)

	repairStackCommon(c, request)
}

func repairStackCommon(c *cli.Context, request model.ClusterRepairV4Request) {
	cbClient := CloudbreakStack(*oauth.NewCloudbreakHTTPClientFromContext(c))
	removeOnly := c.Bool(fl.FlRemoveOnly.Name)
	request.RemoveOnly = removeOnly
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	name := c.String(fl.FlName.Name)
	log.Infof("[RepairStack] repairing stack, id: %s, removeOnly: %t, workspaceId: %d", name, removeOnly, workspaceID)

	_, err := cbClient.Cloudbreak.V4WorkspaceIDStacks.RepairStackInWorkspaceV4(v4stack.NewRepairStackInWorkspaceV4Params().WithWorkspaceID(workspaceID).WithName(name).WithBody(&request))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[RepairStack] stack repaired, name: %s", name)

	if c.Bool(fl.FlWaitOptional.Name) {
		cbClient.waitForOperationToFinish(workspaceID, name, AVAILABLE, AVAILABLE)
	}
}

func RetryCluster(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "retry cluster creation")

	cbClient := CloudbreakStack(*oauth.NewCloudbreakHTTPClientFromContext(c))
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	name := c.String(fl.FlName.Name)
	log.Infof("[RetryCluster] retrying cluster creation, name: %s", name)
	_, err := cbClient.Cloudbreak.V4WorkspaceIDStacks.RetryStackInWorkspaceV4(v4stack.NewRetryStackInWorkspaceV4Params().WithWorkspaceID(workspaceID).WithName(name))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[RetryCluster] cluster creation retried, name: %s", name)

	if c.Bool(fl.FlWaitOptional.Name) {
		cbClient.waitForOperationToFinish(workspaceID, name, AVAILABLE, AVAILABLE)
	}
}

func DeleteStack(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "delete stack")

	cbClient := CloudbreakStack(*oauth.NewCloudbreakHTTPClientFromContext(c))
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	name := c.String(fl.FlName.Name)
	cbClient.deleteStack(workspaceID, name, c.Bool(fl.FlForceOptional.Name))
	log.Infof("[DeleteStack] stack deleted, name: %s", name)

	if c.Bool(fl.FlWaitOptional.Name) {
		cbClient.waitForOperationToFinish(workspaceID, name, DELETE_COMPLETED, SKIP)
	}
}

func ListStacks(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "list stacks")

	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	cbClient := CloudbreakStack(*oauth.NewCloudbreakHTTPClientFromContext(c))
	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	listStacksImpl(cbClient.Cloudbreak.V4WorkspaceIDStacks, output.WriteList, workspaceID)
}

type listStacksByWorkspaceClient interface {
	ListStackInWorkspaceV4(params *v4stack.ListStackInWorkspaceV4Params) (*v4stack.ListStackInWorkspaceV4OK, error)
}

func listStacksImpl(client listStacksByWorkspaceClient, writer func([]string, []commonutils.Row), workspaceID int64) {
	log.Infof("[listStacksImpl] sending stack list request")
	stacksResp, err := client.ListStackInWorkspaceV4(v4stack.NewListStackInWorkspaceV4Params().WithWorkspaceID(workspaceID))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}

	var tableRows []commonutils.Row
	for _, stack := range stacksResp.Payload.Responses {
		tableRows = append(tableRows, convertViewResponseToStack(stack))
	}

	writer(stackHeader, tableRows)
}

func GenerateInventoryFile(c *cli.Context) {
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	clusterName := c.String(fl.FlName.Name)
	outputFileName := c.String(fl.FlOutputFile.Name)
	cbClient := CloudbreakStack(*oauth.NewCloudbreakHTTPClientFromContext(c))
	response, err := cbClient.Cloudbreak.V4WorkspaceIDStacks.GetClusterHostsInventory(v4stack.NewGetClusterHostsInventoryParams().WithWorkspaceID(workspaceID).WithName(clusterName))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	ioutil.WriteFile(outputFileName, []byte(response.Payload), 0644)
}
