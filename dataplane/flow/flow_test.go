package flow

import (
	"strings"
	"testing"
	"time"

	flow_logs "github.com/hortonworks/cb-cli/dataplane/api/client/flow"
	"github.com/hortonworks/cb-cli/dataplane/api/model"
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/aws"
	"github.com/hortonworks/dp-cli-common/utils"
)

const layout = "2006-01-02 15:04:05"
const unixTime = 1568293332

var expectedTime = time.Unix(unixTime/1000, 0)
var expected = "test " + expectedTime.Format(layout) + " FINISHED true SUCCESSFUL empty"

type mockFlowClient struct {
}

var resp = model.FlowLogResponse{
	FlowID:       "test",
	Created:      unixTime,
	CurrentState: "FINISHED",
	Finalized:    true,
	StateStatus:  "SUCCESSFUL",
	NextEvent:    "empty",
}

func (*mockFlowClient) GetLastFlowByResourceName(params *flow_logs.GetLastFlowByResourceNameParams) (*flow_logs.GetLastFlowByResourceNameOK, error) {
	return &flow_logs.GetLastFlowByResourceNameOK{Payload: &resp}, nil
}

func (*mockFlowClient) GetLastFlowByResourceCrn(params *flow_logs.GetLastFlowByResourceCrnParams) (*flow_logs.GetLastFlowByResourceCrnOK, error) {
	return &flow_logs.GetLastFlowByResourceCrnOK{Payload: &resp}, nil
}

func (*mockFlowClient) GetFlowLogsByResourceName(params *flow_logs.GetFlowLogsByResourceNameParams) (*flow_logs.GetFlowLogsByResourceNameOK, error) {
	return &flow_logs.GetFlowLogsByResourceNameOK{Payload: []*model.FlowLogResponse{&resp}}, nil
}

func (*mockFlowClient) GetFlowLogsByResourceCrn(params *flow_logs.GetFlowLogsByResourceCrnParams) (*flow_logs.GetFlowLogsByResourceCrnOK, error) {
	return &flow_logs.GetFlowLogsByResourceCrnOK{Payload: []*model.FlowLogResponse{&resp}}, nil
}

func TestListFlowsByCrnImpl(t *testing.T) {
	var rows []utils.Row
	listFlowsByCrnImpl(new(mockFlowClient), "resource", func(h []string, r []utils.Row) { rows = r })
	if len(rows) != 1 {
		t.Fatalf("row number doesn't match 1 == %d", len(rows))
	}
	for _, r := range rows {
		checkRow(t, r)
	}
}

func TestListFlowsByNameImpl(t *testing.T) {
	var rows []utils.Row
	listFlowsByNameImpl(new(mockFlowClient), "resource", func(h []string, r []utils.Row) { rows = r })
	if len(rows) != 1 {
		t.Fatalf("row number doesn't match 1 == %d", len(rows))
	}
	for _, r := range rows {
		checkRow(t, r)
	}
}

func TestDescribeLastFlowByNameImpl(t *testing.T) {
	var row utils.Row
	describeLastFlowByNameImpl(new(mockFlowClient), "100", func(h []string, r utils.Row) { row = r })
	checkRow(t, row)
}

func TestDescribeLastFlowByCrnImpl(t *testing.T) {
	var row utils.Row
	describeLastFlowByCrnImpl(new(mockFlowClient), "100", func(h []string, r utils.Row) { row = r })
	checkRow(t, row)
}

func checkRow(t *testing.T, row utils.Row) {
	if strings.Join(row.DataAsStringArray(), " ") != expected {
		t.Errorf("row data not match \"%s\" == \"%s\"", expected, strings.Join(row.DataAsStringArray(), " "))
	}
}
