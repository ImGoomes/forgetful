package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gabrielgomes/forgetful/internal/storage"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete <id>",
	Aliases: []string{"rm", "remove"},
	Short:   "Remove a saved command by ID",
	Long:    `Remove a saved command by its ID. Prompts for confirmation before deleting. Use 'forg list' to see IDs.`,
	Example: `  forg delete 3
  forg rm 3
  forg remove 3`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid ID: %q", args[0])
		}

		store, err := storage.Load()
		if err != nil {
			return err
		}

		idx := -1
		for i, e := range store.Entries {
			if e.ID == id {
				idx = i
				break
			}
		}
		if idx == -1 {
			return fmt.Errorf("no command with ID %d found", id)
		}

		entry := store.Entries[idx]
		fmt.Printf("Remove command #%d: %s\n", entry.ID, entry.Command)
		fmt.Print("Confirm? [y/N] ")

		reader := bufio.NewReader(os.Stdin)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))

		if answer != "y" && answer != "yes" {
			fmt.Println("Cancelled.")
			return nil
		}

		store.Entries = append(store.Entries[:idx], store.Entries[idx+1:]...)
		if err := storage.Save(store); err != nil {
			return err
		}

		fmt.Printf("✓ Command #%d removed.\n", id)
		return nil
	},
}
