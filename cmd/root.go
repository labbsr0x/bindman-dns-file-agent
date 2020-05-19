package cmd

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:          "bindman-dns-file-agent",
	Short:        "Bindman DNS File Agent to Bindman DNS Bind9",
	Long:         "The bindmand-dns-file-agent is an agent responsible for persisting in bindman-dns-manager by a file.",
	SilenceUsage: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetEnvPrefix("BINDMAN") // all bindman-dns-file-agent environment variables must be prefixed with BINDMAN_
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.AutomaticEnv() // read in environment variables that match
}
