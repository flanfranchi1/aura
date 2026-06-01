// SPDX-License-Identifier: GPL-3.0-or-later
// Copyright (c) 2026 Fernando Lanfranchi

package main

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/flanfranchi1/aura/internal/adapters/atspi"
	"github.com/godbus/dbus/v5"
)

type stubConn struct{}

func (stubConn) Close() error { return nil }

func restoreDoctorHooks(
	discover func(ctx context.Context) (atspi.DiscoveryResult, error),
	address func(ctx context.Context) (atspi.AccessibilityBusInfo, error),
	connect func(ctx context.Context, address string) (accessibilityConnection, error),
	validate func(ctx context.Context, conn accessibilityConnection) error,
) func() {
	oldDiscover := discoverAccessibilityBus
	oldAddress := getAccessibilityBusAddress
	oldConnect := connectAccessibilityBus
	oldValidate := validateAccessibilityBusConnection

	discoverAccessibilityBus = discover
	getAccessibilityBusAddress = address
	connectAccessibilityBus = connect
	validateAccessibilityBusConnection = validate

	return func() {
		discoverAccessibilityBus = oldDiscover
		getAccessibilityBusAddress = oldAddress
		connectAccessibilityBus = oldConnect
		validateAccessibilityBusConnection = oldValidate
	}
}

func restoreListAppsHooks(
	address func(ctx context.Context) (atspi.AccessibilityBusInfo, error),
	connect func(ctx context.Context, address string) (accessibilityConnection, error),
	children func(ctx context.Context, conn accessibilityConnection, objectPath dbus.ObjectPath) ([]atspi.ObjectReference, error),
	name func(ctx context.Context, conn accessibilityConnection, destination string, objectPath dbus.ObjectPath) (string, error),
	role func(ctx context.Context, conn accessibilityConnection, destination string, objectPath dbus.ObjectPath) (string, error),
) func() {
	oldAddress := getAccessibilityBusAddress
	oldConnect := connectAccessibilityBus
	oldChildren := getChildren
	oldName := getAccessibleName
	oldRole := getAccessibleRoleName

	getAccessibilityBusAddress = address
	connectAccessibilityBus = connect
	getChildren = children
	getAccessibleName = name
	getAccessibleRoleName = role

	return func() {
		getAccessibilityBusAddress = oldAddress
		connectAccessibilityBus = oldConnect
		getChildren = oldChildren
		getAccessibleName = oldName
		getAccessibleRoleName = oldRole
	}
}

func runRootCommand(t *testing.T, args ...string) (string, error) {
	t.Helper()

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetArgs(args)
	rootCmd.SetContext(context.Background())

	err := rootCmd.Execute()
	return buf.String(), err
}

func TestDoctorCommand_AccessibilityBusUnavailable(t *testing.T) {
	restore := restoreDoctorHooks(
		func(ctx context.Context) (atspi.DiscoveryResult, error) {
			return atspi.DiscoveryResult{Found: false}, nil
		},
		getAccessibilityBusAddress,
		connectAccessibilityBus,
		validateAccessibilityBusConnection,
	)
	defer restore()

	output, err := runRootCommand(t, "doctor")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !strings.Contains(output, "Accessibility bus: not found") {
		t.Fatalf("expected missing accessibility bus message, got %q", output)
	}
}

func TestDoctorCommand_AddressRetrievalFails(t *testing.T) {
	restore := restoreDoctorHooks(
		func(ctx context.Context) (atspi.DiscoveryResult, error) {
			return atspi.DiscoveryResult{Found: true, Service: "org.a11y.Bus"}, nil
		},
		func(ctx context.Context) (atspi.AccessibilityBusInfo, error) {
			return atspi.AccessibilityBusInfo{}, fmt.Errorf("address failure")
		},
		connectAccessibilityBus,
		validateAccessibilityBusConnection,
	)
	defer restore()

	_, err := runRootCommand(t, "doctor")
	if err == nil || !strings.Contains(err.Error(), "address failure") {
		t.Fatalf("expected address failure error, got %v", err)
	}
}

func TestDoctorCommand_ValidationFails(t *testing.T) {
	restore := restoreDoctorHooks(
		func(ctx context.Context) (atspi.DiscoveryResult, error) {
			return atspi.DiscoveryResult{Found: true, Service: "org.a11y.Bus"}, nil
		},
		func(ctx context.Context) (atspi.AccessibilityBusInfo, error) {
			return atspi.AccessibilityBusInfo{Service: "org.a11y.Bus", Address: "memory:"}, nil
		},
		func(ctx context.Context, address string) (accessibilityConnection, error) {
			return stubConn{}, nil
		},
		func(ctx context.Context, conn accessibilityConnection) error {
			return fmt.Errorf("validation failed")
		},
	)
	defer restore()

	_, err := runRootCommand(t, "doctor")
	if err == nil || !strings.Contains(err.Error(), "validation failed") {
		t.Fatalf("expected validation failed error, got %v", err)
	}
}

func TestDoctorCommand_Succeeds(t *testing.T) {
	restore := restoreDoctorHooks(
		func(ctx context.Context) (atspi.DiscoveryResult, error) {
			return atspi.DiscoveryResult{Found: true, Service: "org.a11y.Bus"}, nil
		},
		func(ctx context.Context) (atspi.AccessibilityBusInfo, error) {
			return atspi.AccessibilityBusInfo{Service: "org.a11y.Bus", Address: "memory:"}, nil
		},
		func(ctx context.Context, address string) (accessibilityConnection, error) {
			return stubConn{}, nil
		},
		func(ctx context.Context, conn accessibilityConnection) error {
			return nil
		},
	)
	defer restore()

	output, err := runRootCommand(t, "doctor")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !strings.Contains(output, "Accessibility bus validation succeeded") {
		t.Fatalf("expected validation success, got %q", output)
	}
}

func TestListAppsCommand_EmptyResults(t *testing.T) {
	restore := restoreListAppsHooks(
		func(ctx context.Context) (atspi.AccessibilityBusInfo, error) {
			return atspi.AccessibilityBusInfo{Address: "memory:"}, nil
		},
		func(ctx context.Context, address string) (accessibilityConnection, error) {
			return stubConn{}, nil
		},
		func(ctx context.Context, conn accessibilityConnection, objectPath dbus.ObjectPath) ([]atspi.ObjectReference, error) {
			return nil, nil
		},
		getAccessibleName,
		getAccessibleRoleName,
	)
	defer restore()

	output, err := runRootCommand(t, "list-apps")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !strings.Contains(output, "Top-level accessibles: 0") {
		t.Fatalf("expected zero accessible apps, got %q", output)
	}
}

func TestListAppsCommand_SuccessfulEnumeration(t *testing.T) {
	restore := restoreListAppsHooks(
		func(ctx context.Context) (atspi.AccessibilityBusInfo, error) {
			return atspi.AccessibilityBusInfo{Address: "memory:"}, nil
		},
		func(ctx context.Context, address string) (accessibilityConnection, error) {
			return stubConn{}, nil
		},
		func(ctx context.Context, conn accessibilityConnection, objectPath dbus.ObjectPath) ([]atspi.ObjectReference, error) {
			return []atspi.ObjectReference{{BusName: "org.example.App", Path: "/org/example/app"}}, nil
		},
		func(ctx context.Context, conn accessibilityConnection, destination string, objectPath dbus.ObjectPath) (string, error) {
			return "Example App", nil
		},
		func(ctx context.Context, conn accessibilityConnection, destination string, objectPath dbus.ObjectPath) (string, error) {
			return "Frame", nil
		},
	)
	defer restore()

	output, err := runRootCommand(t, "list-apps")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !strings.Contains(output, "Top-level accessibles: 1") {
		t.Fatalf("expected one accessible app, got %q", output)
	}
	if !strings.Contains(output, "Example App (Frame) [org.example.App]") {
		t.Fatalf("unexpected app line, got %q", output)
	}
}
