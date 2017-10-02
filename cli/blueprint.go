package cli

import (
	"encoding/json"
	"time"

	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/hdc-cli/cli/utils"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/blueprints"
	"github.com/hortonworks/hdc-cli/models_cloudbreak"
	"github.com/urfave/cli"
)

type Blueprints struct {
	Name         *string `json:"blueprint_name"`
	StackName    string  `json:"stack_name"`
	StackVersion string  `json:"stack_version"`
}

type BlueprintConfigurations map[string]map[string]interface{}

type BlueprintSettings map[string][]map[string]interface{}

type Component struct {
	Name string `json:"name"`
}

type BlueprintHostGroup struct {
	Name           string           `json:"name"`
	Configurations []Configurations `json:"configurations,omitempty"`
	Components     []Component      `json:"components"`
	Cardinality    string           `json:"cardinality"`
}

type AmbariBlueprint struct {
	Blueprint      Blueprints                `json:"Blueprints"`
	Configurations []BlueprintConfigurations `json:"configurations,omitempty"`
	Settings       []BlueprintSettings       `json:"settings,omitempty"`
	HostGroups     []BlueprintHostGroup      `json:"host_groups"`
}

var BlueprintHeader []string = []string{"Blueprint Name", "Description", "HDP Version", "Hostgroup Count", "Tags"}

type Blueprint struct {
	Name           string `json:"BlueprintName" yaml:"BlueprintName"`
	Description    string `json:"Description" yaml:"Description"`
	HDPVersion     string `json:"HDPVersion" yaml:"HDPVersion"`
	HostgroupCount string `json:"HostgroupCount" yaml:"HostgroupCount"`
	Tags           string `json:"Tags" yaml:"Tags"`
}

func (b *Blueprint) DataAsStringArray() []string {
	return []string{b.Name, b.Description, b.HDPVersion, b.HostgroupCount, b.Tags}
}

type BlueprintsClient interface {
	PostPublicBlueprint(*blueprints.PostPublicBlueprintParams) (*blueprints.PostPublicBlueprintOK, error)
	PostPrivateBlueprint(*blueprints.PostPrivateBlueprintParams) (*blueprints.PostPrivateBlueprintOK, error)
	GetPrivatesBlueprint(*blueprints.GetPrivatesBlueprintParams) (*blueprints.GetPrivatesBlueprintOK, error)
	DeletePrivateBlueprint(*blueprints.DeletePrivateBlueprintParams) error
}

func ListBlueprints(c *cli.Context) {
	checkRequiredFlags(c)
	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	output := utils.Output{Format: c.String(FlOutput.Name)}
	listBlueprintsImpl(cbClient.Cloudbreak.Blueprints, output.WriteList)
}

func listBlueprintsImpl(client BlueprintsClient, writer func([]string, []utils.Row)) {
	defer utils.TimeTrack(time.Now(), "get public blueprints")
	resp, err := client.GetPrivatesBlueprint(blueprints.NewGetPrivatesBlueprintParams())
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	var tableRows []utils.Row
	for _, blueprint := range resp.Payload {
		ambariBlueprint := parseBlueprintJson(blueprint.AmbariBlueprint)
		row := &Blueprint{
			Name:           *blueprint.Name,
			Description:    *blueprint.Description,
			HDPVersion:     ambariBlueprint.Blueprint.StackVersion,
			HostgroupCount: fmt.Sprint(blueprint.HostGroupCount),
			Tags:           blueprint.Status,
		}
		tableRows = append(tableRows, row)
	}

	writer(BlueprintHeader, tableRows)
}

func parseBlueprintJson(blueprintJson string) (ambariBlueprint AmbariBlueprint) {
	err := json.Unmarshal([]byte(blueprintJson), &ambariBlueprint)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	return
}

func CreateBlueprintFromUrl(c *cli.Context) {
	checkRequiredFlags(c)
	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	urlLocation := c.String(FlBlueprintURL.Name)
	createBlueprintImpl(
		cbClient.Cloudbreak.Blueprints,
		c.String(FlName.Name),
		c.String(FlDescription.Name),
		c.Bool(FlPublic.Name),
		readBlueprintFromURL(urlLocation))
}

func CreateBlueprintFromFile(c *cli.Context) {
	checkRequiredFlags(c)
	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	fileLocation := c.String(FlBlueprintFileLocation.Name)
	createBlueprintImpl(
		cbClient.Cloudbreak.Blueprints,
		c.String(FlName.Name),
		c.String(FlDescription.Name),
		c.Bool(FlPublic.Name),
		readBlueprintFromFile(fileLocation))
}

func readBlueprintFromFile(fileLocation string) string {
	log.Infof("[CreateBlueprint] read Ambari blueprint json from file: %s", fileLocation)
	bp := utils.ReadFile(fileLocation)
	return string(bp)
}

func readBlueprintFromURL(urlLocation string) string {
	log.Infof("[CreateBlueprint] read Ambari blueprint json from url: %s", urlLocation)
	bp := utils.ReadFromUrl(urlLocation)
	return string(bp)
}

func createBlueprintImpl(client BlueprintsClient, name string, description string, public bool, ambariBlueprint string) {
	defer utils.TimeTrack(time.Now(), "create blueprint")
	bpRequest := &models_cloudbreak.BlueprintRequest{
		Name:            &name,
		Description:     &description,
		AmbariBlueprint: ambariBlueprint,
		Inputs:          make([]*models_cloudbreak.BlueprintParameter, 0),
	}
	if public {
		resp, err := client.PostPublicBlueprint(blueprints.NewPostPublicBlueprintParams().WithBody(bpRequest))
		if err != nil {
			utils.LogErrorAndExit(err)
		}
		log.Infof("[CreateBlueprint] public blueprint created: %s (id: %d)", resp.Payload.Name, resp.Payload.ID)
	} else {
		resp, err := client.PostPrivateBlueprint(blueprints.NewPostPrivateBlueprintParams().WithBody(bpRequest))
		if err != nil {
			utils.LogErrorAndExit(err)
		}
		log.Infof("[CreateBlueprint] private blueprint created: %s (id: %d)", resp.Payload.Name, resp.Payload.ID)
	}
}

func DeleteBlueprint(c *cli.Context) {
	checkRequiredFlags(c)
	defer utils.TimeTrack(time.Now(), "delete blueprint")

	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	name := c.String(FlName.Name)
	err := cbClient.Cloudbreak.Blueprints.DeletePrivateBlueprint(blueprints.NewDeletePrivateBlueprintParams().WithName(name))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[DeleteBlueprint] blueprint deleted, name: %s", name)
}
