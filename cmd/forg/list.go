package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/gabrielgomes/forgetful/internal/model"
	"github.com/gabrielgomes/forgetful/internal/storage"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List saved commands",
	Long:    `List all saved commands in a table. Use -t to filter by tag.`,
	Example: `  forg list
  forg list -t git
  forg ls -t docker`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tag, _ := cmd.Flags().GetString("tag")

		store, err := storage.Load()
		if err != nil {
			return err
		}

		entries := filter(store.Entries, tag)
		if len(entries) == 0 {
			if tag != "" {
				fmt.Printf("No commands found with tag %q.\n", tag)
			} else {
				fmt.Println("No commands saved yet. Use: forg add -c <command>")
			}
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tTAG\tDESCRIPTION\tCOMMAND")
		fmt.Fprintln(w, "--\t---\t-------\t-----------")
		for _, e := range entries {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", e.ID, e.Tag, e.Description, e.Command)
		}
		w.Flush()
		return nil
	},
}

func filter(entries []*model.Entry, tag string) []*model.Entry {
	if tag == "" {
		return entries
	}
	var result []*model.Entry
	for _, e := range entries {
		if e.Tag == tag {
			result = append(result, e)
		}
	}
	return result
}

func init() {
	listCmd.Flags().StringP("tag", "t", "", "Filter by tag")
}
