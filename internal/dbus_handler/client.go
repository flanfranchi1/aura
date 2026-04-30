package dbus_handler

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

type Client struct {
	conn *dbus.Conn
}

func NewClient() (*Client, error) {
	sessionBus, err := dbus.ConnectSessionBus()
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar no Session Bus: %w", err)
	}
	defer sessionBus.Close() // Fecha o canal de descoberta assim que a função termina

	var a11yAddress string
	obj := sessionBus.Object("org.a11y.Bus", "/org/a11y/bus")
	err = obj.Call("org.a11y.Bus.GetAddress", 0).Store(&a11yAddress)
	if err != nil {
		return nil, fmt.Errorf("falha ao obter endereço A11y: %w", err)
	}

	a11yConn, err := dbus.Connect(a11yAddress)
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar no endereço A11y %s: %w", a11yAddress, err)
	}

	if err = a11yConn.Auth(nil); err != nil {
		a11yConn.Close()
		return nil, fmt.Errorf("falha na autenticação A11y: %w", err)
	}

	return &Client{conn: a11yConn}, nil
}

func (c *Client) subscribe(rule string) (<-chan *dbus.Signal, error) {
	err := c.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, rule).Err
	if err != nil {
		return nil, fmt.Errorf("erro ao adicionar match rule [%s]: %w", rule, err)
	}

	ch := make(chan *dbus.Signal, 100)
	c.conn.Signal(ch)

	return ch, nil
}

func (c *Client) SubscribeToFocusedStateChanged() (<-chan *dbus.Signal, error) {
	rule := "type='signal',interface='org.a11y.atspi.Event.Object',member='StateChanged',arg0='focused'"
	return c.subscribe(rule)
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) getProperty(sender, path, property string) dbus.Variant {
	obj := c.conn.Object(sender, dbus.ObjectPath(path))
	variant, _ := obj.GetProperty(property)
	return variant
}

func (c *Client) GetStringProperty(sender, path, property string) (string, error) {
	variant := c.getProperty(sender, path, property)
	if v, ok := variant.Value().(string); ok {
		return v, nil
	}
	return "", fmt.Errorf("%s property is not a string", property)
}

func (c *Client) GetUint32Property(sender, path, property string) (uint32, error) {
	variant := c.getProperty(sender, path, property)
	if v, ok := variant.Value().(uint32); ok {
		return v, nil
	}
	return 0, fmt.Errorf("%s property is not a uint32", property)
}

func (c *Client) GetStateSetProperty(sender, path, property string) ([]uint32, error) {
	variant := c.getProperty(sender, path, property)
	if v, ok := variant.Value().([]uint32); ok {
		return v, nil
	}
	return nil, fmt.Errorf("%s property is not a StateSet", property)
}

func (c *Client) GetMapProperty(sender, path, property string) (map[string]string, error) {
	variant := c.getProperty(sender, path, property)
	if v, ok := variant.Value().(map[string]string); ok {
		return v, nil
	}
	return nil, fmt.Errorf("%s property is not a map[string]string", property)
}

func (c *Client) ObjectPath(str string) dbus.ObjectPath {
	return dbus.ObjectPath(str)
}
