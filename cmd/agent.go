package cmd

import (
	"fmt"

	"github.com/labbsr0x/bindman-dns-file-agent/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var agentCmd = &cobra.Command{

	Use:     "agent",
	Aliases: []string{"a"},
	Short:   "Starts the sync agent with Bindman",
	Long:    "Starts the sync agent with Bindman",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Making something")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)

	config.AddFlags(agentCmd.Flags())

	err := viper.GetViper().BindPFlags(agentCmd.Flags())
	if err != nil {
		panic(err)
	}
}
