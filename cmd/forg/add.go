package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gabrielgomes/forgetful/internal/model"
	"github.com/gabrielgomes/forgetful/internal/storage"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Save a command",
	Long: `Save a terminal command with an optional tag and description.

If no tag is provided, the first word of the command is used as the tag.
Commands are stored in ~/.forg/commands.json.`,
	Example: `  forg add -c "git config user.name 'Gabriel'" -t git -d "Change local git user"
  forg add -c "docker ps -a" -d "List all containers"
  forg add -c "ffmpeg -i input.mp4 -vn output.mp3"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		command, _ := cmd.Flags().GetString("command")
		tag, _ := cmd.Flags().GetString("tag")
		desc, _ := cmd.Flags().GetString("description")

		if len(args) > 0 {
			extra := strings.Join(args, " ")
			if command == "" {
				command = extra
			} else {
				command = command + " " + extra
			}
		}

		if command == "" {
			return fmt.Errorf("flag --command (-c) is required")
		}

		if tag == "" {
			tag = strings.Fields(command)[0]
		}

		store, err := storage.Load()
		if err != nil {
			return err
		}

		entry := &model.Entry{
			ID:          store.NextID,
			Command:     command,
			Tag:         tag,
			Description: desc,
			CreatedAt:   time.Now(),
		}
		store.Entries = append(store.Entries, entry)
		store.NextID++

		if err := storage.Save(store); err != nil {
			return err
		}

		fmt.Printf("✓ Command saved with ID %d (tag: %s)\n", entry.ID, entry.Tag)
		return nil
	},
}

func init() {
	addCmd.Flags().StringP("command", "c", "", "Command to save (required)")
	addCmd.Flags().StringP("tag", "t", "", "Tag (optional; defaults to the first word of the command)")
	addCmd.Flags().StringP("description", "d", "", "Description of what the command does (optional)")
}
