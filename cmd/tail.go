package cmd

import (
	"github.com/spf13/cobra"
	"github.com/symonk/log-analyse/internal/prof"
)

// tailCmd represents the tail command
var tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Live tailing of log files",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/*
			Parse globs into concrete files
			Flatten the files to remove any duplicates?
		*/
		if profile {
			defer prof.RunProf()()
		}
	},
}

func init() {
	rootCmd.AddCommand(tailCmd)
}
