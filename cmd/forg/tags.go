package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/fatih/color"
	"github.com/imgoomes/forgetful/internal/storage"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
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

		table := tablewriter.NewTable(os.Stdout,
			tablewriter.WithRenderer(renderer.NewColorized(renderer.ColorizedConfig{
				Header: renderer.Tint{FG: renderer.Colors{color.FgHiWhite, color.Bold}},
				Column: renderer.Tint{
					Columns: []renderer.Tint{
						{FG: renderer.Colors{color.FgCyan}},                 // Tag
						{FG: renderer.Colors{color.FgHiYellow, color.Bold}}, // Commands
					},
				},
			})),
			tablewriter.WithConfig(tablewriter.Config{
				Row: tw.CellConfig{Alignment: tw.CellAlignment{Global: tw.AlignLeft}},
			}),
		)

		table.Header([]string{"Tag", "Commands"})
		for _, t := range tags {
			table.Append([]string{t, strconv.Itoa(counts[t])})
		}
		table.Render()

		return nil
	},
}
