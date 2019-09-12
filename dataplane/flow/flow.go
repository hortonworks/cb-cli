package flow

import (
	"strconv"
	"time"

	"github.com/hortonworks/cb-cli/dataplane/oauth"

	"github.com/hortonworks/cb-cli/dataplane/api/client/flow_logs"
	"github.com/hortonworks/cb-cli/dataplane/api/model"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/dp-cli-common/utils"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var flowHeader = []string{"FlowID", "Created", "CurrentState", "Finalized", "StateStatus", "NextEvent"}

type flow struct {
	FlowID       string `json:"FlowID" yaml:"FlowID"`
	Created      string `json:"Created" yaml:"Created"`
	CurrentState string `json:"CurrentState" yaml:"CurrentState"`
	Finalized    bool   `json:"Finalized" yaml:"Finalized"`
	StateStatus  string `json:"StateStatus" yaml:"StateStatus"`
	NextEvent    string `json:"NextEvent" yaml:"NextEvent"`
}

func (f *flow) DataAsStringArray() []string {
	return []string{f.FlowID, string(f.Created), f.CurrentState, strconv.FormatBool(f.Finalized), f.StateStatus, f.NextEvent}
}

type flowClient interface {
	GetLastFlowByResourceName(params *flow_logs.GetLastFlowByResourceNameParams) (*flow_logs.GetLastFlowByResourceNameOK, error)
	GetLastFlowByResourceCrn(params *flow_logs.GetLastFlowByResourceCrnParams) (*flow_logs.GetLastFlowByResourceCrnOK, error)
	GetFlowLogsByResourceName(params *flow_logs.GetFlowLogsByResourceNameParams) (*flow_logs.GetFlowLogsByResourceNameOK, error)
	GetFlowLogsByResourceCrn(params *flow_logs.GetFlowLogsByResourceCrnParams) (*flow_logs.GetFlowLogsByResourceCrnOK, error)
}

func ListByCrn(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "lists the flows by resourceCrn")
	log.Infof("[ListByCrn] Lists the flows by resourceCrn")
	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	resourceCrn := c.String(fl.FlCrn.Name)
	listFlowsByCrnImpl(cbClient.Cloudbreak.FlowLogs, resourceCrn, output.WriteList)
}

func ListByName(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "lists the flows by resource name")
	log.Infof("[ListByName] Lists the flows by resource Name")
	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	resourceName := c.String(fl.FlName.Name)
	listFlowsByNameImpl(cbClient.Cloudbreak.FlowLogs, resourceName, output.WriteList)
}

func listFlowsByCrnImpl(client flowClient, resourceCrn string, writer func([]string, []utils.Row)) {
	resp, err := client.GetFlowLogsByResourceCrn(flow_logs.NewGetFlowLogsByResourceCrnParams().WithResourceCrn(resourceCrn))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	writeOut(resp.Payload, writer)
}

func listFlowsByNameImpl(client flowClient, resourceName string, writer func([]string, []utils.Row)) {
	resp, err := client.GetFlowLogsByResourceName(flow_logs.NewGetFlowLogsByResourceNameParams().WithResourceName(resourceName))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	writeOut(resp.Payload, writer)
}

func DescribeLastFlowByCrn(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "describe the last flow by resourceCrn")
	log.Infof("[DescribeLastFlowByCrn] Describes the last flow by resourceCrn")
	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	resourceName := c.String(fl.FlName.Name)
	describeLastFlowByCrnImpl(cbClient.Cloudbreak.FlowLogs, resourceName, output.Write)
}

func DescribeLastFlowByName(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "describe the last flow by resourceName")
	log.Infof("[DescribeLastFlowByName] Describes the last flow by resourceName")
	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	resourceName := c.String(fl.FlName.Name)
	describeLastFlowByNameImpl(cbClient.Cloudbreak.FlowLogs, resourceName, output.Write)
}

func describeLastFlowByCrnImpl(client flowClient, resourceCrn string, writer func([]string, utils.Row)) {
	resp, err := client.GetLastFlowByResourceCrn(flow_logs.NewGetLastFlowByResourceCrnParams().WithResourceCrn(resourceCrn))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	writeOutOne(resp.Payload, writer)
}

func describeLastFlowByNameImpl(client flowClient, resourceName string, writer func([]string, utils.Row)) {
	resp, err := client.GetLastFlowByResourceName(flow_logs.NewGetLastFlowByResourceNameParams().WithResourceName(resourceName))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	writeOutOne(resp.Payload, writer)
}

func writeOutOne(r *model.FlowLogResponse, writer func([]string, utils.Row)) {
	writer(flowHeader, &flow{
		FlowID:       r.FlowID,
		Created:      unixTimeToString(r.Created),
		CurrentState: r.CurrentState,
		Finalized:    r.Finalized,
		StateStatus:  r.StateStatus,
		NextEvent:    r.NextEvent,
	})
}

func writeOut(responses []*model.FlowLogResponse, writer func([]string, []utils.Row)) {
	var tableRows []utils.Row
	for _, r := range responses {
		row := &flow{
			FlowID:       r.FlowID,
			Created:      unixTimeToString(r.Created),
			CurrentState: r.CurrentState,
			Finalized:    r.Finalized,
			StateStatus:  r.StateStatus,
			NextEvent:    r.NextEvent,
		}
		tableRows = append(tableRows, row)
	}
	writer(flowHeader, tableRows)
}

func unixTimeToString(unix int64) string {
	t := time.Unix(unix/1000, 0)
	return t.Format("2006-01-02 15:04:05")
}
