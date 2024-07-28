package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func parseViper(configFilePath string) error {
	if configFilePath != "" {
		viper.SetConfigFile(configFilePath)
	} else {
		baseDir, err := defaultConfigPath()
		cobra.CheckErr(err)

		viper.AddConfigPath(baseDir)
		viper.SetConfigType(configType)
		viper.SetConfigName(configName)
	}

	if err := viper.ReadInConfig(); err == nil {
		if err := viper.Unmarshal(&GlobalConfig); err != nil {
			return fmt.Errorf("unable to unmarshal config: %w", err)
		}
	} else {
		return fmt.Errorf("no config found: %w", err)
	}
	return nil
}
