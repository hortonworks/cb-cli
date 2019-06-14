package databaseserverconfig

import (
	"github.com/hortonworks/cb-cli/dataplane/oauth"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type ClientRedbeams oauth.Redbeams

func StubFunc(c *cli.Context) {
	// an empty func
	log.Errorf("Not implemented yet")
}
