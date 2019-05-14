package cmd

import (
	"github.com/ritwik310/mini-db/shell"
	"github.com/spf13/cobra"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		backup, err := cmd.Flags().GetBool("backup")   // To persist the data or not
		delay, err := cmd.Flags().GetInt("delay")      // Time delay on data snapshot
		output, err := cmd.Flags().GetString("output") // Data snapshot-file location

		if err != nil {
			return err
		}

		// Running the shell
		shell.Start(backup, delay, output)

		return nil
	},
}

func init() {
	// Flags for Start-Server-Cmd
	shellCmd.Flags().BoolP("backup", "b", false, "To persist data or not (just in-memory)")
	shellCmd.Flags().IntP("delay", "d", 5, "Time delay to for data-snapshot to persistent disk in seconds")
	shellCmd.Flags().StringP("output", "o", "", "Filepath to for saving the persistent data on")

	rootCmd.AddCommand(shellCmd)
}
