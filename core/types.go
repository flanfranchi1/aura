package core

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

// AccessibleObject encapsula os metadados de um elemento da UI
type AccessibleObject struct {
	Path dbus.ObjectPath
	Name string
	Role Role
	ID   string
}

// FetchObject busca os detalhes do objeto que disparou o evento
func FetchObject(conn *dbus.Conn, sender string, path dbus.ObjectPath) (*AccessibleObject, error) {
	obj := conn.Object(sender, path)

	var name string
	// Buscamos o nome (interface org.a11y.atspi.Accessible)
	err := obj.Call("org.a11y.atspi.Accessible.GetName", 0).Store(&name)
	if err != nil {
		name = "unknown"
	}

	var roleID uint32
	// Buscamos o papel (ID numérico que será traduzido pelo seu iota)
	err = obj.Call("org.a11y.atspi.Accessible.GetRole", 0).Store(&roleID)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter role: %v", err)
	}

	return &AccessibleObject{
		Path: path,
		Name: name,
		Role: Role(roleID),
		ID:   sender,
	}, nil
}
