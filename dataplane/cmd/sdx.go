package cmd

import (
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/aws"
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/azure"
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/gcp"
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/openstack"
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/yarn"
	cf "github.com/hortonworks/cb-cli/dataplane/config"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/sdx"
	"github.com/urfave/cli"
)

func init() {
	DataPlaneCommands = append(DataPlaneCommands, cli.Command{
		Name:  "sdx",
		Usage: "create SDX clusters",
		Subcommands: []cli.Command{
			{
				Name:   "runtimes",
				Usage:  "runtimes an SDX cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build(),
				Action: sdx.ListSdxAdvertisedRuntimesSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:        "create",
				Usage:       "creates a new SDX cluster",
				Description: `basic SDX cluster creation`,
				Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:       fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlEnvironmentName, fl.FlClusterShape, fl.FlCloudStorageBaseLocationOptional, fl.FlCloudStorageInstanceProfileOptional, fl.FlCloudStorageManagedIdentityOptional, fl.FlRuntimeOptional, fl.FlCidrOptional, fl.FlWithoutExternalDatabaseOptional, fl.FlWithExternalDatabaseOptional, fl.FlWithNonHaExternalDatabaseOptional, fl.FlSpotPercentage, fl.FlSpotMaxPrice, fl.FlRangerRazEnabled).AddAGlobalFlags().Build(),
				Action:      sdx.CreateSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlEnvironmentName, fl.FlEnvironmentName, fl.FlClusterShape, fl.FlCloudStorageBaseLocationOptional, fl.FlCloudStorageInstanceProfileOptional, fl.FlCloudStorageManagedIdentityOptional, fl.FlRuntimeOptional, fl.FlCidrOptional, fl.FlWithoutExternalDatabaseOptional, fl.FlWithExternalDatabaseOptional, fl.FlWithNonHaExternalDatabaseOptional, fl.FlSpotPercentage, fl.FlSpotMaxPrice, fl.FlRangerRazEnabled).AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "delete",
				Usage:  "deletes an SDX cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlForceOptional).AddNameFlag().AddAGlobalFlags().Build(),
				Action: sdx.DeleteSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlForceOptional).AddNameFlag().AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "describe",
				Usage:  "describes an SDX cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlDetailedOptional).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: sdx.DescribeSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlDetailedOptional).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "start",
				Usage:  "start an SDX cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().AddOutputFlag().Build(),
				Action: sdx.StartSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "stop",
				Usage:  "stop an SDX cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().AddOutputFlag().Build(),
				Action: sdx.StopSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "list",
				Usage:  "list SDX clusters",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentNameOptional).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: sdx.ListSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentNameOptional).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "repair",
				Usage:  "repair SDX cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlHostGroupOptional, fl.FlHostGroupsOptional).AddNameFlag().AddAGlobalFlags().Build(),
				Action: sdx.RepairSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlHostGroupOptional, fl.FlHostGroupsOptional).AddNameFlag().AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "sync",
				Usage:  "sync SDX cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().Build(),
				Action: sdx.SyncSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "retry",
				Usage:  "retry SDX cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().Build(),
				Action: sdx.RetrySdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "upgrade",
				Usage:  "OS or data platform upgrade for the SDX cluster. You need to specify at least either one of imageId, runtime or lockComponents and replaceVms value to proceed!",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().AddFlags(fl.FlImageIdOptional).AddFlags(fl.FlRuntimeOptional).AddFlags(fl.FlLockComponentsOptional).AddFlags(fl.FlDryRunOptional).AddFlags(fl.FlReplaceVms).Build(),
				Action: sdx.SdxClusterkUpgrade,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().AddFlags(fl.FlImageIdOptional).AddFlags(fl.FlRuntimeOptional).AddFlags(fl.FlLockComponentsOptional).AddFlags(fl.FlDryRunOptional).AddFlags(fl.FlReplaceVms).Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:  "diagnostics",
				Usage: "manage diagnostics for an SDX CM cluster",
				Subcommands: []cli.Command{
					{
						Name:        "collect",
						Usage:       "collect SDX CM diagnostics",
						Description: `collect SDX CM diagnostics`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags: fl.NewFlagBuilder().AddFlags(fl.FlCrn, fl.FlCollectionDestination, fl.FlCollectionOnly,
							fl.FlCollectionLabels, fl.FlDescriptionOptional, fl.FlCollectionIssue, fl.FlCollectionExtraLogsFile).AddAGlobalFlags().Build(),
						Action: sdx.CollectDiagnostics,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "logs",
						Usage:       "list default monitored vm logs for SDX CM clusters",
						Description: `list default monitored vm logs for SDX CM clusters`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddAGlobalFlags().Build(),
						Action:      sdx.GetVmLogs,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
		},
	})
}
