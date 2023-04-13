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
				Flags:       fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlEnvironmentName, fl.FlClusterShape, fl.FlCloudStorageBaseLocationOptional, fl.FlCloudStorageInstanceProfileOptional, fl.FlCloudStorageManagedIdentityOptional, fl.FlCloudStorageServiceAccountOptional, fl.FlRuntimeOptional, fl.FlCidrOptional, fl.FlWithoutExternalDatabaseOptional, fl.FlWithExternalDatabaseOptional, fl.FlWithNonHaExternalDatabaseOptional, fl.FlRdsDatabaseEngineOptional, fl.FlSpotPercentage, fl.FlSpotMaxPrice, fl.FlRangerRazEnabled, fl.FlEnableMultiAz, fl.FlAzureDatabaseFlexibleOptional).AddAGlobalFlags().Build(),
				Action:      sdx.CreateSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlEnvironmentName, fl.FlEnvironmentName, fl.FlClusterShape, fl.FlCloudStorageBaseLocationOptional, fl.FlCloudStorageInstanceProfileOptional, fl.FlCloudStorageManagedIdentityOptional, fl.FlCloudStorageServiceAccountOptional, fl.FlRuntimeOptional, fl.FlCidrOptional, fl.FlWithoutExternalDatabaseOptional, fl.FlWithExternalDatabaseOptional, fl.FlWithNonHaExternalDatabaseOptional, fl.FlRdsDatabaseEngineOptional, fl.FlSpotPercentage, fl.FlSpotMaxPrice, fl.FlRangerRazEnabled, fl.FlEnableMultiAz, fl.FlAzureDatabaseFlexibleOptional).AddAGlobalFlags().Build() {
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
				Usage:  "OS or data platform upgrade for the SDX cluster. You need to specify at least either one of imageId, runtime or lockComponents to be able to proceed!",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().AddFlags(fl.FlImageIdOptional).AddFlags(fl.FlRuntimeOptional).AddFlags(fl.FlLockComponentsOptional).AddFlags(fl.FlRollingUpgradeOptional).AddFlags(fl.FlDryRunOptional).AddFlags(fl.FlReplaceVms).AddFlags(fl.FlShowImagesOptional).AddFlags(fl.FlShowLatestImagesOptional).AddFlags(fl.FlUpgradeBackup).Build(),
				Action: sdx.SdxClusterUpgrade,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().AddFlags(fl.FlImageIdOptional).AddFlags(fl.FlRuntimeOptional).AddFlags(fl.FlLockComponentsOptional).AddFlags(fl.FlRollingUpgradeOptional).AddFlags(fl.FlDryRunOptional).AddFlags(fl.FlReplaceVms).AddFlags(fl.FlShowImagesOptional).AddFlags(fl.FlShowLatestImagesOptional).AddFlags(fl.FlUpgradeBackup).Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "prepare-upgrade",
				Usage:  "OS or data platform upgrade preparation for the SDX cluster. You need to specify at least either one of imageId, runtime to be able to proceed!",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().AddFlags(fl.FlImageIdOptional).AddFlags(fl.FlRuntimeOptional).AddFlags(fl.FlDryRunOptional).AddFlags(fl.FlShowImagesOptional).AddFlags(fl.FlShowLatestImagesOptional).Build(),
				Action: sdx.SdxClusterUpgradePrepare,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().AddFlags(fl.FlImageIdOptional).AddFlags(fl.FlRuntimeOptional).AddFlags(fl.FlDryRunOptional).AddFlags(fl.FlShowImagesOptional).AddFlags(fl.FlShowLatestImagesOptional).Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "vertical-scale",
				Usage:  "vertical scales a SDX cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlGroupName, fl.FlWaitOptional, fl.FlInstanceType).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: sdx.VerticalScaleSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlGroupName, fl.FlWaitOptional, fl.FlInstanceType).AddAGlobalFlags().AddOutputFlag().Build() {
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
						Usage:       "collect SDX diagnostics",
						Description: `collect SDX diagnostics`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags: fl.NewFlagBuilder().AddFlags(fl.FlCrn, fl.FlCollectionDestination, fl.FlCollectionOnly,
							fl.FlCollectionLabels, fl.FlDescriptionOptional, fl.FlCollectionIssue, fl.FlCollectionExtraLogsFile,
							fl.FlCollectionHosts, fl.FlCollectionHostGroups, fl.FlIncludeSaltLogs, fl.FlUpdatePackage).AddAGlobalFlags().Build(),
						Action: sdx.CollectDiagnostics,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "list-collections",
						Usage:       "list latest SDX diagnostics collection flows",
						Description: `list latest SDX diagnostics collection flows`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().Build(),
						Action:      sdx.ListDiagnosticsCollections,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "cancel-collections",
						Usage:       "cancel running SDX diagnostics collection flows",
						Description: `cancel running SDX diagnostics collection flows`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().Build(),
						Action:      sdx.CancelDiagnosticsCollections,
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
					{
						Name:        "create-cm-bundle",
						Usage:       "start CM based diagnostics collection for SDX",
						Description: `start CM based diagnostics collection for SDX`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags: fl.NewFlagBuilder().AddFlags(fl.FlCrn, fl.FlCollectionDestination, fl.FlDescriptionOptional, fl.FlCollectionIssue, fl.FlCollectionOnly,
							fl.FlCollectionRoles, fl.FlMonitorMetricsCollection, fl.FlUpdatePackage).AddAGlobalFlags().Build(),
						Action: sdx.CollectCmDiagnostics,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "cm-roles",
						Usage:       "get available roles for SDX CM cluster",
						Description: `get available roles for SDX CM cluster`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().Build(),
						Action:      sdx.GetCMRoles,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:  "operations",
				Usage: "check sdx operations",
				Subcommands: []cli.Command{
					{
						Name:   "last",
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().Build(),
						Usage:  "check progress of the last sdx operation",
						Before: cf.CheckConfigAndCommandFlags,
						Action: sdx.OperationProgress,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:   "rotate-certificates",
				Usage:  "rotate SDX AutoTLS certificates",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().Build(),
				Action: sdx.RotateCertificates,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "recover",
				Usage:  "recover command after failed upgrade to the original version",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().AddFlags(fl.FlRecoverWithoutDataOptional).Build(),
				Action: sdx.SdxClusterRecover,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().AddFlags(fl.FlRecoverWithoutDataOptional).Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "upgrade-database",
				Usage:  "Upgrade database for cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().AddFlags(fl.FlTargetVersionOptional).Build(),
				Action: sdx.UpgradeDatabase,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().AddFlags(fl.FlTargetVersionOptional).Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "sync-component-versions",
				Usage:  "sync SDX cluster component version from CM",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().Build(),
				Action: sdx.SyncComponentVersions,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
		},
	})
}
