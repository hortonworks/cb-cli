package cmd

import (
	cf "github.com/hortonworks/cb-cli/dataplane/config"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/flow"
	"github.com/urfave/cli"
)

func init() {
	DataPlaneCommands = append(DataPlaneCommands, cli.Command{
		Name:   "flow",
		Usage:  "get information about flows",
		Hidden: true,
		Subcommands: []cli.Command{
			{
				Name:  "describelast",
				Usage: "describes the last flow",
				Subcommands: []cli.Command{
					{
						Name:   "by-crn",
						Usage:  "describe the last flow by resource crn",
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAuthenticationFlags().AddOutputFlag().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: flow.DescribeLastFlowByCrn,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAuthenticationFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:   "by-name",
						Usage:  "describe the last flow by resource name",
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAuthenticationFlags().AddOutputFlag().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: flow.DescribeLastFlowByName,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAuthenticationFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:  "list",
				Usage: "list of flows",
				Subcommands: []cli.Command{
					{
						Name:   "by-crn",
						Usage:  "list flows by resource crn",
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAuthenticationFlags().AddOutputFlag().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: flow.ListByCrn,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAuthenticationFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:   "by-name",
						Usage:  "list flows by resource name",
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAuthenticationFlags().AddOutputFlag().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: flow.ListByName,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAuthenticationFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
		},
	})
}
