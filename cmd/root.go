package cmd

import (
	"github.com/RafalSkolasinski/dedent"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kafctl",
	Short: "A simple CLI for testing Kafka configs",
	Long: dedent.Dedent(`
		Helper tool to test various Kafka configurations.
	`),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "List resource",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create resource",
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove resource",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	err := rootCmd.Execute()
	return err
}

func init() {
	// Top Level Flags
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./client.properties", "config file")

	// Commands that represents action on specific resource
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(removeCmd)
}
