package main

import (
	"github.com/spf13/cobra"
	"github.com/shortwavedave/distribution-tooling-for-helm/cmd/dt/annotate"
	"github.com/shortwavedave/distribution-tooling-for-helm/cmd/dt/carvelize"
	"github.com/shortwavedave/distribution-tooling-for-helm/cmd/dt/relocate"
)

var chartCmd = &cobra.Command{
	Use:           "charts",
	Short:         "Helm chart management commands",
	SilenceUsage:  true,
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, _ []string) {
		_ = cmd.Help()
	},
}

func init() {
	chartCmd.AddCommand(relocate.NewCmd(mainConfig), annotate.NewCmd(mainConfig), carvelize.NewCmd(mainConfig))
}
