package telemetry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/hortonworks/cb-cli/dataplane/api-environment/client/v1telemetry"
	"github.com/hortonworks/cb-cli/dataplane/api-environment/model"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/oauth"
	"github.com/hortonworks/dp-cli-common/utils"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var TelemetryHeader = []string{"Features", "Rules"}

type accountTelemetry struct {
	Features features            `json:"Features" yaml:"Features"`
	Rules    []anonymizationRule `json:"Rules" yaml:"Rules"`
}

type anonymizationRule struct {
	Replacement string `json:"Replacement" yaml:"Replacement"`
	Value       string `json:"Value" yaml:"Value"`
}

type features struct {
	ClusterLogsCollection *featureSetting `json:"ClusterLogsCollection,omitempty" yaml:"ClusterLogsCollection"`
	WorkloadAnalytics     *featureSetting `json:"WorkloadAnalytics,omitempty" yaml:"WorkloadAnalytics"`
}

type featureSetting struct {
	Enabled bool `json:"Enabled" yaml:"Enabled"`
}

func UpdateAccountTelemetry(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "update telemetry")
	path := c.String(fl.FlFile.Name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		utils.LogErrorAndExit(err)
	}
	envClient := oauth.NewEnvironmentClientFromContext(c)
	telemetryClient := envClient.Environment.V1telemetry
	content := utils.ReadFile(path)
	var request model.AccountTelemetryRequest
	err := json.Unmarshal(content, &request)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	resp, err := telemetryClient.UpdateAccountTelemetryV1(v1telemetry.NewUpdateAccountTelemetryV1Params().WithBody(&request))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	accountTelemetryResponse := resp.Payload
	if accountTelemetryResponse != nil {
		log.Info("[updateAccountTelemetry] Account level telemetry updated")
	}
}

func DescribeAccountTelemetry(c *cli.Context) error {
	defer utils.TimeTrack(time.Now(), "describe telemetry")
	envClient := oauth.NewEnvironmentClientFromContext(c)
	telemetryClient := envClient.Environment.V1telemetry

	resp, err := telemetryClient.GetAccountTelemetryV1(v1telemetry.NewGetAccountTelemetryV1Params())
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	telemetry := accountTelemetry{Features: *getFeatures(resp.Payload.Features), Rules: *getRules(resp.Payload.Rules)}
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(telemetry); err != nil {
		utils.LogErrorAndExit(err)
	}
	fmt.Println(buf.String())
	return nil
}

func getFeatures(featuresResp *model.FeaturesResponse) *features {
	features := features{}
	if featuresResp.ClusterLogsCollection != nil {
		clusterLogsCollection := featureSetting{Enabled: *featuresResp.ClusterLogsCollection.Enabled}
		features.ClusterLogsCollection = &clusterLogsCollection
	}
	if featuresResp.WorkloadAnalytics != nil {
		workloadAnalytics := featureSetting{Enabled: *featuresResp.WorkloadAnalytics.Enabled}
		features.WorkloadAnalytics = &workloadAnalytics
	}
	return &features
}

func getRules(anonymizationRules []*model.AnonymizationRule) *[]anonymizationRule {
	rules := make([]anonymizationRule, 0)
	for _, v := range anonymizationRules {
		rules = append(rules, anonymizationRule{Value: *v.Value, Replacement: v.Replacement})
	}
	return &rules
}
