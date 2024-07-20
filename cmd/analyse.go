package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// analyseCmd represents the analyse command
var analyseCmd = &cobra.Command{
	Use:   "analyse",
	Short: "Analyses log fails based on the configuration",
	Long:  `Implement`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("analyse called")
	},
}

func init() {
	rootCmd.AddCommand(analyseCmd)
}
