package atspi

import (
	"testing"

	"github.com/godbus/dbus/v5"
)

func TestHasATSPIService(t *testing.T) {
	cases := []struct {
		name      string
		names     []string
		candidate string
		expected  bool
	}{
		{
			name:      "contains atspi bus",
			names:     []string{"org.freedesktop.DBus", "org.a11y.Bus"},
			candidate: "org.a11y.Bus",
			expected:  true,
		},
		{
			name:      "does not contain atspi bus",
			names:     []string{"org.freedesktop.DBus", "org.gnome.Shell"},
			candidate: "org.a11y.Bus",
			expected:  false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := hasATSPIService(tc.names, tc.candidate)
			if actual != tc.expected {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestGetRootAccessiblePath(t *testing.T) {
	expected := dbus.ObjectPath("/org/a11y/atspi/accessible/root")
	actual := GetRootAccessiblePath()
	if actual != expected {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func TestFormatSignalBody(t *testing.T) {
	body := []any{"focused", true, 42, "extra"}
	actual := FormatSignalBody(body)
	expected := "focused, true, 42, extra"
	if actual != expected {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func TestIsFocusStateChangedSignal(t *testing.T) {
	cases := []struct {
		name     string
		signal   *dbus.Signal
		expected bool
	}{
		{
			name:     "nil signal",
			signal:   nil,
			expected: false,
		},
		{
			name:     "wrong event name",
			signal:   &dbus.Signal{Name: "org.a11y.atspi.Event.Object.OtherEvent", Body: []any{"focused", true}},
			expected: false,
		},
		{
			name:     "not focused state",
			signal:   &dbus.Signal{Name: "org.a11y.atspi.Event.Object.StateChanged", Body: []any{"focused", false}},
			expected: false,
		},
		{
			name:     "focused state changed",
			signal:   &dbus.Signal{Name: "org.a11y.atspi.Event.Object.StateChanged", Body: []any{"focused", true}},
			expected: true,
		},
		{
			name:     "focused numeric truthy state",
			signal:   &dbus.Signal{Name: "org.a11y.atspi.Event.Object.StateChanged", Body: []any{"focused", int64(1)}},
			expected: true,
		},
		{
			name:     "focused numeric falsey state",
			signal:   &dbus.Signal{Name: "org.a11y.atspi.Event.Object.StateChanged", Body: []any{"focused", int64(0)}},
			expected: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsFocusStateChangedSignal(tc.signal)
			if actual != tc.expected {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
