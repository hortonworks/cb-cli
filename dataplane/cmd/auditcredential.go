package cmd

import (
	cf "github.com/hortonworks/cb-cli/dataplane/config"
	"github.com/hortonworks/cb-cli/dataplane/auditcredential"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/urfave/cli"
)

func init() {
	DataPlaneCommands = append(DataPlaneCommands, cli.Command{
		Name:  "audit-credential",
		Usage: "audit-credential related operations",
		Subcommands: []cli.Command{
			{
				Name:  "create",
				Usage: "creates a new audit-credential",
				Subcommands: []cli.Command{
					{
						Name:  "aws",
						Usage: "creates a new aws audit-credential",
						Subcommands: []cli.Command{
							{
								Name:   "role-based",
								Usage:  "creates a new role based aws audit-credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlRoleARN).AddAGlobalFlags().Build(),
								Action: auditcredential.CreateAwsCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlRoleARN).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "key-based",
								Usage:  "creates a new key based aws audit-credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlAccessKey, fl.FlSecretKey).AddAGlobalFlags().Build(),
								Action: auditcredential.CreateAwsCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlAccessKey, fl.FlSecretKey).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
					{
						Name:  "azure",
						Usage: "creates a new azure audit-credential",
						Subcommands: []cli.Command{
							{
								Name:   "app-based",
								Usage:  "creates a new app based azure audit-credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlSubscriptionId, fl.FlTenantId, fl.FlAppId, fl.FlAppPassword).AddAGlobalFlags().Build(),
								Action: auditcredential.CreateAzureCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlSubscriptionId, fl.FlTenantId, fl.FlAppId, fl.FlAppPassword).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
					{
						Name:   "from-file",
						Usage:  "creates a new audit-credential from input json file",
						Flags:  fl.NewFlagBuilder().AddResourceFlagsWithOptionalName().AddFlags(fl.FlInputJson).AddAGlobalFlags().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: auditcredential.CreateCredentialFromFile,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddResourceFlagsWithOptionalName().AddFlags(fl.FlInputJson).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:  "modify",
				Usage: "modify an existing audit-credential",
				Subcommands: []cli.Command{
					{
						Name:  "aws",
						Usage: "modify an existing aws audit-credential",
						Subcommands: []cli.Command{
							{
								Name:   "role-based",
								Usage:  "modify a role based aws audit-credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlCredentialVerifyOptional, fl.FlName, fl.FlDescriptionOptional, fl.FlRoleARN).AddAGlobalFlags().Build(),
								Action: auditcredential.ModifyAwsCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredentialVerifyOptional, fl.FlDescriptionOptional, fl.FlRoleARN).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "key-based",
								Usage:  "modify a key based aws audit-credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlCredentialVerifyOptional, fl.FlName, fl.FlDescriptionOptional, fl.FlAccessKey, fl.FlSecretKey).AddAGlobalFlags().Build(),
								Action: auditcredential.ModifyAwsCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredentialVerifyOptional, fl.FlDescriptionOptional, fl.FlAccessKey, fl.FlSecretKey).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
					{
						Name:  "azure",
						Usage: "modify an existing azure audit-credential",
						Subcommands: []cli.Command{
							{
								Name:   "app-based",
								Usage:  "modify an app based azure audit-credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredentialVerifyOptional, fl.FlDescriptionOptional, fl.FlSubscriptionId, fl.FlTenantId, fl.FlAppId, fl.FlAppPassword).AddAGlobalFlags().Build(),
								Action: auditcredential.ModifyAzureCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredentialVerifyOptional, fl.FlDescriptionOptional, fl.FlSubscriptionId, fl.FlTenantId, fl.FlAppId, fl.FlAppPassword).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
					{
						Name:   "from-file",
						Usage:  "modify a audit-credential from input json file",
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlNameOptional, fl.FlInputJson).AddAGlobalFlags().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: auditcredential.ModifyCredentialFromFile,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlNameOptional, fl.FlInputJson).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:   "describe",
				Usage:  "describes a audit-credential",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: auditcredential.DescribeCredential,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "list",
				Usage:  "lists the audit-credentials",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build(),
				Action: auditcredential.ListCredentials,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:  "prerequisites",
				Usage: "get the necessary prerequisites for audit-credential creation",
				Subcommands: []cli.Command{
					{
						Name:   "aws",
						Usage:  "get prerequisites for aws audit-credential creation",
						Before: cf.CheckConfigAndCommandFlags,
						Flags:  fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build(),
						Action: auditcredential.GetAwsCredentialPrerequisites,
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
