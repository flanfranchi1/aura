// SPDX-License-Identifier: GPL-3.0-or-later
// Copyright (c) 2026 Fernando Lanfranchi

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/flanfranchi1/aura/internal/adapters/atspi"
	"github.com/godbus/dbus/v5"
	"github.com/spf13/cobra"
)

var watchEventsCmd = &cobra.Command{
	Use:   "watch-events",
	Short: "Dump raw AT-SPI event signals for runtime diagnosis",
	RunE:  runWatchEvents,
}

func init() {
	rootCmd.AddCommand(watchEventsCmd)
}

func runWatchEvents(cmd *cobra.Command, _ []string) error {
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	busInfo, err := atspi.GetAccessibilityBusAddress(ctx)
	if err != nil {
		return fmt.Errorf("watch-events failed: %w", err)
	}

	accessConn, err := atspi.ConnectAccessibilityBus(ctx, busInfo.Address)
	if err != nil {
		return fmt.Errorf("watch-events failed: %w", err)
	}
	defer accessConn.Close()

	if err := accessConn.AddMatchSignal(
		dbus.WithMatchInterface("org.a11y.atspi.Event.Object"),
	); err != nil {
		return fmt.Errorf("watch-events failed: %w", err)
	}

	signalChannel := make(chan *dbus.Signal, 32)
	accessConn.Signal(signalChannel)
	defer accessConn.RemoveSignal(signalChannel)

	fmt.Fprintf(cmd.OutOrStdout(), "Listening to AT-SPI event signals on accessibility bus\n")
	fmt.Fprintf(cmd.OutOrStdout(), "Press Ctrl+C to stop.\n\n")

	for {
		select {
		case <-ctx.Done():
			return nil
		case sig := <-signalChannel:
			if sig == nil {
				continue
			}

			ts := time.Now().Format("15:04:05.000")
			fmt.Fprintf(cmd.OutOrStdout(), "[%s] %s\n", ts, formatSignalDiagnostics(sig))
		}
	}
}

func formatSignalDiagnostics(sig *dbus.Signal) string {
	return fmt.Sprintf(
		"name=%s sender=%s path=%s body=%s",
		sig.Name,
		sig.Sender,
		sig.Path,
		formatSignalBody(sig.Body),
	)
}

func formatSignalBody(body []any) string {
	if len(body) == 0 {
		return "[]"
	}

	formatted := "["
	for i, item := range body {
		if i > 0 {
			formatted += ", "
		}
		formatted += safeFormatValue(item)
	}
	formatted += "]"
	return formatted
}

func safeFormatValue(item any) string {
	switch v := item.(type) {
	case string:
		return fmt.Sprintf("%q", v)
	case bool:
		return fmt.Sprintf("%v", v)
	case int, int32, int64:
		return fmt.Sprintf("%v", v)
	case uint, uint32, uint64:
		return fmt.Sprintf("%v", v)
	case float32, float64:
		return fmt.Sprintf("%v", v)
	case []string:
		return fmt.Sprintf("%v", v)
	case []interface{}:
		formatted := "["
		for j, elem := range v {
			if j > 0 {
				formatted += ", "
			}
			formatted += safeFormatValue(elem)
		}
		formatted += "]"
		return formatted
	default:
		return fmt.Sprintf("%T(%v)", v, v)
	}
}
