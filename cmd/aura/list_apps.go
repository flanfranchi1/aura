package main

import (
	"fmt"

	"github.com/flanfranchi1/aura/internal/adapters/atspi"
	"github.com/spf13/cobra"
)

var listAppsCmd = &cobra.Command{
	Use:   "list-apps",
	Short: "List top-level accessible desktop applications",
	RunE:  runListApps,
}

func init() {
	rootCmd.AddCommand(listAppsCmd)
}

func runListApps(cmd *cobra.Command, _ []string) error {
	busInfo, err := getAccessibilityBusAddress(cmd.Context())
	if err != nil {
		return fmt.Errorf("list-apps failed: %w", err)
	}

	fmt.Fprintf(cmd.OutOrStdout(), "Accessibility bus address: %s\n", busInfo.Address)

	conn, err := connectAccessibilityBus(cmd.Context(), busInfo.Address)
	if err != nil {
		return fmt.Errorf("list-apps failed: %w", err)
	}
	defer conn.Close()

	rootPath := atspi.GetRootAccessiblePath()
	children, err := getChildren(cmd.Context(), conn, rootPath)
	if err != nil {
		return fmt.Errorf("list-apps failed: %w", err)
	}

	fmt.Fprintf(cmd.OutOrStdout(), "Top-level accessibles: %d\n", len(children))
	for index, child := range children {
		name, err := getAccessibleName(cmd.Context(), conn, child.BusName, child.Path)
		if err != nil {
			name = fmt.Sprintf("<error: %v>", err)
		}
		role, err := getAccessibleRoleName(cmd.Context(), conn, child.BusName, child.Path)
		if err != nil {
			role = fmt.Sprintf("<error: %v>", err)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "%2d: %s (%s) [%s]\n", index+1, name, role, child.BusName)
	}

	return nil
}
