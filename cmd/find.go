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
	RunE:  find,
}

func init() {
	rootCmd.AddCommand(findCmd)
}

func find(cmd *cobra.Command, args []string) error {
	oakRoot := cmd.Context().Value("root").(string)

	query := ""
	if len(args) == 1 {
		query = args[0]
	}

	return helper.Find(oakRoot, query)
}
