package cli

import (
	"encoding/json"
	"strings"
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

var BlueprintHeader []string = []string{"Cluster Type", "HDP Version"}

var BlueprintMap map[string]string

func init() {
	BlueprintMap = make(map[string]string)
	BlueprintMap["hdp25-data-science"] = "Data Science: Apache Spark 1.6, Apache Zeppelin 0.6.0"
	BlueprintMap["hdp25-etl-edw"] = "EDW-ETL: Apache Hive 1.2.1, Apache Spark 1.6"
	BlueprintMap["hdp25-etl-edw-spark2"] = "EDW-ETL: Apache Hive 1.2.1, Apache Spark 2.0"
	BlueprintMap["hdp25-edw-analytics"] = "EDW-Analytics: Apache Hive 2 LLAP, Apache Zeppelin 0.6.0"
	BlueprintMap["hdp25-etl-edw-shared"] = "Enterprise ETL-EDW: Apache Hive 1.2.1"

	BlueprintMap["hdp26-etl-edw"] = "EDW-ETL: Apache Hive 1.2.1, Apache Spark 1.6"
	BlueprintMap["hdp26-data-science"] = "Data Science: Apache Spark 1.6, Apache Zeppelin 0.7.0"
	BlueprintMap["hdp26-edw-analytics"] = "EDW-Analytics: Apache Hive 2 LLAP, Apache Zeppelin 0.7.0"
	BlueprintMap["hdp26-druid-bi"] = "BI: Druid 0.9.2 (Technical Preview)"
	BlueprintMap["hdp26-data-science-spark2"] = "Data Science: Apache Spark 2.1, Apache Zeppelin 0.7.0"
	BlueprintMap["hdp26-etl-edw-spark2"] = "EDW-ETL: Apache Hive 1.2.1, Apache Spark 2.1"

	BlueprintMap["hdp26-shared-services"] = "Enterprise Services: Apache Atlas, Apache Ranger"
	BlueprintMap["hdp26-data-science-spark2-shared"] = "Data Science: Apache Spark 2.1, Apache Zeppelin 0.7.0 Shared"
	BlueprintMap["hdp26-etl-edw-spark2-shared"] = "EDW-ETL: Apache Hive 1.2.1, Apache Spark 2.1 Shared"
	BlueprintMap["hdp26-etl-edw-shared"] = "EDW-ETL: Apache Hive 1.2.1, Apache Spark 1.6 Shared"
	BlueprintMap["hdp26-data-science-shared"] = "Data Science: Apache Spark 1.6, Apache Zeppelin 0.7.0 Shared"
	BlueprintMap["hdp26-druid-bi-shared"] = "BI: Druid 0.9.2 (Technical Preview) Shared"
}

type Blueprint struct {
	ClusterType string `json:"ClusterType" yaml:"ClusterType"`
	HDPVersion  string `json:"HDPVersion" yaml:"HDPVersion"`
}

func (b *Blueprint) DataAsStringArray() []string {
	return []string{b.ClusterType, b.HDPVersion}
}

func ListBlueprints(c *cli.Context) error {
	oAuth2Client := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))

	output := Output{Format: c.String(FlOutput.Name)}
	return listBlueprintsImpl(oAuth2Client.GetPublicBlueprints, output.WriteList)
}

func listBlueprintsImpl(getPublicBlueprints func() []*models_cloudbreak.BlueprintResponse, writer func([]string, []Row)) error {
	respBlueprints := getPublicBlueprints()

	var tableRows []Row
	for _, blueprint := range respBlueprints {
		// this is a workaround, needs to be hidden, by not storing them as public
		if !strings.HasPrefix(*blueprint.Name, "b") {
			ambariBlueprint := parseBlueprintJson(blueprint.AmbariBlueprint)
			row := &Blueprint{ClusterType: getFancyBlueprintName(blueprint), HDPVersion: ambariBlueprint.Blueprint.StackVersion}
			tableRows = append(tableRows, row)
		}
	}

	writer(BlueprintHeader, tableRows)

	return nil
}

func (c *Cloudbreak) GetPublicBlueprints() []*models_cloudbreak.BlueprintResponse {
	defer timeTrack(time.Now(), "get public blueprints")
	resp, err := c.Cloudbreak.Blueprints.GetPublicsBlueprint(blueprints.NewGetPublicsBlueprintParams())
	if err != nil {
		logErrorAndExit(err)
	}
	return resp.Payload
}

func CreateBlueprint(c *cli.Context) {
	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))

	fileLocation, urlLocation := c.String(FlBlueprintFileLocation.Name), c.String(FlBlueprintURL.Name)

	if len(fileLocation) == 0 && len(urlLocation) == 0 {
		logMissingParameterMessageAndExit(c, fmt.Sprintf("--%s or --%s is required", FlBlueprintFileLocation.Name, FlBlueprintURL.Name))
	} else if len(fileLocation) != 0 && len(urlLocation) != 0 {
		logMissingParameterMessageAndExit(c, fmt.Sprintf("only one of --%s and --%s can be set", FlBlueprintFileLocation.Name, FlBlueprintURL.Name))
	}
	var ambariBlueprint string
	if len(fileLocation) != 0 {
		log.Infof("[CreateBlueprint] read Ambari blueprint json from file: %s", fileLocation)
		bp, err := ioutil.ReadFile(fileLocation)
		if err != nil {
			logErrorAndExit(err)
		}
		ambariBlueprint = fmt.Sprintf("%s", bp)
	} else {
		var client http.Client
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
			ambariBlueprint = string(bodyBytes)
		} else {
			logErrorMessageAndExit("Couldn't download blueprint from URL, response code is not 200.")
		}
	}

	createBlueprintImpl(cbClient.Cloudbreak.Blueprints.PostPublicBlueprint, c.String(FlBlueprintName.Name), c.String(FlBlueprintDescription.Name), ambariBlueprint)
}

func createBlueprintImpl(
	postBlueprint func(*blueprints.PostPublicBlueprintParams) (*blueprints.PostPublicBlueprintOK, error),
	name string,
	description string,
	ambariBlueprint string) {

	defer timeTrack(time.Now(), "create blueprint")

	bpRequest := &models_cloudbreak.BlueprintRequest{
		Name:            &name,
		Description:     &description,
		AmbariBlueprint: ambariBlueprint,
		Inputs: 		 make([]*models_cloudbreak.BlueprintParameter, 0),
	}

	resp, err := postBlueprint(blueprints.NewPostPublicBlueprintParams().WithBody(bpRequest))
	if err != nil {
		logErrorAndExit(err)
	}
	log.Infof("[CreateBlueprint] blueprint created, id: %d", resp.Payload.ID)
}

func (c *Cloudbreak) DeleteBlueprint(name string) error {
	defer timeTrack(time.Now(), "delete blueprint")
	log.Infof("[DeleteBlueprint] delete blueprint: %s", name)
	return c.Cloudbreak.Blueprints.DeletePublicBlueprint(blueprints.NewDeletePublicBlueprintParams().WithName(name))
}


func getFancyBlueprintName(blueprint *models_cloudbreak.BlueprintResponse) string {
	var name string
	ambariBlueprint := parseBlueprintJson(blueprint.AmbariBlueprint)
	ambariBpName := *ambariBlueprint.Blueprint.Name
	if len(ambariBpName) > 0 {
		fancyName := BlueprintMap[ambariBpName]
		if len(fancyName) > 0 {
			name = fancyName
		}
	}
	if len(name) == 0 {
		name = *blueprint.Name
	}
	return name
}

func parseBlueprintJson(blueprintJson string) (ambariBlueprint AmbariBlueprint) {
	err := json.Unmarshal([]byte(blueprintJson), &ambariBlueprint)
	if err != nil {
		logErrorAndExit(err)
	}
	return
}
