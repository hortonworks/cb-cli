package credential

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/hortonworks/cb-cli/dataplane/common"
	"github.com/hortonworks/cb-cli/dataplane/oauth"

	v1cred "github.com/hortonworks/cb-cli/dataplane/api-environment/client/v1credentials"

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
	createCredential(c, false)
}

func ModifyAwsCredential(c *cli.Context) {
	cloud.SetProviderType(cloud.AWS)
	modifyCredential(c, false)
}

func CreateAwsGovCredential(c *cli.Context) {
	cloud.SetProviderType(cloud.AWS)
	createCredential(c, true)
}

func ModifyAwsGovCredential(c *cli.Context) {
	cloud.SetProviderType(cloud.AWS)
	modifyCredential(c, true)
}

func CreateAzureCredential(c *cli.Context) {
	cloud.SetProviderType(cloud.AZURE)
	createCredential(c, false)
}

func ModifyAzureCredential(c *cli.Context) {
	cloud.SetProviderType(cloud.AZURE)
	modifyCredential(c, false)
}

func CreateGcpCredential(c *cli.Context) {
	cloud.SetProviderType(cloud.GCP)
	createCredential(c, false)
}

func ModifyGcpCredential(c *cli.Context) {
	cloud.SetProviderType(cloud.GCP)
	modifyCredential(c, false)
}

func CreateOpenstackCredential(c *cli.Context) {
	cloud.SetProviderType(cloud.OPENSTACK)
	createCredential(c, false)
}

func CreateYarnCredential(c *cli.Context) {
	cloud.SetProviderType(cloud.YARN)
	createCredential(c, false)
}

func ModifyOpenstackCredential(c *cli.Context) {
	cloud.SetProviderType(cloud.OPENSTACK)
	modifyCredential(c, false)
}

func CreateCredentialFromFile(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "create credential")

	req := assembleCredentialRequest(c.String(fl.FlInputJson.Name), c.String(fl.FlNameOptional.Name))
	envClient := oauth.NewEnvironmentClientFromContext(c)
	postCredential(envClient.Environment.V1credentials, req)
}

func ModifyCredentialFromFile(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "modify credential from file")

	credReq := assembleCredentialRequest(c.String(fl.FlInputJson.Name), c.String(fl.FlNameOptional.Name))
	envClient := oauth.NewEnvironmentClientFromContext(c)

	putCredential(envClient.Environment.V1credentials, credReq)
}

type verifyCredentialClient interface {
	VerifyCredentialByName(params *v1cred.VerifyCredentialByNameParams) (*v1cred.VerifyCredentialByNameOK, error)
	VerifyCredentialByCrn(params *v1cred.VerifyCredentialByCrnParams) (*v1cred.VerifyCredentialByCrnOK, error)
}

type credentialOutVerify struct {
	*common.CloudResourceOut
	VerificationStatus string `json:"VerificationStatus" yaml:"VerificationStatus"`
}

func (c *credentialOutVerify) DataAsStringArray() []string {
	return append(c.CloudResourceOut.DataAsStringArray(), c.VerificationStatus)
}

func VerifyCredential(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "validate credential")

	envClient := oauth.NewEnvironmentClientFromContext(c)
	name := c.String(fl.FlName.Name)
	crn := c.String(fl.FlCrn.Name)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	verifyCredentialImpl(envClient.Environment.V1credentials, name, crn, output.Write)
}

func verifyCredentialImpl(client verifyCredentialClient, name, crn string, writer func([]string, utils.Row)) {
	cred := verifyCredential(client, name, crn)
	writer(append(common.CloudResourceHeader, "VerificationStatus"),
		&credentialOutVerify{&common.CloudResourceOut{Name: *cred.Name,
			Description:   *cred.Description,
			CloudPlatform: *cred.CloudPlatform,
		}, cred.VerificationStatusText})
}

func assembleCredentialRequest(path, credName string) *model.CredentialV1Request {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		utils.LogErrorAndExit(err)
	}

	log.Infof("[assembleCredentialRequest] read credential json from file: %s", path)
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

func createCredential(c *cli.Context, govCloud bool) {
	defer utils.TimeTrack(time.Now(), "create credential")

	envClient := oauth.NewEnvironmentClientFromContext(c)
	verifyPermissions := c.Bool(fl.FlCredentialVerifyOptional.Name)
	createCredentialImpl(c.String, envClient.Environment.V1credentials, govCloud, verifyPermissions)
}

func modifyCredential(c *cli.Context, govCloud bool) {
	defer utils.TimeTrack(time.Now(), "modify credential")

	envClient := oauth.NewEnvironmentClientFromContext(c)
	verifyPermissions := c.Bool(fl.FlCredentialVerifyOptional.Name)
	modifyCredentialImpl(c.String, envClient.Environment.V1credentials, govCloud, verifyPermissions)
}

type modifyCredentialClient interface {
	GetCredentialByNameV1(params *v1cred.GetCredentialByNameV1Params) (*v1cred.GetCredentialByNameV1OK, error)
	PutCredentialV1(params *v1cred.PutCredentialV1Params) (*v1cred.PutCredentialV1OK, error)
}

func modifyCredentialImpl(stringFinder func(string) string, client modifyCredentialClient, govCloud bool, verifyPermissions bool) *model.CredentialV1Response {
	name := stringFinder(fl.FlName.Name)
	description := stringFinder(fl.FlDescriptionOptional.Name)
	provider := cloud.GetProvider()

	origCredential := getCredential(name, client)
	if *origCredential.CloudPlatform != *provider.GetName() {
		utils.LogErrorAndExit(errors.New("cloud provider cannot be modified"))
	}

	log.Infof("[modifyCredentialImpl] original origCredential found name: %s id: %s", name, origCredential.Crn)
	credReq, err := provider.GetCredentialRequest(stringFinder, govCloud)
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	if len(description) == 0 {
		origDesc := origCredential.Description
		if origDesc != nil && len(*origDesc) > 0 {
			description = *origDesc
		}
	}
	credReq.Description = &description
	credReq.VerifyPermissions = verifyPermissions

	if credReq.Azure != nil {
		if stringFinder("auth-type") == "" {
			credReq.Azure.AppBased.AuthenticationType = origCredential.Azure.AuthenticationType
		}
	}

	return putCredential(client, credReq)
}

func verifyCredential(client verifyCredentialClient, name, crn string) *model.CredentialV1Response {
	defer utils.TimeTrack(time.Now(), "verify credential")

	var response *model.CredentialV1Response
	if len(name) > 0 {
		log.Infof("[verifyCredential] verify credential by name: %s", name)
		responseByName, err := client.VerifyCredentialByName(v1cred.NewVerifyCredentialByNameParams().WithName(name))
		if err != nil {
			utils.LogErrorAndExit(err)
		}
		response = responseByName.Payload
	} else {
		log.Infof("[verifyCredential] verify credential by crn: %s", crn)
		responseByCrn, err := client.VerifyCredentialByCrn(v1cred.NewVerifyCredentialByCrnParams().WithCrn(crn))
		if err != nil {
			utils.LogErrorAndExit(err)
		}
		response = responseByCrn.Payload
	}
	return response
}

func getCredential(name string, client modifyCredentialClient) *model.CredentialV1Response {
	defer utils.TimeTrack(time.Now(), "get credential")

	log.Infof("[getCredential] get credential by name: %s", name)
	response, err := client.GetCredentialByNameV1(v1cred.NewGetCredentialByNameV1Params().WithName(name))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	return response.Payload
}

func putCredential(client modifyCredentialClient, credReq *model.CredentialV1Request) *model.CredentialV1Response {
	var credential *model.CredentialV1Response
	log.Infof("[putCredential] modify public credential: %s", *credReq.Name)

	resp, err := client.PutCredentialV1(v1cred.NewPutCredentialV1Params().WithBody(credReq))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	credential = resp.Payload

	return credential
}

type createCredentialClient interface {
	CreateCredentialV1(*v1cred.CreateCredentialV1Params) (*v1cred.CreateCredentialV1OK, error)
}

func createCredentialImpl(stringFinder func(string) string, client createCredentialClient, govCloud bool, verifyPermissions bool) {
	provider := cloud.GetProvider()
	credReq, err := provider.GetCredentialRequest(stringFinder, govCloud)
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	if credReq.Azure != nil {
		if stringFinder("auth-type") == model.AppBasedV1RequestAuthenticationTypeCERTIFICATE {
			credReq.Azure.AppBased.AuthenticationType = model.AppBasedV1RequestAuthenticationTypeCERTIFICATE
			credReq.Azure.AppBased.GenerateCertificate = true
		} else {
			credReq.Azure.AppBased.AuthenticationType = model.AppBasedV1RequestAuthenticationTypeSECRET
			if len(credReq.Azure.AppBased.SecretKey) == 0 {
				err = errors.New("app-password must be defined for secret based auth, if you intended to use certificate based auth then please specify the --generate-certificate flag")
			}
			if len(credReq.Azure.AppBased.AccessKey) == 0 {
				err = errors.New("app-id must be defined for secret based auth, if you intended to use certificate based auth then please specify the --generate-certificate flag")
			}
		}
	}

	credReq.VerifyPermissions = verifyPermissions
	postCredential(client, credReq)
}

func postCredential(client createCredentialClient, credReq *model.CredentialV1Request) *model.CredentialV1Response {
	var credential *model.CredentialV1Response

	log.Infof("[postCredential] sending create public credential request")
	resp, err := client.CreateCredentialV1(v1cred.NewCreateCredentialV1Params().WithBody(credReq))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	credential = resp.Payload

	log.Infof("[postCredential] credential created: %s (id: %s)", *credential.Name, credential.Crn)
	return credential
}

type credentialOutDescribe struct {
	*common.CloudResourceOut
	CRN                string `json:"CRN" yaml:"CRN"`
	VerificationStatus string `json:"VerificationStatus" yaml:"VerificationStatus"`
}

func (c *credentialOutDescribe) DataAsStringArray() []string {
	return append(c.CloudResourceOut.DataAsStringArray(), c.CRN, c.VerificationStatus)
}

func DescribeCredential(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "describe credential")

	envClient := oauth.NewEnvironmentClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	resp, err := envClient.Environment.V1credentials.GetCredentialByNameV1(v1cred.NewGetCredentialByNameV1Params().WithName(c.String(fl.FlName.Name)))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	cred := resp.Payload
	output.Write(append(common.CloudResourceHeader, "ID", "VerificationStatus"),
		&credentialOutDescribe{&common.CloudResourceOut{Name: *cred.Name, Description: *cred.Description, CloudPlatform: *cred.CloudPlatform},
			cred.Crn, cred.VerificationStatusText})
}

func DeleteCredential(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "delete credential")
	envClient := oauth.NewEnvironmentClientFromContext(c)
	val := c.String(fl.FlNames.Name)
	names := strings.Split(val[1:len(val)-1], ",")
	log.Infof("[DeleteCredential] sending delete credential request with names: %s", names)
	if _, err := envClient.Environment.V1credentials.DeleteCredentialsV1(v1cred.NewDeleteCredentialsV1Params().WithBody(names)); err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[DeleteCredential] credential(s are) deleted, name(s): %s", names)
}

func ListCredentials(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "list credentials")

	envClient := oauth.NewEnvironmentClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	listCredentialsImpl(envClient.Environment.V1credentials, output.WriteList)
}

type credentialListAWSOutDescribe struct {
	*common.CloudResourceOut
	RoleArn *string `json:"RoleArn" yaml:"RoleArn"`
}

type credentialListAzureOutDescribe struct {
	*common.CloudResourceOut
	AzureCredentialProperties *model.AzureCredentialV1ResponseParameters `json:"AzureCredentialProperties" yaml:"AzureCredentialProperties"`
}

type listCredentialsByWorkspaceClient interface {
	ListCredentialsV1(*v1cred.ListCredentialsV1Params) (*v1cred.ListCredentialsV1OK, error)
}

func listCredentialsImpl(client listCredentialsByWorkspaceClient, writer func([]string, []utils.Row)) {
	log.Infof("[listCredentialsImpl] sending credential list request")
	credResp, err := client.ListCredentialsV1(v1cred.NewListCredentialsV1Params())
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	var tableRows []utils.Row
	for _, cred := range credResp.Payload.Responses {
		if cred.Aws != nil && cred.Aws.RoleBased != nil && cred.Aws.RoleBased.RoleArn != nil {
			tableRows = append(tableRows,
				&credentialListAWSOutDescribe{&common.CloudResourceOut{Name: *cred.Name, Description: *cred.Description, CloudPlatform: *cred.CloudPlatform},
					cred.Aws.RoleBased.RoleArn})
		} else if cred.Azure != nil {
			tableRows = append(tableRows,
				&credentialListAzureOutDescribe{&common.CloudResourceOut{Name: *cred.Name, Description: *cred.Description, CloudPlatform: *cred.CloudPlatform},
					cred.Azure})
		} else {
			tableRows = append(tableRows, &common.CloudResourceOut{Name: *cred.Name, Description: *cred.Description, CloudPlatform: *cred.CloudPlatform})
		}
	}

	writer(common.CloudResourceHeader, tableRows)
}

func GetAwsCredentialPrerequisites(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "get credentials prerequisites for Aws")

	envClient := oauth.NewEnvironmentClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}

	log.Infof("[GetAwsCredentialPrerequisites] sending Aws credential prerequisites request")
	prerequisitesForCloudPlatformParams := v1cred.NewGetPrerequisitesForCloudPlatformParams().WithCloudPlatform("aws")
	credPrereqResp, err := envClient.Environment.V1credentials.GetPrerequisitesForCloudPlatform(prerequisitesForCloudPlatformParams)
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
