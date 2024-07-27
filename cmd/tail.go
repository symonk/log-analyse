package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/symonk/log-analyse/internal/config"
	"github.com/symonk/log-analyse/internal/files"
	"github.com/symonk/log-analyse/internal/monitor"
	"github.com/symonk/log-analyse/internal/prof"
)

// tailCmd represents the tail command
var tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "tails log files for configured pattern matches",
	Long: `The tail subcommand is used for live tailing all files
	matched by the glob patterns defined in the configs.  Each matching
	file is monitored for corresponding sets of option(s) and an action
	can be configured on successful matches.
	
	See configuration documentation for more.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if profile {
			defer prof.RunProf()()
		}

		/*
			Fetch the user config
			Resolve the glob(s) into files that exist
			Merge the files that exist from all the globs
			Register monitors on the file(s) to tail them
		*/

		cfg := config.Get()
		fileLocator := files.NewFileLocator(cfg)
		squashedFiles, err := fileLocator.Locate()
		if err != nil {
			return fmt.Errorf("error when resolving glob patterns to files %w", err)
		}
		monitor := monitor.Filemon{}
		done := make(chan struct{})
		for _, f := range squashedFiles {
			go func() {
				monitor.Watch(f, done)
			}()
		}
		<-done
		return nil
	},
}

func init() {
	rootCmd.AddCommand(tailCmd)
}
