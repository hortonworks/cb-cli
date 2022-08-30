package azure

import (
	"github.com/hortonworks/cb-cli/dataplane/api-environment/model"
	"github.com/hortonworks/cb-cli/dataplane/cloud"
	"github.com/hortonworks/cb-cli/dataplane/types"
)

func (p *AzureProvider) GetCredentialRequest(stringFinder func(string) string, govCloud bool) (*model.CredentialV1Request, error) {
	parameters := &model.AzureCredentialV1RequestParameters{
		SubscriptionID: &(&types.S{S: stringFinder("subscription-id")}).S,
		TenantID:       &(&types.S{S: stringFinder("tenant-id")}).S,
		AppBased: &model.AppBasedV1Request{
			AccessKey: &(&types.S{S: stringFinder("app-id")}).S,
			SecretKey: (&types.S{S: stringFinder("app-password")}).S,
		},
	}
	credReq := cloud.CreateBaseCredentialRequest(stringFinder)

	if parameters.AppBased.SecretKey != "" {
		parameters.AppBased.AuthenticationType = model.AppBasedV1RequestAuthenticationTypeSECRET
	} else {
		parameters.AppBased.AuthenticationType = model.AppBasedV1RequestAuthenticationTypeCERTIFICATE
	}

	credReq.Azure = parameters
	return credReq, nil
}
