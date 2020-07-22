package cmd

import (
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/aws"
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/azure"
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/gcp"
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/openstack"
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/yarn"
	cf "github.com/hortonworks/cb-cli/dataplane/config"
	cb "github.com/hortonworks/cb-cli/dataplane/connectors"
	"github.com/hortonworks/cb-cli/dataplane/env"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/urfave/cli"
)

func init() {
	DataPlaneCommands = append(DataPlaneCommands, cli.Command{
		Name:  "cloud",
		Usage: "information about cloud provider resources",
		Subcommands: []cli.Command{
			{
				Name:   "availability-zones",
				Usage:  "lists the available availabilityzones in a region",
				Flags:  fl.NewFlagBuilder().AddAGlobalFlags().AddFlags(fl.FlCredential, fl.FlRegion).AddOutputFlag().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: cb.ListAvailabilityZones,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddAGlobalFlags().AddFlags(fl.FlCredential, fl.FlRegion).AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "resource-groups",
				Usage:  "lists the available resource groups",
				Flags:  fl.NewFlagBuilder().AddAGlobalFlags().AddFlags(fl.FlCredential).AddOutputFlag().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: env.GetResourceGroups,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddAGlobalFlags().AddFlags(fl.FlCredential).AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "instances",
				Usage:  "lists the available instance types",
				Flags:  fl.NewFlagBuilder().AddAGlobalFlags().AddFlags(fl.FlCredential, fl.FlRegion, fl.FlAvailabilityZoneOptional).AddOutputFlag().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: cb.ListInstanceTypes,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddAGlobalFlags().AddFlags(fl.FlCredential, fl.FlRegion, fl.FlAvailabilityZoneOptional).AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "regions",
				Usage:  "lists the available regions",
				Flags:  fl.NewFlagBuilder().AddAGlobalFlags().AddFlags(fl.FlCredential).AddOutputFlag().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: cb.ListRegions,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddAGlobalFlags().AddFlags(fl.FlCredential).AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:  "volumes",
				Usage: "lists the available volume types",
				Subcommands: []cli.Command{
					{
						Name:   "aws",
						Usage:  "list the available aws volume types",
						Flags:  fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: cb.ListAwsVolumeTypes,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:   "azure",
						Usage:  "list the available azure volume types",
						Flags:  fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: cb.ListAzureVolumeTypes,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:   "gcp",
						Usage:  "list the available gcp volume types",
						Flags:  fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: cb.ListGcpVolumeTypes,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
		},
	})
}
