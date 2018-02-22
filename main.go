package main

import (
	"fmt"
	"os"
	"sort"

	"errors"

	log "github.com/Sirupsen/logrus"
	cb "github.com/hortonworks/cb-cli/cli"
	_ "github.com/hortonworks/cb-cli/cli/cloud/aws"
	_ "github.com/hortonworks/cb-cli/cli/cloud/azure"
	_ "github.com/hortonworks/cb-cli/cli/cloud/gcp"
	_ "github.com/hortonworks/cb-cli/cli/cloud/openstack"
	_ "github.com/hortonworks/cb-cli/cli/cloud/yarn"
	"github.com/hortonworks/cb-cli/cli/utils"
	"github.com/urfave/cli"
)

func ConfigRead(c *cli.Context) error {
	args := c.Args()
	if args.Present() {
		name := args.First()
		if k := c.App.Command(name); k != nil {
			// this is a sub-command invocation
			return nil
		}
	}

	server := c.String(cb.FlServerOptional.Name)
	username := c.String(cb.FlUsername.Name)
	password := c.String(cb.FlPassword.Name)
	output := c.String(cb.FlOutputOptional.Name)
	profile := c.String(cb.FlProfileOptional.Name)
	authType := c.String(cb.FlAuthTypeOptional.Name)

	if len(profile) == 0 {
		profile = "default"
	}

	config, err := cb.ReadConfig(cb.GetHomeDirectory(), profile)
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	set := func(name, value string) {
		if err := c.Set(name, value); err != nil {
			log.Debug(err)
		}
	}

	if len(output) == 0 {
		set(cb.FlOutputOptional.Name, config.Output)
	}

	if len(authType) == 0 {
		authType = config.AuthType
		if len(authType) == 0 {
			set(cb.FlAuthTypeOptional.Name, cb.OAUTH2)
		} else {
			if authType != cb.OAUTH2 && authType != cb.BASIC {
				utils.LogErrorAndExit(errors.New(fmt.Sprintf("invalid authentication type, accepted values: [%s, %s]", cb.OAUTH2, cb.BASIC)))
			}
			set(cb.FlAuthTypeOptional.Name, authType)
		}
	}

	if len(server) == 0 || len(username) == 0 || len(password) == 0 {
		if err != nil {
			log.Error(fmt.Sprintf("configuration is not set, see: cb configure --help or provide the following flags: %v",
				[]string{"--" + cb.FlServerOptional.Name, "--" + cb.FlUsername.Name, "--" + cb.FlPassword.Name}))
			os.Exit(1)
		}

		PrintConfig(*config)

		if len(server) == 0 {
			set(cb.FlServerOptional.Name, config.Server)
		}
		if len(username) == 0 {
			set(cb.FlUsername.Name, config.Username)
		}
		if len(password) == 0 {
			set(cb.FlPassword.Name, config.Password)
		}
	}
	return nil
}

func PrintConfig(cfg cb.Config) {
	cfg.Password = "*"
	log.Infof("[ConfigRead] Config read from file, setting as global variable:\n%s", cfg.Yaml())
}

func printFlagCompletion(f cli.Flag) {
	fmt.Printf("--%s\n", f.GetName())
}

func sortByName(commands []cli.Command) {
	sort.Slice(commands, func(i, j int) bool {
		return commands[i].Name < commands[j].Name
	})
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Debug(r)
			os.Exit(1)
		}
	}()

	app := cli.NewApp()
	app.Name = "cb"
	app.HelpName = "Hortonworks Data Cloud command line tool"
	app.Version = cb.Version + "-" + cb.BuildTime
	app.Author = "Hortonworks"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cb.FlDebugOptional,
	}

	app.Before = func(c *cli.Context) error {
		log.SetOutput(os.Stderr)
		log.SetLevel(log.ErrorLevel)
		formatter := &utils.CBFormatter{}
		log.SetFormatter(formatter)
		if c.Bool(cb.FlDebugOptional.Name) {
			log.SetLevel(log.DebugLevel)
		}
		return nil
	}

	cli.AppHelpTemplate = cb.AppHelpTemplate
	cli.HelpPrinter = cb.PrintHelp
	cli.CommandHelpTemplate = cb.CommandHelpTemplate
	cli.SubcommandHelpTemplate = cb.SubCommandHelpTemplate
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Command not found: %q\n", command)
	}

	app.Commands = []cli.Command{
		{
			Name: "configure",
			Description: fmt.Sprintf("it will save the provided server address and credential "+
				"to %s/%s/%s", cb.GetHomeDirectory(), cb.Config_dir, cb.Config_file),
			Usage:  "configure the server address and credentials used to communicate with this server",
			Flags:  cb.NewFlagBuilder().AddFlags(cb.FlServerRequired, cb.FlUsernameRequired, cb.FlPassword, cb.FlProfileOptional, cb.FlAuthTypeOptional).AddOutputFlag().Build(),
			Action: cb.Configure,
			BashComplete: func(c *cli.Context) {
				for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlServerRequired, cb.FlUsernameRequired, cb.FlPassword, cb.FlProfileOptional, cb.FlAuthTypeOptional).AddOutputFlag().Build() {
					printFlagCompletion(f)
				}
			},
		},
		{
			Name:  "blueprint",
			Usage: "blueprint related operations",
			Subcommands: []cli.Command{
				{
					Name:  "create",
					Usage: "adds a new Ambari blueprint from a file or from a URL",
					Subcommands: []cli.Command{
						{
							Name:   "from-url",
							Flags:  cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlURL).AddAuthenticationFlags().Build(),
							Before: ConfigRead,
							Action: cb.CreateBlueprintFromUrl,
							Usage:  "creates a blueprint by downloading it from a URL location",
							BashComplete: func(c *cli.Context) {
								for _, f := range cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlURL).AddAuthenticationFlags().Build() {
									printFlagCompletion(f)
								}
							},
						},
						{
							Name:   "from-file",
							Flags:  cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlFile).AddAuthenticationFlags().Build(),
							Before: ConfigRead,
							Action: cb.CreateBlueprintFromFile,
							Usage:  "creates a blueprint by reading it from a local file",
							BashComplete: func(c *cli.Context) {
								for _, f := range cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlFile).AddAuthenticationFlags().Build() {
									printFlagCompletion(f)
								}
							},
						},
					},
				},
				{
					Name:   "delete",
					Usage:  "deletes a blueprint",
					Flags:  cb.NewFlagBuilder().AddFlags(cb.FlName).AddAuthenticationFlags().Build(),
					Before: ConfigRead,
					Action: cb.DeleteBlueprint,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlName).AddAuthenticationFlags().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "describe",
					Usage:  "describes a blueprint",
					Before: ConfigRead,
					Flags:  cb.NewFlagBuilder().AddFlags(cb.FlName).AddAuthenticationFlags().AddOutputFlag().Build(),
					Action: cb.DescribeBlueprint,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlName).AddAuthenticationFlags().AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "list",
					Usage:  "lists the available blueprints",
					Flags:  cb.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build(),
					Before: ConfigRead,
					Action: cb.ListBlueprints,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
			},
		},
		{
			Name:  "cloud",
			Usage: "information about cloud provider resources",
			Subcommands: []cli.Command{
				{
					Name:   "availability-zones",
					Usage:  "lists the available availabilityzones in a region",
					Flags:  cb.NewFlagBuilder().AddAuthenticationFlags().AddFlags(cb.FlCredential, cb.FlRegion).AddOutputFlag().Build(),
					Before: ConfigRead,
					Action: cb.ListAvailabilityZones,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddAuthenticationFlags().AddFlags(cb.FlCredential, cb.FlRegion).AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "instances",
					Usage:  "lists the available instance types",
					Flags:  cb.NewFlagBuilder().AddAuthenticationFlags().AddFlags(cb.FlCredential, cb.FlRegion, cb.FlAvailabilityZoneOptional).AddOutputFlag().Build(),
					Before: ConfigRead,
					Action: cb.ListInstanceTypes,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddAuthenticationFlags().AddFlags(cb.FlCredential, cb.FlRegion, cb.FlAvailabilityZoneOptional).AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "regions",
					Usage:  "lists the available regions",
					Flags:  cb.NewFlagBuilder().AddAuthenticationFlags().AddFlags(cb.FlCredential).AddOutputFlag().Build(),
					Before: ConfigRead,
					Action: cb.ListRegions,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddAuthenticationFlags().AddFlags(cb.FlCredential).AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:  "volumes",
					Usage: "lists the available volume types",
					Subcommands: []cli.Command{
						{
							Name:   "aws",
							Usage:  "list the available aws volume types",
							Flags:  cb.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build(),
							Before: ConfigRead,
							Action: cb.ListAwsVolumeTypes,
							BashComplete: func(c *cli.Context) {
								for _, f := range cb.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build() {
									printFlagCompletion(f)
								}
							},
						},
						{
							Name:   "azure",
							Usage:  "list the available azure volume types",
							Flags:  cb.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build(),
							Before: ConfigRead,
							Action: cb.ListAzureVolumeTypes,
							BashComplete: func(c *cli.Context) {
								for _, f := range cb.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build() {
									printFlagCompletion(f)
								}
							},
						},
						{
							Name:   "gcp",
							Usage:  "list the available gcp volume types",
							Flags:  cb.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build(),
							Before: ConfigRead,
							Action: cb.ListGcpVolumeTypes,
							BashComplete: func(c *cli.Context) {
								for _, f := range cb.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build() {
									printFlagCompletion(f)
								}
							},
						},
					},
				},
			},
		},
		{
			Name:  "cluster",
			Usage: "cluster related operations",
			Subcommands: []cli.Command{
				{
					Name:   "change-ambari-password",
					Usage:  "changes Ambari password",
					Before: ConfigRead,
					Flags:  cb.NewFlagBuilder().AddFlags(cb.FlName, cb.FlOldPassword, cb.FlNewPassword, cb.FlAmbariUser).AddAuthenticationFlags().AddOutputFlag().Build(),
					Action: cb.ChangeAmbariPassword,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlName, cb.FlOldPassword, cb.FlNewPassword, cb.FlAmbariUser).AddAuthenticationFlags().AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:        "create",
					Usage:       "creates a new cluster",
					Description: `use 'cb cluster generate-template' for cluster request JSON generation`,
					Flags:       cb.NewFlagBuilder().AddResourceFlagsWithOptionalName().AddFlags(cb.FlInputJson, cb.FlAmbariPasswordOptional, cb.FlWaitOptional).AddAuthenticationFlags().Build(),
					Before:      ConfigRead,
					Action:      cb.CreateStack,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddResourceFlagsWithOptionalName().AddFlags(cb.FlInputJson, cb.FlAmbariPasswordOptional, cb.FlWaitOptional).AddAuthenticationFlags().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "delete",
					Usage:  "deletes a cluster",
					Flags:  cb.NewFlagBuilder().AddFlags(cb.FlName, cb.FlForceOptional, cb.FlWaitOptional).AddAuthenticationFlags().Build(),
					Before: ConfigRead,
					Action: cb.DeleteStack,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlName, cb.FlForceOptional, cb.FlWaitOptional).AddAuthenticationFlags().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "describe",
					Usage:  "describes a cluster",
					Before: ConfigRead,
					Flags:  cb.NewFlagBuilder().AddFlags(cb.FlName).AddAuthenticationFlags().AddOutputFlag().Build(),
					Action: cb.DescribeStack,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlName).AddAuthenticationFlags().AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:        "generate-template",
					Usage:       "creates a cluster JSON template",
					Description: cb.StackTemplateDescription,
					Subcommands: []cli.Command{
						{
							Name:   "yarn",
							Usage:  "creates an yarn cluster JSON template",
							Before: ConfigRead,
							Flags:  cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build(),
							Action: cb.GenerateYarnStackTemplate,
							BashComplete: func(c *cli.Context) {
								for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build() {
									printFlagCompletion(f)
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
									Before: ConfigRead,
									Flags:  cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build(),
									Action: cb.GenerateAwsStackTemplate,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build() {
											printFlagCompletion(f)
										}
									},
								},
								{
									Name:   "existing-network",
									Usage:  "creates an aws cluster JSON template with existing network",
									Before: ConfigRead,
									Flags:  cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build(),
									Action: cb.GenerateAwsStackTemplate,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build() {
											printFlagCompletion(f)
										}
									},
								},
								{
									Name:   "existing-subnet",
									Usage:  "creates an aws cluster JSON template with existing subnet",
									Before: ConfigRead,
									Flags:  cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build(),
									Action: cb.GenerateAwsStackTemplate,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build() {
											printFlagCompletion(f)
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
									Before: ConfigRead,
									Flags:  cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build(),
									Action: cb.GenerateAzureStackTemplate,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build() {
											printFlagCompletion(f)
										}
									},
								},
								{
									Name:   "existing-subnet",
									Usage:  "creates an azure cluster JSON template with existing subnet",
									Before: ConfigRead,
									Flags:  cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build(),
									Action: cb.GenerateAzureStackTemplate,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build() {
											printFlagCompletion(f)
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
									Before: ConfigRead,
									Flags:  cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build(),
									Action: cb.GenerateGcpStackTemplate,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build() {
											printFlagCompletion(f)
										}
									},
								},
								{
									Name:   "existing-network",
									Usage:  "creates a gcp cluster JSON template with existing network",
									Before: ConfigRead,
									Flags:  cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build(),
									Action: cb.GenerateGcpStackTemplate,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build() {
											printFlagCompletion(f)
										}
									},
								},
								{
									Name:   "existing-subnet",
									Usage:  "creates a gcp cluster JSON template with existing subnet",
									Before: ConfigRead,
									Flags:  cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build(),
									Action: cb.GenerateGcpStackTemplate,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build() {
											printFlagCompletion(f)
										}
									},
								},
								{
									Name:   "legacy-network",
									Usage:  "creates a gcp cluster JSON template with legacy network",
									Before: ConfigRead,
									Flags:  cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build(),
									Action: cb.GenerateGcpStackTemplate,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build() {
											printFlagCompletion(f)
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
									Before: ConfigRead,
									Flags:  cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build(),
									Action: cb.GenerateOpenstackStackTemplate,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build() {
											printFlagCompletion(f)
										}
									},
								},
								{
									Name:   "existing-network",
									Usage:  "creates a openstack cluster JSON template with existing network",
									Before: ConfigRead,
									Flags:  cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build(),
									Action: cb.GenerateOpenstackStackTemplate,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build() {
											printFlagCompletion(f)
										}
									},
								},
								{
									Name:   "existing-subnet",
									Usage:  "creates a openstack cluster JSON template with existing subnet",
									Before: ConfigRead,
									Flags:  cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build(),
									Action: cb.GenerateOpenstackStackTemplate,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlBlueprintNameOptional, cb.FlBlueprintFileOptional).AddAuthenticationFlags().Build() {
											printFlagCompletion(f)
										}
									},
								},
							},
						},
					},
				},
				{
					Name:   "generate-reinstall-template",
					Usage:  "generates reinstall template",
					Before: ConfigRead,
					Flags:  cb.NewFlagBuilder().AddFlags(cb.FlName, cb.FlBlueprintName).AddAuthenticationFlags().AddOutputFlag().Build(),
					Action: cb.GenerateReinstallTemplate,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlName, cb.FlBlueprintName).AddAuthenticationFlags().AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "list",
					Usage:  "lists the running clusters",
					Flags:  cb.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build(),
					Before: ConfigRead,
					Action: cb.ListStacks,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "reinstall",
					Usage:  "reinstalls a cluster",
					Before: ConfigRead,
					Flags:  cb.NewFlagBuilder().AddFlags(cb.FlName, cb.FlBlueprintNameOptional, cb.FlKerberosPasswordOptional, cb.FlKerberosPrincipalOptional, cb.FlInputJson, cb.FlWaitOptional).AddAuthenticationFlags().AddOutputFlag().Build(),
					Action: cb.ReinstallStack,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlName, cb.FlBlueprintNameOptional, cb.FlKerberosPasswordOptional, cb.FlKerberosPrincipalOptional, cb.FlInputJson, cb.FlWaitOptional).AddAuthenticationFlags().AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "repair",
					Usage:  "repaires a cluster",
					Before: ConfigRead,
					Flags:  cb.NewFlagBuilder().AddFlags(cb.FlName, cb.FlWaitOptional).AddAuthenticationFlags().AddOutputFlag().Build(),
					Action: cb.RepairStack,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlName, cb.FlWaitOptional).AddAuthenticationFlags().AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "scale",
					Usage:  "scales a cluster",
					Before: ConfigRead,
					Flags:  cb.NewFlagBuilder().AddFlags(cb.FlName, cb.FlGroupName, cb.FlDesiredNodeCount, cb.FlWaitOptional).AddAuthenticationFlags().AddOutputFlag().Build(),
					Action: cb.ScaleStack,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlName, cb.FlGroupName, cb.FlDesiredNodeCount, cb.FlWaitOptional).AddAuthenticationFlags().AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "start",
					Usage:  "starts a cluster",
					Before: ConfigRead,
					Flags:  cb.NewFlagBuilder().AddFlags(cb.FlName, cb.FlWaitOptional).AddAuthenticationFlags().AddOutputFlag().Build(),
					Action: cb.StartStack,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlName, cb.FlWaitOptional).AddAuthenticationFlags().AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "stop",
					Usage:  "stops a cluster",
					Before: ConfigRead,
					Flags:  cb.NewFlagBuilder().AddFlags(cb.FlName, cb.FlWaitOptional).AddAuthenticationFlags().AddOutputFlag().Build(),
					Action: cb.StopStack,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlName, cb.FlWaitOptional).AddAuthenticationFlags().AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "sync",
					Usage:  "synchronizes a cluster",
					Before: ConfigRead,
					Flags:  cb.NewFlagBuilder().AddFlags(cb.FlName).AddAuthenticationFlags().AddOutputFlag().Build(),
					Action: cb.SyncStack,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlName).AddAuthenticationFlags().AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "tags",
					Hidden: true,
					Usage:  "default tags related operations",
					Subcommands: []cli.Command{
						{
							Name:  "account",
							Usage: "manipulates default tags for the account",
							Subcommands: []cli.Command{
								{
									Name:   "add",
									Usage:  "adds a new default tag to the account",
									Before: ConfigRead,
									Flags:  cb.NewFlagBuilder().AddFlags(cb.FlKey, cb.FlValue).AddAuthenticationFlags().AddOutputFlag().Build(),
									Action: cb.AddAccountTag,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlKey, cb.FlValue).AddAuthenticationFlags().AddOutputFlag().Build() {
											printFlagCompletion(f)
										}
									},
								},
								{
									Name:   "delete",
									Usage:  "deletes a default tag of the account",
									Before: ConfigRead,
									Flags:  cb.NewFlagBuilder().AddFlags(cb.FlKey).AddAuthenticationFlags().AddOutputFlag().Build(),
									Action: cb.DeleteAccountTag,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlKey).AddAuthenticationFlags().AddOutputFlag().Build() {
											printFlagCompletion(f)
										}
									},
								},
								{
									Name:   "list",
									Usage:  "lists the default tags for the account",
									Before: ConfigRead,
									Flags:  cb.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build(),
									Action: cb.ListAccountTags,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build() {
											printFlagCompletion(f)
										}
									},
								},
							},
						},
					},
				},
			},
		},
		{
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
									Before: ConfigRead,
									Flags:  cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlRoleARN).AddAuthenticationFlags().Build(),
									Action: cb.CreateAwsCredential,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlRoleARN).AddAuthenticationFlags().Build() {
											printFlagCompletion(f)
										}
									},
								},
								{
									Name:   "key-based",
									Usage:  "creates a new key based aws credential",
									Before: ConfigRead,
									Flags:  cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlAccessKey, cb.FlSecretKey).AddAuthenticationFlags().Build(),
									Action: cb.CreateAwsCredential,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlAccessKey, cb.FlSecretKey).AddAuthenticationFlags().Build() {
											printFlagCompletion(f)
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
									Before: ConfigRead,
									Flags:  cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlSubscriptionId, cb.FlTenantId, cb.FlAppId, cb.FlAppPassword).AddAuthenticationFlags().Build(),
									Action: cb.CreateAzureCredential,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlSubscriptionId, cb.FlTenantId, cb.FlAppId, cb.FlAppPassword).AddAuthenticationFlags().Build() {
											printFlagCompletion(f)
										}
									},
								},
							},
						},
						{
							Name:   "gcp",
							Usage:  "creates a new gcp credential",
							Before: ConfigRead,
							Flags:  cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlProjectId, cb.FlServiceAccountId, cb.FlServiceAccountPrivateKeyFile).AddAuthenticationFlags().Build(),
							Action: cb.CreateGcpCredential,
							BashComplete: func(c *cli.Context) {
								for _, f := range cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlProjectId, cb.FlServiceAccountId, cb.FlServiceAccountPrivateKeyFile).AddAuthenticationFlags().Build() {
									printFlagCompletion(f)
								}
							},
						},
						{
							Name:  "openstack",
							Usage: "creates a new openstack credential",
							Subcommands: []cli.Command{
								{
									Name:   "keystone-v2",
									Usage:  "creates a new keystone version 2 openstack credential",
									Before: ConfigRead,
									Flags:  cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlTenantUser, cb.FlTenantPassword, cb.FlTenantName, cb.FlEndpoint, cb.FlFacingOptional).AddAuthenticationFlags().Build(),
									Action: cb.CreateOpenstackCredential,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlTenantUser, cb.FlTenantPassword, cb.FlTenantName, cb.FlEndpoint, cb.FlFacingOptional).AddAuthenticationFlags().Build() {
											printFlagCompletion(f)
										}
									},
								},
								{
									Name:   "keystone-v3",
									Usage:  "creates a new keystone version 3 openstack credential",
									Before: ConfigRead,
									Flags: cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlTenantUser, cb.FlTenantPassword,
										cb.FlUserDomain, cb.FlProjectDomainNameOptional, cb.FlProjectNameOptional, cb.FlDomainNameOptional, cb.FlKeystoneScopeOptional, cb.FlEndpoint, cb.FlFacingOptional).AddAuthenticationFlags().Build(),
									Action: cb.CreateOpenstackCredential,
									BashComplete: func(c *cli.Context) {
										for _, f := range cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlTenantUser, cb.FlTenantPassword,
											cb.FlUserDomain, cb.FlProjectDomainNameOptional, cb.FlProjectNameOptional, cb.FlDomainNameOptional, cb.FlKeystoneScopeOptional, cb.FlEndpoint, cb.FlFacingOptional).AddAuthenticationFlags().Build() {
											printFlagCompletion(f)
										}
									},
								},
							},
						},
						{
							Name:   "from-file",
							Usage:  "creates a new credential from input json file",
							Flags:  cb.NewFlagBuilder().AddResourceFlagsWithOptionalName().AddFlags(cb.FlInputJson).AddAuthenticationFlags().Build(),
							Before: ConfigRead,
							Action: cb.CreateCredentialFromFile,
							BashComplete: func(c *cli.Context) {
								for _, f := range cb.NewFlagBuilder().AddResourceFlagsWithOptionalName().AddFlags(cb.FlInputJson).AddAuthenticationFlags().Build() {
									printFlagCompletion(f)
								}
							},
						},
					},
				},
				{
					Name:   "delete",
					Usage:  "deletes a credential",
					Flags:  cb.NewFlagBuilder().AddFlags(cb.FlName).AddAuthenticationFlags().Build(),
					Before: ConfigRead,
					Action: cb.DeleteCredential,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlName).AddAuthenticationFlags().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "describe",
					Usage:  "describes a credential",
					Before: ConfigRead,
					Flags:  cb.NewFlagBuilder().AddFlags(cb.FlName).AddAuthenticationFlags().AddOutputFlag().Build(),
					Action: cb.DescribeCredential,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlName).AddAuthenticationFlags().AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "list",
					Usage:  "lists the credentials",
					Before: ConfigRead,
					Flags:  cb.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build(),
					Action: cb.ListCredentials,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
			},
		},
		{
			Name:   "ldap",
			Usage:  "ldap related operations",
			Hidden: true,
			Subcommands: []cli.Command{
				{
					Name:  "create",
					Usage: "creates a new LDAP",
					Flags: cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlLdapServer, cb.FlLdapDomain,
						cb.FlLdapBindDN, cb.FlLdapBindPassword, cb.FlLdapDirectoryType, cb.FlLdapUserSearchBase,
						cb.FlLdapUserNameAttribute, cb.FlLdapUserObjectClass, cb.FlLdapGroupMemberAttribute,
						cb.FlLdapGroupNameAttribute, cb.FlLdapGroupObjectClass, cb.FlLdapGroupSearchBase, cb.FlLdapAdminGroup).AddAuthenticationFlags().Build(),
					Before: ConfigRead,
					Action: cb.CreateLDAP,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlLdapServer, cb.FlLdapDomain,
							cb.FlLdapBindDN, cb.FlLdapBindPassword, cb.FlLdapDirectoryType, cb.FlLdapUserSearchBase,
							cb.FlLdapUserNameAttribute, cb.FlLdapUserObjectClass, cb.FlLdapGroupMemberAttribute,
							cb.FlLdapGroupNameAttribute, cb.FlLdapGroupObjectClass, cb.FlLdapGroupSearchBase, cb.FlLdapAdminGroup).AddAuthenticationFlags().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "delete",
					Usage:  "deletes an LDAP",
					Flags:  cb.NewFlagBuilder().AddFlags(cb.FlName).AddOutputFlag().AddAuthenticationFlags().Build(),
					Before: ConfigRead,
					Action: cb.DeleteLdap,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlName).AddOutputFlag().AddAuthenticationFlags().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "list",
					Usage:  "list the available ldaps",
					Flags:  cb.NewFlagBuilder().AddOutputFlag().AddAuthenticationFlags().Build(),
					Before: ConfigRead,
					Action: cb.ListLdaps,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddOutputFlag().AddAuthenticationFlags().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:  "user",
					Usage: "manage LDAP users",
					Subcommands: []cli.Command{
						{
							Name:  "create",
							Usage: "create a new LDAP user in the given base",
							Flags: cb.NewFlagBuilder().AddFlags(cb.FlLdapServer,
								cb.FlLdapBindDN, cb.FlLdapBindPassword, cb.FlLdapUserToCreate, cb.FlLdapUserToCreatePassword,
								cb.FlLdapUserToCreateBase, cb.FlLdapUserToCreateGroups, cb.FlLdapDirectoryType).Build(),
							Action: cb.CreateLdapUser,
							BashComplete: func(c *cli.Context) {
								for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlLdapServer,
									cb.FlLdapBindDN, cb.FlLdapBindPassword, cb.FlLdapUserToCreate, cb.FlLdapUserToCreatePassword,
									cb.FlLdapUserToCreateBase, cb.FlLdapUserToCreateGroups, cb.FlLdapDirectoryType).Build() {
									printFlagCompletion(f)
								}
							},
						},
						{
							Name:  "delete",
							Usage: "delete a user from LDAP",
							Flags: cb.NewFlagBuilder().AddFlags(cb.FlLdapServer,
								cb.FlLdapBindDN, cb.FlLdapBindPassword, cb.FlLdapUserToDelete,
								cb.FlLdapUserToDeleteBase, cb.FlLdapDirectoryType).Build(),
							Action: cb.DeleteLdapUser,
							BashComplete: func(c *cli.Context) {
								for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlLdapServer,
									cb.FlLdapBindDN, cb.FlLdapBindPassword, cb.FlLdapUserToDelete,
									cb.FlLdapUserToDeleteBase, cb.FlLdapDirectoryType).Build() {
									printFlagCompletion(f)
								}
							},
						},
					},
				},
			},
		},
		{
			Name:  "imagecatalog",
			Usage: "imagecatalog related operations",
			Subcommands: []cli.Command{
				{
					Name:   "create",
					Usage:  "creates a new imagecatalog from a URL",
					Flags:  cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlURL).AddAuthenticationFlags().Build(),
					Before: ConfigRead,
					Action: cb.CreateImagecatalogFromUrl,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlURL).AddAuthenticationFlags().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "delete",
					Usage:  "deletes an imagecatalog",
					Flags:  cb.NewFlagBuilder().AddFlags(cb.FlName).AddOutputFlag().AddAuthenticationFlags().Build(),
					Before: ConfigRead,
					Action: cb.DeleteImagecatalog,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlName).AddOutputFlag().AddAuthenticationFlags().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:  "images",
					Usage: "lists available images from imagecatalog",
					Subcommands: []cli.Command{
						{
							Name:   "aws",
							Usage:  "lists available aws images from an imagecatalog",
							Flags:  cb.NewFlagBuilder().AddFlags(cb.FlImageCatalog).AddOutputFlag().AddAuthenticationFlags().Build(),
							Before: ConfigRead,
							Action: cb.ListAwsImages,
							BashComplete: func(c *cli.Context) {
								for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlImageCatalog).AddOutputFlag().AddAuthenticationFlags().Build() {
									printFlagCompletion(f)
								}
							},
						},
						{
							Name:   "azure",
							Usage:  "lists available azure images from an imagecatalog",
							Flags:  cb.NewFlagBuilder().AddFlags(cb.FlImageCatalog).AddOutputFlag().AddAuthenticationFlags().Build(),
							Before: ConfigRead,
							Action: cb.ListAzureImages,
							BashComplete: func(c *cli.Context) {
								for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlImageCatalog).AddOutputFlag().AddAuthenticationFlags().Build() {
									printFlagCompletion(f)
								}
							},
						},
						{
							Name:   "gcp",
							Usage:  "lists available gcp images from an imagecatalog",
							Flags:  cb.NewFlagBuilder().AddFlags(cb.FlImageCatalog).AddOutputFlag().AddAuthenticationFlags().Build(),
							Before: ConfigRead,
							Action: cb.ListGcpImages,
							BashComplete: func(c *cli.Context) {
								for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlImageCatalog).AddOutputFlag().AddAuthenticationFlags().Build() {
									printFlagCompletion(f)
								}
							},
						},
						{
							Name:   "openstack",
							Usage:  "lists available openstack images from an imagecatalog",
							Flags:  cb.NewFlagBuilder().AddFlags(cb.FlImageCatalog).AddOutputFlag().AddAuthenticationFlags().Build(),
							Before: ConfigRead,
							Action: cb.ListOpenstackImages,
							BashComplete: func(c *cli.Context) {
								for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlImageCatalog).AddOutputFlag().AddAuthenticationFlags().Build() {
									printFlagCompletion(f)
								}
							},
						},
					},
				},
				{
					Name:   "list",
					Usage:  "lists the available imagecatalogs",
					Flags:  cb.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build(),
					Before: ConfigRead,
					Action: cb.ListImagecatalogs,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "set-default",
					Usage:  "sets the default imagecatalog",
					Flags:  cb.NewFlagBuilder().AddFlags(cb.FlName).AddOutputFlag().AddAuthenticationFlags().Build(),
					Before: ConfigRead,
					Action: cb.SetDefaultImagecatalog,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlName).AddOutputFlag().AddAuthenticationFlags().Build() {
							printFlagCompletion(f)
						}
					},
				},
			},
		},
		{
			Name:  "recipe",
			Usage: "recipe related operations",
			Subcommands: []cli.Command{
				{
					Name:  "create",
					Usage: "adds a new recipe from a file or from a URL",
					Subcommands: []cli.Command{
						{
							Name:   "from-url",
							Flags:  cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlExecutionType, cb.FlURL).AddAuthenticationFlags().Build(),
							Before: ConfigRead,
							Action: cb.CreateRecipeFromUrl,
							Usage:  "creates a recipe by downloading it from a URL location",
							BashComplete: func(c *cli.Context) {
								for _, f := range cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlExecutionType, cb.FlURL).AddAuthenticationFlags().Build() {
									printFlagCompletion(f)
								}
							},
						},
						{
							Name:   "from-file",
							Flags:  cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlExecutionType, cb.FlFile).AddAuthenticationFlags().Build(),
							Before: ConfigRead,
							Action: cb.CreateRecipeFromFile,
							Usage:  "creates a recipe by reading it from a local file",
							BashComplete: func(c *cli.Context) {
								for _, f := range cb.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(cb.FlExecutionType, cb.FlFile).AddAuthenticationFlags().Build() {
									printFlagCompletion(f)
								}
							},
						},
					},
				},
				{
					Name:   "delete",
					Usage:  "deletes a recipe",
					Flags:  cb.NewFlagBuilder().AddFlags(cb.FlName).AddOutputFlag().AddAuthenticationFlags().Build(),
					Before: ConfigRead,
					Action: cb.DeleteRecipe,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlName).AddOutputFlag().AddAuthenticationFlags().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "describe",
					Usage:  "describes a recipe",
					Before: ConfigRead,
					Flags:  cb.NewFlagBuilder().AddFlags(cb.FlName).AddAuthenticationFlags().AddOutputFlag().Build(),
					Action: cb.DescribeRecipe,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddFlags(cb.FlName).AddAuthenticationFlags().AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
				{
					Name:   "list",
					Usage:  "lists the available recipes",
					Flags:  cb.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build(),
					Before: ConfigRead,
					Action: cb.ListRecipes,
					BashComplete: func(c *cli.Context) {
						for _, f := range cb.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build() {
							printFlagCompletion(f)
						}
					},
				},
			},
		},
	}
	sortByName(app.Commands)

	// internal commands
	app.Commands = append(app.Commands, []cli.Command{
		{
			Name:   "internal",
			Usage:  "shows the internal commands",
			Hidden: true,
			Action: cb.ShowHiddenCommands,
		},
	}...)

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
