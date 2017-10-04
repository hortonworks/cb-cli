package openstack

import "github.com/hortonworks/hdc-cli/cli/cloud"

var name string

func init() {
	name = string(cloud.OPENSTACK)
	cloud.CloudProviders[cloud.OPENSTACK] = new(OpenstackProvider)
}

type OpenstackProvider struct {
}

func (p *OpenstackProvider) GetName() *string {
	return &name
}
