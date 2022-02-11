package cmd

import (
	cf "github.com/hortonworks/cb-cli/dataplane/config"
	"github.com/hortonworks/cb-cli/dataplane/distrox"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/urfave/cli"
)

var distroXTemplateDescription = `Template parameters to fill in the generated template:
		userName:	Name of the CM user
		password:	Password of the CM user
		name:	Name of the DistroX cluster
		region:	Region of the DistroX cluster
		availabilityZone:	Availability zone of the DistroX cluster, on AZURE it is the same as the region
		blueprintName:	Name of the selected blueprint
		credentialName:	Name of the selected credential
		instanceGroups.group:	Name of the instance group
		instanceGroups.nodeCount:	Number of nodes in the group
		instanceGroups.template.instanceType:	Name of the selected template
		instanceGroups.template.volumeCount:	Number of volumes
		instanceGroups.template.volumeSize:	Size of Volumes in Gb
		stackAuthentication.publicKey:	Public key
`

func init() {
	DataPlaneCommands = append(DataPlaneCommands, cli.Command{
		Name:  "distrox",
		Usage: "DistroX related operations",
		Subcommands: []cli.Command{
			{
				Name:   "change-image",
				Usage:  "changes image of the DistroX cluster - will be used when creating new instances or repairing failed ones",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlImageId, fl.FlImageCatalogOptional).AddAGlobalFlags().Build(),
				Action: distrox.ChangeImage,
				Hidden: true,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlImageId).AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:        "create",
				Usage:       "creates a new DistroX cluster",
				Description: `use 'dp cluster generate-template' for cluster request JSON generation`,
				Flags:       fl.NewFlagBuilder().AddResourceFlagsWithOptionalName().AddFlags(fl.FlInputJson, fl.FlCMUserOptional, fl.FlCMPasswordOptional, fl.FlWaitOptional).AddAGlobalFlags().Build(),
				Before:      cf.CheckConfigAndCommandFlags,
				Action:      distrox.CreateDistroX,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddResourceFlagsWithOptionalName().AddFlags(fl.FlInputJson, fl.FlCMUserOptional, fl.FlCMPasswordOptional, fl.FlWaitOptional).AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "delete",
				Usage:  "deletes a DistroX cluster",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlForceOptional, fl.FlWaitOptional).AddAGlobalFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: distrox.DeleteDistroX,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlForceOptional, fl.FlWaitOptional).AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "multi-delete",
				Usage:  "deletes multiple DistroX clusters by name",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlNames, fl.FlForceOptional).AddAGlobalFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: distrox.DeleteMultipleDistroxClusters,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlNames, fl.FlForceOptional).AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "describe",
				Usage:  "describes a DistroX cluster",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: distrox.DescribeDistroX,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "list",
				Usage:  "lists the running DistroX clusters",
				Flags:  fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: distrox.ListDistroXs,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:        "repair",
				Usage:       "repairs a DistroX",
				Description: "repairs a DistroX",
				Subcommands: []cli.Command{
					{
						Name:   "host-groups",
						Usage:  "repairs host-groups of a cluster",
						Before: cf.CheckConfigAndCommandFlags,
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlHostGroups, fl.FlRemoveOnly, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build(),
						Action: distrox.RepairDistroXHostGroups,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlHostGroups, fl.FlRemoveOnly, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:   "nodes",
						Usage:  "repairs nodes of a DistroX",
						Before: cf.CheckConfigAndCommandFlags,
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlNodes, fl.FlDeleteVolumes, fl.FlRemoveOnly, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build(),
						Action: distrox.RepairDistroXNodes,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlNodes, fl.FlDeleteVolumes, fl.FlRemoveOnly, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:   "retry",
				Usage:  "retries the creation of a cluster",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: distrox.RetryDistroX,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "scale",
				Usage:  "scales a DistroX cluster",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlGroupName, fl.FlDesiredNodeCount, fl.FlWaitOptional, fl.FlPreferredSubnetIds).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: distrox.ScaleDistroX,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlGroupName, fl.FlDesiredNodeCount, fl.FlWaitOptional, fl.FlPreferredSubnetIds).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "vertical-scale",
				Usage:  "vertical scales a DistroX cluster",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlGroupName, fl.FlWaitOptional, fl.FlInstanceType).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: distrox.VerticalScaleDistroX,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlGroupName, fl.FlWaitOptional, fl.FlInstanceType).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "start",
				Usage:  "starts a DistroX cluster",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: distrox.StartDistroX,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "stop",
				Usage:  "stops a DistroX cluster",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: distrox.StopDistroX,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "sync",
				Usage:  "synchronizes a DistroX cluster",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: distrox.SyncDistroX,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "get-request-from-name",
				Usage:  "gets the CDP CLI request json for DistroX cluster by name",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: distrox.GetRequestByName,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:  "diagnostics",
				Usage: "manage diagnostics for a DistroX CM cluster",
				Subcommands: []cli.Command{
					{
						Name:        "collect",
						Usage:       "collect DistroX cluster diagnostics",
						Description: `collect DistroX cluster diagnostics`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags: fl.NewFlagBuilder().AddFlags(fl.FlCrn, fl.FlCollectionDestination, fl.FlCollectionOnly,
							fl.FlCollectionLabels, fl.FlDescriptionOptional, fl.FlCollectionIssue, fl.FlCollectionExtraLogsFile,
							fl.FlCollectionHosts, fl.FlCollectionHostGroups, fl.FlIncludeSaltLogs, fl.FlUpdatePackage).AddAGlobalFlags().Build(),
						Action: distrox.CollectDiagnostics,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "list-collections",
						Usage:       "list latest DistroX diagnostics collection flows",
						Description: `list latest DistroX diagnostics collection flows`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().Build(),
						Action:      distrox.ListDiagnosticsCollections,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "cancel-collections",
						Usage:       "cancel runnong DistroX diagnostics collection flows",
						Description: `cancel running DistroX diagnostics collection flows`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().Build(),
						Action:      distrox.CancelDiagnosticsCollections,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "logs",
						Usage:       "list default monitored vm logs for DistroX CM clusters",
						Description: `list default monitored vm logs for DistroX CM clusters`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddAGlobalFlags().Build(),
						Action:      distrox.GetVmLogs,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "create-cm-bundle",
						Usage:       "start CM based diagnostics collection for DistroX",
						Description: `start CM based diagnostics collection for DistroX`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags: fl.NewFlagBuilder().AddFlags(fl.FlCrn, fl.FlCollectionDestination, fl.FlDescriptionOptional, fl.FlCollectionIssue, fl.FlCollectionOnly,
							fl.FlCollectionRoles, fl.FlMonitorMetricsCollection, fl.FlUpdatePackage).AddAGlobalFlags().Build(),
						Action: distrox.CollectCmDiagnostics,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "cm-roles",
						Usage:       "get available roles for DistroX CM cluster",
						Description: `get available roles for DistroX CM cluster`,
						Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:       fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().Build(),
						Action:      distrox.GetCMRoles,
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
				Usage: "check distrox operations",
				Subcommands: []cli.Command{
					{
						Name:   "last",
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().Build(),
						Usage:  "check progress of the last distrox operation",
						Before: cf.CheckConfigAndCommandFlags,
						Action: distrox.OperationProgress,
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
				Usage:  "rotate distrox AutoTLS certificates",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().Build(),
				Action: distrox.RotateCertificates,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "upgrade",
				Usage:  "OS or data platform upgrade for the DistroX cluster. You need to specify at least either one of imageId, runtime or lockComponents to be able to proceed!",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().AddFlags(fl.FlImageIdOptional).AddFlags(fl.FlRuntimeOptional).AddFlags(fl.FlLockComponentsOptional).AddFlags(fl.FlDryRunOptional).AddFlags(fl.FlReplaceVms).AddFlags(fl.FlShowImagesOptional).AddFlags(fl.FlShowLatestImagesOptional).Build(),
				Action: distrox.DistroxClusterkUpgrade,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddNameFlag().AddAGlobalFlags().AddFlags(fl.FlImageIdOptional).AddFlags(fl.FlRuntimeOptional).AddFlags(fl.FlLockComponentsOptional).AddFlags(fl.FlDryRunOptional).AddFlags(fl.FlReplaceVms).AddFlags(fl.FlShowImagesOptional).AddFlags(fl.FlShowLatestImagesOptional).Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
		},
	})
}
