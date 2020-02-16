package tag

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/hortonworks/cb-cli/dataplane/common"
	"github.com/hortonworks/cb-cli/dataplane/oauth"

	v1tags "github.com/hortonworks/cb-cli/dataplane/api-environment/client/v1tags"
	"github.com/hortonworks/cb-cli/dataplane/api-environment/model"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/dp-cli-common/utils"
	log "github.com/sirupsen/logrus"
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

func CreateAccountTagsFromFile(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "create account tags")

	req := assembleAccountTagsRequest(c.String(fl.FlInputJson.Name))
	envClient := oauth.NewEnvironmentClientFromContext(c)
	postAccountTags(envClient.Environment.V1tags, req)
}

func assembleAccountTagsRequest(path string) *model.AccountTagRequests {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		utils.LogErrorAndExit(err)
	}

	log.Infof("[assembleAccountTagsRequest] read AccountTagRequests json from file: %s", path)
	content := utils.ReadFile(path)

	var req model.AccountTagRequests
	err := json.Unmarshal(content, &req)
	if err != nil {
		msg := fmt.Sprintf(`Invalid json format: %s. Please make sure that the json is valid (check for commas and double quotes).`, err.Error())
		utils.LogErrorMessageAndExit(msg)
	}

	return &req
}

type createAccountTagsClient interface {
	PutTagsV1(*v1tags.PutTagsV1Params) (*v1tags.PutTagsV1OK, error)
}

func postAccountTags(client createAccountTagsClient, tagReq *model.AccountTagRequests) *model.AccountTagResponses {
	var tagResp *model.AccountTagResponses

	log.Infof("[postCredential] sending create public credential request")
	resp, err := client.PutTagsV1(v1tags.NewPutTagsV1Params().WithBody(tagReq))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	tagResp = resp.Payload
	return tagResp
}

func ListAccountTags(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "list account tags")

	envClient := oauth.NewEnvironmentClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	listAccountTagsImpl(envClient.Environment.V1tags, output.WriteList)
}

type accountTagsListOutDescribe struct {
	*common.CloudResourceOut
}

type listAccountTagsClient interface {
	ListTagsV1(*v1tags.ListTagsV1Params) (*v1tags.ListTagsV1OK, error)
}

func listAccountTagsImpl(client listAccountTagsClient, writer func([]string, []utils.Row)) {
	log.Infof("[listAccountTagsImpl] sending account tags list request")
	tagResp, err := client.ListTagsV1(v1tags.NewListTagsV1Params())
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	var tableRows []utils.Row
	for _, tag := range tagResp.Payload.Responses {
		tableRows = append(tableRows, &accountTagOut{Key: tag.Key, Value: tag.Value})
	}

	writer(accountTagHeader, tableRows)
}
