package cli

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hortonworks/cb-cli/cli/cloud"
	"github.com/hortonworks/cb-cli/cli/types"
	"github.com/hortonworks/cb-cli/cli/utils"
	"github.com/hortonworks/cb-cli/models_cloudbreak"
	"github.com/urfave/cli"
)

type node struct {
	name      string
	groupType string
	count     int32
}

var defaultNodes []node = []node{
	node{"master", models_cloudbreak.InstanceGroupResponseTypeGATEWAY, 1},
	node{"slave", models_cloudbreak.InstanceGroupResponseTypeCORE, 3},
}

var maxCardinality map[string]int = map[string]int{
	"1":   1,
	"0-1": 1,
	"1-2": 2,
	"0+":  9,
	"1+":  9,
	"ALL": 9,
}

var getBlueprintClient func(string, string, string) getPublicBlueprint = func(server, userName, password string) getPublicBlueprint {
	cbClient := NewCloudbreakOAuth2HTTPClient(server, userName, password)
	return cbClient.Cloudbreak.V1blueprints
}

func GenerateAwsStackTemplate(c *cli.Context) error {
	cloud.SetProviderType(cloud.AWS)
	return generateStackTemplateImpl(getNetworkMode(c), c.String, getBlueprintClient)
}

func GenerateAzureStackTemplate(c *cli.Context) error {
	cloud.SetProviderType(cloud.AZURE)
	return generateStackTemplateImpl(getNetworkMode(c), c.String, getBlueprintClient)
}

func GenerateGcpStackTemplate(c *cli.Context) error {
	cloud.SetProviderType(cloud.GCP)
	return generateStackTemplateImpl(getNetworkMode(c), c.String, getBlueprintClient)
}

func GenerateOpenstackStackTemplate(c *cli.Context) error {
	cloud.SetProviderType(cloud.OPENSTACK)
	return generateStackTemplateImpl(getNetworkMode(c), c.String, getBlueprintClient)
}

func getNetworkMode(c *cli.Context) cloud.NetworkMode {
	switch c.Command.Names()[0] {
	case "new-network":
		return cloud.NEW_NETWORK_NEW_SUBNET
	case "existing-network":
		return cloud.EXISTING_NETWORK_NEW_SUBNET
	case "existing-subnet":
		return cloud.EXISTING_NETWORK_EXISTING_SUBNET
	case "legacy-network":
		return cloud.LEGACY_NETWORK
	default:
		utils.LogErrorMessage("network mode not found")
		panic(2)
	}
}

func generateStackTemplateImpl(mode cloud.NetworkMode, stringFinder func(string) string, getBlueprintClient func(string, string, string) getPublicBlueprint) error {
	provider := cloud.GetProvider()

	template := models_cloudbreak.StackV2Request{
		ClusterRequest: &models_cloudbreak.ClusterV2Request{
			AmbariRequest: &models_cloudbreak.AmbariV2Request{
				BlueprintName: "____",
				UserName:      &(&types.S{S: "____"}).S,
				Password:      &(&types.S{S: ""}).S,
			},
		},
		Name:             &(&types.S{S: "____"}).S,
		CredentialName:   "____",
		AvailabilityZone: "____",
		Region:           "____",
		Orchestrator:     &models_cloudbreak.OrchestratorRequest{Type: &(&types.S{S: "SALT"}).S},
		InstanceGroups:   []*models_cloudbreak.InstanceGroupsV2{},
		Network: &models_cloudbreak.NetworkV2Request{
			Parameters: provider.GetNetworkParamatersTemplate(mode),
		},
		StackAuthentication: &models_cloudbreak.StackAuthentication{PublicKey: "____"},
	}

	nodes := defaultNodes
	if bpName := stringFinder(FlBlueprintName.Name); len(bpName) != 0 {
		bpResp := fetchBlueprint(bpName, getBlueprintClient(stringFinder(FlServer.Name), stringFinder(FlUsername.Name), stringFinder(FlPassword.Name)))
		bp, err := base64.StdEncoding.DecodeString(bpResp.AmbariBlueprint)
		if err != nil {
			utils.LogErrorAndExit(err)
		}
		nodes = getNodesByBlueprint(bp)
	} else if bpFile := stringFinder(FlBlueprintFile.Name); len(bpFile) != 0 {
		bp := utils.ReadFile(bpFile)
		nodes = getNodesByBlueprint(bp)
	}
	for _, n := range nodes {
		template.InstanceGroups = append(template.InstanceGroups, convertNodeToInstanceGroup(n))
	}

	if mode != cloud.EXISTING_NETWORK_EXISTING_SUBNET && mode != cloud.LEGACY_NETWORK {
		template.Network.SubnetCIDR = "10.0.0.0/16"
	}

	if params := provider.GetParamatersTemplate(); params != nil {
		template.Parameters = params
	}

	resp, err := json.MarshalIndent(template, "", "\t")
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	fmt.Printf("%s\n", string(resp))
	return nil
}

func getNodesByBlueprint(bp []byte) []node {
	var bpJson map[string]interface{}
	json.Unmarshal(bp, &bpJson)
	nodes := []*node{}
	if bpJson["host_groups"] == nil {
		utils.LogErrorMessageAndExit("host_groups not found in blueprint")
	}
	var gateway *node
	for _, e := range bpJson["host_groups"].([]interface{}) {
		var hg map[string]interface{} = e.(map[string]interface{})
		var count int = 1
		if hg["cardinality"] != nil {
			cardinality := hg["cardinality"].(string)
			count = maxCardinality[cardinality]
			if count == 0 {
				var err error
				count, err = strconv.Atoi(cardinality)
				if err != nil {
					utils.LogErrorMessageAndExit("Unable to parse as number: " + cardinality)
				}
			}
		}
		if hg["name"] == nil {
			utils.LogErrorMessageAndExit("host group name not found in blueprint")
		}
		node := node{hg["name"].(string), models_cloudbreak.InstanceGroupResponseTypeCORE, int32(count)}
		nodes = append(nodes, &node)
		if gateway == nil || gateway.count > node.count {
			gateway = &node
		}
	}
	gateway.groupType = models_cloudbreak.InstanceGroupResponseTypeGATEWAY
	resp := []node{}
	for _, n := range nodes {
		resp = append(resp, *n)
	}
	return resp
}

func convertNodeToHostGroup(node node) *models_cloudbreak.HostGroupRequest {
	return &models_cloudbreak.HostGroupRequest{
		Name:         &node.name,
		RecoveryMode: "AUTO",
		Constraint: &models_cloudbreak.Constraint{
			InstanceGroupName: node.name,
			HostCount:         &node.count,
		},
	}
}

func convertNodeToInstanceGroup(node node) *models_cloudbreak.InstanceGroupsV2 {
	ig := &models_cloudbreak.InstanceGroupsV2{
		Template: &models_cloudbreak.TemplateV2Request{
			InstanceType: &(&types.S{S: "____"}).S,
			VolumeType:   "____",
			VolumeCount:  1,
			VolumeSize:   10,
		},
		Group:     &node.name,
		NodeCount: &node.count,
		Type:      node.groupType,
		SecurityGroup: &models_cloudbreak.SecurityGroupV2Request{
			SecurityRules: []*models_cloudbreak.SecurityRuleRequest{
				getDefaultSecurityRule("22"),
			},
		},
	}
	if node.groupType == models_cloudbreak.InstanceGroupResponseTypeGATEWAY {
		ig.SecurityGroup.SecurityRules = append(ig.SecurityGroup.SecurityRules, getDefaultSecurityRule("443"))
		ig.SecurityGroup.SecurityRules = append(ig.SecurityGroup.SecurityRules, getDefaultSecurityRule("9443"))
	}
	return ig
}

func getDefaultSecurityRule(port string) *models_cloudbreak.SecurityRuleRequest {
	return &models_cloudbreak.SecurityRuleRequest{
		Subnet:   &(&types.S{S: "0.0.0.0/0"}).S,
		Protocol: &(&types.S{S: "tcp"}).S,
		Ports:    &port,
	}
}
