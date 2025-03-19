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
	RunE:  add,
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringSliceP("tag", "t", []string{}, "Tags for the note")
}

func add(cmd *cobra.Command, args []string) error {
	title := args[0]
	tags, _ := cmd.Flags().GetStringSlice("tag")

	oakRoot := cmd.Context().Value("root").(string)

	path := filepath.Join(oakRoot, title+".md")
	return helper.OpenNote(path, title, tags)

}
