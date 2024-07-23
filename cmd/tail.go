package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tailCmd represents the tail command
var tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Live tailing of log files",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tail called")
	},
}

func init() {
	rootCmd.AddCommand(tailCmd)
}
