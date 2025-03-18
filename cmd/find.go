/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/cdkini/oak/v2/helper"
	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find <query?>",
	Short: "Find notes by title",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		oakRoot := cmd.Context().Value("root").(string)

		var query string
		if len(args) == 0 {
			query = ""
		} else {
			query = args[0]
		}

		return find(oakRoot, query)
	},
}

func init() {
	rootCmd.AddCommand(findCmd)
}

func find(root string, query string) error {
	return helper.Find(root, query)
}
