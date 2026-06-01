// SPDX-License-Identifier: GPL-3.0-or-later
// Copyright (c) 2026 Fernando Lanfranchi

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aura",
	Short: "Aura Linux screen reader bootstrap",
	Long:  "Aura is a lightweight screen reader runtime for GNOME/KDE desktops.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
