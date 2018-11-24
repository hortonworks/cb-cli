package userinfo

import (
	"time"

	//"github.com/hortonworks/cb-cli/cloudbreak/users"

	log "github.com/Sirupsen/logrus"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	fl "github.com/hortonworks/cb-cli/cloudbreak/flags"
	"github.com/hortonworks/cb-cli/cloudbreak/oauth"
	"github.com/hortonworks/cb-cli/dataplane/api/client/oidc"
	"github.com/hortonworks/cb-cli/dataplane/api/client/users"
	"github.com/hortonworks/cb-cli/dataplane/api/model"
	"github.com/hortonworks/cb-cli/dps-common/utils"
	"github.com/urfave/cli"
)

var userInfoHeader = []string{"ID", "TenantID", "PrefferedUserName", "Name", "Email"}
var userListHeader = []string{"ID", "TenantID", "PrefferedUserName", "Name", "Email"}

type UsersInfoOut struct {
	User *model.UserInfo
}

func (u *UsersInfoOut) DataAsStringArray() []string {
	return []string{
		strfmt.UUID.String(u.User.ID),
		strfmt.UUID.String(*u.User.TenantID),
		swag.StringValue(u.User.PreferredUsername),
		swag.StringValue(u.User.Name),
		u.User.Email,
	}
}

type userListOut struct {
	User *model.UserWithRoles
}

func (u *userListOut) DataAsStringArray() []string {
	return []string{
		strfmt.UUID.String(u.User.ID),
		strfmt.UUID.String(*u.User.TenantID),
		swag.StringValue(u.User.PreferredUsername),
		swag.StringValue(u.User.Name),
	}
}

func LoggedInUserinfo(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "show logged in userinfo")
	log.Info("[LoggedInUserinfo] List information of logged in user")
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	dpClient := oauth.NewDataplaneHTTPClientFromContext(c)
	userID := c.String(fl.FlUserIDOptional.Name)
	if userID == "" {
		resp, err := dpClient.Dataplane.Oidc.LoggedInUserInfo(oidc.NewLoggedInUserInfoParams())
		if err != nil {
			utils.LogErrorAndExit(err)
		}
		output.Write(userInfoHeader, &UsersInfoOut{resp.Payload})
	} else {
		resp, err := dpClient.Dataplane.Users.GetAllUsers(users.NewGetAllUsersParams().WithID(&userID))
		if err != nil {
			utils.LogErrorAndExit(err)
		}
		tableRows := []utils.Row{}
		for _, user := range resp.Payload {
			tableRows = append(tableRows, &userListOut{user})
		}
		output.WriteList(userListHeader, tableRows)
	}
}
