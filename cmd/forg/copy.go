package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/imgoomes/forgetful/internal/storage"
	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:     "copy <id>",
	Aliases: []string{"cp"},
	Short:   "Copy to clipboard a saved command by ID",
	Long:    `Copy to clipboard a saved command directly by its ID. Use 'forg list' to see available IDs.`,
	Example: `  forg copy 1
  forg cp 5`,
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
				if err := writeToClipboard(e.Command); err != nil {
					return fmt.Errorf("clipboard error: %w", err)
				}
				fmt.Printf("$ %s\n", e.Command)
				fmt.Println("Copied to clipboard!")
			}
		}

		if !found {
			return fmt.Errorf("no command with ID %d found", id)
		}
		return nil
	},
}

func writeToClipboard(text string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "linux":
		if _, err := exec.LookPath("xclip"); err == nil {
			cmd = exec.Command("xclip", "-selection", "clipboard")
		} else if _, err := exec.LookPath("xsel"); err == nil {
			cmd = exec.Command("xsel", "--clipboard", "--input")
		} else {
			return fmt.Errorf("no clipboard tool found; install xclip or xsel")
		}
	default:
		return fmt.Errorf("clipboard not supported on %s", runtime.GOOS)
	}
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}
