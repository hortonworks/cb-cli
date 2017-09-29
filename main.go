package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	hdc "github.com/hortonworks/hdc-cli/cli"
	_ "github.com/hortonworks/hdc-cli/cli/cloud/aws"
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

	server := c.String(hdc.FlServer.Name)
	username := c.String(hdc.FlUsername.Name)
	password := c.String(hdc.FlPassword.Name)
	output := c.String(hdc.FlOutput.Name)

	config, err := hdc.ReadConfig(hdc.GetHomeDirectory())
	if err == nil {
		if len(output) == 0 {
			c.Set(hdc.FlOutput.Name, config.Output)
		}
	}

	if len(server) == 0 || len(username) == 0 || len(password) == 0 {
		if err != nil {
			log.Error(fmt.Sprintf("configuration is not set, see: hdc configure --help or provide the following flags: %v",
				[]string{"--" + hdc.FlServer.Name, "--" + hdc.FlUsername.Name, "--" + hdc.FlPassword.Name}))
			os.Exit(1)
		}

		PrintConfig(*config)

		if len(server) == 0 {
			c.Set(hdc.FlServer.Name, config.Server)
		}
		if len(username) == 0 {
			c.Set(hdc.FlUsername.Name, config.Username)
		}
		if len(password) == 0 {
			c.Set(hdc.FlPassword.Name, config.Password)
		}
	}
	return nil
}

func PrintConfig(cfg hdc.Config) {
	cfg.Password = "*"
	log.Infof("[ConfigRead] Config read from file, setting as global variable:\n%s", cfg.Yaml())
}

func printFlagCompletion(f cli.Flag) {
	fmt.Printf("--%s\n", f.GetName())
}

func main() {

	app := cli.NewApp()
	app.Name = "hdc"
	app.HelpName = "Hortonworks Data Cloud command line tool"
	app.Version = hdc.Version + "-" + hdc.BuildTime
	app.Author = "Hortonworks"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		hdc.FlDebug,
	}

	app.Before = func(c *cli.Context) error {
		log.SetOutput(os.Stderr)
		log.SetLevel(log.ErrorLevel)
		formatter := &hdc.CBFormatter{}
		log.SetFormatter(formatter)
		if c.Bool(hdc.FlDebug.Name) {
			log.SetLevel(log.DebugLevel)
		}
		return nil
	}

	cli.AppHelpTemplate = hdc.AppHelpTemplate
	cli.HelpPrinter = hdc.PrintHelp
	cli.CommandHelpTemplate = hdc.CommandHelpTemplate
	cli.SubcommandHelpTemplate = hdc.SubCommandHelpTemplate

	app.Commands = []cli.Command{
		{
			Name: "configure",
			Description: fmt.Sprintf("it will save the provided server address and credential "+
				"to %s/%s/%s", hdc.GetHomeDirectory(), hdc.Hdc_dir, hdc.Config_file),
			Usage:  "configure the server address and credentials used to communicate with this server",
			Flags:  hdc.NewFlagBuilder().AddFlags(hdc.FlServerRequired, hdc.FlUsernameRequired, hdc.FlPassword).AddOutputFlag().Build(),
			Action: hdc.Configure,
			BashComplete: func(c *cli.Context) {
				for _, f := range hdc.NewFlagBuilder().AddFlags(hdc.FlServerRequired, hdc.FlUsernameRequired, hdc.FlPassword).AddOutputFlag().Build() {
					printFlagCompletion(f)
				}
			},
		},
		{
			Name:  "create-credential",
			Usage: "create a new credential",
			Subcommands: []cli.Command{
				{
					Name:  "aws",
					Usage: "create a new aws credential",
					Subcommands: []cli.Command{
						{
							Name:   "role-based",
							Usage:  "create a new role based aws credential",
							Before: ConfigRead,
							Flags:  hdc.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(hdc.FlRoleARN).AddAuthenticationFlags().Build(),
							Action: hdc.CreateAwsCredential,
							BashComplete: func(c *cli.Context) {
								for _, f := range hdc.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(hdc.FlRoleARN).AddAuthenticationFlags().Build() {
									printFlagCompletion(f)
								}
							},
						},
						{
							Name:   "key-based",
							Usage:  "create a new key based aws credential",
							Before: ConfigRead,
							Flags:  hdc.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(hdc.FlAccessKey, hdc.FlSecretKey).AddAuthenticationFlags().Build(),
							Action: hdc.CreateAwsCredential,
							BashComplete: func(c *cli.Context) {
								for _, f := range hdc.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(hdc.FlAccessKey, hdc.FlSecretKey).AddAuthenticationFlags().Build() {
									printFlagCompletion(f)
								}
							},
						},
					},
				},
			},
		},
		{
			Name:   "list-credentials",
			Usage:  "list the credentials",
			Before: ConfigRead,
			Flags:  hdc.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build(),
			Action: hdc.ListPrivateCredentials,
			BashComplete: func(c *cli.Context) {
				for _, f := range hdc.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build() {
					printFlagCompletion(f)
				}
			},
		},
		{
			Name:   "delete-credential",
			Usage:  "delete a credential",
			Flags:  hdc.NewFlagBuilder().AddFlags(hdc.FlName).AddAuthenticationFlags().Build(),
			Before: ConfigRead,
			Action: hdc.DeleteCredential,
			BashComplete: func(c *cli.Context) {
				for _, f := range hdc.NewFlagBuilder().AddFlags(hdc.FlName).AddAuthenticationFlags().Build() {
					printFlagCompletion(f)
				}
			},
		},
	}

	// internal commands
	app.Commands = append(app.Commands, []cli.Command{
		{
			Name:   "internal",
			Usage:  "shows the internal commands",
			Hidden: true,
			Action: hdc.ShowHiddenCommands,
		},
	}...)

	app.Run(os.Args)
}
