package cmd

import (
	cf "github.com/hortonworks/cb-cli/dataplane/config"
	"github.com/hortonworks/cb-cli/dataplane/credential"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/urfave/cli"
)

func init() {
	DataPlaneCommands = append(DataPlaneCommands, cli.Command{
		Name:  "credential",
		Usage: "credential related operations",
		Subcommands: []cli.Command{
			{
				Name:  "create",
				Usage: "creates a new credential",
				Subcommands: []cli.Command{
					{
						Name:  "aws",
						Usage: "creates a new aws credential",
						Subcommands: []cli.Command{
							{
								Name:   "role-based",
								Usage:  "creates a new role based aws credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlRoleARN).AddAGlobalFlags().Build(),
								Action: credential.CreateAwsCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlRoleARN).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "key-based",
								Usage:  "creates a new key based aws credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlAccessKey, fl.FlSecretKey).AddAGlobalFlags().Build(),
								Action: credential.CreateAwsCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlAccessKey, fl.FlSecretKey).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
					{
						Name:  "aws-gov",
						Usage: "creates a new aws govcloud credential",
						Subcommands: []cli.Command{
							{
								Name:   "key-based",
								Usage:  "creates a new key based aws govcloud credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlAccessKey, fl.FlSecretKey).AddAGlobalFlags().Build(),
								Action: credential.CreateAwsGovCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlAccessKey, fl.FlSecretKey).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "role-based",
								Usage:  "creates a new role based aws govcloud credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlRoleARN).AddAGlobalFlags().Build(),
								Action: credential.CreateAwsGovCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlRoleARN).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
					{
						Name:  "azure",
						Usage: "creates a new azure credential",
						Subcommands: []cli.Command{
							{
								Name:   "app-based",
								Usage:  "creates a new app based azure credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlSubscriptionId, fl.FlTenantId, fl.FlAppIdOptional, fl.FlAppPasswordOptional, fl.FlAzureAuthTypeOptional).AddAGlobalFlags().Build(),
								Action: credential.CreateAzureCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlSubscriptionId, fl.FlTenantId, fl.FlAppIdOptional, fl.FlAppPasswordOptional, fl.FlAzureAuthTypeOptional).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
					{
						Name:  "gcp",
						Usage: "creates a new gcp credential",
						Subcommands: []cli.Command{
							{
								Name:   "p12-based",
								Usage:  "creates a new P12 based gcp credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlProjectId, fl.FlServiceAccountId, fl.FlServiceAccountPrivateKeyFile).AddAGlobalFlags().Build(),
								Action: credential.CreateGcpCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlProjectId, fl.FlServiceAccountId, fl.FlServiceAccountPrivateKeyFile).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "json-based",
								Usage:  "creates a new JSON based gcp credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlServiceAccountJsonFile).AddAGlobalFlags().Build(),
								Action: credential.CreateGcpCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlServiceAccountJsonFile).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
					{
						Name:  "openstack",
						Usage: "creates a new openstack credential",
						Subcommands: []cli.Command{
							{
								Name:   "keystone-v2",
								Usage:  "creates a new keystone version 2 openstack credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlTenantUser, fl.FlTenantPassword, fl.FlTenantName, fl.FlEndpoint, fl.FlFacingOptional).AddAGlobalFlags().Build(),
								Action: credential.CreateOpenstackCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlTenantUser, fl.FlTenantPassword, fl.FlTenantName, fl.FlEndpoint, fl.FlFacingOptional).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "keystone-v3",
								Usage:  "creates a new keystone version 3 openstack credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags: fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlTenantUser, fl.FlTenantPassword,
									fl.FlUserDomain, fl.FlProjectDomainNameOptional, fl.FlProjectNameOptional, fl.FlDomainNameOptional, fl.FlKeystoneScopeOptional, fl.FlEndpoint, fl.FlFacingOptional).AddAGlobalFlags().Build(),
								Action: credential.CreateOpenstackCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlTenantUser, fl.FlTenantPassword,
										fl.FlUserDomain, fl.FlProjectDomainNameOptional, fl.FlProjectNameOptional, fl.FlDomainNameOptional, fl.FlKeystoneScopeOptional, fl.FlEndpoint, fl.FlFacingOptional).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
					{
						Name:   "ycloud",
						Usage:  "creates a new ycloud credential",
						Before: cf.CheckConfigAndCommandFlags,
						Flags:  fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlEndpoint).AddAGlobalFlags().Build(),
						Action: credential.CreateYarnCredential,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlCredentialVerifyOptional, fl.FlEndpoint).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:   "from-file",
						Usage:  "creates a new credential from input json file",
						Flags:  fl.NewFlagBuilder().AddResourceFlagsWithOptionalName().AddFlags(fl.FlInputJson).AddAGlobalFlags().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: credential.CreateCredentialFromFile,
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
				Usage: "modify an existing credential",
				Subcommands: []cli.Command{
					{
						Name:  "aws",
						Usage: "modify an existing aws credential",
						Subcommands: []cli.Command{
							{
								Name:   "role-based",
								Usage:  "modify a role based aws credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlCredentialVerifyOptional, fl.FlName, fl.FlDescriptionOptional, fl.FlRoleARN).AddAGlobalFlags().Build(),
								Action: credential.ModifyAwsCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredentialVerifyOptional, fl.FlDescriptionOptional, fl.FlRoleARN).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "key-based",
								Usage:  "modify a key based aws credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlCredentialVerifyOptional, fl.FlName, fl.FlDescriptionOptional, fl.FlAccessKey, fl.FlSecretKey).AddAGlobalFlags().Build(),
								Action: credential.ModifyAwsCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredentialVerifyOptional, fl.FlDescriptionOptional, fl.FlAccessKey, fl.FlSecretKey).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
					{
						Name:  "aws-gov",
						Usage: "modify an existing aws govcloud credential",
						Subcommands: []cli.Command{
							{
								Name:   "role-based",
								Usage:  "modify a role based aws govcloud credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredentialVerifyOptional, fl.FlDescriptionOptional, fl.FlRoleARN).AddAGlobalFlags().Build(),
								Action: credential.ModifyAwsGovCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredentialVerifyOptional, fl.FlDescriptionOptional, fl.FlRoleARN).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "key-based",
								Usage:  "modify a key based aws govcloud credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredentialVerifyOptional, fl.FlDescriptionOptional, fl.FlAccessKey, fl.FlSecretKey).AddAGlobalFlags().Build(),
								Action: credential.ModifyAwsGovCredential,
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
						Usage: "modify an existing azure credential",
						Subcommands: []cli.Command{
							{
								Name:   "app-based",
								Usage:  "modify an app based azure credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredentialVerifyOptional, fl.FlDescriptionOptional, fl.FlSubscriptionId, fl.FlTenantId, fl.FlAppIdOptional, fl.FlAppPasswordOptional, fl.FlAzureAuthTypeOptional, fl.FlGenerateCertAuthOptional).AddAGlobalFlags().Build(),
								Action: credential.ModifyAzureCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredentialVerifyOptional, fl.FlDescriptionOptional, fl.FlSubscriptionId, fl.FlTenantId, fl.FlAppIdOptional, fl.FlAppPasswordOptional, fl.FlAzureAuthTypeOptional, fl.FlGenerateCertAuthOptional).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
					{
						Name:  "gcp",
						Usage: "modify an existing gcp credential",
						Subcommands: []cli.Command{
							{
								Name:   "p12-based",
								Usage:  "modify a P12 based gcp credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredentialVerifyOptional, fl.FlDescriptionOptional, fl.FlProjectId, fl.FlServiceAccountId, fl.FlServiceAccountPrivateKeyFile).AddAGlobalFlags().Build(),
								Action: credential.ModifyGcpCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredentialVerifyOptional, fl.FlDescriptionOptional, fl.FlProjectId, fl.FlServiceAccountId, fl.FlServiceAccountPrivateKeyFile).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "json-based",
								Usage:  "modify a JSON based gcp credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredentialVerifyOptional, fl.FlDescriptionOptional, fl.FlServiceAccountJsonFile).AddAGlobalFlags().Build(),
								Action: credential.ModifyGcpCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredentialVerifyOptional, fl.FlDescriptionOptional, fl.FlServiceAccountJsonFile).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
					{
						Name:  "openstack",
						Usage: "modify an existing openstack credential",
						Subcommands: []cli.Command{
							{
								Name:   "keystone-v2",
								Usage:  "modify a keystone version 2 openstack credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredentialVerifyOptional, fl.FlDescriptionOptional, fl.FlTenantUser, fl.FlTenantPassword, fl.FlTenantName, fl.FlEndpoint, fl.FlFacingOptional).AddAGlobalFlags().Build(),
								Action: credential.ModifyOpenstackCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredentialVerifyOptional, fl.FlDescriptionOptional, fl.FlTenantUser, fl.FlTenantPassword, fl.FlTenantName, fl.FlEndpoint, fl.FlFacingOptional).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "keystone-v3",
								Usage:  "modify a keystone version 3 openstack credential",
								Before: cf.CheckConfigAndCommandFlags,
								Flags: fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredentialVerifyOptional, fl.FlDescriptionOptional, fl.FlTenantUser, fl.FlTenantPassword,
									fl.FlUserDomain, fl.FlProjectDomainNameOptional, fl.FlProjectNameOptional, fl.FlDomainNameOptional, fl.FlKeystoneScopeOptional, fl.FlEndpoint, fl.FlFacingOptional).AddAGlobalFlags().Build(),
								Action: credential.ModifyOpenstackCredential,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName, fl.FlCredentialVerifyOptional, fl.FlDescriptionOptional, fl.FlTenantUser, fl.FlTenantPassword,
										fl.FlUserDomain, fl.FlProjectDomainNameOptional, fl.FlProjectNameOptional, fl.FlDomainNameOptional, fl.FlKeystoneScopeOptional, fl.FlEndpoint, fl.FlFacingOptional).AddAGlobalFlags().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
					{
						Name:   "from-file",
						Usage:  "modify a credential from input json file",
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlNameOptional, fl.FlInputJson).AddAGlobalFlags().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: credential.ModifyCredentialFromFile,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlNameOptional, fl.FlInputJson).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:   "delete",
				Usage:  "deletes an credential or more if names are spearated by commas",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlNames).AddAGlobalFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: credential.DeleteCredential,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlNames).AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "describe",
				Usage:  "describes a credential",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build(),
				Action: credential.DescribeCredential,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "list",
				Usage:  "lists the credentials",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build(),
				Action: credential.ListCredentials,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:  "prerequisites",
				Usage: "get the necessary prerequisites for credential creation",
				Subcommands: []cli.Command{
					{
						Name:   "aws",
						Usage:  "get prerequisites for aws credential creation",
						Before: cf.CheckConfigAndCommandFlags,
						Flags:  fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build(),
						Action: credential.GetAwsCredentialPrerequisites,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:  "verify",
				Usage: "verify credentials",
				Subcommands: []cli.Command{
					{
						Name:   "by-name",
						Usage:  "verify credential by name",
						Before: cf.CheckConfigAndCommandFlags,
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build(),
						Action: credential.VerifyCredential,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:   "by-crn",
						Usage:  "verify credential by crn",
						Before: cf.CheckConfigAndCommandFlags,
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().AddOutputFlag().Build(),
						Action: credential.VerifyCredential,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
		},
	})
}
