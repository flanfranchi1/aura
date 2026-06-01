package atspi

import (
	"strings"
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

func findRootApp(children []ObjectReference, path dbus.ObjectPath) (ObjectReference, bool) {
	var bestMatch ObjectReference
	var bestLen int
	found := false

	for _, child := range children {
		if strings.HasPrefix(string(path), string(child.Path)) {
			if len(child.Path) > bestLen {
				bestMatch = child
				bestLen = len(child.Path)
				found = true
			}
		}
	}
	return bestMatch, found
}

func TestFindRootApp(t *testing.T) {
	children := []ObjectReference{
		{BusName: ":1.5", Path: "/org/a11y/atspi/accessible/100"},
		{BusName: ":1.6", Path: "/org/a11y/atspi/accessible/100/child"},
		{BusName: ":1.7", Path: "/org/a11y/atspi/accessible/200"},
	}

	// Deeper path should match the longest prefix.
	obj, ok := findRootApp(children, dbus.ObjectPath("/org/a11y/atspi/accessible/100/child/leaf"))
	if !ok || obj.BusName != ":1.6" {
		t.Fatalf("expected to find %s, got %v ok=%v", ":1.6", obj, ok)
	}

	// No match should return false.
	_, ok = findRootApp(children, dbus.ObjectPath("/some/other/path"))
	if ok {
		t.Fatalf("expected no match for unrelated path")
	}
}
