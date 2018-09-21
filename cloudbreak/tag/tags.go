package tag

import (
	"github.com/hortonworks/cb-cli/cloudbreak/oauth"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/cb-cli/cloudbreak/api/client/v1accountpreferences"
	"github.com/hortonworks/cb-cli/cloudbreak/api/model"
	fl "github.com/hortonworks/cb-cli/cloudbreak/flags"
	"github.com/hortonworks/cb-cli/utils"
	"github.com/urfave/cli"
)

var accountTagHeader = []string{"Key", "Value"}

type accountTagOut struct {
	Key   string `json:"Key" yaml:"Key"`
	Value string `json:"Value" yaml:"Value"`
}

func (r *accountTagOut) DataAsStringArray() []string {
	return []string{r.Key, r.Value}
}

func ListAccountTags(c *cli.Context) {
	fl.CheckRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "list default tags for account")

	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	listAccountTagsImpl(cbClient.Cloudbreak.V1accountpreferences, output.WriteList)
}

func AddAccountTag(c *cli.Context) {
	fl.CheckRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "add a default tag for account")

	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	addAccountTagsImpl(cbClient.Cloudbreak.V1accountpreferences, c.String(fl.FlKey.Name), c.String(fl.FlValue.Name))
}

func DeleteAccountTag(c *cli.Context) {
	fl.CheckRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "delete a default tag of the account")

	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	deleteAccountTagsImpl(cbClient.Cloudbreak.V1accountpreferences, c.String(fl.FlKey.Name))
}

type accountTagsClient interface {
	GetAccountPreferencesEndpoint(params *v1accountpreferences.GetAccountPreferencesEndpointParams) (*v1accountpreferences.GetAccountPreferencesEndpointOK, error)
	PutAccountPreferencesEndpoint(params *v1accountpreferences.PutAccountPreferencesEndpointParams) (*v1accountpreferences.PutAccountPreferencesEndpointOK, error)
}

func listAccountTagsImpl(client accountTagsClient, writer func([]string, []utils.Row)) {
	log.Infof("[listAccountTagsImpl] sending default tags for account list request")
	accountprefResp, err := client.GetAccountPreferencesEndpoint(v1accountpreferences.NewGetAccountPreferencesEndpointParams())
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	tableRows := []utils.Row{}
	for key, val := range accountprefResp.Payload.DefaultTags {
		tableRows = append(tableRows, &accountTagOut{key, val})
	}
	writer(accountTagHeader, tableRows)
}

func addAccountTagsImpl(client accountTagsClient, key string, value string) {
	log.Infof("[addAccountTagsImpl] sending add default tags for account request")
	accountprefResp, err := client.GetAccountPreferencesEndpoint(v1accountpreferences.NewGetAccountPreferencesEndpointParams())
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	accountPref := accountprefResp.Payload
	if accountPref.DefaultTags == nil {
		accountPref.DefaultTags = make(map[string]string)
	}
	accountPrefReq := &model.AccountPreferencesRequest{
		AllowedInstanceTypes:       accountPref.AllowedInstanceTypes,
		ClusterTimeToLive:          accountPref.ClusterTimeToLive,
		DefaultTags:                accountPref.DefaultTags,
		MaxNumberOfClusters:        accountPref.MaxNumberOfClusters,
		MaxNumberOfClustersPerUser: accountPref.MaxNumberOfClustersPerUser,
		MaxNumberOfNodesPerCluster: accountPref.MaxNumberOfNodesPerCluster,
		Platforms:                  accountPref.Platforms,
		SmartsenseEnabled:          accountPref.SmartsenseEnabled,
		UserTimeToLive:             accountPref.UserTimeToLive,
	}
	accountPref.DefaultTags[key] = value
	_, err = client.PutAccountPreferencesEndpoint(v1accountpreferences.NewPutAccountPreferencesEndpointParams().WithBody(accountPrefReq))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[addAccountTagsImpl] default tag added to account, key: %s, value: %s", key, value)
}

func deleteAccountTagsImpl(client accountTagsClient, key string) {
	log.Infof("[deleteAccountTagsImpl] sending delete default tags for account request")
	accountprefResp, err := client.GetAccountPreferencesEndpoint(v1accountpreferences.NewGetAccountPreferencesEndpointParams())
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	accountPref := accountprefResp.Payload
	accountPrefReq := &model.AccountPreferencesRequest{
		AllowedInstanceTypes:       accountPref.AllowedInstanceTypes,
		ClusterTimeToLive:          accountPref.ClusterTimeToLive,
		DefaultTags:                accountPref.DefaultTags,
		MaxNumberOfClusters:        accountPref.MaxNumberOfClusters,
		MaxNumberOfClustersPerUser: accountPref.MaxNumberOfClustersPerUser,
		MaxNumberOfNodesPerCluster: accountPref.MaxNumberOfNodesPerCluster,
		Platforms:                  accountPref.Platforms,
		SmartsenseEnabled:          accountPref.SmartsenseEnabled,
		UserTimeToLive:             accountPref.UserTimeToLive,
	}
	delete(accountPref.DefaultTags, key)
	_, err = client.PutAccountPreferencesEndpoint(v1accountpreferences.NewPutAccountPreferencesEndpointParams().WithBody(accountPrefReq))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[deleteAccountTagsImpl] default tag deleted from account, key: %s", key)
}
