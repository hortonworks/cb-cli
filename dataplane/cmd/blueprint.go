package cmd

import (
	bp "github.com/hortonworks/cb-cli/dataplane/blueprint"
	cf "github.com/hortonworks/cb-cli/dataplane/config"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/urfave/cli"
)

func init() {
	DataPlaneCommands = append(DataPlaneCommands, cli.Command{
		Name:  "blueprint",
		Usage: "blueprint related operations",
		Subcommands: []cli.Command{
			{
				Name:  "create",
				Usage: "adds a new blueprint from a file or from a URL",
				Subcommands: []cli.Command{
					{
						Name:   "from-url",
						Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlURL, fl.FlDlOptional).AddAGlobalFlags().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: bp.CreateBlueprintFromUrl,
						Usage:  "creates a blueprint by downloading it from a URL location",
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlURL, fl.FlDlOptional).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:   "from-file",
						Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlFile, fl.FlDlOptional).AddAGlobalFlags().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: bp.CreateBlueprintFromFile,
						Usage:  "creates a blueprint by reading it from a local file",
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlFile, fl.FlDlOptional).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:   "delete",
				Usage:  "deletes one or more blueprints",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlNames).AddAGlobalFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: bp.DeleteBlueprints,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlNames).AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "describe",
				Usage:  "describes a blueprint",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: bp.DescribeBlueprint,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "list",
				Usage:  "lists the available blueprints",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlWithSdxOptional).AddAGlobalFlags().AddOutputFlag().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: bp.ListBlueprints,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlWithSdxOptional).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
		},
	})
}
