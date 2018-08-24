package cli

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/cb-cli/cli/types"
	"github.com/hortonworks/cb-cli/cli/utils"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v3_organization_id_stack"
	"github.com/hortonworks/cb-cli/models_cloudbreak"
	"github.com/urfave/cli"
)

func ChangeAmbariPassword(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "update ambari password")

	cbClient := NewCloudbreakHTTPClientFromContext(c)
	orgID := c.Int64(FlOrganizationOptional.Name)
	name := c.String(FlName.Name)
	log.Infof("[ChangeAmbariPassword] updating ambari password, name: %s", name)
	req := &models_cloudbreak.UserNamePassword{
		OldPassword: &(&types.S{S: c.String(FlOldPassword.Name)}).S,
		Password:    &(&types.S{S: c.String(FlNewPassword.Name)}).S,
		UserName:    &(&types.S{S: c.String(FlAmbariUser.Name)}).S,
	}
	err := cbClient.Cloudbreak.V3OrganizationIDStack.PutpasswordStackV3(v3_organization_id_stack.NewPutpasswordStackV3Params().WithOrganizationID(orgID).WithName(name).WithBody(req))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[RepairStack] ambari password updated, name: %s", name)
}
