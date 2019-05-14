package cmd

import (
	"github.com/ritwik310/mini-db/server"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts a TCP-server in which database is accesable",
	Long: `Starts a TCP-server on specified port,
client can communicate to this the address.

Example: mini-db server -p 8000 --backup
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		port, err := cmd.Flags().GetString("port")     // Port
		backup, err := cmd.Flags().GetBool("backup")   // To persist the data or not
		delay, err := cmd.Flags().GetInt("delay")      // Time delay on data snapshot
		output, err := cmd.Flags().GetString("output") // Data snapshot-file location

		if err != nil {
			return err
		}

		// Running the server
		server.Start(port, backup, delay, output)

		return nil
	},
}

func init() {
	// Flags for Start-Server-Cmd
	serverCmd.Flags().StringP("port", "p", "8080", "The port to run the server")
	serverCmd.Flags().BoolP("backup", "b", false, "To persist data or not (just in-memory)")
	serverCmd.Flags().IntP("delay", "d", 5, "Time delay to for data-snapshot to persistent disk in seconds")
	serverCmd.Flags().StringP("output", "o", "", "Filepath to for saving the persistent data on")

	rootCmd.AddCommand(serverCmd)
}
