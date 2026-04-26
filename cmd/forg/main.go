package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "forg",
	Short:        "forg — never forget a terminal command again",
	SilenceUsage: true,
	Long: `forg saves terminal commands with tags and descriptions
so you can find and run them whenever you need.`,
}

func main() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(tagsCmd)
	rootCmd.AddCommand(copyCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
