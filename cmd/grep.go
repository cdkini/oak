/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/cdkini/oak/v2/helper"
	"github.com/spf13/cobra"
)

var grepCmd = &cobra.Command{
	Use:   "grep <query?>",
	Short: "Search notes by content",
	Args:  cobra.MaximumNArgs(1), // TODO: Make this a passthrough for grep/rg flags
	RunE: func(cmd *cobra.Command, args []string) error {
		oakRoot := cmd.Context().Value("root").(string)

		var query string
		if len(args) == 1 {
			query = args[0]
		}

		return grep(oakRoot, query)
	},
}

func init() {
	rootCmd.AddCommand(grepCmd)
}

func grep(root string, query string) error {
	return helper.Grep(root, query)
}
