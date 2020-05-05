package cmd

import (
	"github.com/hortonworks/cb-cli/dataplane/cluster"
	cf "github.com/hortonworks/cb-cli/dataplane/config"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/stack"
	"github.com/urfave/cli"
)

var stackTemplateDescription = `Template parameters to fill in the generated template:
		userName:	Name of the Ambari user
		password:	Password of the Ambari user
		name:	Name of the cluster
		region:	Region of the cluster
		availabilityZone:	Availability zone of the cluster, on AZURE it is the same as the region
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
		Name:  "cluster",
		Usage: "cluster related operations",
		Subcommands: []cli.Command{
			{
				Name:   "change-ambari-password",
				Usage:  "changes Ambari password",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlOldPassword, fl.FlNewPassword, fl.FlAmbariUser).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: cluster.ChangeAmbariPassword,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlOldPassword, fl.FlNewPassword, fl.FlAmbariUser).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "change-image",
				Usage:  "changes image of the cluster - will be used when creating new instances or repairing failed ones",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlImageId, fl.FlImageCatalogOptional).AddAGlobalFlags().Build(),
				Action: stack.ChangeImage,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlImageId).AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:        "create",
				Usage:       "creates a new cluster",
				Description: `use 'dp cluster generate-template' for cluster request JSON generation`,
				Flags:       fl.NewFlagBuilder().AddResourceFlagsWithOptionalName().AddFlags(fl.FlInputJson, fl.FlEnvironmentNameOptional, fl.FlCMUserOptional, fl.FlCMPasswordOptional, fl.FlWaitOptional).AddAGlobalFlags().Build(),
				Before:      cf.CheckConfigAndCommandFlags,
				Action:      stack.CreateStack,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddResourceFlagsWithOptionalName().AddFlags(fl.FlInputJson, fl.FlEnvironmentNameOptional, fl.FlCMUserOptional, fl.FlCMPasswordOptional, fl.FlWaitOptional).AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "delete",
				Usage:  "deletes a cluster",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlForceOptional, fl.FlWaitOptional).AddAGlobalFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: stack.DeleteStack,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlForceOptional, fl.FlWaitOptional).AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "describe",
				Usage:  "describes a cluster",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: stack.DescribeStack,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "describe-instances",
				Usage:  "describes cluster instances",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: stack.DescribeStackInstances,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "generate-inventory",
				Usage:  "creates an ansible compatible ini file from the cluster details",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlOutputFile).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: stack.GenerateInventoryFile,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlOutputFile).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:        "generate-template",
				Usage:       "creates a cluster JSON template",
				Description: stackTemplateDescription,
				Subcommands: []cli.Command{
					{
						Name:   "yarn",
						Usage:  "creates an yarn cluster JSON template",
						Before: cf.CheckConfigAndCommandFlags,
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional).AddAGlobalFlags().AddTemplateFlags().Build(),
						Action: stack.GenerateYarnStackTemplate,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional).AddAGlobalFlags().AddTemplateFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:  "aws",
						Usage: "creates an aws cluster JSON template",
						Subcommands: []cli.Command{
							{
								Name:   "new-network",
								Usage:  "creates an aws cluster JSON template with new network",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional, fl.FlDefaultEncryptionOptional, fl.FlCustomEncryptionOptional).AddAGlobalFlags().AddTemplateFlags().Build(),
								Action: stack.GenerateAwsStackTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional, fl.FlDefaultEncryptionOptional, fl.FlCustomEncryptionOptional).AddAGlobalFlags().AddTemplateFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "existing-network",
								Usage:  "creates an aws cluster JSON template with existing network",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional, fl.FlDefaultEncryptionOptional, fl.FlCustomEncryptionOptional).AddAGlobalFlags().AddTemplateFlags().Build(),
								Action: stack.GenerateAwsStackTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional, fl.FlDefaultEncryptionOptional, fl.FlCustomEncryptionOptional).AddAGlobalFlags().AddTemplateFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "existing-subnet",
								Usage:  "creates an aws cluster JSON template with existing subnet",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional, fl.FlDefaultEncryptionOptional, fl.FlCustomEncryptionOptional).AddAGlobalFlags().AddTemplateFlags().Build(),
								Action: stack.GenerateAwsStackTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional, fl.FlDefaultEncryptionOptional, fl.FlCustomEncryptionOptional).AddAGlobalFlags().AddTemplateFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
					{
						Name:  "azure",
						Usage: "creates an azure cluster JSON template",
						Subcommands: []cli.Command{
							{
								Name:   "new-network",
								Usage:  "creates an azure cluster JSON template with new network",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional).AddAGlobalFlags().AddTemplateFlags().Build(),
								Action: stack.GenerateAzureStackTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional).AddAGlobalFlags().AddTemplateFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "existing-subnet",
								Usage:  "creates an azure cluster JSON template with existing subnet",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional).AddAGlobalFlags().AddTemplateFlags().Build(),
								Action: stack.GenerateAzureStackTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional).AddAGlobalFlags().AddTemplateFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
					{
						Name:  "gcp",
						Usage: "creates an gcp cluster JSON template",
						Subcommands: []cli.Command{
							{
								Name:   "new-network",
								Usage:  "creates a gcp cluster JSON template with new network",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional, fl.FlRawEncryptionOptional, fl.FlRsaEncryptionOptional, fl.FlKmsEncryptionOptional).AddAGlobalFlags().AddTemplateFlags().Build(),
								Action: stack.GenerateGcpStackTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional, fl.FlRawEncryptionOptional, fl.FlRsaEncryptionOptional, fl.FlKmsEncryptionOptional).AddAGlobalFlags().AddTemplateFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "existing-network",
								Usage:  "creates a gcp cluster JSON template with existing network",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional, fl.FlRawEncryptionOptional, fl.FlRsaEncryptionOptional, fl.FlKmsEncryptionOptional).AddAGlobalFlags().AddTemplateFlags().Build(),
								Action: stack.GenerateGcpStackTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional, fl.FlRawEncryptionOptional, fl.FlRsaEncryptionOptional, fl.FlKmsEncryptionOptional).AddAGlobalFlags().AddTemplateFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "existing-subnet",
								Usage:  "creates a gcp cluster JSON template with existing subnet",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional, fl.FlRawEncryptionOptional, fl.FlRsaEncryptionOptional, fl.FlKmsEncryptionOptional).AddAGlobalFlags().AddTemplateFlags().Build(),
								Action: stack.GenerateGcpStackTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional, fl.FlRawEncryptionOptional, fl.FlRsaEncryptionOptional, fl.FlKmsEncryptionOptional).AddAGlobalFlags().AddTemplateFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "legacy-network",
								Usage:  "creates a gcp cluster JSON template with legacy network",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional).AddAGlobalFlags().AddTemplateFlags().Build(),
								Action: stack.GenerateGcpStackTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional).AddAGlobalFlags().AddTemplateFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "shared-network",
								Usage:  "creates a gcp cluster JSON template with shared network",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional).AddAGlobalFlags().AddTemplateFlags().Build(),
								Action: stack.GenerateGcpStackTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional).AddAGlobalFlags().AddTemplateFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
					{
						Name:  "openstack",
						Usage: "creates an openstack cluster JSON template",
						Subcommands: []cli.Command{
							{
								Name:   "new-network",
								Usage:  "creates a openstack cluster JSON template with new network",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional).AddAGlobalFlags().AddTemplateFlags().Build(),
								Action: stack.GenerateOpenstackStackTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional).AddAGlobalFlags().AddTemplateFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "existing-network",
								Usage:  "creates a openstack cluster JSON template with existing network",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional).AddAGlobalFlags().AddTemplateFlags().Build(),
								Action: stack.GenerateOpenstackStackTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional).AddAGlobalFlags().AddTemplateFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "existing-subnet",
								Usage:  "creates a openstack cluster JSON template with existing subnet",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional).AddAGlobalFlags().AddTemplateFlags().Build(),
								Action: stack.GenerateOpenstackStackTemplate,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlBlueprintNameOptional, fl.FlBlueprintFileOptional).AddAGlobalFlags().AddTemplateFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
				},
			},
			{
				Name:   "generate-attached-cluster-template",
				Usage:  "generates attached cluster template",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlWithSourceCluster, fl.FlBlueprintName, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: stack.GenerateAtachedStackTemplate,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlWithSourceCluster, fl.FlBlueprintName, fl.FlBlueprintFileOptional, fl.FlCloudStorageTypeOptional).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "list",
				Usage:  "lists the running clusters",
				Flags:  fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: stack.ListStacks,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:        "repair",
				Usage:       "repairs a cluster",
				Description: "repairs a cluster",
				Subcommands: []cli.Command{
					{
						Name:   "host-groups",
						Usage:  "repairs host-groups of a cluster",
						Before: cf.CheckConfigAndCommandFlags,
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlHostGroups, fl.FlRemoveOnly, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build(),
						Action: stack.RepairStackHostGroups,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlHostGroups, fl.FlRemoveOnly, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:   "nodes",
						Usage:  "repairs nodes of a cluster",
						Before: cf.CheckConfigAndCommandFlags,
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlNodes, fl.FlDeleteVolumes, fl.FlRemoveOnly, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build(),
						Action: stack.RepairStackNodes,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlNodes, fl.FlDeleteVolumes, fl.FlRemoveOnly, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:   "request",
				Usage:  "print the request for a cluster",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: stack.GetStackRequest,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "retry",
				Usage:  "retries the creation of a cluster",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: stack.RetryCluster,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "scale",
				Usage:  "scales a cluster",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlGroupName, fl.FlDesiredNodeCount, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: stack.ScaleStack,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlGroupName, fl.FlDesiredNodeCount, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "start",
				Usage:  "starts a cluster",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: stack.StartStack,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "stop",
				Usage:  "stops a cluster",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: stack.StopStack,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlWaitOptional).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "sync",
				Usage:  "synchronizes a cluster",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: stack.SyncStack,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "update-salt",
				Usage:  "update salt states on cluster",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: stack.UpdateSalt,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
		},
		Hidden: true,
	})
}
