package auditcredential

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/hortonworks/cb-cli/dataplane/common"
	"github.com/hortonworks/cb-cli/dataplane/oauth"

	v1auditcred "github.com/hortonworks/cb-cli/dataplane/api-environment/client/v1credentialsaudit"

	"errors"

	"github.com/hortonworks/cb-cli/dataplane/api-environment/model"
	"github.com/hortonworks/cb-cli/dataplane/cloud"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/dp-cli-common/utils"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var AwsPrerequisiteOutputHeader = []string{"Account Id", "CloudPlatform", "External Id", "Policy JSON"}

type awsPrerequisiteOutput struct {
	AccountId     string `json:"AccountId" yaml:"AccountId"`
	CloudPlatform string `json:"CloudPlatform" yaml:"CloudPlatform"`
	ExternalId    string `json:"ExternalId" yaml:"ExternalId"`
	PolicyJSON    string `json:"PolicyJSON" yaml:"PolicyJSON"`
}

func (p *awsPrerequisiteOutput) DataAsStringArray() []string {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(p.PolicyJSON), &raw)
	if err != nil {
		return []string{p.CloudPlatform, p.AccountId, p.ExternalId, string(err.Error())}
	}

	policyJSON, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return []string{p.CloudPlatform, p.AccountId, p.ExternalId, string(err.Error())}
	}
	return []string{p.AccountId, p.CloudPlatform, p.ExternalId, string(policyJSON)}
}

func CreateAwsCredential(c *cli.Context) {
	cloud.SetProviderType(cloud.AWS)
	createAuditCredential(c, false)
}

func ModifyAwsCredential(c *cli.Context) {
	cloud.SetProviderType(cloud.AWS)
	modifyAuditCredential(c, false)
}

func CreateAzureCredential(c *cli.Context) {
	cloud.SetProviderType(cloud.AZURE)
	createAuditCredential(c, false)
}

func ModifyAzureCredential(c *cli.Context) {
	cloud.SetProviderType(cloud.AZURE)
	modifyAuditCredential(c, false)
}

func CreateCredentialFromFile(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "create audit credential")

	req := assembleCredentialRequest(c.String(fl.FlInputJson.Name), c.String(fl.FlNameOptional.Name))
	envClient := oauth.NewEnvironmentClientFromContext(c)
	postAuditCredential(envClient.Environment.V1credentialsaudit, req)
}

func ModifyCredentialFromFile(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "modify audit credential from file")

	credReq := assembleCredentialRequest(c.String(fl.FlInputJson.Name), c.String(fl.FlNameOptional.Name))
	envClient := oauth.NewEnvironmentClientFromContext(c)

	putAuditCredential(envClient.Environment.V1credentialsaudit, credReq)
}

func assembleCredentialRequest(path, credName string) *model.CredentialV1Request {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		utils.LogErrorAndExit(err)
	}

	log.Infof("[assembleCredentialRequest] read audit credential json from file: %s", path)
	content := utils.ReadFile(path)

	var req model.CredentialV1Request
	err := json.Unmarshal(content, &req)
	if err != nil {
		msg := fmt.Sprintf(`Invalid json format: %s. Please make sure that the json is valid (check for commas and double quotes).`, err.Error())
		utils.LogErrorMessageAndExit(msg)
	}

	if len(credName) != 0 {
		req.Name = &credName
	}
	if req.Name == nil || len(*req.Name) == 0 {
		utils.LogErrorMessageAndExit("Name of the credential must be set either in the template or with the --name command line option.")
	}

	return &req
}

func createAuditCredential(c *cli.Context, govCloud bool) {
	defer utils.TimeTrack(time.Now(), "create audit credential")

	envClient := oauth.NewEnvironmentClientFromContext(c)
	verifyPermissions := c.Bool(fl.FlCredentialVerifyOptional.Name)
	createAuditCredentialImpl(c.String, envClient.Environment.V1credentialsaudit, govCloud, verifyPermissions)
}

func modifyAuditCredential(c *cli.Context, govCloud bool) {
	defer utils.TimeTrack(time.Now(), "modify audit credential")

	envClient := oauth.NewEnvironmentClientFromContext(c)
	verifyPermissions := c.Bool(fl.FlCredentialVerifyOptional.Name)
	modifyAuditCredentialImpl(c.String, envClient.Environment.V1credentialsaudit, govCloud, verifyPermissions)
}

type modifyAuditCredentialClient interface {
	GetAuditCredentialByResourceCrn(params *v1auditcred.GetAuditCredentialByResourceCrnV1Params) (*v1auditcred.GetAuditCredentialByResourceCrnV1OK, error)
	PutAuditCredential(params *v1auditcred.PutAuditCredentialV1Params) (*v1auditcred.PutAuditCredentialV1OK, error)
}

func modifyAuditCredentialImpl(stringFinder func(string) string, client modifyAuditCredentialClient, govCloud bool, verifyPermissions bool) *model.CredentialV1Response {
	name := stringFinder(fl.FlName.Name)
	description := stringFinder(fl.FlDescriptionOptional.Name)
	provider := cloud.GetProvider()

	credential := getAuditCredential(name, client)
	if *credential.CloudPlatform != *provider.GetName() {
		utils.LogErrorAndExit(errors.New("cloud provider cannot be modified"))
	}

	log.Infof("[modifyCredentialImpl] original credential found name: %s id: %s", name, credential.Crn)
	credReq, err := provider.GetCredentialRequest(stringFinder, govCloud)
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	if len(description) == 0 {
		origDesc := credential.Description
		if origDesc != nil && len(*origDesc) > 0 {
			description = *origDesc
		}
	}
	credReq.Description = &description
	credReq.VerifyPermissions = verifyPermissions

	return putAuditCredential(client, credReq)
}

func getAuditCredential(crn string, client modifyAuditCredentialClient) *model.CredentialV1Response {
	defer utils.TimeTrack(time.Now(), "get credential")

	log.Infof("[getCredential] get credential by crn: %s", crn)
	response, err := client.GetAuditCredentialByResourceCrn(v1auditcred.NewGetAuditCredentialByResourceCrnV1Params().WithCrn(crn))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	return response.Payload
}

func putAuditCredential(client modifyAuditCredentialClient, credReq *model.CredentialV1Request) *model.CredentialV1Response {
	var credential *model.CredentialV1Response
	log.Infof("[putCredential] modify public credential: %s", *credReq.Name)

	resp, err := client.PutAuditCredential(v1auditcred.NewPutAuditCredentialV1Params().WithBody(credReq))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	credential = resp.Payload

	return credential
}

type createAuditCredentialClient interface {
	CreateAuditCredentialV1(*v1auditcred.CreateAuditCredentialV1Params) (*v1auditcred.CreateAuditCredentialV1OK, error)
}

func createAuditCredentialImpl(stringFinder func(string) string, client createAuditCredentialClient, govCloud bool, verifyPermissions bool) {
	provider := cloud.GetProvider()
	credReq, err := provider.GetCredentialRequest(stringFinder, govCloud)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	credReq.VerifyPermissions = verifyPermissions
	postAuditCredential(client, credReq)
}

func postAuditCredential(client createAuditCredentialClient, credReq *model.CredentialV1Request) *model.CredentialV1Response {
	var credential *model.CredentialV1Response

	log.Infof("[postAuditCredential] sending create public audit credential request")
	resp, err := client.CreateAuditCredentialV1(v1auditcred.NewCreateAuditCredentialV1Params().WithBody(credReq))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	credential = resp.Payload

	log.Infof("[postAuditCredential] audit credential created: %s (id: %s)", *credential.Name, credential.Crn)
	return credential
}

type auditCredentialOutDescribe struct {
	*common.CloudResourceOut
	CRN                string `json:"CRN" yaml:"CRN"`
	VerificationStatus string `json:"VerificationStatus" yaml:"VerificationStatus"`
}

func (c *auditCredentialOutDescribe) DataAsStringArray() []string {
	return append(c.CloudResourceOut.DataAsStringArray(), c.CRN, c.VerificationStatus)
}

func DescribeAuditCredential(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "describe audit credential")

	envClient := oauth.NewEnvironmentClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	resp, err := envClient.Environment.V1credentialsaudit.GetAuditCredentialByResourceCrnV1(v1auditcred.NewGetAuditCredentialByResourceCrnV1Params().WithCrn(c.String(fl.FlCrn.Name)))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	cred := resp.Payload
	output.Write(append(common.CloudResourceHeader, "ID", "VerificationStatus"),
		&auditCredentialOutDescribe{&common.CloudResourceOut{Name: *cred.Name, Description: *cred.Description, CloudPlatform: *cred.CloudPlatform},
			cred.Crn, cred.VerificationStatusText})
}

func ListAuditCredentials(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "list credentials")

	envClient := oauth.NewEnvironmentClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	listAuditCredentialsImpl(envClient.Environment.V1credentialsaudit, output.WriteList)
}

type auditcredentialListOutDescribe struct {
	*common.CloudResourceOut
	RoleArn *string `json:"RoleArn" yaml:"RoleArn"`
}

type listAuditCredentialsByWorkspaceClient interface {
	ListAuditCredentialsV1(*v1auditcred.ListAuditCredentialsV1Params) (*v1auditcred.ListAuditCredentialsV1OK, error)
}

func listAuditCredentialsImpl(client listAuditCredentialsByWorkspaceClient, writer func([]string, []utils.Row)) {
	log.Infof("[listAuditCredentialsImpl] sending credential list request")
	credResp, err := client.ListAuditCredentialsV1(v1auditcred.NewListAuditCredentialsV1Params())
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	var tableRows []utils.Row
	for _, cred := range credResp.Payload.Responses {
		if cred.Aws != nil && cred.Aws.RoleBased != nil && cred.Aws.RoleBased.RoleArn != nil {
			tableRows = append(tableRows,
				&auditcredentialListOutDescribe{&common.CloudResourceOut{Name: *cred.Name, Description: *cred.Description, CloudPlatform: *cred.CloudPlatform},
					cred.Aws.RoleBased.RoleArn})
		} else {
			tableRows = append(tableRows, &common.CloudResourceOut{Name: *cred.Name, Description: *cred.Description, CloudPlatform: *cred.CloudPlatform})
		}
	}

	writer(common.CloudResourceHeader, tableRows)
}

func GetAwsAuditCredentialPrerequisites(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "get audit credentials prerequisites for Aws")

	envClient := oauth.NewEnvironmentClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}

	log.Infof("[GetAwsAuditCredentialPrerequisites] sending Aws credential prerequisites request")
	prerequisitesForCloudPlatformParams := v1auditcred.NewGetAuditPrerequisitesForCloudPlatformParams().WithCloudPlatform("aws")
	credPrereqResp, err := envClient.Environment.V1credentialsaudit.GetAuditPrerequisitesForCloudPlatform(prerequisitesForCloudPlatformParams)
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	prerequisites := credPrereqResp.Payload
	output.Write(AwsPrerequisiteOutputHeader, convertAwsPrerequisitesToJSON(prerequisites))
}

func convertAwsPrerequisitesToJSON(prerequisites *model.CredentialPrerequisitesResponse) *awsPrerequisiteOutput {
	policyJsonEncoded := utils.SafeStringConvert(prerequisites.Aws.PolicyJSON)
	policyJson, err := base64.StdEncoding.DecodeString(policyJsonEncoded)
	if err != nil {
		utils.LogErrorMessageAndExit("Could not parse AWS cb-policy.json from the response.")
	}
	return &awsPrerequisiteOutput{
		AccountId:     prerequisites.AccountID,
		CloudPlatform: utils.SafeStringConvert(prerequisites.CloudPlatform),
		ExternalId:    utils.SafeStringConvert(prerequisites.Aws.ExternalID),
		PolicyJSON:    string(policyJson),
	}
}
