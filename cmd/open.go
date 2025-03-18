/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open <note(s)>",
	Short: "Open note(s)",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil // TODO
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
