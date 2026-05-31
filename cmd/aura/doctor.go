package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check DBus and AT-SPI environment readiness",
	RunE:  runDoctor,
}

func runDoctor(cmd *cobra.Command, _ []string) error {
	result, err := discoverAccessibilityBus(cmd.Context())
	if err != nil {
		return fmt.Errorf("doctor failed: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), "DBus session bus: connected")
	if !result.Found {
		fmt.Fprintln(cmd.OutOrStdout(), "Accessibility bus: not found")
		return nil
	}

	fmt.Fprintf(cmd.OutOrStdout(), "Accessibility bus: found (%s)\n", result.Service)

	busInfo, err := getAccessibilityBusAddress(cmd.Context())
	if err != nil {
		return fmt.Errorf("doctor failed: %w", err)
	}
	fmt.Fprintf(cmd.OutOrStdout(), "Accessibility bus address: %s\n", busInfo.Address)

	conn, err := connectAccessibilityBus(cmd.Context(), busInfo.Address)
	if err != nil {
		return fmt.Errorf("doctor failed: %w", err)
	}
	defer conn.Close()
	fmt.Fprintln(cmd.OutOrStdout(), "Connected to accessibility bus")

	if err := validateAccessibilityBusConnection(cmd.Context(), conn); err != nil {
		return fmt.Errorf("doctor failed: %w", err)
	}
	fmt.Fprintln(cmd.OutOrStdout(), "Accessibility bus validation succeeded")

	return nil
}
