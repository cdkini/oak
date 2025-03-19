/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"path/filepath"
	"time"

	"github.com/cdkini/oak/v2/helper"
	"github.com/spf13/cobra"
)

var dailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "Add a new daily note",
	Args:  cobra.NoArgs,
	RunE:  daily,
}

func init() {
	rootCmd.AddCommand(dailyCmd)
}

func daily(cmd *cobra.Command, args []string) error {
	oakRoot := cmd.Context().Value("root").(string)

	title := time.Now().Format("2006-01-02")
	path := filepath.Join(oakRoot, title+".md")

	return helper.OpenNote(path, title, []string{"daily"})
}
