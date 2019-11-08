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
				Name:        "create",
				Usage:       "creates a new SDX cluster",
				Description: `basic SDX cluster creation`,
				Before:      cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:       fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlEnvironmentName, fl.FlClusterShape, fl.FlCloudStorageBaseLocationOptional, fl.FlCloudStorageInstanceProfileOptional, fl.FlCloudStorageManagedIdentityOptional, fl.FlCidrOptional, fl.FlWithoutExternalDatabaseOptional, fl.FlWithExternalDatabaseOptional).AddAuthenticationFlags().Build(),
				Action:      sdx.CreateSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlEnvironmentName, fl.FlEnvironmentName, fl.FlClusterShape, fl.FlCloudStorageBaseLocationOptional, fl.FlCloudStorageInstanceProfileOptional, fl.FlCloudStorageManagedIdentityOptional, fl.FlCidrOptional, fl.FlWithoutExternalDatabaseOptional, fl.FlWithExternalDatabaseOptional).AddAuthenticationFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "delete",
				Usage:  "deletes an SDX cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlForceOptional).AddNameFlag().AddAuthenticationFlags().Build(),
				Action: sdx.DeleteSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlForceOptional).AddNameFlag().AddAuthenticationFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "describe",
				Usage:  "describes an SDX cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlDetailedOptional).AddAuthenticationFlags().AddOutputFlag().Build(),
				Action: sdx.DescribeSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlDetailedOptional).AddAuthenticationFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "start",
				Usage:  "start an SDX cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddNameFlag().AddAuthenticationFlags().AddOutputFlag().Build(),
				Action: sdx.StartSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddNameFlag().AddAuthenticationFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "stop",
				Usage:  "stop an SDX cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddNameFlag().AddAuthenticationFlags().AddOutputFlag().Build(),
				Action: sdx.StopSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddNameFlag().AddAuthenticationFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "list",
				Usage:  "list SDX clusters",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentNameOptional).AddAuthenticationFlags().AddOutputFlag().Build(),
				Action: sdx.ListSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentNameOptional).AddAuthenticationFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "repair",
				Usage:  "repair SDX cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlHostGroupOptional, fl.FlHostGroupsOptional).AddNameFlag().AddAuthenticationFlags().Build(),
				Action: sdx.RepairSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlHostGroupOptional, fl.FlHostGroupsOptional).AddNameFlag().AddAuthenticationFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "sync",
				Usage:  "sync SDX cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddNameFlag().AddAuthenticationFlags().Build(),
				Action: sdx.SyncSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddNameFlag().AddAuthenticationFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "retry",
				Usage:  "retry SDX cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddNameFlag().AddAuthenticationFlags().Build(),
				Action: sdx.RetrySdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddNameFlag().AddAuthenticationFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "upgrade",
				Usage:  "OS upgrade for the SDX cluster",
				Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
				Flags:  fl.NewFlagBuilder().AddNameFlag().AddAuthenticationFlags().AddFlags(fl.FlDryRunOptional).Build(),
				Action: sdx.UpgradeSdx,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddNameFlag().AddAuthenticationFlags().AddFlags(fl.FlDryRunOptional).Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
		},
	})
}
