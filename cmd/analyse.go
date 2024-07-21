package cmd

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"

	"github.com/spf13/cobra"
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
		// TODO: check files exist, do what we can or add a strict flag

		// TODO: Asynchronously process all files, all lines scaling out massively
		// TODO: matching patterns in the config file
		for _, f := range flattened {
			opened, err := os.Open(f.Path)
			if err != nil {
				panic(err)
			}
			// TODO: don't defer in the loop!
			defer opened.Close()

			scanner := bufio.NewScanner(opened)
			for scanner.Scan() {
				line := scanner.Text()
				for _, pattern := range f.Threshold.Patterns {
					ok, err := regexp.Match(pattern, []byte(line))
					if err != nil {
						slog.Error("error matching line with pattern", slog.String("line", line), slog.String("pattern", pattern))
					}
					if ok {
						slog.Info("matched", slog.String("line", line), slog.String("pattern", pattern))
					}
				}
			}
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
