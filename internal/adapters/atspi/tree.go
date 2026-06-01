package atspi

import (
	"context"
	"fmt"

	"github.com/godbus/dbus/v5"
)

const rootAccessiblePath = "/org/a11y/atspi/accessible/root"

type ObjectReference struct {
	BusName string
	Path    dbus.ObjectPath
}

func GetChildren(ctx context.Context, conn *dbus.Conn, objectPath dbus.ObjectPath) ([]ObjectReference, error) {
	obj := conn.Object("org.a11y.atspi.Registry", objectPath)
	var raw []struct {
		BusName string
		Path    dbus.ObjectPath
	}
	if err := obj.CallWithContext(ctx, "org.a11y.atspi.Accessible.GetChildren", 0).Store(&raw); err != nil {
		return nil, fmt.Errorf("call GetChildren: %w", err)
	}

	children := make([]ObjectReference, len(raw))
	for i, item := range raw {
		children[i] = ObjectReference{BusName: item.BusName, Path: item.Path}
	}
	return children, nil
}

func GetAccessibleName(ctx context.Context, conn *dbus.Conn, destination string, objectPath dbus.ObjectPath) (string, error) {
	obj := conn.Object(destination, objectPath)
	variant, err := obj.GetProperty("org.a11y.atspi.Accessible.Name")
	if err != nil {
		return "", fmt.Errorf("get accessible name: %w", err)
	}
	var name string
	if err := variant.Store(&name); err != nil {
		return "", fmt.Errorf("decode accessible name: %w", err)
	}
	return name, nil
}

func GetAccessibleDescription(ctx context.Context, conn *dbus.Conn, destination string, objectPath dbus.ObjectPath) (string, error) {
	obj := conn.Object(destination, objectPath)
	variant, err := obj.GetProperty("org.a11y.atspi.Accessible.Description")
	if err != nil {
		return "", fmt.Errorf("get accessible description: %w", err)
	}
	var description string
	if err := variant.Store(&description); err != nil {
		return "", fmt.Errorf("decode accessible description: %w", err)
	}
	return description, nil
}

func GetAccessibleRoleName(ctx context.Context, conn *dbus.Conn, destination string, objectPath dbus.ObjectPath) (string, error) {
	obj := conn.Object(destination, objectPath)
	var roleName string
	if err := obj.CallWithContext(ctx, "org.a11y.atspi.Accessible.GetRoleName", 0).Store(&roleName); err != nil {
		return "", fmt.Errorf("call GetRoleName: %w", err)
	}
	return roleName, nil
}

func GetRootAccessiblePath() dbus.ObjectPath {
	return dbus.ObjectPath(rootAccessiblePath)
}
