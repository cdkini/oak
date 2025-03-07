/*
Copyright © 2025 Chetan Kini ckini123@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "Open daily note",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("daily called")
	},
}

func init() {
	rootCmd.AddCommand(dailyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dailyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dailyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
