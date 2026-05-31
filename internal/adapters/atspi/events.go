package atspi

import (
	"context"
	"fmt"
	"strings"

	"github.com/godbus/dbus/v5"
)

const (
	FocusEventName = "Object:StateChanged:Focused"
)

func GetConnectionUniqueName(conn *dbus.Conn) (string, error) {
	if conn == nil {
		return "", fmt.Errorf("connection is nil")
	}

	names := conn.Names()
	if len(names) == 0 {
		return "", fmt.Errorf("connection has no bus name")
	}

	for _, name := range names {
		if strings.HasPrefix(name, ":") {
			return name, nil
		}
	}
	return names[0], nil
}

func RegisterEventListener(ctx context.Context, conn *dbus.Conn, event string, properties []string, appBusName string) error {
	if conn == nil {
		return fmt.Errorf("connection is nil")
	}
	if appBusName == "" {
		return fmt.Errorf("app bus name is empty")
	}

	registry := conn.Object("org.a11y.atspi.Registry", "/org/a11y/atspi/registry")
	call := registry.CallWithContext(ctx, "org.a11y.atspi.Registry.RegisterEvent", 0, event, properties, appBusName)
	if call.Err != nil {
		return fmt.Errorf("register event %q: %w", event, call.Err)
	}
	return nil
}

func DeregisterEventListener(ctx context.Context, conn *dbus.Conn, event string) error {
	if conn == nil {
		return fmt.Errorf("connection is nil")
	}

	registry := conn.Object("org.a11y.atspi.Registry", "/org/a11y/atspi/registry")
	call := registry.CallWithContext(ctx, "org.a11y.atspi.Registry.DeregisterEvent", 0, event)
	if call.Err != nil {
		return fmt.Errorf("deregister event %q: %w", event, call.Err)
	}
	return nil
}

func FormatSignalBody(body []any) string {
	parts := make([]string, len(body))
	for i, item := range body {
		parts[i] = fmt.Sprintf("%v", item)
	}
	return strings.Join(parts, ", ")
}

func IsFocusStateChangedSignal(sig *dbus.Signal) bool {
	if sig == nil {
		return false
	}
	if sig.Name != "org.a11y.atspi.Event.Object.StateChanged" {
		return false
	}
	if len(sig.Body) < 2 {
		return false
	}

	state, ok := sig.Body[0].(string)
	if !ok || state != "focused" {
		return false
	}

	return isTruthyStateValue(sig.Body[1])
}

func isTruthyStateValue(value any) bool {
	switch v := value.(type) {
	case bool:
		return v
	case int:
		return v != 0
	case int8:
		return v != 0
	case int16:
		return v != 0
	case int32:
		return v != 0
	case int64:
		return v != 0
	case uint:
		return v != 0
	case uint8:
		return v != 0
	case uint16:
		return v != 0
	case uint32:
		return v != 0
	case uint64:
		return v != 0
	case float32:
		return v != 0
	case float64:
		return v != 0
	case dbus.Variant:
		return isTruthyStateValue(v.Value())
	default:
		return false
	}
}
