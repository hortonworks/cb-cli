package cmd

import (
	cf "github.com/hortonworks/cb-cli/cloudbreak/config"
	fl "github.com/hortonworks/cb-cli/cloudbreak/flags"
	"github.com/hortonworks/cb-cli/cloudbreak/mpack"
	"github.com/urfave/cli"
)

func init() {
	CloudbreakCommands = append(CloudbreakCommands, cli.Command{
		Name:  "mpack",
		Usage: "management pack related operations",
		Subcommands: []cli.Command{
			{
				Name:  "create",
				Usage: "create a new management pack",
				Flags: fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FLMpackURL,
					fl.FLMpackPurge, fl.FLMpackPurgeList, fl.FLMpackForce).AddAuthenticationFlags().Build(),
				Before: cf.ConfigRead,
				Action: mpack.CreateMpack,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FLMpackURL,
						fl.FLMpackPurge, fl.FLMpackPurgeList, fl.FLMpackForce).AddAuthenticationFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "delete",
				Usage:  "deletes a management pack",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAuthenticationFlags().Build(),
				Before: cf.ConfigRead,
				Action: mpack.DeleteMpack,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAuthenticationFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "list",
				Usage:  "list the available management packs",
				Flags:  fl.NewFlagBuilder().AddOutputFlag().AddAuthenticationFlags().Build(),
				Before: cf.ConfigRead,
				Action: mpack.ListMpacks,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddOutputFlag().AddAuthenticationFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
		},
	})
}
