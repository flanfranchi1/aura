package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/flanfranchi1/aura/internal/adapters/atspi"
	"github.com/flanfranchi1/aura/internal/speech"
	"github.com/godbus/dbus/v5"
	"github.com/spf13/cobra"
)

var watchFocusCmd = &cobra.Command{
	Use:   "watch-focus",
	Short: "Watch AT-SPI focus events in real time",
	RunE:  runWatchFocus,
}

// containsAppName checks if the given text already contains the app name
func containsAppName(text, app string) bool {
	return len(text) > len(app) && text[:len(app)] == app
}

func init() {
	rootCmd.AddCommand(watchFocusCmd)
}

func runWatchFocus(cmd *cobra.Command, _ []string) error {
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	busInfo, err := atspi.GetAccessibilityBusAddress(ctx)
	if err != nil {
		return fmt.Errorf("watch-focus failed: %w", err)
	}

	accessConn, err := atspi.ConnectAccessibilityBus(ctx, busInfo.Address)
	if err != nil {
		return fmt.Errorf("watch-focus failed: %w", err)
	}
	defer accessConn.Close()

	if err := accessConn.AddMatchSignal(
		dbus.WithMatchInterface("org.a11y.atspi.Event.Object"),
		dbus.WithMatchMember("StateChanged"),
	); err != nil {
		return fmt.Errorf("watch-focus failed: %w", err)
	}

	signalChannel := make(chan *dbus.Signal, 16)
	accessConn.Signal(signalChannel)
	defer accessConn.RemoveSignal(signalChannel)

	appBusName, err := atspi.GetConnectionUniqueName(accessConn)
	if err != nil {
		return fmt.Errorf("watch-focus failed: %w", err)
	}

	if err := atspi.RegisterEventListener(ctx, accessConn, atspi.FocusEventName, nil, appBusName); err != nil {
		return fmt.Errorf("watch-focus failed: %w", err)
	}
	defer func() {
		_ = atspi.DeregisterEventListener(context.Background(), accessConn, atspi.FocusEventName)
	}()

	fmt.Fprintf(cmd.OutOrStdout(), "Watching %s events on accessibility bus as %s\n", atspi.FocusEventName, appBusName)
	fmt.Fprintln(cmd.OutOrStdout(), "Press Ctrl+C to stop.")

	var lastSpokenText string
	
	for {
		select {
		case <-ctx.Done():
			return nil
		case sig := <-signalChannel:
			if !atspi.IsFocusStateChangedSignal(sig) {
				continue
			}

			// Enrich with accessible metadata to get human-readable app name
			app, role, name, _ := atspi.EnrichAccessible(ctx, accessConn, sig.Sender, sig.Path)
			// Generate speech text if we have useful info
			if app != "" || name != "" || role != "" {
				var speechText string
				
				// Optional improvement: Only repeat app name if it changed
				if app != "" && (lastSpokenText == "" || !containsAppName(lastSpokenText, app)) {
					speechText += app + ". "
				}
				
				if name != "" {
					speechText += name + ". "
				}
				if role != "" {
					speechText += role + "."
				}
				
				// Only speak if the text is different from last time
				if speechText != "" && speechText != lastSpokenText {
					speech.Speak(ctx, speechText)
					lastSpokenText = speechText
				}
			}
			// Print enriched information
			fmt.Fprintln(cmd.OutOrStdout(), "[FOCUS]")
			fmt.Fprintf(cmd.OutOrStdout(), "app=%s\n", app)
			fmt.Fprintf(cmd.OutOrStdout(), "role=%s\n", role)
			fmt.Fprintf(cmd.OutOrStdout(), "name=%s\n", name)
			fmt.Fprintf(cmd.OutOrStdout(), "sender=%s\n", sig.Sender)
			fmt.Fprintf(cmd.OutOrStdout(), "path=%s\n", sig.Path)
			// Print state information from the signal body
			// IsFocusStateChangedSignal already ensures len(sig.Body) >= 2
			fmt.Fprintf(cmd.OutOrStdout(), "state=%v\n", sig.Body[0])
			fmt.Fprintf(cmd.OutOrStdout(), "enabled=%v\n", sig.Body[1])
		}
	}
}
