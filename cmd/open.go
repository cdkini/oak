/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/cdkini/oak/v2/helper"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open <note(s)>",
	Short: "Open note(s)",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		oakRoot := cmd.Context().Value("root").(string)

		var query string
		if len(args) == 0 {
			query = ""
		} else {
			query = args[0]
		}

		return open(oakRoot, query)
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}

func open(root string, query string) error {
	return helper.FZFOpen(root, query)
}
