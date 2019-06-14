package cmd

import (
	cf "github.com/hortonworks/cb-cli/dataplane/config"
	"github.com/hortonworks/cb-cli/dataplane/databaseconfig"
	"github.com/hortonworks/cb-cli/dataplane/databaseserverconfig"
	// databaseconfig "github.com/hortonworks/cb-cli/dataplane/databaseconfig"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/urfave/cli"
)

func init() {
	DataPlaneCommands = append(DataPlaneCommands, cli.Command{
		Name:  "database-redbeams",
		Usage: "database-server and database config management related operations",
		Subcommands: []cli.Command{
			{
				Name:  "database",
				Usage: "database config related operations",
				Subcommands: []cli.Command{
					{
						Name:  "register",
						Usage: "register a new database configuration",
						Flags: fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlRdsUserName, fl.FlRdsPassword, fl.FlRdsURL,
							fl.FlRdsType, fl.FlRdsConnectorJarURLOptional).AddAuthenticationFlags().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: databaseconfig.EmptyFunc,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlRdsUserName, fl.FlRdsPassword, fl.FlRdsURL,
								fl.FlRdsType, fl.FlRdsConnectorJarURLOptional).AddAuthenticationFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},

					{
						Name:  "create",
						Usage: "create a new database configuration",
						Flags: fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlRdsUserName, fl.FlRdsPassword, fl.FlRdsURL,
							fl.FlRdsType, fl.FlRdsConnectorJarURLOptional).AddAuthenticationFlags().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: databaseconfig.EmptyFunc,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlRdsUserName, fl.FlRdsPassword, fl.FlRdsURL,
								fl.FlRdsType, fl.FlRdsConnectorJarURLOptional).AddAuthenticationFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},

					{
						Name:   "list",
						Usage:  "list database configs",
						Flags:  fl.NewFlagBuilder().AddAuthenticationFlags().AddFlags(fl.FlEnvironmentCrn).AddOutputFlag().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: databaseconfig.EmptyFunc,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlEnvironmentCrn).AddOutputFlag().AddAuthenticationFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},

					{
						Name:   "describe",
						Usage:  "describes a database configuration",
						Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlEnvironmentCrn).AddAuthenticationFlags().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: databaseconfig.EmptyFunc,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlEnvironmentCrn).AddAuthenticationFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},

					{
						Name:  "test",
						Usage: "test a database configuration",
						Subcommands: []cli.Command{
							{
								Name:   "by-name",
								Usage:  "test a stored database configuration by name",
								Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlName).AddAuthenticationFlags().Build(),
								Before: cf.CheckConfigAndCommandFlags,
								Action: databaseconfig.EmptyFunc,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlName).AddAuthenticationFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:  "by-param",
								Usage: "test database connection parameters",
								Flags: fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlRdsUserName, fl.FlRdsPassword, fl.FlRdsURL,
									fl.FlRdsType, fl.FlRdsConnectorJarURLOptional).AddAuthenticationFlags().Build(),
								Before: cf.CheckConfigAndCommandFlags,
								Action: databaseconfig.EmptyFunc,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlRdsUserName, fl.FlRdsPassword, fl.FlRdsURL,
										fl.FlRdsType, fl.FlRdsConnectorJarURLOptional).AddAuthenticationFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},

					{
						Name:   "delete",
						Usage:  "deletes one or more database configuration(s)",
						Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlNames).AddAuthenticationFlags().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: databaseconfig.EmptyFunc,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlNames).AddAuthenticationFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:  "database-server",
				Usage: "database server config related operations",
				Subcommands: []cli.Command{
					{
						Name:  "register",
						Usage: "register a database server",
						Flags: fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlRdsUserName, fl.FlRdsPassword, fl.FlRdsURL,
							fl.FlRdsType, fl.FlRdsConnectorJarURLOptional).AddAuthenticationFlags().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: databaseserverconfig.StubFunc,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlRdsUserName, fl.FlRdsPassword, fl.FlRdsURL,
								fl.FlRdsType, fl.FlRdsConnectorJarURLOptional).AddAuthenticationFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},

					{
						Name:  "create",
						Usage: "create a database server",
						Flags: fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlRdsUserName, fl.FlRdsPassword, fl.FlRdsURL,
							fl.FlRdsType, fl.FlRdsConnectorJarURLOptional).AddAuthenticationFlags().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: databaseserverconfig.StubFunc,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlRdsUserName, fl.FlRdsPassword, fl.FlRdsURL,
								fl.FlRdsType, fl.FlRdsConnectorJarURLOptional).AddAuthenticationFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},

					{
						Name:   "list",
						Usage:  "list database servers",
						Flags:  fl.NewFlagBuilder().AddAuthenticationFlags().AddFlags(fl.FlEnvironmentCrn).Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: databaseserverconfig.StubFunc,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlEnvironmentCrn).AddAuthenticationFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},

					{
						Name:   "describe",
						Usage:  "describe a database server",
						Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlEnvironmentCrn).AddAuthenticationFlags().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: databaseserverconfig.StubFunc,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlEnvironmentCrn).AddAuthenticationFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},

					{
						Name:  "create-database-on-server",
						Usage: "create a database on a database server",
						Flags: fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlRdsUserName, fl.FlRdsPassword, fl.FlRdsURL,
							fl.FlRdsType, fl.FlRdsConnectorJarURLOptional).AddAuthenticationFlags().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: databaseserverconfig.StubFunc,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlRdsUserName, fl.FlRdsPassword, fl.FlRdsURL,
								fl.FlRdsType, fl.FlRdsConnectorJarURLOptional).AddAuthenticationFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},

					{
						Name:  "test",
						Usage: "test a database server",
						Subcommands: []cli.Command{
							{
								Name:   "by-name",
								Usage:  "test a stored database server configuration by name",
								Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlName).AddAuthenticationFlags().Build(),
								Before: cf.CheckConfigAndCommandFlags,
								Action: databaseserverconfig.StubFunc,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlName).AddAuthenticationFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:  "by-param",
								Usage: "test database server connection parameters",
								Flags: fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlRdsUserName, fl.FlRdsPassword, fl.FlRdsURL,
									fl.FlRdsType, fl.FlRdsConnectorJarURLOptional).AddAuthenticationFlags().Build(),
								Before: cf.CheckConfigAndCommandFlags,
								Action: databaseserverconfig.StubFunc,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlRdsUserName, fl.FlRdsPassword, fl.FlRdsURL,
										fl.FlRdsType, fl.FlRdsConnectorJarURLOptional).AddAuthenticationFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},

					{
						Name:   "delete",
						Usage:  "delete database server(s)",
						Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddAuthenticationFlags().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: databaseserverconfig.StubFunc,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlRdsUserName, fl.FlRdsPassword, fl.FlRdsURL,
								fl.FlRdsType, fl.FlRdsConnectorJarURLOptional).AddAuthenticationFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
		},
	})
}
