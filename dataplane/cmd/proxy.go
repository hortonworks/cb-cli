package cmd

import (
	cf "github.com/hortonworks/cb-cli/dataplane/config"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/proxy"
	"github.com/urfave/cli"
)

func init() {
	DataPlaneCommands = append(DataPlaneCommands, cli.Command{
		Name:  "proxy",
		Usage: "proxy related operations",
		Subcommands: []cli.Command{
			{
				Name:  "create",
				Usage: "creates a new proxy",
				Flags: fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlProxyHost, fl.FlProxyPort,
					fl.FlProxyProtocol, fl.FlProxyUser, fl.FlProxyPassword, fl.FlProxyNoProxyHosts).AddAGlobalFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: proxy.CreateProxy,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddResourceDefaultFlags().AddFlags(fl.FlProxyHost, fl.FlProxyPort,
						fl.FlProxyProtocol, fl.FlProxyUser, fl.FlProxyPassword, fl.FlProxyNoProxyHosts).AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "delete",
				Usage:  "deletes a proxy",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlName).AddOutputFlag().AddAGlobalFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: proxy.DeleteProxy,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlName).AddOutputFlag().AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "list",
				Usage:  "list the available proxies",
				Flags:  fl.NewFlagBuilder().AddOutputFlag().AddAGlobalFlags().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: proxy.ListProxies,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddOutputFlag().AddAGlobalFlags().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
		},
	})
}
