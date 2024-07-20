package cmd

import (
	"github.com/spf13/cobra"
)

// analyseCmd represents the analyse command
var analyseCmd = &cobra.Command{
	Use:   "analyse",
	Short: "Analyses log fails based on the configuration",
	Long:  `Implement`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: pull out the locations from the config

		// TODO: flatten down to single files in terms of thresholding
		// TODO: because folders can be set, or explicit files

		// TODO: check files exist, do what we can or add a strict flag

		// TODO: Asynchronously process all files, all lines scaling out massively
		// TODO: matching patterns in the config file

		// TODO: Collect matches for each of the thresholds

		// TODO: live notify as we go or fan in all results
		// TODO: notify at the end of the run
	},
}

func init() {
	rootCmd.AddCommand(analyseCmd)
	analyseCmd.Flags().BoolP("strict", "s", false, "if set any log files in the config that do not exist on disk will cause an exit")
}
