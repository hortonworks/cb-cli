package yarn

import "github.com/hortonworks/cb-cli/cloudbreak/cloud"

func (p *YarnProvider) GetNetworkParamatersTemplate(mode cloud.NetworkMode) map[string]interface{} {
	return nil
}

func (p *YarnProvider) GetParamatersTemplate() map[string]interface{} {
	return map[string]interface{}{"yarnQueue": "default-developers"}
}

func (p *YarnProvider) GetInstanceGroupParamatersTemplate(node cloud.Node) map[string]interface{} {
	return nil
}
