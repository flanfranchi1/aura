package atspi

import (
	"context"
	"fmt"

	"github.com/godbus/dbus/v5"
)

var knownATSPIServiceNames = []string{
	"org.a11y.Bus",
	"org.a11y.atspi.Registry",
}

type DiscoveryResult struct {
	Found   bool
	Service string
}

type AccessibilityBusInfo struct {
	Service string
	Address string
}

func DiscoverAccessibilityBus(ctx context.Context) (DiscoveryResult, error) {
	conn, err := connectSessionBus(ctx)
	if err != nil {
		return DiscoveryResult{}, fmt.Errorf("connect session bus: %w", err)
	}
	defer conn.Close()

	names, err := listDBusNames(ctx, conn)
	if err != nil {
		return DiscoveryResult{}, fmt.Errorf("list DBus names: %w", err)
	}

	for _, candidate := range knownATSPIServiceNames {
		if hasATSPIService(names, candidate) {
			return DiscoveryResult{Found: true, Service: candidate}, nil
		}
	}

	return DiscoveryResult{Found: false}, nil
}

// GetAccessibilityBusAddress retrieves the address of the accessibility bus
// by calling org.a11y.Bus.GetAddress on the session bus.
func GetAccessibilityBusAddress(ctx context.Context) (AccessibilityBusInfo, error) {
	sessionConn, err := connectSessionBus(ctx)
	if err != nil {
		return AccessibilityBusInfo{}, fmt.Errorf("connect session bus: %w", err)
	}
	defer sessionConn.Close()

	obj := sessionConn.Object("org.a11y.Bus", "/org/a11y/bus")
	var address string
	if err := obj.CallWithContext(ctx, "org.a11y.Bus.GetAddress", 0).Store(&address); err != nil {
		return AccessibilityBusInfo{}, fmt.Errorf("call GetAddress: %w", err)
	}

	return AccessibilityBusInfo{Service: "org.a11y.Bus", Address: address}, nil
}

// ConnectAccessibilityBus connects to the accessibility bus using the provided address.
func ConnectAccessibilityBus(_ context.Context, address string) (*dbus.Conn, error) {
	conn, err := dbus.Connect(address)
	if err != nil {
		return nil, fmt.Errorf("connect to accessibility bus: %w", err)
	}
	return conn, nil
}

func ConnectSessionBus(_ context.Context) (*dbus.Conn, error) {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return nil, fmt.Errorf("connect to session bus: %w", err)
	}
	return conn, nil
}

// ValidateAccessibilityBusConnection verifies that we can communicate on the accessibility bus.
func ValidateAccessibilityBusConnection(ctx context.Context, conn *dbus.Conn) error {
	if conn == nil {
		return fmt.Errorf("connection is nil")
	}

	obj := conn.Object("org.freedesktop.DBus", "/org/freedesktop/DBus")
	var names []string
	if err := obj.CallWithContext(ctx, "org.freedesktop.DBus.ListNames", 0).Store(&names); err != nil {
		return fmt.Errorf("validate accessibility bus connection: %w", err)
	}
	if len(names) == 0 {
		return fmt.Errorf("validate accessibility bus connection: no names returned")
	}
	return nil
}

func connectSessionBus(_ context.Context) (*dbus.Conn, error) {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func listDBusNames(ctx context.Context, conn *dbus.Conn) ([]string, error) {
	obj := conn.Object("org.freedesktop.DBus", "/org/freedesktop/DBus")
	var names []string
	if err := obj.CallWithContext(ctx, "org.freedesktop.DBus.ListNames", 0).Store(&names); err != nil {
		return nil, err
	}
	return names, nil
}

func hasATSPIService(names []string, candidate string) bool {
	for _, name := range names {
		if name == candidate {
			return true
		}
	}
	return false
}
