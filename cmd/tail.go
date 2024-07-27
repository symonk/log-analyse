package cmd

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"
	"github.com/symonk/log-analyse/internal/files"
	"github.com/symonk/log-analyse/internal/prof"
)

// tailCmd represents the tail command
var tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Live tailing of log files",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		/*
			Parse globs into concrete files
			Flatten the files to remove any duplicates?
		*/
		if profile {
			defer prof.RunProf()()
		}
		fileLocator := files.NewFileLocator(cfg)
		squashedFiles, err := fileLocator.Locate()
		if err != nil {
			return fmt.Errorf("error when resolving glob patterns to files %w", err)
		}
		var wg sync.WaitGroup
		wg.Add(len(squashedFiles))
		for _, f := range squashedFiles {
			fmt.Println(f.Path)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(tailCmd)
}
