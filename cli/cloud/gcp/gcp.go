package gcp

import "github.com/hortonworks/hdc-cli/cli/cloud"

var name string

func init() {
	name = string(cloud.GCP)
	cloud.CloudProviders[cloud.GCP] = new(GcpProvider)
}

type GcpProvider struct {
}

func (p *GcpProvider) GetName() *string {
	return &name
}