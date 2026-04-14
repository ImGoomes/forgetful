package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/gabrielgomes/forgetful/internal/storage"
	"github.com/spf13/cobra"
)

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "List all tags and how many commands each has",
	Long:  `Display all existing tags and the number of commands associated with each. Useful alongside 'forg list -t <tag>'.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := storage.Load()
		if err != nil {
			return err
		}

		counts := map[string]int{}
		for _, e := range store.Entries {
			counts[e.Tag]++
		}

		if len(counts) == 0 {
			fmt.Println("No tags found. Use: forg add -c <command>")
			return nil
		}

		tags := make([]string, 0, len(counts))
		for t := range counts {
			tags = append(tags, t)
		}
		sort.Strings(tags)

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "TAG\tCOMMANDS")
		fmt.Fprintln(w, "---\t--------")
		for _, t := range tags {
			fmt.Fprintf(w, "%s\t%d\n", t, counts[t])
		}
		w.Flush()
		return nil
	},
}
