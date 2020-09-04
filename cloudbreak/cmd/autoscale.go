package cmd

import (
	as "github.com/hortonworks/cb-cli/cloudbreak/autoscale"
	cf "github.com/hortonworks/cb-cli/cloudbreak/config"
	fl "github.com/hortonworks/cb-cli/cloudbreak/flags"
	"github.com/urfave/cli"
)

func init() {
	CloudbreakCommands = append(CloudbreakCommands, cli.Command{
		Name:  "as",
		Usage: "autoscale related operations",
		Subcommands: []cli.Command{
			{
				Name:   "enable",
				Usage:  "enables autoscale for cluster",
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlClusterID).AddAuthenticationFlags().AddOutputFlag().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: as.Enable,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlClusterID).AddAuthenticationFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "disable",
				Usage:  "disables autoscale for cluster",
				Before: cf.CheckConfigAndCommandFlags,
				Flags:  fl.NewFlagBuilder().AddFlags(fl.FlClusterID).AddAuthenticationFlags().AddOutputFlag().Build(),
				Action: as.Disable,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlClusterID).AddAuthenticationFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:   "clusters",
				Usage:  "lists autoscale clusters",
				Flags:  fl.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build(),
				Before: cf.CheckConfigAndCommandFlags,
				Action: as.ListClusters,
				BashComplete: func(c *cli.Context) {
					for _, f := range fl.NewFlagBuilder().AddAuthenticationFlags().AddOutputFlag().Build() {
						fl.PrintFlagCompletion(f)
					}
				},
			},
			{
				Name:  "policy",
				Usage: "autoscale policy configuration",
				Subcommands: []cli.Command{
					{
						Name:   "list",
						Usage:  "list autoscale policies",
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlAsID).AddAuthenticationFlags().AddOutputFlag().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: as.ListPolicies,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlAsID).AddAuthenticationFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:   "create",
						Usage:  "create autoscale policy",
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlAsID, fl.FlName, fl.FlAlertID, fl.FlHostgroup, fl.FlScalingAdjustment, fl.FlAdjustmentType).AddAuthenticationFlags().AddOutputFlag().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: as.CreatePolicy,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlAsID, fl.FlName, fl.FlAlertID, fl.FlHostgroup, fl.FlScalingAdjustment, fl.FlAdjustmentType).AddAuthenticationFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:   "update",
						Usage:  "update autoscale policy",
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlAsID, fl.FlPolicyID, fl.FlName, fl.FlAlertID, fl.FlHostgroup, fl.FlScalingAdjustment, fl.FlAdjustmentType).AddAuthenticationFlags().AddOutputFlag().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: as.UpdatePolicy,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlAsID, fl.FlPolicyID, fl.FlName, fl.FlAlertID, fl.FlHostgroup, fl.FlScalingAdjustment, fl.FlAdjustmentType).AddAuthenticationFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:   "delete",
						Usage:  "delete autoscale policy",
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlAsID, fl.FlPolicyID).AddAuthenticationFlags().AddOutputFlag().Build(),
						Before: cf.CheckConfigAndCommandFlags,
						Action: as.DeletePolicy,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlAsID, fl.FlPolicyID).AddAuthenticationFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:  "alert",
				Usage: "autoscale alert configuration",
				Subcommands: []cli.Command{
					{
						Name:  "metric",
						Usage: "metric alert commands",
						Subcommands: []cli.Command{
							{
								Name:   "list",
								Usage:  "lists metric alerts",
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlAsID).AddAuthenticationFlags().AddOutputFlag().Build(),
								Before: cf.CheckConfigAndCommandFlags,
								Action: as.ListMetricAlerts,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlAsID).AddAuthenticationFlags().AddOutputFlag().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "create",
								Usage:  "creates metric alert",
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlAsID, fl.FlName, fl.FlAlertDefinition, fl.FlAlertStateOptional, fl.FlAlertPeriodOptional).AddAuthenticationFlags().AddOutputFlag().Build(),
								Before: cf.CheckConfigAndCommandFlags,
								Action: as.CreateMetricAlert,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlAsID, fl.FlName, fl.FlAlertDefinition, fl.FlAlertStateOptional, fl.FlAlertPeriodOptional).AddAuthenticationFlags().AddOutputFlag().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "delete",
								Usage:  "deletes a metric alert",
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlAsID, fl.FlAlertID).AddAuthenticationFlags().AddOutputFlag().Build(),
								Before: cf.CheckConfigAndCommandFlags,
								Action: as.DeleteMetricAlert,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlAsID, fl.FlAlertID).AddAuthenticationFlags().AddOutputFlag().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "definitions",
								Usage:  "lists metric alert definitions",
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlAsID).AddAuthenticationFlags().AddOutputFlag().Build(),
								Before: cf.CheckConfigAndCommandFlags,
								Action: as.ListMetricAlertDefinitions,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlAsID).AddAuthenticationFlags().AddOutputFlag().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
					{
						Name:  "time",
						Usage: "time alert commands",
						Subcommands: []cli.Command{
							{
								Name:   "list",
								Usage:  "lists time alerts",
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlAsID).AddAuthenticationFlags().AddOutputFlag().Build(),
								Before: cf.CheckConfigAndCommandFlags,
								Action: as.ListTimeAlerts,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlAsID).AddAuthenticationFlags().AddOutputFlag().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "create",
								Usage:  "creates time alert",
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlAsID, fl.FlName, fl.FlCron, fl.FlTimezoneOptional).AddAuthenticationFlags().AddOutputFlag().Build(),
								Before: cf.CheckConfigAndCommandFlags,
								Action: as.CreateTimeAlert,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlAsID, fl.FlName, fl.FlCron, fl.FlTimezoneOptional).AddAuthenticationFlags().AddOutputFlag().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
							{
								Name:   "delete",
								Usage:  "deletes a time alert",
								Flags:  fl.NewFlagBuilder().AddFlags(fl.FlAsID, fl.FlAlertID).AddAuthenticationFlags().AddOutputFlag().Build(),
								Before: cf.CheckConfigAndCommandFlags,
								Action: as.DeleteTimeAlert,
								BashComplete: func(c *cli.Context) {
									for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlAsID, fl.FlAlertID).AddAuthenticationFlags().AddOutputFlag().Build() {
										fl.PrintFlagCompletion(f)
									}
								},
							},
						},
					},
				},
			},
		},
	})
}
