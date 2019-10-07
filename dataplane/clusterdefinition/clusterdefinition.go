package clusterdefinition

import (
	"time"

	"github.com/hortonworks/cb-cli/dataplane/oauth"

	"encoding/json"
	"fmt"

	v4cd "github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_clustertemplates"
	"github.com/hortonworks/cb-cli/dataplane/api/model"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/dp-cli-common/utils"
	commonutils "github.com/hortonworks/dp-cli-common/utils"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var clusterDefinitionHeader = []string{"Name", "Description", "Type", "Cloud platform", "Platform Version", "Compatible environments"}

type clusterDefinitionOut struct {
	Name                   string `json:"Name" yaml:"Name"`
	Description            string `json:"Description" yaml:"Description"`
	Type                   string `json:"Type" yaml:"Type"`
	CloudPlatform          string `json:"CloudPlatform" yaml:"CloudPlatform"`
	PlatformVersion        string `json:"PlatformVersion" yaml:"PlatformVersion"`
	CompatibleEnvironments string `json:"CompatibleEnvironments" yaml:"CompatibleEnvironments"`
}

type clusterDefinitionOutJSONDescribe struct {
	*clusterDefinitionOut
	DatalakeRequired string `json:"DatalakeRequired" yaml:"DatalakeRequired"`
	ResourceStatus   string `json:"ResourceStatus" yaml:"ResourceStatus"`
}

type clusterDefinitionOutTableDescribe struct {
	*clusterDefinitionOut
	DatalakeRequired string `json:"DatalakeRequired" yaml:"DatalakeRequired"`
	ResourceStatus   string `json:"ResourceStatus" yaml:"ResourceStatus"`
}

func (cd *clusterDefinitionOut) DataAsStringArray() []string {
	return []string{cd.Name, cd.Description, cd.Type, cd.CloudPlatform, cd.PlatformVersion, cd.CompatibleEnvironments}
}

func (cd *clusterDefinitionOutTableDescribe) DataAsStringArray() []string {
	return append(cd.clusterDefinitionOut.DataAsStringArray(), cd.DatalakeRequired, cd.ResourceStatus)
}

func (cd *clusterDefinitionOutJSONDescribe) DataAsStringArray() []string {
	return append(cd.clusterDefinitionOut.DataAsStringArray(), cd.DatalakeRequired, cd.ResourceStatus)
}

func CreateClusterDefinitionFromFile(c *cli.Context) {
	log.Infof("[CreateClusterDefinitionFromFile] creating cluster definition from a file")

	fileLocation := c.String(fl.FlFile.Name)
	content := utils.ReadFile(fileLocation)
	var req model.ClusterTemplateV4Request

	err := json.Unmarshal(content, &req)
	if err != nil {
		msg := fmt.Sprintf(`Invalid JSON format: %s. Please make sure that the json is valid (check for commas and double quotes).`, err.Error())
		utils.LogErrorMessageAndExit(msg)
	}
	name := c.String(fl.FlName.Name)
	req.Name = &name
	createClusterDefinitionImpl(c, &req)
}

func createClusterDefinitionImpl(c *cli.Context, template *model.ClusterTemplateV4Request) {
	defer utils.TimeTrack(time.Now(), "create cluster template")

	client := oauth.NewCloudbreakHTTPClientFromContext(c)
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	log.Infof("[createClusterDefinitionImpl] sending create cluster definition request")
	resp, err := client.Cloudbreak.V4WorkspaceIDClustertemplates.CreateClusterTemplateInWorkspace(v4cd.NewCreateClusterTemplateInWorkspaceParams().WithWorkspaceID(workspaceID).WithBody(template))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	result := resp.Payload

	log.Infof("[createClusterDefinitionImpl] a cluster template - with the type of \"%s\" and with name \"%s\" - has created for the given environment: %s", *result.Type, *result.Name, result.EnvironmentName)
}

func DescribeClusterDefinition(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "describe cluster definition")

	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	name := c.String(fl.FlName.Name)
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)

	resp, err := cbClient.Cloudbreak.V4WorkspaceIDClustertemplates.GetClusterTemplateByNameInWorkspace(v4cd.NewGetClusterTemplateByNameInWorkspaceParams().WithName(name).WithWorkspaceID(workspaceID))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	output.Write(append(clusterDefinitionHeader, "DATALAKE_REQUIRED", "STATUS"), convertResponseWithAdditionalData(resp.Payload))
}

func DeleteClusterDefinitions(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "delete cluster definitions")

	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	deleteClusterDefinitionsImpl(cbClient.Cloudbreak.V4WorkspaceIDClustertemplates, c.Int64(fl.FlWorkspaceOptional.Name), c.StringSlice(fl.FlNames.Name))
}

func deleteClusterDefinitionsImpl(client clusterDefinitionClient, workspace int64, names []string) {
	log.Infof("[deleteClusterDefinitionsImpl] sending delete cluster definitions request with names: %s", names)

	_, err := client.DeleteClusterTemplatesInWorkspace(v4cd.NewDeleteClusterTemplatesInWorkspaceParams().WithWorkspaceID(workspace).WithBody(names))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[deleteClusterDefinitionsImpl] cluster definitions deleted, names: %s", names)
}

func ListClusterDefinitions(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "get cluster definitions")

	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	workspace := fl.FlWorkspaceOptional.Name
	listClusterDefinitionsImpl(c, c.Int64(workspace), cbClient.Cloudbreak.V4WorkspaceIDClustertemplates)
}

func listClusterDefinitionsImpl(c *cli.Context, workspace int64, client clusterDefinitionClient) {
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	log.Infof("[listclusterDefinitionsImpl] sending cluster template list request")
	resp, err := client.ListClusterTemplatesByWorkspace(v4cd.NewListClusterTemplatesByWorkspaceParams().WithWorkspaceID(workspace))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	var tableRows []commonutils.Row
	for _, cd := range resp.Payload.Responses {
		tableRows = append(tableRows, convertResponseToClusterDefinition(cd))
	}

	output.WriteList(clusterDefinitionHeader, tableRows)
}

type clusterDefinitionClient interface {
	CreateClusterTemplateInWorkspace(params *v4cd.CreateClusterTemplateInWorkspaceParams) (*v4cd.CreateClusterTemplateInWorkspaceOK, error)
	ListClusterTemplatesByWorkspace(params *v4cd.ListClusterTemplatesByWorkspaceParams) (*v4cd.ListClusterTemplatesByWorkspaceOK, error)
	DeleteClusterTemplatesInWorkspace(params *v4cd.DeleteClusterTemplatesInWorkspaceParams) (*v4cd.DeleteClusterTemplatesInWorkspaceOK, error)
}

func convertResponseToClusterDefinition(ct *model.ClusterTemplateViewV4Response) *clusterDefinitionOut {
	return &clusterDefinitionOut{
		Name:                   *ct.Name,
		Description:            utils.SafeStringConvert(ct.Description),
		Type:                   ct.Type,
		CloudPlatform:          ct.CloudPlatform,
		PlatformVersion:        ct.StackVersion,
		CompatibleEnvironments: getCompatibleEnvironments(ct.EnvironmentName),
	}
}

func convertResponseWithAdditionalData(cd *model.ClusterTemplateV4Response) *clusterDefinitionOutJSONDescribe {
	return &clusterDefinitionOutJSONDescribe{
		clusterDefinitionOut: &clusterDefinitionOut{
			Name:                   *cd.Name,
			Description:            *cd.Description,
			Type:                   *cd.Type,
			CloudPlatform:          cd.CloudPlatform,
			CompatibleEnvironments: getCompatibleEnvironments(cd.EnvironmentName),
		},
		DatalakeRequired: cd.DatalakeRequired,
		ResourceStatus:   cd.Status,
	}
}

func getCompatibleEnvironments(envName string) string {
	var finalEnvName string
	if len(envName) > 0 {
		finalEnvName = envName
	} else {
		finalEnvName = "ALL"
	}
	return finalEnvName
}
