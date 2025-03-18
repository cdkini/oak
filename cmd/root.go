/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/cdkini/oak/v2/helper"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "oak",
	Short: "A simple command line note taking app for developers",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		oakRoot, err := helper.GetOakRoot()
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
