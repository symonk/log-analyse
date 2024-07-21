package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/symonk/log-analyse/internal/files"
)

// analyseCmd represents the analyse command
var analyseCmd = &cobra.Command{
	Use:   "analyse",
	Short: "Analyses log fails based on the configuration",
	Long:  `Implement`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("Detected globs", slog.Any("globs", cfg.Globs()))
		// instantiate something that can locate files on disk
		locator := files.NewFileLocator()
		condenser := files.NewCondenser(locator)
		// flatten the config to individual files
		// requires alot of testing
		flattened, err := condenser.Flatten(cfg.Files)
		_, _ = flattened, err

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
