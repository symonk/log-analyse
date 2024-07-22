package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/symonk/log-analyse/internal/analyser"
	"github.com/symonk/log-analyse/internal/files"
)

// analyseCmd represents the analyse command
// TODO: This implementation is very WIP and is a piece of rubbish atm.
var analyseCmd = &cobra.Command{
	Use:   "analyse",
	Short: "Analyses log fails based on the configuration",
	Long:  `Implement`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("Detected globs", slog.Any("globs", cfg.Globs()))
		// instantiate something that can locate files on disk
		locator := files.NewFileLocator(cfg)
		flattened, err := locator.Locate()
		if err != nil {
			slog.Error("unable to parse files", slog.Any("error", err))
		}
		for _, path := range flattened {
			slog.Info("Will scan file", slog.Any("file", path.Path))
		}
		fAnalyser := analyser.NewFileAnalyser(flattened, analyser.WithBounds(0))
		matches, err := fAnalyser.Analyse()
		if err != nil {
			slog.Error("error analysing", slog.Any("error", err))
		}
		for match := range matches {
			slog.Info("Match", slog.Any("match", match))
		}

		// TODO: Collect matches for each of the thresholds

		// TODO: live notify as we go or fan in all results
		// TODO: notify at the end of the run
	},
}

func init() {
	rootCmd.AddCommand(analyseCmd)
	analyseCmd.Flags().BoolP("strict", "s", false, "if set any log files in the config that do not exist on disk will cause an exit")
}
