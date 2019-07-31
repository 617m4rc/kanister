package kando

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/kanisterio/kanister/pkg/version"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	root := newRootCommand()
	if err := root.Execute(); err != nil {
		log.Errorf("%+v", err)
		os.Exit(1)
	}
}

func newRootCommand() *cobra.Command {
	// RootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:     "kando <command>",
		Short:   "A set of tools used from Kanister Blueprints",
		Version: version.VersionString(),
	}
	rootCmd.AddCommand(newLocationCommand())
	rootCmd.AddCommand(newOutputCommand())
	rootCmd.AddCommand(newChronicleCommand())
	return rootCmd
}
