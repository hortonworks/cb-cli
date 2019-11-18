package cmd

import (
	cf "github.com/hortonworks/cb-cli/dataplane/config"
	"github.com/hortonworks/cb-cli/dataplane/env"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/urfave/cli"
)

func init() {
	DataPlaneCommands = append(DataPlaneCommands, cli.Command{
		Name:  "env",
		Usage: "environment related operations",
		Subcommands: []cli.Command{
			{
				Name:  "create",
				Usage: "creates a new Environment",
				Subcommands: []cli.Command{
					{
						Name:   "from-file",
						Usage:  "creates a new Environment from JSON template",
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentTemplateFile, fl.FlNameOptional).AddAuthenticationFlags().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: env.CreateEnvironmentFromTemplate,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentTemplateFile, fl.FlNameOptional).AddAuthenticationFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:   "list",
				Usage:  "list the available environments",
				Flags:  fl.NewFlagBuilder().AddOutputFlag().AddAuthenticationFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: env.ListEnvironments,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddOutputFlag().AddAuthenticationFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "describe",
				Usage:  "describes an environment",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddOutputFlag().AddAuthenticationFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: env.DescribeEnvironment,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddOutputFlag().AddAuthenticationFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "stop",
				Usage:  "stop an environment and all related datahubs, datalake and freeipa",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAuthenticationFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: env.StopEnvironment,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAuthenticationFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "start",
				Usage:  "start an environment and all related datahubs, datalake and freeipa",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAuthenticationFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: env.StartEnvironment,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAuthenticationFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "delete",
				Usage:  "deletes an environment or more if names are spearated by commas",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlNames).AddOutputFlag().AddAuthenticationFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: env.DeleteEnvironment,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlNames).AddOutputFlag().AddAuthenticationFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "change-cred",
				Usage:  "change the credential of an environment. also changes the credential of the clusters in the environment.",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredential).AddOutputFlag().AddAuthenticationFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: env.ChangeCredential,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredential).AddOutputFlag().AddAuthenticationFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:  "edit",
				Usage: "edit an environment. description, network, regions and location can be changed.",
				Subcommands: []cli.Command{
					{
						Name:   "from-file",
						Usage:  "edits an Environment from JSON template",
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentEditTemplateFile, fl.FlName).AddAuthenticationFlags().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: env.EditEnvironmentFromTemplate,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentEditTemplateFile, fl.FlName).AddAuthenticationFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:  "generate-template",
				Usage: "creates an environment JSON template",
				Subcommands: []cli.Command{
					{
						Name:  "aws",
						Usage: "creates an aws specific environment JSON template",
						Subcommands: []cli.Command{
							{
								Name:   "use-existing-network",
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlNetworkId, fl.FlSubnetIds).AddAuthenticationFlags().Build(),
								Usage:  "attach existing vpc and subnets",
								Before: cf.CheckConfigAndCommandFlags,
								Action: env.GenerateAwsEnvironmentTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlNetworkId, fl.FlSubnetIds).AddAuthenticationFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "create-new-network",
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlNetworkCidr, fl.FlSubnetCidrs).AddAuthenticationFlags().Build(),
								Usage:  "create new vpc and subnets",
								Before: cf.CheckConfigAndCommandFlags,
								Action: env.GenerateAwsEnvironmentTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlNetworkCidr, fl.FlSubnetCidrs).AddAuthenticationFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
					{
						Name:  "azure",
						Usage: "creates an azure specific environment JSON template",
						Subcommands: []cli.Command{
							{
								Name:   "use-existing-network",
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlNetworkId, fl.FlSubnetIds, fl.FlResourceGroupName).AddAuthenticationFlags().Build(),
								Usage:  "attach existing network and subnets",
								Before: cf.CheckConfigAndCommandFlags,
								Action: env.GenerateAzureEnvironmentTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlNetworkId, fl.FlSubnetIds, fl.FlResourceGroupName).AddAuthenticationFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "create-new-network",
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlNetworkCidr, fl.FlSubnetCidrs).AddAuthenticationFlags().Build(),
								Usage:  "create new network and subnets",
								Before: cf.CheckConfigAndCommandFlags,
								Action: env.GenerateAzureEnvironmentTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlNetworkCidr, fl.FlSubnetCidrs).AddAuthenticationFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
				},
			},
		},
	})
}
