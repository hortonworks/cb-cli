package openstack

import (
	envmodel "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
	"github.com/hortonworks/cb-cli/dataplane/api/model"
	"github.com/hortonworks/cb-cli/dataplane/cloud"
)

func (p *OpenstackProvider) GenerateDefaultNetwork(mode cloud.NetworkMode) *model.NetworkV4Request {
	return nil
}

func (p *OpenstackProvider) GenerateDefaultNetworkWithParams(getFlags func(string) string, mode cloud.NetworkMode) *envmodel.EnvironmentNetworkV1Request {
	return &envmodel.EnvironmentNetworkV1Request{}
}

func (p *OpenstackProvider) SetParametersTemplate(request *model.StackV4Request) {
}

func (p *OpenstackProvider) SetInstanceGroupParametersTemplate(request *model.InstanceGroupV4Request, node cloud.Node) {
}

func (p *OpenstackProvider) GenerateNetworkRequestFromNetworkResponse(response *model.NetworkV4Response) *model.NetworkV4Request {
	return nil
}
