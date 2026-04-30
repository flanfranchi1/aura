package main

import (
	"aura/core"
	"aura/internal/dbus_handler"
	"aura/internal/speech"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Starting Aura...")

	sp, err := speech.NewClient()
	if err != nil {
		log.Printf("[ERROr] Speech Dispatcher offline. Outputs in terminal only: %v\n", err)
	} else {
		defer sp.Close()
		sp.SetPriority("TEXT")
		sp.Speak("Ready to go!")
	}
	conn, err := dbus_handler.NewClient()
	if err != nil {
		log.Fatalf("Failed to connect to D-Bus: %v", err)
	}
	FocusMonitor, err := conn.SubscribeToFocusedStateChanged()
	if err != nil {
		log.Fatalf("Failed to subscribe to FocusedStateChanged: %v", err)
	}
	defer conn.Close()

	for signal := range FocusMonitor {
		obj, err := core.FetchObject(conn, signal.Sender, signal.Body[0].(string))
		if err != nil {
			log.Printf("Failed to fetch object details: %v", err)
			continue
		}
		output := fmt.Sprintf("Focused Object: %s (Role: %s, States: %s)", obj.Name, obj.Role.String(), obj.PrepareStatesAsText())
		fmt.Println(output)
		sp.Speak(output)
	}
}
