package cmd

import (
	"fmt"

	cobra "github.com/spf13/cobra"
)

const cmdName = "ml_practice_ctl"

var usage string

var RootCmd = &cobra.Command{
	Use: cmdName,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(usage)
	},
}

func init() {
	cobra.OnInitialize()
	usage = RootCmd.UsageString()
}
