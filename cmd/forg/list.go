package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/imgoomes/forgetful/internal/model"
	"github.com/imgoomes/forgetful/internal/storage"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
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

		table := tablewriter.NewTable(os.Stdout,
			tablewriter.WithRenderer(renderer.NewColorized(renderer.ColorizedConfig{
				Header: renderer.Tint{FG: renderer.Colors{color.FgHiWhite, color.Bold}},
				Column: renderer.Tint{
					Columns: []renderer.Tint{
						{FG: renderer.Colors{color.FgHiYellow, color.Bold}}, // ID
						{FG: renderer.Colors{color.FgCyan}},                 // Tag
						{},                                                  // Description
						{FG: renderer.Colors{color.FgHiGreen}},              // Command
					},
				},
			})),
			tablewriter.WithConfig(tablewriter.Config{
				Row: tw.CellConfig{Alignment: tw.CellAlignment{Global: tw.AlignLeft}},
			}),
		)

		table.Header([]string{"ID", "Tag", "Description", "Command"})
		for _, e := range entries {
			table.Append([]string{strconv.Itoa(e.ID), e.Tag, e.Description, e.Command})
		}
		table.Render()

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
