package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "charon",
	Short: "Charon is an open source content moderator.",
	Long: `Charon is an open source content moderator that allows
			users to define the way they want to filter content,
			allowing their communities to stay positive, curious, and growing.`,
}

// Execute runs the command provided
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
