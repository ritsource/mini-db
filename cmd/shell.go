package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shell called")
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
}
