package main

import (
	"aura/core"
	"aura/internal/speech"
	"fmt"
	"log"

	"github.com/godbus/dbus/v5"
)

func main() {
	fmt.Println("--- Iniciando Diagnóstico do Aura ---")

	sp, err := speech.NewClient()
	if err != nil {
		log.Printf("[ERRO] Speech Dispatcher offline. Aura rodará apenas no terminal: %v\n", err)
	} else {
		defer sp.Close()
		sp.SetPriority("TEXT")
		sp.Speak("Aura em vigília")
	}

	// 1. Conexão com o Session Bus
	sessionConn, err := dbus.ConnectSessionBus()
	if err != nil {
		log.Fatalf("[ERRO] Falha ao conectar no Session Bus: %v", err)
	}

	atspiBusLocationPath := dbus.ObjectPath(core.AtspiLocation)
	var address string
	obj := sessionConn.Object(core.AtspiBusName, atspiBusLocationPath)
	err = obj.Call("org.a11y.Bus.GetAddress", 0).Store(&address)
	if err != nil {
		log.Fatalf("[ERRO] Falha ao obter endereço A11y: %v", err)
	}
	sessionConn.Close()

	// 2. Conexão com o Barramento de Acessibilidade
	conn, err := dbus.Connect(address)
	if err != nil {
		log.Fatalf("[ERRO] Falha ao conectar no barramento A11y: %v", err)
	}
	defer conn.Close()

	// 3. Registro no Registry
	registryPath := dbus.ObjectPath(core.AtspiRegistryPath)
	registry := conn.Object("org.a11y.atspi.Registry", registryPath)

	regCall := registry.Call("org.a11y.atspi.Registry.RegisterEvent", 0, "object:state-changed", []string{}, "")
	if regCall.Err != nil {
		fmt.Printf("[ALERTA] Erro no registro do Registry: %v\n", regCall.Err)
	}

	// 4. Configuração do Match Rule (Erro de Shadowing corrigido)
	rule := "type='signal',interface='org.a11y.atspi.Event.Object',member='StateChanged',arg0='focused'"
	matchCall := conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, rule)
	if matchCall.Err != nil {
		fmt.Printf("[ALERTA] Falha ao adicionar Match Rule: %v\n", matchCall.Err)
	} else {
		fmt.Println("[OK] Match Rule aplicada.")
	}

	// 5. Setup do Canal de Escuta
	c := make(chan *dbus.Signal, 100)
	conn.Signal(c)
	fmt.Println("--- Aura em Vigília: Aguardando Eventos ---")

	for v := range c {
		focusObj, err := core.FetchObject(conn, v.Sender, v.Path)
		if err != nil {
			continue // Ignora silenciosamente para não poluir o log
		}

		// Lógica corrigida: Se tem nome, fala o nome. Se não, fala a função.
		var msg string
		if focusObj.Name != "unknown" && focusObj.Name != "" {
			msg = fmt.Sprintf("%s, %s", focusObj.Name, focusObj.Role.String())
		} else {
			msg = focusObj.Role.String()
		}

		if sp != nil {
			sp.Cancel()
			sp.Speak(msg)
		}

		fmt.Printf("\n[AURA] Foco em: %s\n", msg)
	}
}
