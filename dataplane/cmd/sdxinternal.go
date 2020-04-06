package cmd

import (
	cf "github.com/hortonworks/cb-cli/dataplane/config"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/sdx"
	"github.com/urfave/cli"
)

func init() {
	DataPlaneCommands = append(DataPlaneCommands, cli.Command{
		Name:   "sdx-internal",
		Usage:  "create internal SDX clusters",
		Hidden: true,
		Subcommands: []cli.Command{
			{
				Name:        "create",
				Usage:       "creates a new SDX cluster",
				Description: `basic SDX cluster creation`,
				Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:       fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlEnvironmentName, fl.FlInputJson, fl.FlWithExternalDatabaseOptional, fl.FlWithNonHaExternalDatabaseOptional, fl.FlWithoutExternalDatabaseOptional, fl.FlRuntimeOptional, fl.FlCloudStorageBaseLocationOptional, fl.FlCloudStorageInstanceProfileOptional, fl.FlSpotPercentage).AddAGlobalFlags().Build(),
				Action:      sdx.CreateSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddResourceFlagsWithOptionalName().AddFlags(fl.FlEnvironmentName, fl.FlEnvironmentName, fl.FlInputJson, fl.FlWithExternalDatabaseOptional, fl.FlWithNonHaExternalDatabaseOptional, fl.FlWithoutExternalDatabaseOptional, fl.FlRuntimeOptional, fl.FlCloudStorageBaseLocationOptional, fl.FlCloudStorageInstanceProfileOptional, fl.FlSpotPercentage).AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
		},
	})
}
