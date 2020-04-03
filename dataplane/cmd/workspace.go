package cmd

import (
	cf "github.com/hortonworks/cb-cli/dataplane/config"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/workspace"
	"github.com/urfave/cli"
)

func init() {
	DataPlaneCommands = append(DataPlaneCommands, cli.Command{
		Name:  "workspace",
		Usage: "workspace related operations",
		Subcommands: []cli.Command{
			{
				Name:   "list",
				Usage:  "list workspaces",
				Flags:  fl.NewFlagBuilder().AddOutputFlag().AddGlobalFlagsWithoutWorkspace().Build(),
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Action: workspace.ListWorkspaces,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddOutputFlag().AddGlobalFlagsWithoutWorkspace().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "describe",
				Usage:  "describes a workspace",
				Flags:  fl.NewFlagBuilder().AddOutputFlag().AddFlags(fl.FlName).AddGlobalFlagsWithoutWorkspace().Build(),
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Action: workspace.DescribeWorkspace,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddOutputFlag().AddFlags(fl.FlName).AddGlobalFlagsWithoutWorkspace().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
		},
	})
}
