package core

import (
	"aura/internal/dbus_handler"
	"strings"
)

// AccessibleObject encapsula os metadados de um elemento da UI
type AccessibleObject struct {
	ParentObjectName string
	Name             string
	Role             Role
	State            StateSet
	ID               string
}

// FetchObject busca os detalhes do objeto que disparou o evento
func FetchObject(c *dbus_handler.Client, sender, path string) (*AccessibleObject, error) {
	name, err := c.GetStringProperty(sender, path, "Name")
	role, err := c.GetUint32Property(sender, path, "Role")
	state, err := c.GetStateSetProperty(sender, path, "State")
	return &AccessibleObject{
		Name:  name,
		Role:  Role(role),
		State: StateSet(state),
	}, err
}

func (a *AccessibleObject) CurrentStates() []string {
	states := make([]string, 0, 20)

	for i := uint32(0); i < 32; i++ {
		if a.State[0]&(1<<i) != 0 {
			states = append(states, State(i).String())
		}

	}

	for i := uint32(0); i < 32; i++ {
		if a.State[1]&(1<<i) != 0 {
			states = append(states, State(i+32).String())
		}
	}

	return states
}

func (a *AccessibleObject) PrepareStatesAsText() string {
	states := a.CurrentStates()
	if len(states) == 0 {
		return ""
	}
	return strings.Join(states, ", ")
}
