package cli

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/hdc-cli/client_autoscale/alerts"
	"github.com/hortonworks/hdc-cli/client_autoscale/clusters"
	"github.com/hortonworks/hdc-cli/client_autoscale/configurations"
	"github.com/hortonworks/hdc-cli/client_autoscale/policies"
	"github.com/hortonworks/hdc-cli/models_autoscale"
	"github.com/hortonworks/hdc-cli/models_cloudbreak"
	"github.com/urfave/cli"
	"time"
)

var AlertListHeader = []string{"Name", "Label"}

type AlertListElement struct {
	Name  string `json:"Name" yaml:"Name"`
	Label string `json:"Label" yaml:"Label"`
}

func (c *AlertListElement) DataAsStringArray() []string {
	return []string{c.Name, c.Label}
}

func ListDefinitions(c *cli.Context) error {
	defer timeTrack(time.Now(), "list scaling definitions")

	cbClient, asClient := NewOAuth2HTTPClients(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	output := Output{Format: c.String(FlOutput.Name)}

	clusterName := c.String(FlClusterName.Name)
	if len(clusterName) == 0 {
		logMissingParameterAndExit(c, []string{FlClusterName.Name})
	}

	return listDefinitionsImpl(clusterName, asClient.getAlertDefinitions, cbClient.GetClusterByName, output.WriteList)
}

func listDefinitionsImpl(clusterName string,
	getAlertDefinitions func(string, func(name string) *models_cloudbreak.StackResponse) []*models_autoscale.AlertRuleDefinitionEntry,
	getCluster func(name string) *models_cloudbreak.StackResponse,
	writer func([]string, []Row)) error {

	alertDefinitions := getAlertDefinitions(clusterName, getCluster)

	tableRows := make([]Row, len(alertDefinitions))
	for i, a := range alertDefinitions {
		tableRows[i] = &AlertListElement{
			Name:  SafeStringConvert(a.Name),
			Label: SafeStringConvert(a.Label),
		}
	}
	writer(AlertListHeader, tableRows)

	return nil
}

func (autosScaling *Autoscaling) getAlertDefinitions(clusterName string,
	getCluster func(name string) *models_cloudbreak.StackResponse) []*models_autoscale.AlertRuleDefinitionEntry {

	log.Infof("[getAlertDefinitions] get cluster by name: %s", clusterName)
	clusterId := getCluster(clusterName).ID
	log.Infof("[getAlertDefinitions] found cluster, id: %d", *clusterId)

	alertDefinitions, err := autosScaling.AutoScaling.Alerts.GetPrometheusDefinitions(&alerts.GetPrometheusDefinitionsParams{ClusterID: *clusterId})
	if err != nil {
		logErrorAndExit(err)
	}

	return alertDefinitions.Payload
}

func (autosScaling *Autoscaling) createBaseAutoscalingCluster(stackId int64) int64 {
	defer timeTrack(time.Now(), "create base autoscaling cluster")

	log.Infof("[createBaseAutoscalingCluster] add cluster to autoscaling, id: %d", stackId)
	state := "PENDING"
	resp, err := autosScaling.AutoScaling.Clusters.AddCluster(
		&clusters.AddClusterParams{
			Body: &models_autoscale.AmbariConnectionDetails{
				ClusterState: &state,
				StackID:      &stackId,
			}})

	if err != nil {
		logErrorAndExit(err)
	}

	log.Infof("[createBaseAutoscalingCluster] base autoscale cluster created, id: %d", *resp.Payload.ID)
	return *resp.Payload.ID
}

func (autosScaling *Autoscaling) setConfiguration(asClusterId int64, config AutoscalingConfiguration) {
	defer timeTrack(time.Now(), "set autoscaling configuration")

	log.Infof("[setConfiguration] set autoscaling cluster configuration for cluster: %d, config: %v", asClusterId, config)
	_, err := autosScaling.AutoScaling.Configurations.SetScalingConfiguration(
		&configurations.SetScalingConfigurationParams{
			Body: &models_autoscale.ScalingConfiguration{
				Cooldown: &config.CooldownTime,
				MinSize:  &config.ClusterMinSize,
				MaxSize:  &config.ClusterMaxSize,
			},
			ClusterID: asClusterId,
		})

	if err != nil {
		logErrorAndExit(err)
	}
}

func (autosScaling *Autoscaling) addPrometheusAlert(asClusterId int64, policy AutoscalingPolicy) int64 {
	defer timeTrack(time.Now(), "add prometheus alert")

	log.Infof("[addPrometheusAlert] add prometheus alert to cluster: %d, name: %v", asClusterId, policy)
	operatorString := policy.getOperatorName()
	resp, err := autosScaling.AutoScaling.Alerts.CreatePrometheusAlert(
		&alerts.CreatePrometheusAlertParams{
			Body: &models_autoscale.PromhetheusAlert{
				AlertName:     &policy.Name,
				AlertOperator: &operatorString,
				AlertRuleName: &policy.ScalingDefinition,
				Period:        &policy.Period,
				Threshold:     &policy.Threshold,
			},
			ClusterID: asClusterId,
		})
	if err != nil {
		logErrorAndExit(err)
	}

	log.Infof("[addPrometheusAlert] prometheus alert added, id: %d", *resp.Payload.ID)
	return *resp.Payload.ID
}

func (autosScaling *Autoscaling) addScalingPolicy(asClusterId int64, policy AutoscalingPolicy, alertId int64) int64 {
	defer timeTrack(time.Now(), "add scaling policy")

	log.Infof("[addPolicy] add scaling policy: %v to alert id: %d", policy, alertId)

	adjustmentType := "NODE_COUNT"
	resp, err := autosScaling.AutoScaling.Policies.AddScaling(
		&policies.AddScalingParams{
			Body: &models_autoscale.ScalingPolicy{
				AdjustmentType:    &adjustmentType,
				AlertID:           &alertId,
				HostGroup:         &policy.NodeType,
				Name:              &policy.Name,
				ScalingAdjustment: &policy.ScalingAdjustment,
			},
			ClusterID: asClusterId,
		})
	if err != nil {
		logErrorAndExit(err)
	}

	log.Infof("[addScalingPolicy] scaling policy added, id: %d", *resp.Payload.ID)
	return *resp.Payload.ID
}

func (p *AutoscalingPolicy) getOperatorName() string {
	switch p.Operator {
	case ">":
		return "MORE_THAN"
	case "<":
		return "LESS_THAN"
	default:
		logErrorAndExit(errors.New("Unrecognised operator: " + p.Operator))
		return ""
	}
}
