package main

import (
	"context"
	"fmt"

	"github.com/flanfranchi1/aura/internal/adapters/atspi"
	"github.com/godbus/dbus/v5"
)

type accessibilityConnection interface {
	Close() error
}

var (
	discoverAccessibilityBus   = atspi.DiscoverAccessibilityBus
	getAccessibilityBusAddress = atspi.GetAccessibilityBusAddress
	connectAccessibilityBus    = func(ctx context.Context, address string) (accessibilityConnection, error) {
		return atspi.ConnectAccessibilityBus(ctx, address)
	}
	validateAccessibilityBusConnection = func(ctx context.Context, conn accessibilityConnection) error {
		dbusConn, ok := conn.(*dbus.Conn)
		if !ok {
			return fmt.Errorf("unsupported connection type")
		}
		return atspi.ValidateAccessibilityBusConnection(ctx, dbusConn)
	}
	getChildren = func(ctx context.Context, conn accessibilityConnection, objectPath dbus.ObjectPath) ([]atspi.ObjectReference, error) {
		return atspi.GetChildren(ctx, conn.(*dbus.Conn), objectPath)
	}
	getAccessibleName = func(ctx context.Context, conn accessibilityConnection, destination string, objectPath dbus.ObjectPath) (string, error) {
		return atspi.GetAccessibleName(ctx, conn.(*dbus.Conn), destination, objectPath)
	}
	getAccessibleRoleName = func(ctx context.Context, conn accessibilityConnection, destination string, objectPath dbus.ObjectPath) (string, error) {
		return atspi.GetAccessibleRoleName(ctx, conn.(*dbus.Conn), destination, objectPath)
	}
)
