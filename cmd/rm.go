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
	RunE:  rm,
}

func init() {
	rootCmd.AddCommand(rmCmd)
}

func rm(cmd *cobra.Command, args []string) error {
	oakRoot := getOakRootFromCmd(cmd)

	query := ""
	if len(args) == 1 {
		query = args[0]
	}

	files, err := helper.FZFSelect(oakRoot, query)
	if err != nil {
		return err
	}

	trash, err := prepTrash(oakRoot)
	if err != nil {
		return err
	}

	return deleteFiles(oakRoot, trash, files)
}

func prepTrash(root string) (string, error) {
	trash := filepath.Join(root, TRASH)
	return trash, os.MkdirAll(trash, os.ModePerm)
}

func deleteFiles(root string, trash string, files []string) error {
	for _, file := range files {
		oldPath := filepath.Join(root, file)
		newPath := filepath.Join(trash, file)

		os.Remove(newPath) // Remove from trash if it already exists

		if err := os.Rename(oldPath, newPath); err != nil {
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
