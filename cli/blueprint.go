package cli

import (
	"encoding/json"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/blueprints"
	"github.com/hortonworks/hdc-cli/models_cloudbreak"
	"github.com/urfave/cli"
	"fmt"
	"io/ioutil"
	"net/http"
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
	PostPrivateBlueprint(params *blueprints.PostPrivateBlueprintParams) (*blueprints.PostPrivateBlueprintOK, error)
	GetPrivatesBlueprint(params *blueprints.GetPrivatesBlueprintParams) (*blueprints.GetPrivatesBlueprintOK, error)
	DeletePrivateBlueprint(params *blueprints.DeletePrivateBlueprintParams) error
}

func ListBlueprints(c *cli.Context) {
	checkRequiredFlags(c)
	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	output := Output{Format: c.String(FlOutput.Name)}
	listBlueprintsImpl(cbClient.Cloudbreak.Blueprints, output.WriteList)
}

func listBlueprintsImpl(client BlueprintsClient, writer func([]string, []Row)) {
	defer timeTrack(time.Now(), "get public blueprints")
	resp, err := client.GetPrivatesBlueprint(blueprints.NewGetPrivatesBlueprintParams())
	if err != nil {
		logErrorAndExit(err)
	}
	var tableRows []Row
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
		logErrorAndExit(err)
	}
	return
}

func CreateBlueprintFromUrl(c *cli.Context) {
	checkRequiredFlags(c)
	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	urlLocation := c.String(FlBlueprintURL.Name)
	createBlueprintImpl(
		cbClient.Cloudbreak.Blueprints,
		c.String(FlBlueprintName.Name),
		c.String(FlBlueprintDescription.Name),
		readBlueprintFromURL(urlLocation, new(http.Client)))
}

func CreateBlueprintFromFile(c *cli.Context) {
	checkRequiredFlags(c)
	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	fileLocation := c.String(FlBlueprintFileLocation.Name)
	createBlueprintImpl(
		cbClient.Cloudbreak.Blueprints,
		c.String(FlBlueprintName.Name),
		c.String(FlBlueprintDescription.Name),
		readBlueprintFromFile(fileLocation))
}

func readBlueprintFromFile(fileLocation string) string {
	log.Infof("[CreateBlueprint] read Ambari blueprint json from file: %s", fileLocation)
	bp, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		logErrorAndExit(err)
	}
	return fmt.Sprintf("%s", bp)
}

func readBlueprintFromURL(urlLocation string, client *http.Client) string {
	resp, httperr := client.Get(urlLocation)
	if httperr != nil {
		logErrorAndExit(httperr)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		bodyBytes, ioerr := ioutil.ReadAll(resp.Body)
		if ioerr != nil {
			logErrorAndExit(ioerr)
		}
		return string(bodyBytes)
	} else {
		logErrorMessageAndExit("Couldn't download blueprint from URL, response code is not 200.")
	}
	return ""
}

func createBlueprintImpl(client BlueprintsClient, name string, description string, ambariBlueprint string) {
	defer timeTrack(time.Now(), "create blueprint")
	bpRequest := &models_cloudbreak.BlueprintRequest{
		Name:            &name,
		Description:     &description,
		AmbariBlueprint: ambariBlueprint,
		Inputs:          make([]*models_cloudbreak.BlueprintParameter, 0),
	}
	resp, err := client.PostPrivateBlueprint(blueprints.NewPostPrivateBlueprintParams().WithBody(bpRequest))
	if err != nil {
		logErrorAndExit(err)
	}
	log.Infof("[CreateBlueprint] blueprint created: %s (id: %d)", resp.Payload.Name, resp.Payload.ID)
}

func DeleteBlueprint(c *cli.Context) {
	checkRequiredFlags(c)
	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	deleteBlueprintsImpl(cbClient.Cloudbreak.Blueprints, c.String(FlBlueprintName.Name))
}

func deleteBlueprintsImpl(client BlueprintsClient, name string) {
	defer timeTrack(time.Now(), "delete blueprint")
	err := client.DeletePrivateBlueprint(blueprints.NewDeletePrivateBlueprintParams().WithName(name))
	if err != nil {
		logErrorAndExit(err)
	}
	log.Infof("[DeleteBlueprint] blueprint deleted, name: %s", name)
}
