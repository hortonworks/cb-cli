package cmd

import (
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/aws"
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/azure"
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/gcp"
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/openstack"
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/yarn"
	cf "github.com/hortonworks/cb-cli/dataplane/config"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/freeipa"
	"github.com/urfave/cli"
)

func init() {
	DataPlaneCommands = append(DataPlaneCommands, cli.Command{
		Name:  "freeipa",
		Usage: "create FreeIpa clusters",
		Subcommands: []cli.Command{
			{
				Name:        "create",
				Usage:       "creates a new FreeIpa cluster",
				Description: `basic FreeIpa cluster creation`,
				Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:       fl.NewFlagBuilder().AddFlags(fl.FlNameOptional, fl.FlInputJson).AddAGlobalFlags().AddOutputFlag().Build(),
				Action:      freeipa.CreateFreeIpa,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlNameOptional, fl.FlInputJson).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "delete",
				Usage:  "deletes a FreeIpa cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: freeipa.DeleteFreeIpa,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "describe",
				Usage:  "describes a FreeIpa cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: freeipa.DescribeFreeIpa,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:  "diagnostics",
				Usage: "manage dianostics for a FreeIpa cluster",
				Subcommands: []cli.Command{
					{
						Name:        "collect",
						Usage:       "collect FreeIpa diagnostics",
						Description: `collect FreeIpa diagnostics`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags: fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlCollectionDestination, fl.FlCollectionOnly,
							fl.FlCollectionLabels, fl.FlDescriptionOptional, fl.FlCollectionIssue, fl.FlCollectionExtraLogsFile,
							fl.FlCollectionHosts, fl.FlCollectionHostGroups, fl.FlIncludeSaltLogs, fl.FlUpdatePackage).AddAGlobalFlags().Build(),
						Action: freeipa.CollectDiagnostics,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "logs",
						Usage:       "list default monitored vm logs for FreeIpa clusters",
						Description: `list default monitored vm logs for FreeIpa clusters`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddAGlobalFlags().Build(),
						Action:      freeipa.GetVmLogs,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "list-collections",
						Usage:       "list latest diagnostics collection flows",
						Description: `list latest diagnostics collection flows`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().Build(),
						Action:      freeipa.ListDiagnosticsCollections,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "cancel-collections",
						Usage:       "cancel running diagnostics collection flows",
						Description: `cancel running diagnostics collection flows`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().Build(),
						Action:      freeipa.CancelDiagnosticsCollections,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:   "start",
				Usage:  "start a FreeIpa cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: freeipa.StartFreeIpa,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "stop",
				Usage:  "stop a FreeIpa cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: freeipa.StopFreeIpa,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "reboot",
				Usage:  "reboot a FreeIpa cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlNodesOptional, fl.FlForceOptional, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: freeipa.RebootFreeIpa,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlNodesOptional, fl.FlForceOptional, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "repair",
				Usage:  "repair a FreeIpa cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlNodesOptional, fl.FlForceOptional, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: freeipa.RepairFreeIpa,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlNodesOptional, fl.FlForceOptional, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "upscale",
				Usage:  "upscales a FreeIpa cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlScaleTargetFormFactor).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: freeipa.UpscaleFreeIpa,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlScaleTargetFormFactor).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "list",
				Usage:  "list FreeIpa clusters",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build(),
				Action: freeipa.ListFreeIpa,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:  "user",
				Usage: "manage users on FreeIpa cluster",
				Subcommands: []cli.Command{
					{
						Name:        "syncall",
						Usage:       "syncs all users to FreeIpa clusters",
						Description: `syncs all users to FreeIpa clusters`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddFlags(fl.FlIpaUserCrnsSlice, fl.FlIpaMachineUserCrnsSlice, fl.FlIpaEnvironmentCrnsSlice, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build(),
						Action:      freeipa.SynchronizeAllUsers,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlIpaUserCrnsSlice, fl.FlIpaMachineUserCrnsSlice, fl.FlIpaEnvironmentCrnsSlice, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "sync",
						Usage:       "syncs current user to FreeIpa clusters",
						Description: `syncs current user to FreeIpa clusters`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddFlags(fl.FlIpaEnvironmentCrnsSlice, fl.FlIpaEnvironmentNamesOptionalSlice, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build(),
						Action:      freeipa.SynchronizeCurrentUser,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlIpaEnvironmentCrnsSlice, fl.FlIpaEnvironmentNamesOptionalSlice, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "passwd",
						Usage:       "sets password in FreeIpa clusters",
						Description: `sets password in FreeIpa clusters`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddFlags(fl.FlIpaUserPassword, fl.FlIpaEnvironmentCrnsSlice, fl.FlIpaEnvironmentNamesOptionalSlice, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build(),
						Action:      freeipa.SetPassword,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlIpaUserPassword, fl.FlIpaEnvironmentCrnsSlice, fl.FlIpaEnvironmentNamesOptionalSlice, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "status",
						Usage:       "gets operation status",
						Description: `gets operation status`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddFlags(fl.FlIpaSyncOperationId, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build(),
						Action:      freeipa.GetSyncOperationStatus,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlIpaSyncOperationId, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "state",
						Usage:       "gets user synchronization state for an environment",
						Description: `gets user synchronization state for an environment`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().AddOutputFlag().Build(),
						Action:      freeipa.GetEnvironmentSyncState,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:  "dns",
				Usage: "Manages DNS related resources",
				Subcommands: []cli.Command{
					{
						Name:        "add-a-record",
						Usage:       "Adds an A record to FreeIPA",
						Description: "Adds an A record to FreeIPA",
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlDnsZone, fl.FlDnsHostname, fl.FlDnsIp, fl.FlDnsCreateReverse).AddAGlobalFlags().AddOutputFlag().Build(),
						Action:      freeipa.AddDnsARecord,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlDnsZone, fl.FlDnsHostname, fl.FlDnsIp, fl.FlDnsCreateReverse).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "add-cname-record",
						Usage:       "Adds a CNAME record to FreeIPA",
						Description: "Adds a CNAME record to FreeIPA",
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlDnsZone, fl.FlDnsCname, fl.FlDnsTargetFqdn).AddAGlobalFlags().AddOutputFlag().Build(),
						Action:      freeipa.AddDnsCnameRecord,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlDnsZone, fl.FlDnsCname, fl.FlDnsTargetFqdn).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "delete-a-record",
						Usage:       "Deletes an A record from FreeIPA",
						Description: "Deletes an A record from FreeIPA",
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlDnsZone, fl.FlDnsHostname).AddAGlobalFlags().AddOutputFlag().Build(),
						Action:      freeipa.DeleteDnsARecord,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlDnsZone, fl.FlDnsHostname).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "delete-cname-record",
						Usage:       "Deletes a CNAME record from FreeIPA",
						Description: "Deletes a CNAME record from FreeIPA",
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlDnsZone, fl.FlDnsCname).AddAGlobalFlags().AddOutputFlag().Build(),
						Action:      freeipa.DeleteDnsCnameRecord,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlDnsZone, fl.FlDnsCname).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:   "change-image",
				Usage:  "changes image of freeipa - will be used when creating new instances or repairing failed ones",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlImageIdOptional, fl.FlImageCatalogOptional).AddAGlobalFlags().Build(),
				Action: freeipa.ChangeImage,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlImageIdOptional, fl.FlImageCatalogOptional).AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "update-salt",
				Usage:  "update salt states on freeipa",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: freeipa.UpdateSalt,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:  "operations",
				Usage: "check environment operations",
				Subcommands: []cli.Command{
					{
						Name:   "last",
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().Build(),
						Usage:  "check progress of last freeipa operation",
						Before: cf.CheckConfigAndCommandFlags,
						Action: freeipa.OperationProgress,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:   "upgrade",
				Usage:  "upgrades freeipa to new image",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlImageCatalogOptional, fl.FlImageIdOptional).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: freeipa.Upgrade,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlImageCatalogOptional, fl.FlImageIdOptional).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "upgrade-options",
				Usage:  "list available images for upgrade",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlImageCatalogOptional).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: freeipa.UpgradeOptions,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName, fl.FlImageCatalogOptional).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "retry",
				Usage:  "retry last failed flow if it's retryable",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: freeipa.Retry,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "list-retryable-flows",
				Usage:  "List retryable failed flows",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: freeipa.ListRetryableFlows,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentName).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
		},
	})
}
