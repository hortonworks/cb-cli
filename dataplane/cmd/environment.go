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
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentTemplateFile, fl.FlNameOptional).AddAGlobalFlags().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: env.CreateEnvironmentFromTemplate,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentTemplateFile, fl.FlNameOptional).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:   "list",
				Usage:  "list the available environments",
				Flags:  fl.NewFlagBuilder().AddOutputFlag().AddAGlobalFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: env.ListEnvironments,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddOutputFlag().AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "describe",
				Usage:  "describes an environment",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddOutputFlag().AddAGlobalFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: env.DescribeEnvironment,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddOutputFlag().AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "stop",
				Usage:  "stop an environment and all related datahubs, datalake and freeipa",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: env.StopEnvironment,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "start",
				Usage:  "start an environment and all related datahubs, datalake and freeipa",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: env.StartEnvironment,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "delete",
				Usage:  "deletes an environment or more if names are separated by commas",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlNames).AddFlags(fl.FlForceOptional).AddFlags(fl.FlCascadeOptional).AddOutputFlag().AddAGlobalFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: env.DeleteEnvironment,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlNames).AddFlags(fl.FlForceOptional).AddFlags(fl.FlCascadeOptional).AddOutputFlag().AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "verify-policy",
				Usage:  "verify policy an environment or more if names are separated by commas",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddFlags(fl.FlServiceNames).AddOutputFlag().AddAGlobalFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: env.VerifyPolicyForEnvironment,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddFlags(fl.FlServiceNames).AddOutputFlag().AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "change-cred",
				Usage:  "change the credential of an environment. also changes the credential of the clusters in the environment.",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredential).AddOutputFlag().AddAGlobalFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: env.ChangeCredential,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredential).AddOutputFlag().AddAGlobalFlags().Build() {
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
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentEditTemplateFile, fl.FlName).AddAGlobalFlags().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: env.EditEnvironmentFromTemplate,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentEditTemplateFile, fl.FlName).AddAGlobalFlags().Build() {
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
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlNetworkId, fl.FlSubnetIds).AddAGlobalFlags().Build(),
								Usage:  "attach existing vpc and subnets",
								Before: cf.CheckConfigAndCommandFlags,
								Action: env.GenerateAwsEnvironmentTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlNetworkId, fl.FlSubnetIds).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "create-new-network",
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlNetworkCidr, fl.FlSubnetCidrs).AddAGlobalFlags().Build(),
								Usage:  "create new vpc and subnets",
								Before: cf.CheckConfigAndCommandFlags,
								Action: env.GenerateAwsEnvironmentTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlNetworkCidr, fl.FlSubnetCidrs).AddAGlobalFlags().Build() {
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
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlNetworkId, fl.FlSubnetIds, fl.FlResourceGroupName).AddAGlobalFlags().Build(),
								Usage:  "attach existing network and subnets",
								Before: cf.CheckConfigAndCommandFlags,
								Action: env.GenerateAzureEnvironmentTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlNetworkId, fl.FlSubnetIds, fl.FlResourceGroupName).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "create-new-network",
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlNetworkCidr, fl.FlSubnetCidrs).AddAGlobalFlags().Build(),
								Usage:  "create new network and subnets",
								Before: cf.CheckConfigAndCommandFlags,
								Action: env.GenerateAzureEnvironmentTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlNetworkCidr, fl.FlSubnetCidrs).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
				},
			},
			{
				Name:  "operations",
				Usage: "check environment operations",
				Subcommands: []cli.Command{
					{
						Name:   "last",
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().Build(),
						Usage:  "check progress of last environment operation",
						Before: cf.CheckConfigAndCommandFlags,
						Action: env.OperationProgress,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:  "telemetry",
				Usage: "environment level telemetry related operations.",
				Subcommands: []cli.Command{
					{
						Name:  "features",
						Usage: "Environment level telemetry feature operations",
						Subcommands: []cli.Command{
							{
								Name:   "edit",
								Usage:  "Update environment level telemetry features",
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlEnvironmentEditTemplateFile).AddAGlobalFlags().Build(),
								Before: cf.CheckConfigAndCommandFlags,
								Action: env.EditEnvironmentTelemetryFeaturesFromTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().Build() {
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
