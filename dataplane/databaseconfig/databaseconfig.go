package databaseconfig

import (
	"github.com/hortonworks/cb-cli/dataplane/oauth"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type ClientRedbeams oauth.Redbeams

func EmptyFunc(c *cli.Context) {
	// this is empty
	log.Errorf("Not implemented yet ")
}
