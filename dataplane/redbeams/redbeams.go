package redbeams

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	model "github.com/hortonworks/cb-cli/dataplane/api-redbeams"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/oauth"
	commonutils "github.com/hortonworks/dp-cli-common/utils"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type ClientRedbeams oauth.Redbeams

func _nilsafe(s *string) string {
	if s == nil {
		return "<nil>"
	} else {
		return *s
	}
}

var serverListHeader = []string{"Name", "Description", "Crn", "EnvironmentCrn", "Status", "ResourceStatus", "DatabaseVendor", "Host", "Port", "CreationDate"}

type serverDetails struct {
	Name           string `json:"Name" yaml:"Name"`
	Description    string `json:"Description" yaml:"Description"`
	CRN            string `json:"CRN" yaml:"CRN"`
	EnvironmentCrn string `json:"EnvironmentCrn" yaml:"EnvironmentCrn"`
	Status         string `json:"Status" yaml:"Status"`
	StatusReason   string `json:"StatusReason" yaml:"StatusReason"`
	ResourceStatus string `json:"ResourceStatus" yaml:"ResourceStatus"`
	DatabaseVendor string `json:"DatabaseVendor" yaml:"DatabaseVendor"`
	Host           string `json:"Host" yaml:"Host"`
	Port           int32  `json:"Port" yaml:"Port"`
	CreationDate   int64  `json:"CreationDate" yaml:"CreationDate"`
}

func (server *serverDetails) DataAsStringArray() []string {
	return []string{server.Name, server.Description, server.CRN, server.EnvironmentCrn, server.Status, server.ResourceStatus, server.DatabaseVendor, server.Host, strconv.FormatInt(int64(server.Port), 10), strconv.FormatInt(server.CreationDate, 10)}
}

func NewDetailsFromResponse(r *model.DatabaseServerV4Response) *serverDetails {
	details := &serverDetails{
		Name:           r.Name,
		CRN:            *r.Crn,
		EnvironmentCrn: r.EnvironmentCrn,
		Status:         *r.Status,
		StatusReason:   *r.StatusReason,
		ResourceStatus: *r.ResourceStatus,
		CreationDate:   *r.CreationDate,
	}

	// ternary operator, where art thou?
	if r.Description != nil {
		details.Description = *r.Description
	}
	if &r.DatabaseVendor != nil {
		details.DatabaseVendor = r.DatabaseVendor
	}
	if &r.Host != nil {
		details.Host = r.Host
	}
	if &r.Port != nil {
		details.Port = r.Port
	}

	return details
}

var statusListHeader = []string{"Name", "Crn", "EnvironmentCrn", "Status"}

type serverStatusDetails struct {
	Name           string `json:"Name" yaml:"Name"`
	CRN            string `json:"CRN" yaml:"CRN"`
	EnvironmentCrn string `json:"EnvironmentCrn" yaml:"EnvironmentCrn"`
	Status         string `json:"Status" yaml:"Status"`
}

func (status *serverStatusDetails) DataAsStringArray() []string {
	return []string{status.Name, status.CRN, status.EnvironmentCrn, status.Status}
}

func NewDetailsFromStatusResponse(r *model.DatabaseServerStatusV4Response) *serverStatusDetails {
	details := &serverStatusDetails{
		Name:           r.Name,
		CRN:            r.ResourceCrn,
		EnvironmentCrn: r.EnvironmentCrn,
		Status:         r.Status,
	}

	return details
}

func ListDatabaseServers(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "List database servers in environment")
	envCrn := c.String(fl.FlEnvironmentCrn.Name)
	redbeamsDbServerClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.DatabaseServersApi

	log.Infof("[ListDBServers] Listing database servers in environment: %s", envCrn)
	resp, _, err := redbeamsDbServerClient.ListDatabaseServersExecute(redbeamsDbServerClient.ListDatabaseServers(context.Background()).EnvironmentCrn(envCrn))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}

	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}

	var tableRows []commonutils.Row

	for _, response := range resp.GetResponses() {
		row := NewDetailsFromResponse(&response)
		tableRows = append(tableRows, row)
	}
	output.WriteList(serverListHeader, tableRows)
}

func GetDatabaseServer(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Get a database server")
	redbeamsDbServerClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.DatabaseServersApi

	var server *model.DatabaseServerV4Response
	crn := c.String(fl.FlCrn.Name)
	if len(crn) != 0 {
		log.Infof("[GetDBServer] Getting database server with CRN: %s", crn)
		resp, _, err := redbeamsDbServerClient.GetDatabaseServerByCrnExecute(redbeamsDbServerClient.GetDatabaseServerByCrn(context.Background(), crn))
		if err != nil {
			commonutils.LogErrorAndExit(err)
		}
		server = resp
	} else {
		envCrn := c.String(fl.FlEnvironmentCrn.Name)
		name := c.String(fl.FlName.Name)
		log.Infof("[GetDBServer] Getting database server with name: %s", name)
		resp, _, err := redbeamsDbServerClient.GetDatabaseServerByNameExecute(redbeamsDbServerClient.GetDatabaseServerByName(context.Background(), name).EnvironmentCrn(envCrn))
		if err != nil {
			commonutils.LogErrorAndExit(err)
		}
		server = resp
	}

	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	row := NewDetailsFromResponse(server)
	output.Write(serverListHeader, row)
}

func CreateManagedDatabaseServer(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Create a managed database server")
	fileLocation := c.String(fl.FlDatabaseServerCreationFile.Name)

	log.Infof("[CreateManagedDBServer] Creating database server from file: %s", fileLocation)
	content := commonutils.ReadFile(fileLocation)
	var req model.AllocateDatabaseServerV4Request
	err := json.Unmarshal(content, &req)
	if err != nil {
		msg := fmt.Sprintf(`Invalid JSON: %s`, err.Error())
		commonutils.LogErrorMessageAndExit(msg)
	}

	log.Infof("[CreateManagedDBServer] JSON read, creating database server with name: %s", req.Name)
	redbeamsDbServerClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.DatabaseServersApi
	resp, _, err := redbeamsDbServerClient.CreateDatabaseServerExecute(redbeamsDbServerClient.CreateDatabaseServer(context.Background()).AllocateDatabaseServerV4Request(req))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	status := resp

	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	row := NewDetailsFromStatusResponse(status)
	output.Write(statusListHeader, row)
}

func ReleaseManagedDatabaseServer(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Release a managed database server")
	crn := c.String(fl.FlCrn.Name)
	redbeamsDbServerClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.DatabaseServersApi

	log.Infof("[ReleaseDBServer] Releasing database server with CRN: %s", crn)
	resp, _, err := redbeamsDbServerClient.ReleaseManagedDatabaseServerExecute(redbeamsDbServerClient.ReleaseManagedDatabaseServer(context.Background(), crn))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	server := resp

	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	row := NewDetailsFromResponse(server)
	output.Write(serverListHeader, row)
}

func RegisterDatabaseServer(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Register database server")
	fileLocation := c.String(fl.FlDatabaseServerRegistrationFile.Name)

	log.Infof("[RegisterDBServer] Registrating database server from file: %s", fileLocation)
	content := commonutils.ReadFile(fileLocation)
	var req model.DatabaseServerV4Request
	err := json.Unmarshal(content, &req)
	if err != nil {
		msg := fmt.Sprintf(`Invalid JSON: %s`, err.Error())
		commonutils.LogErrorMessageAndExit(msg)
	}

	log.Infof("[RegisterDBServer] JSON read, registering database server with name: %s", _nilsafe(&req.Name))
	redbeamsDbServerClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.DatabaseServersApi
	resp, _, err := redbeamsDbServerClient.RegisterDatabaseServerExecute(redbeamsDbServerClient.RegisterDatabaseServer(context.Background()).DatabaseServerV4Request(req))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	server := resp

	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	row := NewDetailsFromResponse(server)
	output.Write(serverListHeader, row)
}

func DeleteDatabaseServer(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Terminate and/or delete a database server")
	crn := c.String(fl.FlCrn.Name)
	force := c.Bool(fl.FlForceOptional.Name)
	redbeamsDbServerClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.DatabaseServersApi

	log.Infof("[DeleteDBServer] Deleting database server with CRN: %s force: %t", crn, force)
	resp, _, err := redbeamsDbServerClient.DeleteDatabaseServerByCrnExecute(redbeamsDbServerClient.DeleteDatabaseServerByCrn(context.Background(), crn).Force(force))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}

	server := resp
	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	row := NewDetailsFromResponse(server)
	output.Write(serverListHeader, row)
}

var dbListHeader = []string{"Name", "Type", "Description", "Crn", "EnvironmentCrn", "DatabaseVendor", "ConnectionURL", "ResourceStatus", "CreationDate"}

type dbDetails struct {
	Name           string `json:"Name" yaml:"Name"`
	Type           string `json:"Type" yaml:"Type"`
	Description    string `json:"Description" yaml:"Description"`
	CRN            string `json:"CRN" yaml:"CRN"`
	EnvironmentCrn string `json:"EnvironmentCrn" yaml:"EnvironmentCrn"`
	DatabaseEngine string `json:"DatabaseEngine" yaml:"DatabaseEngine"`
	ConnectionURL  string `json:"ConnectionURL" yaml:"ConnectionURL"`
	ResourceStatus string `json:"ResourceStatus" yaml:"ResourceStatus"`
	CreationDate   int64  `json:"CreationDate" yaml:"CreationDate"`
}

func (db *dbDetails) DataAsStringArray() []string {
	return []string{db.Name, db.Type, db.Description, db.CRN, db.EnvironmentCrn, db.DatabaseEngine, db.ConnectionURL, db.ResourceStatus, strconv.FormatInt(db.CreationDate, 10)}
}

func NewDetailsFromDbResponse(r *model.DatabaseV4Response) *dbDetails {
	details := &dbDetails{
		Name:           r.Name,
		Type:           r.Type,
		CRN:            *r.Crn,
		EnvironmentCrn: r.EnvironmentCrn,
		ConnectionURL:  r.ConnectionURL,
		ResourceStatus: *r.ResourceStatus,
		CreationDate:   *r.CreationDate,
	}

	if r.Description != nil {
		details.Description = *r.Description
	}
	if &r.DatabaseEngine != nil {
		details.DatabaseEngine = r.DatabaseEngine
	}

	return details
}

func ListDatabases(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "List databases in environment")
	envCrn := c.String(fl.FlEnvironmentCrn.Name)
	redbeamsDbClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.DatabasesApi

	log.Infof("[ListDBs] Listing databases in environment: %s", envCrn)
	resp, _, err := redbeamsDbClient.ListDatabasesExecute(redbeamsDbClient.ListDatabases(context.Background()).EnvironmentCrn(envCrn))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}

	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}

	var tableRows []commonutils.Row

	for _, response := range resp.Responses {
		row := NewDetailsFromDbResponse(&response)
		tableRows = append(tableRows, row)
	}
	output.WriteList(dbListHeader, tableRows)
}

func GetDatabase(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Get a database")
	redbeamsDbClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.DatabasesApi

	var db *model.DatabaseV4Response
	crn := c.String(fl.FlCrn.Name)
	if len(crn) != 0 {
		log.Infof("[GetDB] Getting database with CRN: %s", crn)
		resp, _, err := redbeamsDbClient.GetDatabaseByCrnExecute(redbeamsDbClient.GetDatabaseByCrn(context.Background(), crn))
		if err != nil {
			commonutils.LogErrorAndExit(err)
		}
		db = resp
	} else {
		envCrn := c.String(fl.FlEnvironmentCrn.Name)
		name := c.String(fl.FlName.Name)
		log.Infof("[GetDB] Getting database with name: %s", name)
		resp, _, err := redbeamsDbClient.GetDatabaseByNameExecute(redbeamsDbClient.GetDatabaseByName(context.Background(), name).EnvironmentCrn(envCrn))
		if err != nil {
			commonutils.LogErrorAndExit(err)
		}
		db = resp
	}

	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	row := NewDetailsFromDbResponse(db)
	output.Write(dbListHeader, row)
}

func CreateDatabase(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Create a database")
	fileLocation := c.String(fl.FlDatabaseCreationFile.Name)

	log.Infof("[CreateDB] Creating database from file: %s", fileLocation)
	content := commonutils.ReadFile(fileLocation)
	var req model.CreateDatabaseV4Request
	err := json.Unmarshal(content, &req)
	if err != nil {
		msg := fmt.Sprintf(`Invalid JSON: %s`, err.Error())
		commonutils.LogErrorMessageAndExit(msg)
	}

	log.Infof("[CreateDB] JSON read, creating database with name: %s", _nilsafe(&req.DatabaseName))
	redbeamsDbServerClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.DatabaseServersApi
	result, _, err := redbeamsDbServerClient.CreateDatabaseOnServerExecute(redbeamsDbServerClient.CreateDatabaseOnServer(context.Background()).CreateDatabaseV4Request(req))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}

	log.Infof("[DeleteDBServer] Created database with name: %s Details: %s", _nilsafe(&req.DatabaseName), result)
}

func RegisterDatabase(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Register database")
	fileLocation := c.String(fl.FlDatabaseRegistrationFile.Name)

	log.Infof("[RegisterDB] Registering database from file: %s", fileLocation)
	content := commonutils.ReadFile(fileLocation)
	var req model.DatabaseV4Request
	err := json.Unmarshal(content, &req)
	if err != nil {
		msg := fmt.Sprintf(`Invalid JSON: %s`, err.Error())
		commonutils.LogErrorMessageAndExit(msg)
	}

	log.Infof("[RegisterDB] JSON read, registering database with name: %s", _nilsafe(&req.Name))
	redbeamsDbClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.DatabasesApi
	resp, _, err := redbeamsDbClient.RegisterDatabaseExecute(redbeamsDbClient.RegisterDatabase(context.Background()).DatabaseV4Request(req))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	db := resp

	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	row := NewDetailsFromDbResponse(db)
	output.Write(dbListHeader, row)
}

func DeleteDatabase(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Delete a registered database")
	redbeamsDbClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.DatabasesApi
	crn := c.String(fl.FlCrn.Name)
	if len(crn) != 0 {
		log.Infof("[DeleteDB] Deleting database with CRN: %s", crn)
		result, _, err := redbeamsDbClient.DeleteDatabaseByCrnExecute(redbeamsDbClient.DeleteDatabaseByCrn(context.Background(), crn))
		if err != nil {
			commonutils.LogErrorAndExit(err)
		}
		log.Infof("[DeleteDB] Deleted database with CRN: %s Details: %s", crn, result)
	} else {
		envCrn := c.String(fl.FlEnvironmentCrn.Name)
		name := c.String(fl.FlName.Name)
		log.Infof("[DeleteDB] Deleting database with name: %s", name)
		result, _, err := redbeamsDbClient.DeleteDatabaseByNameExecute(redbeamsDbClient.DeleteDatabaseByName(context.Background(), name).EnvironmentCrn(envCrn))
		if err != nil {
			commonutils.LogErrorAndExit(err)
		}
		log.Infof("[DeleteDB] Deleted database with name: %s Details: %s", name, result)
	}
}
