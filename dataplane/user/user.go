package user

import (
	"github.com/hortonworks/cb-cli/dataplane/oauth"
	"time"

	"github.com/go-openapi/swag"
	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/cb-cli/dataplane/oauthapi/client/oidc"
	"github.com/hortonworks/cb-cli/dataplane/oauthapi/client/roles"
	"github.com/hortonworks/cb-cli/dataplane/oauthapi/client/users"
	"github.com/hortonworks/cb-cli/dataplane/oauthapi/model"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/dp-cli-common/utils"
	"github.com/urfave/cli"
)

/* var userListHeader = []string{"User ID", "Name"}

type userListOut struct {
	User *model.UserResponseJSON
}

func (u *userListOut) DataAsStringArray() []string {
	return []string{u.User.UserID, u.User.Username}
}

func ListUsers(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "list users")
	log.Infof("[ListUsers] List all users in a tenant")
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}

	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)

	resp, err := cbClient.Cloudbreak.V1users.GetAllUsers(v1users.NewGetAllUsersParams())
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	tableRows := []utils.Row{}
	for _, user := range resp.Payload {
		tableRows = append(tableRows, &userListOut{user})
	}
	output.WriteList(userListHeader, tableRows)
} */

var userListHeader = []string{"ID", "PreferredUsername", "Name", "Email"}
var userInfoHeader = []string{"ID", "PrefferedUserName", "Name", "Email"}
var roleDetailsHeader = []string{"ID", "Name", "DisplayName", "Service"}

type UsersInfoOut struct {
	User *model.UserInfo
}

func (u *UsersInfoOut) DataAsStringArray() []string {
	return []string{
		u.User.ID.String(),
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
		u.User.ID.String(),
		swag.StringValue(u.User.PreferredUsername),
		swag.StringValue(u.User.Name),
		u.User.Email,
	}
}

type roleDetailsOut struct {
	Role *model.RoleWithPermissionDetails
}

func (r *roleDetailsOut) DataAsStringArray() []string {
	return []string{
		r.Role.ID.String(),
		swag.StringValue(r.Role.Name),
		swag.StringValue(r.Role.DisplayName),
		swag.StringValue(r.Role.Service.Name),
	}
}

type roleOut struct {
	Role *model.RoleWithPermission
}

type userClient interface {
	GetAllUsers(params *users.GetAllUsersParams) (*users.GetAllUsersOK, error)
}

// ListUsers :
func ListUsers(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "list users")
	log.Infof("[ListUsers] List all users in a tenant")
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	dpClient := oauth.NewDataplaneHTTPClientFromContext(c)
	listUsersImpl(dpClient.Dataplane.Users, output.WriteList)
}

func listUsersImpl(client userClient, writer func([]string, []utils.Row)) {
	resp, err := client.GetAllUsers(users.NewGetAllUsersParams())
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	tableRows := []utils.Row{}
	for _, user := range resp.Payload {
		tableRows = append(tableRows, &userListOut{user})
	}
	writer(userListHeader, tableRows)
}

// ListRoles :
func ListRoles(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "list roles")
	log.Infof("[ListUsers] List all roles of the user")
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	dpClient := oauth.NewDataplaneHTTPClientFromContext(c)
	resp, err := dpClient.Dataplane.Roles.GetRoles(roles.NewGetRolesParams())
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	tableRows := []utils.Row{}
	for _, role := range resp.Payload {
		tableRows = append(tableRows, &roleDetailsOut{role})
	}
	output.WriteList(roleDetailsHeader, tableRows)
}

// Userinfo :
func Userinfo(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "list user information")
	log.Infof("[ListUsers] List information for connected user")
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	dpClient := oauth.NewDataplaneHTTPClientFromContext(c)
	resp, err := dpClient.Dataplane.Oidc.LoggedInUserInfo(oidc.NewLoggedInUserInfoParams())
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	output.Write(userInfoHeader, &UsersInfoOut{resp.Payload})
}
