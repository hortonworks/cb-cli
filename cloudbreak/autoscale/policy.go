package autoscale

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hortonworks/cb-cli/cloudbreak/api-as/client/v1policies"
	"github.com/hortonworks/cb-cli/cloudbreak/api-as/model"
	"github.com/hortonworks/cb-cli/cloudbreak/oauth"

	log "github.com/Sirupsen/logrus"
	fl "github.com/hortonworks/cb-cli/cloudbreak/flags"
	"github.com/hortonworks/cb-cli/utils"
	"github.com/urfave/cli"
)

var policyHeader = []string{"ID", "Name", "AlertID", "HostGroup", "ScalingAdjustment", "AdjustmentType"}

type policyOut struct {
	ID                string `json:"ID" yaml:"ID"`
	Name              string `json:"Name" yaml:"Name"`
	AlertID           string `json:"AlertID" yaml:"AlertID"`
	HostGroup         string `json:"HostGroup" yaml:"HostGroup"`
	ScalingAdjustment string `json:"ScalingAdjustment" yaml:"ScalingAdjustment"`
	AdjustmentType    string `json:"AdjustmentType" yaml:"AdjustmentType"`
}

func (r *policyOut) DataAsStringArray() []string {
	return []string{r.ID, r.Name, r.AlertID, r.HostGroup, r.ScalingAdjustment, r.AdjustmentType}
}

type autoscalePolicyClient interface {
	AddScalingPolicy(params *v1policies.AddScalingPolicyParams) (*v1policies.AddScalingPolicyOK, error)
	DeleteScalingPolicy(params *v1policies.DeleteScalingPolicyParams) error
	GetScalingPolicies(params *v1policies.GetScalingPoliciesParams) (*v1policies.GetScalingPoliciesOK, error)
	UpdateScalingPolicy(params *v1policies.UpdateScalingPolicyParams) (*v1policies.UpdateScalingPolicyOK, error)
}

func ListPolicies(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "list policies")
	asClient := oauth.NewAutoscaleHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	clusterID, err := strconv.ParseInt(c.String(fl.FlAsID.Name), 10, 64)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	listPoliciesImpl(asClient.Autoscale.V1policies, output.WriteList, clusterID)
}

func CreatePolicy(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "create autoscaling policy")
	asClient := oauth.NewAutoscaleHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	clusterID, err := strconv.ParseInt(c.String(fl.FlAsID.Name), 10, 64)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	alertID, err := strconv.ParseInt(c.String(fl.FlAlertID.Name), 10, 64)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	adjustment, err := strconv.ParseInt(c.String(fl.FlScalingAdjustment.Name), 10, 32)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	name := c.String(fl.FlName.Name)
	hostgroup := c.String(fl.FlHostgroup.Name)
	adjustmentType := c.String(fl.FlAdjustmentType.Name)
	createPolicyImpl(asClient.Autoscale.V1policies, output.Write, clusterID, alertID, adjustment, name, hostgroup, adjustmentType)
}

func UpdatePolicy(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "update autoscaling policy")
	asClient := oauth.NewAutoscaleHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	clusterID, err := strconv.ParseInt(c.String(fl.FlAsID.Name), 10, 64)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	policyID, err := strconv.ParseInt(c.String(fl.FlPolicyID.Name), 10, 64)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	alertID, err := strconv.ParseInt(c.String(fl.FlAlertID.Name), 10, 64)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	adjustment, err := strconv.ParseInt(c.String(fl.FlScalingAdjustment.Name), 10, 32)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	name := c.String(fl.FlName.Name)
	hostgroup := c.String(fl.FlHostgroup.Name)
	adjustmentType := c.String(fl.FlAdjustmentType.Name)
	updatePolicyImpl(asClient.Autoscale.V1policies, output.Write, clusterID, policyID, alertID, adjustment, name, hostgroup, adjustmentType)
}

func DeletePolicy(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "delete scaling policy")
	asClient := oauth.NewAutoscaleHTTPClientFromContext(c)
	clusterID, err := strconv.ParseInt(c.String(fl.FlAsID.Name), 10, 64)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	policyID, err := strconv.ParseInt(c.String(fl.FlPolicyID.Name), 10, 64)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	deletePolicyImpl(asClient.Autoscale.V1policies, clusterID, policyID)
}

func listPoliciesImpl(client autoscalePolicyClient, writer func([]string, []utils.Row), clusterID int64) {
	log.Infof("[listPoliciesImpl] sending autoscale list policies request")
	clusterResp, err := client.GetScalingPolicies(v1policies.NewGetScalingPoliciesParams().WithClusterID(clusterID))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	tableRows := []utils.Row{}
	for _, policy := range clusterResp.Payload {
		tableRows = append(tableRows, &policyOut{fmt.Sprintf("%v", policy.ID), policy.Name, fmt.Sprintf("%v", policy.AlertID), policy.HostGroup, fmt.Sprintf("%v", policy.ScalingAdjustment), policy.AdjustmentType})
	}
	writer(policyHeader, tableRows)
}

func createPolicyImpl(client autoscalePolicyClient, writer func([]string, utils.Row), clusterID, alertID, adjustment int64, name, hostgroup, adjustmentType string) {
	log.Infof("[createPolicyImpl] sending create policy request")
	clusterResp, err := client.AddScalingPolicy(v1policies.NewAddScalingPolicyParams().WithClusterID(clusterID).WithBody(&model.ScalingPolicyRequest{
		Name:              name,
		AlertID:           alertID,
		HostGroup:         hostgroup,
		ScalingAdjustment: int32(adjustment),
		AdjustmentType:    adjustmentType,
	}))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	policy := clusterResp.Payload
	tableRows := &policyOut{fmt.Sprintf("%v", policy.ID), policy.Name, fmt.Sprintf("%v", policy.AlertID), policy.HostGroup, fmt.Sprintf("%v", policy.ScalingAdjustment), policy.AdjustmentType}
	writer(policyHeader, tableRows)
}

func updatePolicyImpl(client autoscalePolicyClient, writer func([]string, utils.Row), clusterID, policyID, alertID, adjustment int64, name, hostgroup, adjustmentType string) {
	log.Infof("[createPolicyImpl] sending create policy request")
	clusterResp, err := client.UpdateScalingPolicy(v1policies.NewUpdateScalingPolicyParams().WithClusterID(clusterID).WithPolicyID(policyID).WithBody(&model.ScalingPolicyRequest{
		Name:              name,
		AlertID:           alertID,
		HostGroup:         hostgroup,
		ScalingAdjustment: int32(adjustment),
		AdjustmentType:    adjustmentType,
	}))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	policy := clusterResp.Payload
	tableRows := &policyOut{fmt.Sprintf("%v", policy.ID), policy.Name, fmt.Sprintf("%v", policy.AlertID), policy.HostGroup, fmt.Sprintf("%v", policy.ScalingAdjustment), policy.AdjustmentType}
	writer(policyHeader, tableRows)
}

func deletePolicyImpl(client autoscalePolicyClient, clusterID, policyID int64) {
	log.Infof("[deletePolicyImpl] sending delete policy request")
	err := client.DeleteScalingPolicy(v1policies.NewDeleteScalingPolicyParams().WithClusterID(clusterID).WithPolicyID(policyID))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
}
