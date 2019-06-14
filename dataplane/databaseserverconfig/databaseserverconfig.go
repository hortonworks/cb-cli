package databaseserverconfig

import (
	"strconv"

	"github.com/hortonworks/cb-cli/dataplane/api-redbeams/client/v4databaseservers"
	"github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/oauth"
	"github.com/hortonworks/dp-cli-common/utils"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type ClientRedbeams oauth.Redbeams

var databaseServerHeader = []string{"Name", "Description", "Host", "Port", "Crn", "Database Vendor", "Connection Driver", "Connector jar URL", "Environment ID"}

type databaseserverRow struct {
	Name             string `json:"Name" yaml:"Name"`
	Description      string `json:"Description" yaml:"Description"`
	Host             string `json:"Host" yaml:"Host"`
	Port             int32  `json:"Port" yaml:"Port"`
	Crn              string `json:"Crn" yaml:"Crn"`
	DatabaseVendor   string `json:"Database vendor" yaml:"Database vendor"`
	ConnectionDriver string `json:"ConnectionDriver" yaml:"ConnectionDriver"`
	ConnectorJarURL  string `json:"Connector jar url" yaml:"Connector jar url"`
	EnvironmentID    string `json:"Environment id" yaml:"Environment id"`
}

func (ds databaseserverRow) DataAsStringArray() []string {
	return []string{ds.Name, ds.Description, ds.Host, strconv.FormatInt(int64(ds.Port), 10), ds.Crn, ds.DatabaseVendor, ds.ConnectionDriver, ds.ConnectorJarURL, ds.EnvironmentID}
}

func StubFunc(c *cli.Context) {
	// an empty func
	log.Errorf("Not implemented yet")
}

func List(c *cli.Context) {
	redbeamsClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams
	environmentID := c.String(flags.FlEnvironmentCrn.Name)
	resp, err := redbeamsClient.V4databaseservers.ListDatabaseServers(v4databaseservers.NewListDatabaseServersParams().WithEnvironmentID(environmentID))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	output := utils.Output{Format: c.String(flags.FlOutputOptional.Name)}
	var tableRows []utils.Row

	for _, ds := range resp.Payload.Responses {
		row := &databaseserverRow{
			Name:             *ds.Name,
			Description:      getValue(ds.Description),
			Host:             *ds.Host,
			Port:             *ds.Port,
			Crn:              ds.Crn,
			DatabaseVendor:   *ds.DatabaseVendor,
			ConnectionDriver: *ds.ConnectionDriver,
			ConnectorJarURL:  getValue(ds.ConnectorJarURL),
			EnvironmentID:    *ds.EnvironmentID,
		}
		tableRows = append(tableRows, row)
	}
	output.WriteList(databaseServerHeader, tableRows)
}

func getValue(valuePtr *string) string {
	if valuePtr == nil {
		return ""
	}
	return *valuePtr
}
