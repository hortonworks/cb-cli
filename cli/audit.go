package cli

import (
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/cb-cli/cli/utils"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v3_organization_id_audits"
	"github.com/hortonworks/cb-cli/models_cloudbreak"
	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

var auditListHeader = []string{"AuditID", "EventType", "TimeStamp", "ResourceId", "ResourceName", "ResourceType", "UserName", "Status", "Duration"}

type auditListOut struct {
	Audit *models_cloudbreak.AuditEvent `json:"Audit" yaml:"Audit"`
}

func (a *auditListOut) DataAsStringArray() []string {
	return []string{strconv.FormatInt(a.Audit.AuditID, 10), a.Audit.Operation.EventType, a.Audit.Operation.ZonedDateTime.String(), strconv.FormatInt(a.Audit.Operation.ResourceID, 10), a.Audit.Operation.ResourceName, a.Audit.Operation.ResourceType, a.Audit.Operation.UserName, a.Audit.Status, strconv.FormatInt(a.Audit.Duration, 10)}
}

var auditHeader = []string{"Audit"}

type auditOut struct {
	Audit *models_cloudbreak.AuditEvent `json:"Audit" yaml:"Audit"`
}

func (a *auditOut) DataAsStringArray() []string {
	if a.Audit.RawFlowEvent.Blueprint != nil {
		a.Audit.RawFlowEvent.Blueprint.BlueprintJSON = "---TRUNCATED---"
	}
	auditYAML, err := yaml.Marshal(a.Audit)
	if err != nil {
		return []string{err.Error()}
	}
	return []string{string(auditYAML)}
}

func ListBlueprintAudits(c *cli.Context) {
	listAudits("blueprints", c)
}

func ListClusterAudits(c *cli.Context) {
	listAudits("stacks", c)
}

func ListCredentialAudits(c *cli.Context) {
	listAudits("credentials", c)
}

func ListDatabaseAudits(c *cli.Context) {
	listAudits("rdsconfigs", c)
}

func ListImagecatalogAudits(c *cli.Context) {
	listAudits("imagecatalogs", c)
}

func ListLdapAudits(c *cli.Context) {
	listAudits("ldapconfigs", c)
}

func ListRecipeAudits(c *cli.Context) {
	listAudits("recipes", c)
}

func listAudits(resourceType string, c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "list audits")
	log.Infof("[ListAudits] List all audits for a resource identified by Resource ID")
	orgID := c.Int64(FlOrganizationOptional.Name)
	output := utils.Output{Format: c.String(FlOutputOptional.Name)}
	resourceID, err := strconv.ParseInt(c.String(FlResourceID.Name), 10, 64)
	if err != nil {
		utils.LogErrorMessageAndExit("Unable to parse as number: " + c.String(FlResourceID.Name))
	}

	cbClient := NewCloudbreakHTTPClientFromContext(c)

	resp, err := cbClient.Cloudbreak.V3OrganizationIDAudits.GetAuditEventsInOrganization(v3_organization_id_audits.NewGetAuditEventsInOrganizationParams().WithOrganizationID(orgID).WithResourceType(resourceType).WithResourceID(resourceID))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	tableRows := []utils.Row{}
	for _, audit := range resp.Payload {
		tableRows = append(tableRows, &auditListOut{audit})
	}
	output.WriteList(auditListHeader, tableRows)
}

func DescribeAudit(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "describe audit")
	log.Infof("[DescribeAudit] Show audit entry identified by Audit ID")
	orgID := c.Int64(FlOrganizationOptional.Name)
	output := utils.Output{Format: c.String(FlOutputOptional.Name)}
	auditID, err := strconv.ParseInt(c.String(FlAuditID.Name), 10, 64)
	if err != nil {
		utils.LogErrorMessageAndExit("Unable to parse as number: " + c.String(FlAuditID.Name))
	}

	cbClient := NewCloudbreakHTTPClientFromContext(c)

	resp, err := cbClient.Cloudbreak.V3OrganizationIDAudits.GetAuditEventByOrganization(v3_organization_id_audits.NewGetAuditEventByOrganizationParams().WithOrganizationID(orgID).WithAuditID(auditID))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	tableRows := []utils.Row{}
	tableRows = append(tableRows, &auditOut{resp.Payload})

	output.Write(auditHeader, &auditOut{resp.Payload})
}
