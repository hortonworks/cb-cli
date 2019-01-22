package rds

import (
	"fmt"
	"github.com/hortonworks/cb-cli/cloudbreak/oauth"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/cb-cli/cloudbreak/api/client/v3_workspace_id_rdsconfigs"
	"github.com/hortonworks/cb-cli/cloudbreak/api/model"
	fl "github.com/hortonworks/cb-cli/cloudbreak/flags"
	"github.com/hortonworks/cb-cli/cloudbreak/types"
	"github.com/hortonworks/cb-cli/utils"
	"github.com/urfave/cli"
)

var rdsHeader = []string{"Name", "ConnectionURL", "DatabaseEngine", "Type", "Driver"}

type rds struct {
	Name           string `json:"Name" yaml:"Name"`
	URL            string `json:"ConnectionURL" yaml:"ConnectionURL"`
	DatabaseEngine string `json:"DatabaseEngine" yaml:"DatabaseEngine"`
	Type           string `json:"Type" yaml:"Type"`
	Driver         string `json:"Driver" yaml:"Driver"`
}

type rdsOutDescribe struct {
	*rds
	ID string `json:"ID" yaml:"ID"`
}

func (r *rds) DataAsStringArray() []string {
	return []string{r.Name, r.URL, r.DatabaseEngine, r.Type, r.Driver}
}

func (r *rdsOutDescribe) DataAsStringArray() []string {
	return append(r.rds.DataAsStringArray(), r.ID)
}

type rdsClient interface {
	ListRdsConfigsByWorkspace(params *v3_workspace_id_rdsconfigs.ListRdsConfigsByWorkspaceParams) (*v3_workspace_id_rdsconfigs.ListRdsConfigsByWorkspaceOK, error)
	CreateRdsConfigInWorkspace(params *v3_workspace_id_rdsconfigs.CreateRdsConfigInWorkspaceParams) (*v3_workspace_id_rdsconfigs.CreateRdsConfigInWorkspaceOK, error)
	TestRdsConnectionInWorkspace(params *v3_workspace_id_rdsconfigs.TestRdsConnectionInWorkspaceParams) (*v3_workspace_id_rdsconfigs.TestRdsConnectionInWorkspaceOK, error)
}

func TestRdsByName(c *cli.Context) {
	log.Infof("[TestRdsByParams] test a database configuration")
	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	testRdsByNameImpl(
		cbClient.Cloudbreak.V3WorkspaceIDRdsconfigs,
		c.Int64(fl.FlWorkspaceOptional.Name),
		c.String(fl.FlName.Name))
}

func TestRdsByParams(c *cli.Context) {
	log.Infof("[TestRdsByParams] test a database configuration")
	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	testRdsByParamsImpl(
		cbClient.Cloudbreak.V3WorkspaceIDRdsconfigs,
		c.Int64(fl.FlWorkspaceOptional.Name),
		c.String(fl.FlRdsUserName.Name),
		c.String(fl.FlRdsPassword.Name),
		c.String(fl.FlRdsURL.Name),
		c.String(fl.FlRdsConnectorJarURLOptional.Name))
}

func testRdsByNameImpl(client rdsClient, workspaceID int64, name string) {
	defer utils.TimeTrack(time.Now(), "test database configuration by name")
	rdsRequest := &model.RdsTestRequest{
		Name: name,
	}
	log.Infof("[testRdsByParamsImpl] sending test database configuration by parameters request")
	resp, err := client.TestRdsConnectionInWorkspace(v3_workspace_id_rdsconfigs.NewTestRdsConnectionInWorkspaceParams().WithWorkspaceID(workspaceID).WithBody(rdsRequest))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	if responseText := getEmptyIfNil(resp.Payload.ConnectionResult); responseText != "connected" {
		utils.LogErrorMessageAndExit(fmt.Sprintf("database configuration test result: %s", responseText))
	}
}

func testRdsByParamsImpl(client rdsClient, workspaceID int64, username string, password string, URL string, jarURL string) {
	defer utils.TimeTrack(time.Now(), "test database configuration by parameters")
	rdsRequest := &model.RdsTestRequest{
		RdsConfig: &model.RdsConfig{
			Name:               &(&types.S{S: "testconnection"}).S,
			ConnectionUserName: &username,
			ConnectionPassword: &password,
			ConnectionURL:      &URL,
			Type:               &(&types.S{S: "testtype"}).S,
			ConnectorJarURL:    jarURL,
		},
	}
	log.Infof("[testRdsByParamsImpl] sending test database configuration by parameters request")
	resp, err := client.TestRdsConnectionInWorkspace(v3_workspace_id_rdsconfigs.NewTestRdsConnectionInWorkspaceParams().WithWorkspaceID(workspaceID).WithBody(rdsRequest))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	if responseText := getEmptyIfNil(resp.Payload.ConnectionResult); responseText != "connected" {
		utils.LogErrorMessageAndExit(fmt.Sprintf("database configuration test result: %s", responseText))
	}
}

func CreateRdsOracle11(c *cli.Context) {
	log.Infof("[CreateRdsOracle11] creating a database configuration")
	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	createRdsImpl(
		cbClient.Cloudbreak.V3WorkspaceIDRdsconfigs,
		c.Int64(fl.FlWorkspaceOptional.Name),
		c.String(fl.FlName.Name),
		c.String(fl.FlRdsUserName.Name),
		c.String(fl.FlRdsPassword.Name),
		c.String(fl.FlRdsURL.Name),
		c.String(fl.FlRdsType.Name),
		c.String(fl.FlRdsConnectorJarURLOptional.Name),
		&model.Oracle{
			Version: &(&types.S{S: "11"}).S,
		})
}

func CreateRdsOracle12(c *cli.Context) {
	log.Infof("[CreateRdsOracle12] creating a database configuration")
	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	createRdsImpl(
		cbClient.Cloudbreak.V3WorkspaceIDRdsconfigs,
		c.Int64(fl.FlWorkspaceOptional.Name),
		c.String(fl.FlName.Name),
		c.String(fl.FlRdsUserName.Name),
		c.String(fl.FlRdsPassword.Name),
		c.String(fl.FlRdsURL.Name),
		c.String(fl.FlRdsType.Name),
		c.String(fl.FlRdsConnectorJarURLOptional.Name),
		&model.Oracle{
			Version: &(&types.S{S: "12"}).S,
		})
}

func CreateRds(c *cli.Context) {
	log.Infof("[CreateRds] creating a database configuration")
	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	createRdsImpl(
		cbClient.Cloudbreak.V3WorkspaceIDRdsconfigs,
		c.Int64(fl.FlWorkspaceOptional.Name),
		c.String(fl.FlName.Name),
		c.String(fl.FlRdsUserName.Name),
		c.String(fl.FlRdsPassword.Name),
		c.String(fl.FlRdsURL.Name),
		c.String(fl.FlRdsType.Name),
		c.String(fl.FlRdsConnectorJarURLOptional.Name),
		nil)
}

func createRdsImpl(client rdsClient, workspaceID int64, name string, username string, password string, URL string, rdsType string, jarURL string, oracle *model.Oracle) {
	defer utils.TimeTrack(time.Now(), "create database")
	rdsRequest := &model.RdsConfig{
		Name:               &name,
		ConnectionUserName: &username,
		ConnectionPassword: &password,
		ConnectionURL:      &URL,
		Type:               &rdsType,
		ConnectorJarURL:    jarURL,
	}
	if oracle != nil {
		rdsRequest.Oracle = oracle
	}
	var rdsResponse *model.RDSConfigResponse
	log.Infof("[createRdsImpl] sending create database request")
	resp, err := client.CreateRdsConfigInWorkspace(v3_workspace_id_rdsconfigs.NewCreateRdsConfigInWorkspaceParams().WithWorkspaceID(workspaceID).WithBody(rdsRequest))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	rdsResponse = resp.Payload

	log.Infof("[createRdsImpl] rds created: %s (id: %d)", *rdsResponse.Name, rdsResponse.ID)
}

func DeleteRds(c *cli.Context) error {
	defer utils.TimeTrack(time.Now(), "delete a database configuration")

	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	rdsName := c.String(fl.FlName.Name)
	log.Infof("[DeleteRds] delete database configuration by name: %s", rdsName)
	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)

	if _, err := cbClient.Cloudbreak.V3WorkspaceIDRdsconfigs.DeleteRdsConfigInWorkspace(v3_workspace_id_rdsconfigs.NewDeleteRdsConfigInWorkspaceParams().WithWorkspaceID(workspaceID).WithName(rdsName)); err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[DeleteRds] database configuration deleted: %s", rdsName)
	return nil
}

func DescribeRds(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "describe a database configuration")
	log.Infof("[DescribeRds] Describes a database configuration")
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}

	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	rdsName := c.String(fl.FlName.Name)
	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)

	resp, err := cbClient.Cloudbreak.V3WorkspaceIDRdsconfigs.GetRdsConfigInWorkspace(v3_workspace_id_rdsconfigs.NewGetRdsConfigInWorkspaceParams().WithWorkspaceID(workspaceID).WithName(rdsName))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	r := resp.Payload

	output.Write(append(rdsHeader, "ID"), &rdsOutDescribe{
		&rds{
			Name:           *r.Name,
			URL:            *r.ConnectionURL,
			DatabaseEngine: *r.DatabaseEngine,
			Type:           getEmptyIfNil(r.Type),
			Driver:         getEmptyIfNil(r.ConnectionDriver),
		}, strconv.FormatInt(r.ID, 10)})

}

func ListAllRds(c *cli.Context) error {
	defer utils.TimeTrack(time.Now(), "list database configurations")

	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)

	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	return listAllRdsImpl(cbClient.Cloudbreak.V3WorkspaceIDRdsconfigs, output.WriteList, workspaceID)
}

func listAllRdsImpl(rdsClient rdsClient, writer func([]string, []utils.Row), workspaceID int64) error {
	resp, err := rdsClient.ListRdsConfigsByWorkspace(v3_workspace_id_rdsconfigs.NewListRdsConfigsByWorkspaceParams().WithWorkspaceID(workspaceID))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	var tableRows []utils.Row
	for _, r := range resp.Payload {
		row := &rds{
			Name:           *r.Name,
			URL:            *r.ConnectionURL,
			DatabaseEngine: *r.DatabaseEngine,
			Type:           getEmptyIfNil(r.Type),
			Driver:         getEmptyIfNil(r.ConnectionDriver),
		}
		tableRows = append(tableRows, row)
	}

	writer(rdsHeader, tableRows)
	return nil
}

func getEmptyIfNil(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
