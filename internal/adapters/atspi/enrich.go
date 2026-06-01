// SPDX-License-Identifier: GPL-3.0-or-later
// Copyright (c) 2026 Fernando Lanfranchi

package atspi

import (
	"context"

	"github.com/godbus/dbus/v5"
)

// EnrichAccessible performs lightweight metadata lookups for an accessible
// object. It returns application name (when inexpensive to derive), role
// name and accessible name. Errors are non-fatal: callers should treat empty
// strings as absent metadata.
func EnrichAccessible(ctx context.Context, conn *dbus.Conn, destination string, objectPath dbus.ObjectPath) (app string, role string, name string, err error) {
	// Try to fetch name and role from the object directly. These are cheap
	// synchronous calls against the accessibility bus and are already used
	// elsewhere in the codebase.
	if n, e := GetAccessibleName(ctx, conn, destination, objectPath); e == nil {
		name = n
	}
	// If name is empty, try accessible description as a fallback.
	if name == "" {
		if d, e := GetAccessibleDescription(ctx, conn, destination, objectPath); e == nil && d != "" {
			name = d
		}
	}
	if r, e := GetAccessibleRoleName(ctx, conn, destination, objectPath); e == nil {
		role = r
	}

	// Get application name directly from the bus name to avoid expensive lookups
	// This is a cheap and deterministic approach
	app, err = GetAccessibleName(ctx, conn, destination, GetRootAccessiblePath())
	if err != nil {
		// Fallback to using the bus name if we can't get the application name
		app = destination
	}

	return app, role, name, nil
}
