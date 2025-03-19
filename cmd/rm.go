/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cdkini/oak/v2/helper"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const TRASH = ".trash"

var rmCmd = &cobra.Command{
	Use:   "rm <note(s)>",
	Short: "Delete note(s)",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		oakRoot := cmd.Context().Value("root").(string)

		var query string
		if len(args) == 1 {
			query = args[0]
		}

		return rm(oakRoot, query)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}

func rm(root string, query string) error {
	files, err := helper.FZFSelect(root, query)
	if err != nil {
		return err
	}

	trash := filepath.Join(root, TRASH)
	if err = os.MkdirAll(trash, os.ModePerm); err != nil {
		return err
	}

	for _, file := range files {
		oldPath := filepath.Join(root, file)
		newPath := filepath.Join(trash, file)

		os.Remove(newPath) // Remove from trash if it already exists

		if err = os.Rename(oldPath, newPath); err != nil {
			return err
		}

		prettyPrintDeleted(oldPath)
	}
	return nil
}

func prettyPrintDeleted(path string) {
	green := color.New(color.FgGreen)
	green.Print("[DELETED] ")
	fmt.Println(path)
}
