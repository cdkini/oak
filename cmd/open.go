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
	RunE:  open,
}

func init() {
	rootCmd.AddCommand(openCmd)
}

func open(cmd *cobra.Command, args []string) error {
	oakRoot := getOakRootFromCmd(cmd)

	query := ""
	if len(args) == 1 {
		query = args[0]
	}

	return helper.FZFOpen(oakRoot, query)
}
