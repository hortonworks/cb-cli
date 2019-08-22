package redbeams

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hortonworks/cb-cli/dataplane/api-redbeams/client/database_servers"
	"github.com/hortonworks/cb-cli/dataplane/api-redbeams/client/databases"
	"github.com/hortonworks/cb-cli/dataplane/api-redbeams/model"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/oauth"
	commonutils "github.com/hortonworks/dp-cli-common/utils"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type ClientRedbeams oauth.Redbeams

var serverListHeader = []string{"Name", "Description", "Crn", "EnvironmentCrn", "Status", "ResourceStatus", "DatabaseVendor", "Host", "Port"}

type serverDetails struct {
	Name           string `json:"name" yaml:"name"`
	Description    string `json:"description" yaml:"description"`
	CRN            string `json:"crn" yaml:"crn"`
	EnvironmentCrn string `json:"environmentCrn" yaml:"environmentCrn"`
	Status         string `json:"status" yaml:"status"`
	StatusReason   string `json:"statusReason" yaml:"statusReason"`
	ResourceStatus string `json:"resourceStatus" yaml:"resourceStatus"`
	DatabaseVendor string `json:"databaseVendor" yaml:"databaseVendor"`
	Host           string `json:"host" yaml:"host"`
	Port           int32  `json:"port" yaml:"port"`
	CreationDate   int64  `json:"creationDate" yaml:"creationDate"`
}

func (server *serverDetails) DataAsStringArray() []string {
	return []string{server.Name, server.Description, server.CRN, server.EnvironmentCrn, server.Status, server.DatabaseVendor, server.Host, string(server.Port)}
}

func NewDetailsFromResponse(r *model.DatabaseServerV4Response) *serverDetails {
	details := &serverDetails{
		Name:           *r.Name,
		CRN:            r.Crn,
		EnvironmentCrn: *r.EnvironmentCrn,
		Status:         r.Status,
		StatusReason:   r.StatusReason,
		ResourceStatus: r.ResourceStatus,
		CreationDate:   r.CreationDate,
	}

	// ternary operator, where art thou?
	if r.Description != nil {
		details.Description = *r.Description
	}
	if r.DatabaseVendor != nil {
		details.DatabaseVendor = *r.DatabaseVendor
	}
	if r.Host != nil {
		details.Host = *r.Host
	}
	if r.Port != nil {
		details.Port = *r.Port
	}

	return details
}

var statusListHeader = []string{"Name", "Crn", "EnvironmentCrn", "Status"}

type serverStatusDetails struct {
	Name           string `json:"name" yaml:"name"`
	CRN            string `json:"crn" yaml:"crn"`
	EnvironmentCrn string `json:"environmentCrn" yaml:"environmentCrn"`
	Status         string `json:"status" yaml:"status"`
}

func (status *serverStatusDetails) DataAsStringArray() []string {
	return []string{status.Name, status.CRN, status.EnvironmentCrn, status.Status}
}

func NewDetailsFromStatusResponse(r *model.DatabaseServerStatusV4Response) *serverStatusDetails {
	details := &serverStatusDetails{
		Name:           *r.Name,
		CRN:            *r.ResourceCrn,
		EnvironmentCrn: *r.EnvironmentCrn,
		Status:         *r.Status,
	}

	return details
}

func ListDatabaseServers(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "List database servers in environment")
	envCrn := c.String(fl.FlEnvironmentCrn.Name)
	redbeamsDbServerClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.DatabaseServers

	log.Infof("[ListDBServers] Listing database servers in environment: %s", envCrn)
	resp, err := redbeamsDbServerClient.ListDatabaseServers(database_servers.NewListDatabaseServersParams().WithEnvironmentCrn(envCrn), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}

	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}

	var tableRows []commonutils.Row

	for _, response := range resp.Payload.Responses {
		row := NewDetailsFromResponse(response)
		tableRows = append(tableRows, row)
	}
	output.WriteList(serverListHeader, tableRows)
}

func GetDatabaseServer(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Get a database server")
	redbeamsDbServerClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.DatabaseServers

	var server *model.DatabaseServerV4Response
	crn := c.String(fl.FlCrn.Name)
	if len(crn) != 0 {
		log.Infof("[GetDBServer] Getting database server with CRN: %s", crn)
		resp, err := redbeamsDbServerClient.GetDatabaseServerByCrn(database_servers.NewGetDatabaseServerByCrnParams().WithCrn(crn), nil)
		if err != nil {
			commonutils.LogErrorAndExit(err)
		}
		server = resp.Payload
	} else {
		envCrn := c.String(fl.FlEnvironmentCrn.Name)
		name := c.String(fl.FlName.Name)
		log.Infof("[GetDBServer] Getting database server with name: %s", name)
		resp, err := redbeamsDbServerClient.GetDatabaseServerByName(database_servers.NewGetDatabaseServerByNameParams().WithEnvironmentCrn(envCrn).WithName(name), nil)
		if err != nil {
			commonutils.LogErrorAndExit(err)
		}
		server = resp.Payload
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
	redbeamsDbServerClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.DatabaseServers
	resp, err := redbeamsDbServerClient.CreateDatabaseServer(database_servers.NewCreateDatabaseServerParams().WithBody(&req), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	status := resp.Payload

	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	row := NewDetailsFromStatusResponse(status)
	output.Write(statusListHeader, row)
}

func ReleaseManagedDatabaseServer(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Release a managed database server")
	crn := c.String(fl.FlCrn.Name)
	redbeamsDbServerClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.DatabaseServers

	log.Infof("[ReleaseDBServer] Releasing database server with CRN: %s", crn)
	resp, err := redbeamsDbServerClient.ReleaseManagedDatabaseServer(database_servers.NewReleaseManagedDatabaseServerParams().WithCrn(crn), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	server := resp.Payload

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

	log.Infof("[RegisterDBServer] JSON read, registering database server with name: %s", *req.Name)
	redbeamsDbServerClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.DatabaseServers
	resp, err := redbeamsDbServerClient.RegisterDatabaseServer(database_servers.NewRegisterDatabaseServerParams().WithBody(&req), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	server := resp.Payload

	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	row := NewDetailsFromResponse(server)
	output.Write(serverListHeader, row)
}

func DeleteDatabaseServer(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Terminate and/or delete a database server")
	crn := c.String(fl.FlCrn.Name)
	force := c.Bool(fl.FlForceOptional.Name)
	redbeamsDbServerClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.DatabaseServers

	log.Infof("[DeleteDBServer] Deleting database server with CRN: %s force: %t", crn, force)
	resp, err := redbeamsDbServerClient.DeleteDatabaseServerByCrn(database_servers.NewDeleteDatabaseServerByCrnParams().WithCrn(crn).WithForce(&force), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}

	server := resp.Payload
	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	row := NewDetailsFromResponse(server)
	output.Write(serverListHeader, row)
}

var dbListHeader = []string{"Name", "Description", "Crn", "EnvironmentCrn", "DatabaseVendor", "ConnectionURL"}

type dbDetails struct {
	Name           string `json:"name" yaml:"name"`
	Description    string `json:"description" yaml:"description"`
	CRN            string `json:"crn" yaml:"crn"`
	EnvironmentCrn string `json:"environmentCrn" yaml:"environmentCrn"`
	DatabaseEngine string `json:"databaseEngine" yaml:"databaseEngine"`
	ConnectionURL  string `json:"connectionURL" yaml:"connectionURL"`
	CreationDate   int64  `json:"creationDate" yaml:"creationDate"`
}

func (db *dbDetails) DataAsStringArray() []string {
	return []string{db.Name, db.Description, db.CRN, db.EnvironmentCrn, db.DatabaseEngine, db.ConnectionURL, string(db.CreationDate)}
}

func NewDetailsFromDbResponse(r *model.DatabaseV4Response) *dbDetails {
	details := &dbDetails{
		Name:           *r.Name,
		CRN:            r.Crn,
		EnvironmentCrn: *r.EnvironmentCrn,
		ConnectionURL:  *r.ConnectionURL,
		CreationDate:   r.CreationDate,
	}

	if r.Description != nil {
		details.Description = *r.Description
	}
	if r.DatabaseEngine != nil {
		details.DatabaseEngine = *r.DatabaseEngine
	}

	return details
}

func ListDatabases(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "List databases in environment")
	envCrn := c.String(fl.FlEnvironmentCrn.Name)
	redbeamsDbClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.Databases

	log.Infof("[ListDBs] Listing databases in environment: %s", envCrn)
	resp, err := redbeamsDbClient.ListDatabases(databases.NewListDatabasesParams().WithEnvironmentCrn(envCrn), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}

	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}

	var tableRows []commonutils.Row

	for _, response := range resp.Payload.Responses {
		row := NewDetailsFromDbResponse(response)
		tableRows = append(tableRows, row)
	}
	output.WriteList(dbListHeader, tableRows)
}

func GetDatabase(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Get a database")
	redbeamsDbClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.Databases

	var db *model.DatabaseV4Response
	crn := c.String(fl.FlCrn.Name)
	if len(crn) != 0 {
		log.Infof("[GetDB] Getting database with CRN: %s", crn)
		resp, err := redbeamsDbClient.GetDatabaseByCrn(databases.NewGetDatabaseByCrnParams().WithCrn(crn), nil)
		if err != nil {
			commonutils.LogErrorAndExit(err)
		}
		db = resp.Payload
	} else {
		envCrn := c.String(fl.FlEnvironmentCrn.Name)
		name := c.String(fl.FlName.Name)
		log.Infof("[GetDB] Getting database with name: %s", name)
		resp, err := redbeamsDbClient.GetDatabaseByName(databases.NewGetDatabaseByNameParams().WithEnvironmentCrn(envCrn).WithName(name), nil)
		if err != nil {
			commonutils.LogErrorAndExit(err)
		}
		db = resp.Payload
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

	log.Infof("[CreateDB] JSON read, creating database with name: %s", *req.DatabaseName)
	redbeamsDbServerClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.DatabaseServers
	result, err := redbeamsDbServerClient.CreateDatabaseOnServer(database_servers.NewCreateDatabaseOnServerParams().WithBody(&req), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}

	log.Infof("[DeleteDBServer] Created database with name: %s Details: %s", *req.DatabaseName, result)
}

func RegisterDatabase(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Register database")
	fileLocation := c.String(fl.FlDatabaseRegistrationFile.Name)

	log.Infof("[RegisterDB] Registering database server from file: %s", fileLocation)
	content := commonutils.ReadFile(fileLocation)
	var req model.DatabaseV4Request
	err := json.Unmarshal(content, &req)
	if err != nil {
		msg := fmt.Sprintf(`Invalid JSON: %s`, err.Error())
		commonutils.LogErrorMessageAndExit(msg)
	}

	log.Infof("[RegisterDB] JSON read, registering database with name: %s", *req.Name)
	redbeamsDbClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.Databases
	resp, err := redbeamsDbClient.RegisterDatabase(databases.NewRegisterDatabaseParams().WithBody(&req), nil)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	db := resp.Payload

	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	row := NewDetailsFromDbResponse(db)
	output.Write(dbListHeader, row)
}

func DeleteDatabase(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "Delete a registered database")
	redbeamsDbClient := ClientRedbeams(*oauth.NewRedbeamsClientFromContext(c)).Redbeams.Databases
	crn := c.String(fl.FlCrn.Name)
	if len(crn) != 0 {
		log.Infof("[DeleteDB] Deleting database with CRN: %s", crn)
		result, err := redbeamsDbClient.DeleteDatabaseByCrn(databases.NewDeleteDatabaseByCrnParams().WithCrn(crn), nil)
		if err != nil {
			commonutils.LogErrorAndExit(err)
		}
		log.Infof("[DeleteDB] Deleted database with CRN: %s Details: %s", crn, result)
	} else {
		envCrn := c.String(fl.FlEnvironmentCrn.Name)
		name := c.String(fl.FlName.Name)
		log.Infof("[DeleteDB] Deleting database with name: %s", name)
		result, err := redbeamsDbClient.DeleteDatabaseByName(databases.NewDeleteDatabaseByNameParams().WithEnvironmentCrn(envCrn).WithName(name), nil)
		if err != nil {
			commonutils.LogErrorAndExit(err)
		}
		log.Infof("[DeleteDB] Deleted database with name: %s Details: %s", name, result)
	}
}
