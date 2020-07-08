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
		},
	})
}
