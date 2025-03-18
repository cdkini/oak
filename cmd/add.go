/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"path/filepath"

	"github.com/cdkini/oak/v2/helper"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <title>",
	Short: "Add a new note",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		tags, _ := cmd.Flags().GetStringSlice("tag")

		oakRoot := cmd.Context().Value("root").(string)
		if err := add(oakRoot, title, tags); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringSliceP("tag", "t", []string{}, "Tags for the note")
}

func add(root string, title string, tags []string) error {
	path := filepath.Join(root, title+".md")
	return helper.OpenNote(path, title, tags)
}
