package cli

import (
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/cb-cli/cli/utils"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v3_organization_id_mpacks"
	"github.com/hortonworks/cb-cli/models_cloudbreak"
	"github.com/urfave/cli"
)

type mpackClient interface {
	CreateManagementPackInOrganization(params *v3_organization_id_mpacks.CreateManagementPackInOrganizationParams) (*v3_organization_id_mpacks.CreateManagementPackInOrganizationOK, error)
	ListManagementPacksByOrganization(params *v3_organization_id_mpacks.ListManagementPacksByOrganizationParams) (*v3_organization_id_mpacks.ListManagementPacksByOrganizationOK, error)
}

var mpackHeader = []string{"Name", "Description", "URL", "Purge", "PurgeList", "Force"}

type mpack struct {
	Name        string `json:"Name" yaml:"Name"`
	Description string `json:"Description" yaml:"Description"`
	URL         string `json:"URL" yaml:"URL"`
	Purge       string `json:"Purge" yaml:"Purge"`
	PurgeList   string `json:"PurgeList" yaml:"PurgeList"`
	Force       string `json:"Force" yaml:"Force"`
}

func (m *mpack) DataAsStringArray() []string {
	return []string{m.Name, m.Description, m.URL, m.Purge, m.PurgeList, m.Force}
}

func CreateMpack(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	cbClient := NewCloudbreakHTTPClientFromContext(c)
	createMpackImpl(
		cbClient.Cloudbreak.V3OrganizationIDMpacks,
		c.Int64(FlOrganizationOptional.Name),
		c.String(FlName.Name),
		c.String(FlDescriptionOptional.Name),
		c.String(FLMpackURL.Name),
		c.Bool(FLMpackPurge.Name),
		c.String(FLMpackPurgeList.Name),
		c.Bool(FLMpackForce.Name))
}

func createMpackImpl(client mpackClient, orgID int64, name, description, url string, purge bool, purgeList string, force bool) {
	defer utils.TimeTrack(time.Now(), "create management pack")
	pList := make([]string, 0)
	if len(purgeList) > 0 {
		pList = strings.Split(purgeList, ",")
	}

	req := &models_cloudbreak.ManagementPackRequest{
		Name:        &name,
		Description: &description,
		MpackURL:    &url,
		Purge:       &purge,
		PurgeList:   pList,
		Force:       &force,
	}

	var mpackResponse *models_cloudbreak.ManagementPackResponse
	log.Infof("[createMpackImpl] sending create public management pack request with name: %s", name)
	resp, err := client.CreateManagementPackInOrganization(v3_organization_id_mpacks.NewCreateManagementPackInOrganizationParams().WithOrganizationID(orgID).WithBody(req))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	mpackResponse = resp.Payload
	log.Infof("[createMpackImpl] management pack created: %s (id: %d)", *mpackResponse.Name, mpackResponse.ID)
}

func DeleteMpack(c *cli.Context) error {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "delete a management pack")

	orgID := c.Int64(FlOrganizationOptional.Name)
	mpackName := c.String(FlName.Name)
	log.Infof("[DeleteMpack] delete management pack by name: %s", mpackName)
	cbClient := NewCloudbreakHTTPClientFromContext(c)

	if _, err := cbClient.Cloudbreak.V3OrganizationIDMpacks.DeleteManagementPackInOrganization(v3_organization_id_mpacks.NewDeleteManagementPackInOrganizationParams().WithOrganizationID(orgID).WithName(mpackName)); err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[DeleteMpack] management pack deleted: %s", mpackName)
	return nil
}

func ListMpacks(c *cli.Context) error {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "list management pack")

	cbClient := NewCloudbreakHTTPClientFromContext(c)

	output := utils.Output{Format: c.String(FlOutputOptional.Name)}
	orgID := c.Int64(FlOrganizationOptional.Name)
	return listMpacksImpl(cbClient.Cloudbreak.V3OrganizationIDMpacks, output.WriteList, orgID)
}

func listMpacksImpl(mpackClient mpackClient, writer func([]string, []utils.Row), orgID int64) error {
	resp, err := mpackClient.ListManagementPacksByOrganization(v3_organization_id_mpacks.NewListManagementPacksByOrganizationParams().WithOrganizationID(orgID))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	var tableRows []utils.Row
	for _, m := range resp.Payload {
		purgeString := "false"
		forceString := "false"
		desc := ""
		if m.Purge != nil && *m.Purge {
			purgeString = "true"
		}
		if m.Force != nil && *m.Force {
			forceString = "true"
		}
		if m.Description != nil && len(*m.Description) > 0 {
			desc = *m.Description
		}
		row := &mpack{
			Name:        *m.Name,
			Description: desc,
			URL:         *m.MpackURL,
			Purge:       purgeString,
			PurgeList:   strings.Join(m.PurgeList, ","),
			Force:       forceString,
		}
		tableRows = append(tableRows, row)
	}

	writer(mpackHeader, tableRows)
	return nil
}
