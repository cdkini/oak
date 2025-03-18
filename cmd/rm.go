/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm <note(s)>",
	Short: "Delete note(s)",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil // TODO
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
