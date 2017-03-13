package cli

import (
	"strconv"
	"strings"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/blueprints"
	"github.com/hortonworks/hdc-cli/models_cloudbreak"
	"github.com/urfave/cli"
)

var BlueprintHeader []string = []string{"Cluster Type", "HDP Version"}

var BlueprintMap map[string]string

func init() {
	BlueprintMap = make(map[string]string)
	BlueprintMap["hdp-data-science"] = "Data Science: Apache Spark 1.6, Apache Zeppelin 0.6.0"
	BlueprintMap["hdp26-data-science"] = "Data Science: Apache Spark 1.6, Apache Zeppelin 0.7.0"
	BlueprintMap["hdp-etl-edw"] = "EDW-ETL: Apache Hive 1.2.1, Apache Spark 1.6"
	BlueprintMap["hdp-etl-edw-spark2"] = "EDW-ETL: Apache Hive 1.2.1, Apache Spark 2.0"
	BlueprintMap["hdp-edw-analytics"] = "EDW-Analytics: Apache Hive 2 LLAP, Apache Zeppelin 0.6.0"
	BlueprintMap["hdp26-edw-analytics"] = "EDW-Analytics: Apache Hive 2 LLAP, Apache Zeppelin 0.7.0"
	BlueprintMap["shared-services"] = "Enterprise Services: Apache Atlas, Apache Ranger"
	BlueprintMap["hdp25-etl-edw-shared"] = "Enterprise ETL-EDW: Apache Hive 1.2.1"
	BlueprintMap["hdp-etl-edw-tp"] = "EDW-ETL: Apache Hive 1.2.1, Apache Spark 2.0"
	BlueprintMap["hdp26-etl-edw-spark2"] = "EDW-ETL: Apache Hive 1.2.1, Apache Spark 2.1"
	BlueprintMap["hdp-druid-bi"] = "BI: Druid 0.9.2 (Technical Preview)"
	BlueprintMap["hdp-data-science-spark2"] = "Data Science: Apache Spark 2.1, Apache Zeppelin 0.7.0"
}

func getDefaultClusterType() string {
	for _, v := range BlueprintMap {
		if !strings.Contains(v, "Enterprise") && !strings.Contains(v, "Druid") && !strings.Contains(v, "Spark 2.1") {
			return v
		}
	}
	return ""
}

type Blueprint struct {
	ClusterType string `json:"ClusterType" yaml:"ClusterType"`
	HDPVersion  string `json:"HDPVersion" yaml:"HDPVersion"`
}

func (b *Blueprint) DataAsStringArray() []string {
	return []string{b.ClusterType, b.HDPVersion}
}

func ListBlueprints(c *cli.Context) error {
	oAuth2Client := NewOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))

	output := Output{Format: c.String(FlOutput.Name)}
	return listBlueprintsImpl(oAuth2Client.GetPublicBlueprints, output.WriteList)
}

func listBlueprintsImpl(getPublicBlueprints func() []*models_cloudbreak.BlueprintResponse, writer func([]string, []Row)) error {
	respBlueprints := getPublicBlueprints()

	var tableRows []Row
	for _, blueprint := range respBlueprints {
		// this is a workaround, needs to be hidden, by not storing them as public
		if !strings.HasPrefix(blueprint.Name, "b") {
			row := &Blueprint{ClusterType: getFancyBlueprintName(blueprint), HDPVersion: blueprint.AmbariBlueprint.Blueprint.StackVersion}
			tableRows = append(tableRows, row)
		}
	}

	writer(BlueprintHeader, tableRows)

	return nil
}

func (c *Cloudbreak) GetPublicBlueprints() []*models_cloudbreak.BlueprintResponse {
	defer timeTrack(time.Now(), "get public blueprints")
	resp, err := c.Cloudbreak.Blueprints.GetPublics(&blueprints.GetPublicsParams{})
	if err != nil {
		logErrorAndExit(err)
	}
	return resp.Payload
}

func (c *Cloudbreak) DeleteBlueprint(name string) error {
	defer timeTrack(time.Now(), "delete blueprint")
	log.Infof("[DeleteBlueprint] delete blueprint: %s", name)
	return c.Cloudbreak.Blueprints.DeletePublic(&blueprints.DeletePublicParams{Name: name})
}

func (c *Cloudbreak) GetBlueprintByName(name string) *models_cloudbreak.BlueprintResponse {
	defer timeTrack(time.Now(), "get blueprint by name")
	log.Infof("[GetBlueprintByName] get blueprint by name: %s", name)

	resp, err := c.Cloudbreak.Blueprints.GetPublic(&blueprints.GetPublicParams{Name: getRealBlueprintName(name)})
	if err != nil {
		logErrorAndExit(err)
	}

	id64 := *resp.Payload.ID
	log.Infof("[GetBlueprintByName] '%s' blueprint id: %d", name, id64)
	return resp.Payload
}

func (c *Cloudbreak) CreateBlueprint(skeleton ClusterSkeleton, blueprint *models_cloudbreak.BlueprintResponse, channel chan int64, wg *sync.WaitGroup) {
	defer wg.Done()
	createBlueprintImpl(skeleton, blueprint, channel, c.Cloudbreak.Blueprints.PostPublic)
}

func createBlueprintImpl(skeleton ClusterSkeleton, blueprint *models_cloudbreak.BlueprintResponse, channel chan int64, postPublicBlueprint func(*blueprints.PostPublicParams) (*blueprints.PostPublicOK, error)) {
	if len(skeleton.Configurations) == 0 {
		log.Info("[CreateBlueprint] there are no custom configurations, use the default blueprint")
		channel <- *blueprint.ID
		return
	}

	defer timeTrack(time.Now(), "create blueprint")

	bpRequest := createBlueprintRequest(skeleton, blueprint)

	resp, err := postPublicBlueprint(&blueprints.PostPublicParams{Body: bpRequest})

	if err != nil {
		logErrorAndExit(err)
	}

	log.Infof("[CreateBlueprint] blueprint created, id: %d", resp.Payload.ID)
	channel <- *resp.Payload.ID
}

func createBlueprintRequest(skeleton ClusterSkeleton, blueprint *models_cloudbreak.BlueprintResponse) *models_cloudbreak.BlueprintRequest {
	blueprintName := "b" + strconv.FormatInt(time.Now().UnixNano(), 10)

	bpRequest := models_cloudbreak.BlueprintRequest{
		Name:            blueprintName,
		AmbariBlueprint: blueprint.AmbariBlueprint,
		Properties:      skeleton.Configurations,
	}

	return &bpRequest
}

func getFancyBlueprintName(blueprint *models_cloudbreak.BlueprintResponse) string {
	var name string
	ambariBpName := *blueprint.AmbariBlueprint.Blueprint.Name
	if len(ambariBpName) > 0 {
		fancyName := BlueprintMap[ambariBpName]
		if len(fancyName) > 0 {
			name = fancyName
		}
	}
	if len(name) == 0 {
		name = blueprint.Name
	}
	return name
}

func getRealBlueprintName(name string) string {
	realName := BlueprintMap[name]
	if len(realName) > 0 {
		return realName
	}
	return name
}
