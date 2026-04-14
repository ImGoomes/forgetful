package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/gabrielgomes/forgetful/internal/storage"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:          "run <id>",
	Short:        "Execute a saved command by ID",
	Long:         `Execute a saved command directly by its ID. Use 'forg list' to see available IDs.`,
	Example:      `  forg run 1
  forg run 5`,
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid ID: %q", args[0])
		}

		store, err := storage.Load()
		if err != nil {
			return err
		}

		var found bool
		for _, e := range store.Entries {
			if e.ID == id {
				found = true
				fmt.Printf("$ %s\n", e.Command)

				var c *exec.Cmd
				if runtime.GOOS == "windows" {
					c = exec.Command("cmd", "/C", e.Command)
				} else {
					c = exec.Command("sh", "-c", e.Command)
				}
				c.Stdout = os.Stdout
				c.Stderr = os.Stderr
				c.Stdin = os.Stdin
				return c.Run()
			}
		}

		if !found {
			return fmt.Errorf("no command with ID %d found", id)
		}
		return nil
	},
}
