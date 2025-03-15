/*
Copyright © 2025 Chetan Kini ckini123@gmail.com
*/
package cmd

import (
	"time"

	"github.com/cdkini/oak/v2/helper"
	"github.com/spf13/cobra"
)

var dailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "Open daily note",
	Run: func(cmd *cobra.Command, args []string) {
		oakRoot := helper.GetOakRoot()
		addDailyNote(oakRoot)
	},
}

func init() {
	rootCmd.AddCommand(dailyCmd)
}

func addDailyNote(root string) {
	title := time.Now().Format("2006-01-02")
	helper.OpenNote(root, title, []string{"daily"})
}
