/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/symonk/log-analyse/internal/config"
)

var (
	cfgFile string
	cfg     *config.Config
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
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/.loganalyse/loganalyse.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "display more information to stdout")
	rootCmd.PersistentFlags().BoolVarP(&profile, "profile", "p", false, "if set enables cpu profiling as a development aid")
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
		viper.SetConfigName("log-analyse")
	}

	// viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		if err := viper.Unmarshal(&cfg); err != nil {
			slog.Error("configuration file was not valid", slog.String("config", viper.ConfigFileUsed()), slog.Any("error", err))
			os.Exit(2)
		}
	} else {
		slog.Error("no config file could be found: ", slog.Any("error", err))
		os.Exit(1)
	}
	slog.Info("Successfully built a config")
}
