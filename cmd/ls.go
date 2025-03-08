/*
Copyright © 2025 Chetan Kini ckini123@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/cdkini/oak/v2/helper"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List notes",
	Run: func(cmd *cobra.Command, args []string) {
		oakRoot := helper.GetOakRoot()
		if err := listMarkdownFiles(oakRoot); err != nil {
			helper.Error("Failed to list notes: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

type fileInfo struct {
	name      string
	createdAt time.Time
	updatedAt time.Time
}

func listMarkdownFiles(dir string) error {
	files, err := collectMarkdownFiles(dir)
	if err != nil {
		helper.Error("Failed to collect markdown files: %v", err)
	}

	// Sort files by modification time (most recent first)
	sort.Slice(files, func(i, j int) bool {
		return files[i].updatedAt.After(files[j].updatedAt)
	})

	printMarkdownFiles(files)

	return nil
}

func collectMarkdownFiles(dir string) ([]fileInfo, error) {
	var files []fileInfo

	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return files, err
	}

	for _, entry := range dirEntries {
		if entry.IsDir() {
			continue // Skip nested directories
		}
		if filepath.Ext(entry.Name()) == ".md" {
			name := strings.TrimSuffix(entry.Name(), ".md")

			stat, err := os.Stat(filepath.Join(dir, entry.Name()))
			if err != nil {
				return files, err
			}
			createdAt := stat.ModTime()

			info, err := entry.Info()
			if err != nil {
				return files, err
			}
			updatedAt := info.ModTime()

			files = append(files, fileInfo{name: name, createdAt: createdAt, updatedAt: updatedAt})
		}
	}

	return files, nil
}

func printMarkdownFiles(files []fileInfo) {
	maxNameLen := 0
	for _, file := range files {
		if len(file.name) > maxNameLen {
			maxNameLen = len(file.name)
		}
	}

	fmt.Printf("%-*s | %s | %s\n", maxNameLen, "Filename", "Created At", "Last Modified")
	fmt.Println(strings.Repeat("-", maxNameLen+35))
	for _, file := range files {
		fmt.Printf("%-*s | %s | %s\n", maxNameLen, file.name, file.updatedAt.Format("01/02/2006 15:04"), file.createdAt.Format("01/02/2006 15:04"))
	}
}
