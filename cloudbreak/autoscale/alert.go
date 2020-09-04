package autoscale

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hortonworks/cb-cli/cloudbreak/api-as/client/v1alerts"
	"github.com/hortonworks/cb-cli/cloudbreak/api-as/model"
	"github.com/hortonworks/cb-cli/cloudbreak/oauth"

	log "github.com/Sirupsen/logrus"
	fl "github.com/hortonworks/cb-cli/cloudbreak/flags"
	"github.com/hortonworks/cb-cli/utils"
	"github.com/urfave/cli"
)

var metricDefHeader = []string{"Name", "Description", "Enabled"}
var metricAlertHeader = []string{"ID", "Name", "AlertState", "Period", "AlertDefinition"}
var timeAlertHeader = []string{"ID", "Name", "Cron", "Timezone"}

type metricDefOut struct {
	Name        string `json:"Name" yaml:"Name"`
	Description string `json:"Description" yaml:"Description"`
	Enabled     string `json:"Enabled" yaml:"Enabled"`
}

type metricAlertOut struct {
	ID              string `json:"ID" yaml:"ID"`
	Name            string `json:"Name" yaml:"Name"`
	AlertState      string `json:"AlertState" yaml:"AlertState"`
	Period          string `json:"Period" yaml:"Period"`
	AlertDefinition string `json:"AlertDefinition" yaml:"AlertDefinition"`
}

type timeAlertOut struct {
	ID       string `json:"ID" yaml:"ID"`
	Name     string `json:"Name" yaml:"Name"`
	Cron     string `json:"Cron" yaml:"Cron"`
	Timezone string `json:"Timezone" yaml:"Timezone"`
}

func (r *metricDefOut) DataAsStringArray() []string {
	return []string{r.Name, r.Description, r.Enabled}
}

func (r *metricAlertOut) DataAsStringArray() []string {
	return []string{r.ID, r.Name, r.AlertState, r.Period, r.AlertDefinition}
}

func (r *timeAlertOut) DataAsStringArray() []string {
	return []string{r.ID, r.Name, r.Cron, r.Timezone}
}

type autoscaleAlertClient interface {
	CreateMetricAlerts(params *v1alerts.CreateMetricAlertsParams) (*v1alerts.CreateMetricAlertsOK, error)
	CreateTimeAlert(params *v1alerts.CreateTimeAlertParams) (*v1alerts.CreateTimeAlertOK, error)
	DeleteMetricAlarm(params *v1alerts.DeleteMetricAlarmParams) error
	DeleteTimeAlert(params *v1alerts.DeleteTimeAlertParams) error
	GetAlertDefinitions(params *v1alerts.GetAlertDefinitionsParams) (*v1alerts.GetAlertDefinitionsOK, error)
	GetMetricAlerts(params *v1alerts.GetMetricAlertsParams) (*v1alerts.GetMetricAlertsOK, error)
	GetTimeAlerts(params *v1alerts.GetTimeAlertsParams) (*v1alerts.GetTimeAlertsOK, error)
}

func ListMetricAlertDefinitions(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "list metric alert definitions")
	asClient := oauth.NewAutoscaleHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	clusterID, err := strconv.ParseInt(c.String(fl.FlAsID.Name), 10, 64)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	listMetricAlertDefinitionsImpl(asClient.Autoscale.V1alerts, output.WriteList, clusterID)
}

func ListMetricAlerts(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "list metric alerts")
	asClient := oauth.NewAutoscaleHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	clusterID, err := strconv.ParseInt(c.String(fl.FlAsID.Name), 10, 64)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	listMetricAlertsImpl(asClient.Autoscale.V1alerts, output.WriteList, clusterID)
}

func ListTimeAlerts(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "list time alerts")
	asClient := oauth.NewAutoscaleHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	clusterID, err := strconv.ParseInt(c.String(fl.FlAsID.Name), 10, 64)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	listTimeAlertsImpl(asClient.Autoscale.V1alerts, output.WriteList, clusterID)
}

func CreateMetricAlert(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "create metric alerts")
	asClient := oauth.NewAutoscaleHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	clusterID, err := strconv.ParseInt(c.String(fl.FlAsID.Name), 10, 64)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	name := c.String(fl.FlName.Name)
	definition := c.String(fl.FlAlertDefinition.Name)
	state := c.String(fl.FlAlertStateOptional.Name)
	period := c.String(fl.FlAlertPeriodOptional.Name)
	createMetricAlertImpl(asClient.Autoscale.V1alerts, output.Write, clusterID, name, definition, state, period)
}

func CreateTimeAlert(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "create time alerts")
	asClient := oauth.NewAutoscaleHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	clusterID, err := strconv.ParseInt(c.String(fl.FlAsID.Name), 10, 64)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	name := c.String(fl.FlName.Name)
	cron := c.String(fl.FlCron.Name)
	timezone := c.String(fl.FlTimezoneOptional.Name)
	createTimeAlertImpl(asClient.Autoscale.V1alerts, output.Write, clusterID, name, cron, timezone)
}

func DeleteMetricAlert(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "delete metric alert")
	asClient := oauth.NewAutoscaleHTTPClientFromContext(c)
	clusterID, err := strconv.ParseInt(c.String(fl.FlAsID.Name), 10, 64)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	alertID, err := strconv.ParseInt(c.String(fl.FlAlertID.Name), 10, 64)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	deleteMetricAlertImpl(asClient.Autoscale.V1alerts, clusterID, alertID)
}

func DeleteTimeAlert(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "delete time alert")
	asClient := oauth.NewAutoscaleHTTPClientFromContext(c)
	clusterID, err := strconv.ParseInt(c.String(fl.FlAsID.Name), 10, 64)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	alertID, err := strconv.ParseInt(c.String(fl.FlAlertID.Name), 10, 64)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	deleteTimeAlertImpl(asClient.Autoscale.V1alerts, clusterID, alertID)
}

func listMetricAlertDefinitionsImpl(client autoscaleAlertClient, writer func([]string, []utils.Row), clusterID int64) {
	log.Infof("[listMetricAlertDefinitionsImpl] sending autoscale list metric alert definition request")
	clusterResp, err := client.GetAlertDefinitions(v1alerts.NewGetAlertDefinitionsParams().WithClusterID(clusterID))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	tableRows := []utils.Row{}
	for _, cluster := range clusterResp.Payload {
		tableRows = append(tableRows, &metricDefOut{fmt.Sprintf("%v", cluster["name"]), fmt.Sprintf("%v", cluster["description"]), fmt.Sprintf("%v", cluster["enabled"])})
	}
	writer(metricDefHeader, tableRows)
}

func listMetricAlertsImpl(client autoscaleAlertClient, writer func([]string, []utils.Row), clusterID int64) {
	log.Infof("[listMetricAlertsImpl] sending autoscale list metric alerts request")
	clusterResp, err := client.GetMetricAlerts(v1alerts.NewGetMetricAlertsParams().WithClusterID(clusterID))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	tableRows := []utils.Row{}
	for _, alert := range clusterResp.Payload {
		tableRows = append(tableRows, &metricAlertOut{fmt.Sprintf("%v", alert.ID), alert.AlertName, alert.AlertState, fmt.Sprintf("%v", alert.Period), alert.AlertDefinition})
	}
	writer(metricAlertHeader, tableRows)
}

func listTimeAlertsImpl(client autoscaleAlertClient, writer func([]string, []utils.Row), clusterID int64) {
	log.Infof("[listTimeAlertsImpl] sending autoscale list time alerts request")
	clusterResp, err := client.GetTimeAlerts(v1alerts.NewGetTimeAlertsParams().WithClusterID(clusterID))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	tableRows := []utils.Row{}
	for _, alert := range clusterResp.Payload {
		tableRows = append(tableRows, &timeAlertOut{fmt.Sprintf("%v", alert.ID), alert.AlertName, alert.Cron, alert.TimeZone})
	}
	writer(timeAlertHeader, tableRows)
}

func createMetricAlertImpl(client autoscaleAlertClient, writer func([]string, utils.Row), clusterID int64, name, definition, state, period string) {
	log.Infof("[createMetricAlertImpl] sending create metric alert request")
	if len(state) == 0 {
		state = model.MetricAlertRequestAlertStateCRITICAL
	}
	if len(period) == 0 {
		period = "5"
	}
	periodInt, err := strconv.ParseInt(period, 10, 32)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	clusterResp, err := client.CreateMetricAlerts(v1alerts.NewCreateMetricAlertsParams().WithClusterID(clusterID).WithBody(&model.MetricAlertRequest{
		AlertName:       name,
		AlertDefinition: definition,
		AlertState:      state,
		Period:          int32(periodInt),
	}))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	alert := clusterResp.Payload
	tableRows := &metricAlertOut{fmt.Sprintf("%v", alert.ID), alert.AlertName, alert.AlertState, fmt.Sprintf("%v", alert.Period), alert.AlertDefinition}
	writer(metricAlertHeader, tableRows)
}

func createTimeAlertImpl(client autoscaleAlertClient, writer func([]string, utils.Row), clusterID int64, name, cron, timezone string) {
	log.Infof("[createTimeAlertImpl] sending create time alert request")
	if len(timezone) == 0 {
		timezone = "Etc/Greenwich"
	}
	clusterResp, err := client.CreateTimeAlert(v1alerts.NewCreateTimeAlertParams().WithClusterID(clusterID).WithBody(&model.TimeAlertRequest{
		AlertName: name,
		Cron:      cron,
		TimeZone:  timezone,
	}))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	alert := clusterResp.Payload
	tableRows := &timeAlertOut{fmt.Sprintf("%v", alert.ID), alert.AlertName, alert.Cron, alert.TimeZone}
	writer(timeAlertHeader, tableRows)
}

func deleteMetricAlertImpl(client autoscaleAlertClient, clusterID, alertID int64) {
	log.Infof("[deleteMetricAlertImpl] sending delete metric alert request")
	err := client.DeleteMetricAlarm(v1alerts.NewDeleteMetricAlarmParams().WithClusterID(clusterID).WithAlertID(alertID))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
}

func deleteTimeAlertImpl(client autoscaleAlertClient, clusterID, alertID int64) {
	log.Infof("[deleteTimeAlertImpl] sending delete time alert request")
	err := client.DeleteTimeAlert(v1alerts.NewDeleteTimeAlertParams().WithClusterID(clusterID).WithAlertID(alertID))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
}
