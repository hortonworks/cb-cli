package cmd

import (
	cf "github.com/hortonworks/cb-cli/dataplane/config"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/telemetry"
	"github.com/urfave/cli"
)

func init() {
	DataPlaneCommands = append(DataPlaneCommands, cli.Command{
		Name:  "telemetry",
		Usage: "account telemetry related operations",
		Subcommands: []cli.Command{
			{
				Name:        "update",
				Usage:       "update account telemetry settings",
				Description: "update or create account level telemetry settings.",
				Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:       fl.NewFlagBuilder().AddFlags(fl.FlFile).AddAuthenticationFlags().Build(),
				Action:      telemetry.UpdateAccountTelemetry,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlFile).AddAuthenticationFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:        "describe",
				Usage:       "describe account telemetry settings",
				Description: "describe account level telemetry settings (feature flags and anonymization rules)",
				Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:       fl.NewFlagBuilder().AddAuthenticationFlags().Build(),
				Action:      telemetry.DescribeAccountTelemetry,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddAuthenticationFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
		},
	})
}
