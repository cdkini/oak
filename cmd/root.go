/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const oakRootEnv = "OAK_ROOT"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "oak",
	Short: "A simple command line note taking app for developers",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		oakRoot, err := getOakRoot()
		if err != nil {
			return err
		}

		ctx := context.WithValue(cmd.Context(), "root", oakRoot)
		cmd.SetContext(ctx)
		return nil
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		pprintError(err)
	}
}

func pprintError(err error) {
	color.New(color.FgRed).Print("[ERROR] ")
	fmt.Println(err.Error())
	os.Exit(1)
}

func getOakRoot() (string, error) {
	root := os.Getenv(oakRootEnv)
	if root == "" {
		return "", errors.New("Please set the OAK_ROOT environment variable!")
	}
	if stat, err := os.Stat(root); err != nil || !stat.IsDir() {
		return "", errors.New("OAK_ROOT must be a valid directory!")
	}

	return root, nil
}
