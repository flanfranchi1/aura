package main

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "github.com/flanfranchi1/aura/internal/adapters/atspi"
    "github.com/godbus/dbus/v5"
    "github.com/spf13/cobra"
)

var watchFocusCmd = &cobra.Command{
    Use:   "watch-focus",
    Short: "Watch AT-SPI focus events in real time",
    RunE:  runWatchFocus,
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

    for {
        select {
        case <-ctx.Done():
            return nil
        case sig := <-signalChannel:
            if !atspi.IsFocusStateChangedSignal(sig) {
                continue
            }

            fmt.Fprintf(cmd.OutOrStdout(), "focus event sender=%s path=%s body=%s\n",
                sig.Sender, sig.Path, atspi.FormatSignalBody(sig.Body))
        }
    }
}
