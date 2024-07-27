/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/symonk/log-analyse/internal/config"
)

var (
	cfgFile string
	verbose bool
	profile bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "log-analyse",
	Short: "Highly scalable log analysis with builtin alerting",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() { config.Init(cfgFile) })
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/.loganalyse/loganalyse.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "display more information to stdout")
	rootCmd.PersistentFlags().BoolVarP(&profile, "profile", "p", false, "enable CPU profiler (temporary dev aid)")
}
