/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/cdkini/oak/v2/helper"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all notes (in reverse chronological update order)",
	RunE: func(cmd *cobra.Command, args []string) error {
		oakRoot := cmd.Context().Value("root").(string)

		return list(oakRoot)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

func list(root string) error {
	metadata, err := collectFileMetadata(root)
	if err != nil {
		return err
	}

	padding := [...]int{35, 20, 20, 5} // Should probably be calculated dynamically
	format := "%-*s %-*s %-*s %-*s\n"

	color.Blue(format, padding[0], "Title", padding[1], "Updated At", padding[2], "Created At", padding[3], "Tags")
	for _, m := range metadata {
		fmt.Printf(format, padding[0], m.Title, padding[1], *m.UpdatedAt, padding[2], m.CreatedAt, padding[3], m.Tags)
	}

	return nil
}

func collectFileMetadata(dir string) ([]*helper.FileMetadata, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var metadata []*helper.FileMetadata
	for _, entry := range entries {
		if filepath.Ext(entry.Name()) != ".md" {
			continue
		}

		path := filepath.Join(dir, entry.Name())
		m, err := helper.ParseFileMetadata(path)
		if err != nil {
			continue
		}
		metadata = append(metadata, m)
	}

	sort.Slice(metadata, func(i, j int) bool {
		return *metadata[i].UpdatedAt > *metadata[j].UpdatedAt
	})

	return metadata, nil
}
