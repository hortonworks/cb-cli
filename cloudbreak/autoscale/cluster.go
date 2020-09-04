package autoscale

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hortonworks/cb-cli/cloudbreak/api-as/client/v1clusters"
	"github.com/hortonworks/cb-cli/cloudbreak/api-as/client/v2clusters"
	"github.com/hortonworks/cb-cli/cloudbreak/oauth"

	log "github.com/Sirupsen/logrus"
	fl "github.com/hortonworks/cb-cli/cloudbreak/flags"
	"github.com/hortonworks/cb-cli/utils"
	"github.com/urfave/cli"
)

var clusterHeader = []string{"AsID", "ClusterID", "Host", "Port", "Enabled"}

type clusterOut struct {
	AsID      string `json:"AsID" yaml:"AsID"`
	ClusterID string `json:"ClusterID" yaml:"ClusterID"`
	Host      string `json:"Host" yaml:"Host"`
	Port      string `json:"Port" yaml:"Port"`
	Enabled   bool   `json:"Enabled" yaml:"Enabled"`
}

func (r *clusterOut) DataAsStringArray() []string {
	return []string{r.AsID, r.ClusterID, r.Host, r.Port, strconv.FormatBool(r.Enabled)}
}

type autoscaleListClient interface {
	GetClusters(params *v1clusters.GetClustersParams) (*v1clusters.GetClustersOK, error)
}

type autoscaleClient interface {
	EnableAutoscaleStateByCloudbreakCluster(params *v2clusters.EnableAutoscaleStateByCloudbreakClusterParams) (*v2clusters.EnableAutoscaleStateByCloudbreakClusterOK, error)
	DisableAutoscaleStateByCloudbreakCluster(params *v2clusters.DisableAutoscaleStateByCloudbreakClusterParams) (*v2clusters.DisableAutoscaleStateByCloudbreakClusterOK, error)
}

func ListClusters(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "list clusters")

	asClient := oauth.NewAutoscaleHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	listClustersImpl(asClient.Autoscale.V1clusters, output.WriteList)
}

func Enable(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "enable autoscale for cluster")

	asClient := oauth.NewAutoscaleHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	clusterID, err := strconv.ParseInt(c.String(fl.FlClusterID.Name), 10, 64)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	enableImpl(asClient.Autoscale.V2clusters, output.Write, clusterID)
}

func Disable(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "disable autoscale for cluster")

	asClient := oauth.NewAutoscaleHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	clusterID, err := strconv.ParseInt(c.String(fl.FlClusterID.Name), 10, 64)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	disableImpl(asClient.Autoscale.V2clusters, output.Write, clusterID)
}

func listClustersImpl(client autoscaleListClient, writer func([]string, []utils.Row)) {
	log.Infof("[listClustersImpl] sending autoscale cluster list request")
	clusterResp, err := client.GetClusters(v1clusters.NewGetClustersParams())
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	tableRows := []utils.Row{}
	for _, cluster := range clusterResp.Payload {
		tableRows = append(tableRows, &clusterOut{fmt.Sprintf("%v", cluster.ID), fmt.Sprintf("%v", *cluster.StackID), cluster.Host, cluster.Port, *cluster.AutoscalingEnabled})
	}

	writer(clusterHeader, tableRows)
}

func enableImpl(client autoscaleClient, writer func([]string, utils.Row), clusterID int64) {
	log.Infof("[enableImpl] sending enable autoscale for cluster request")
	clusterResp, err := client.EnableAutoscaleStateByCloudbreakCluster(v2clusters.NewEnableAutoscaleStateByCloudbreakClusterParams().WithCbClusterID(clusterID))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	cluster := clusterResp.Payload
	tableRows := &clusterOut{fmt.Sprintf("%v", cluster.ID), fmt.Sprintf("%v", *cluster.StackID), cluster.Host, cluster.Port, *cluster.AutoscalingEnabled}
	writer(clusterHeader, tableRows)
}

func disableImpl(client autoscaleClient, writer func([]string, utils.Row), clusterID int64) {
	log.Infof("[disableImpl] sending disable autoscale for cluster request")
	clusterResp, err := client.DisableAutoscaleStateByCloudbreakCluster(v2clusters.NewDisableAutoscaleStateByCloudbreakClusterParams().WithCbClusterID(clusterID))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	cluster := clusterResp.Payload
	tableRows := &clusterOut{fmt.Sprintf("%v", cluster.ID), fmt.Sprintf("%v", *cluster.StackID), cluster.Host, cluster.Port, *cluster.AutoscalingEnabled}
	writer(clusterHeader, tableRows)
}
