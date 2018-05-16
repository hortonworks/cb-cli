package azure

import (
	"github.com/hortonworks/cb-cli/cli/cloud"
	"github.com/hortonworks/cb-cli/cli/utils"
	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

func (p *AzureProvider) GetNetworkParamatersTemplate(mode cloud.NetworkMode) map[string]interface{} {
	switch mode {
	case cloud.EXISTING_NETWORK_EXISTING_SUBNET:
		return map[string]interface{}{"resourceGroupName": "____", "networkId": "____", "subnetId": "____", "noPublicIp": false, "noFirewallRules": false}
	default:
		return nil
	}
}

func (p *AzureProvider) GetParamatersTemplate() map[string]interface{} {
	return map[string]interface{}{"encryptStorage": "false"}
}

func (p *AzureProvider) GetInstanceGroupParamatersTemplate(node cloud.Node) map[string]interface{} {
	if node.Count > 1 {
		return map[string]interface{}{
			"availabilitySet": map[string]interface{}{
				"faultDomainCount":  2,
				"updateDomainCount": 20,
				"name":              node.Name + "-" + utils.RandStr(20) + "-as",
			},
		}
	}
	return nil
}

func (p *AzureProvider) GenerateNetworkRequestFromNetworkResponse(response *models_cloudbreak.NetworkResponse) *models_cloudbreak.NetworkV2Request {
	request := &models_cloudbreak.NetworkV2Request{
		Parameters: utils.CopyToByTargets(response.Parameters, "resourceGroupName", "networkId", "subnetId", "noPublicIp", "noFirewallRules"),
	}
	return request
}
