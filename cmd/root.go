/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/symonk/log-analyse/internal/config"
)

var cfgFile string

var cfg *config.Config

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
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/.loganalyse/loganalyse.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		baseDir, err := config.ConfigDefaultFolder()
		cobra.CheckErr(err)

		viper.AddConfigPath(baseDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".log-analyse")
	}

	// viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		if err := viper.Unmarshal(&cfg); err != nil {
			fmt.Fprintf(os.Stderr, "yaml file is not valid config: %s because %s", viper.ConfigFileUsed(), err)
			os.Exit(2)
		}
	} else {
		fmt.Fprintln(os.Stderr, "no config file could be found.")
		os.Exit(1)
	}
	fmt.Println("Successfully built a config!")

}
