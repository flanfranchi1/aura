package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/godbus/dbus/v5"
)

type ATSPIElement struct {
	Name       string
	Role       string
	States     []uint32
	Attributes map[string]string
	Address    string
	Path       dbus.ObjectPath
}

type ATSPIReference struct {
	Name string
	Path dbus.ObjectPath
}

type Extractor struct {
	SessionConn *dbus.Conn
	A11yConn    *dbus.Conn
	Writer      *os.File
}

func NewExtractor(outputPath string) (*Extractor, error) {
	file, err := os.Create(outputPath)
	if err != nil {
		return nil, err
	}

	sessionConn, err := dbus.ConnectSessionBus()
	if err != nil {
		file.Close()
		return nil, err
	}

	return &Extractor{
		SessionConn: sessionConn,
		Writer:      file,
	}, nil
}

func (e *Extractor) ConnectA11yBus() error {
	obj := e.SessionConn.Object("org.a11y.Bus", "/org/a11y/bus")

	var address string
	err := obj.Call("org.a11y.Bus.GetAddress", 0).Store(&address)
	if err != nil {
		return err
	}

	a11yConn, err := dbus.Dial(address)
	if err != nil {
		return err
	}

	err = a11yConn.Auth(nil)
	if err != nil {
		a11yConn.Close()
		return err
	}

	err = a11yConn.Hello()
	if err != nil {
		a11yConn.Close()
		return err
	}

	e.A11yConn = a11yConn
	return nil
}

func (e *Extractor) FetchRootChildren() ([]ATSPIReference, error) {
	obj := e.A11yConn.Object("org.a11y.atspi.Registry", "/org/a11y/atspi/accessible/root")

	var children []ATSPIReference

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := obj.CallWithContext(ctx, "org.a11y.atspi.Accessible.GetChildren", 0).Store(&children)
	if err != nil {
		return nil, err
	}

	return children, nil
}

func (e *Extractor) InspectNode(serviceName string, path dbus.ObjectPath) (*ATSPIElement, error) {
	obj := e.A11yConn.Object(serviceName, path)
	element := &ATSPIElement{
		Address: serviceName,
		Path:    path,
	}

	err := obj.Call("org.a11y.atspi.Accessible.GetName", 0).Store(&element.Name)
	if err != nil {
		return nil, err
	}

	var role uint32
	obj.Call("org.a11y.atspi.Accessible.GetRole", 0).Store(&role)
	element.Role = fmt.Sprintf("%d", role)

	obj.Call("org.a11y.atspi.Accessible.GetState", 0).Store(&element.States)
	obj.Call("org.a11y.atspi.Accessible.GetAttributes", 0).Store(&element.Attributes)

	return element, nil
}

func (e *Extractor) WriteElement(el *ATSPIElement) error {
	line := fmt.Sprintf("Path: %s | Name: %s | Role: %s | States: %v | Attributes: %v\n",
		el.Path, el.Name, el.Role, el.States, el.Attributes)

	_, err := e.Writer.WriteString(line)
	return err
}

func (e *Extractor) Close() {
	if e.A11yConn != nil {
		e.A11yConn.Close()
	}
	if e.SessionConn != nil {
		e.SessionConn.Close()
	}
	if e.Writer != nil {
		e.Writer.Close()
	}
}

func main() {
	extractor, err := NewExtractor("atspi_dump.txt")
	if err != nil {
		fmt.Printf("Initialization failed: %v\n", err)
		os.Exit(1)
	}
	defer extractor.Close()

	err = extractor.ConnectA11yBus()
	if err != nil {
		fmt.Printf("A11y bus connection failed: %v\n", err)
		os.Exit(1)
	}

	children, err := extractor.FetchRootChildren()
	if err != nil {
		fmt.Printf("Root fetch failed: %v\n", err)
		os.Exit(1)
	}

	extractor.Writer.WriteString("AT-SPI2 Root Children Extraction\n================================\n")

	for _, ref := range children {
		el, err := extractor.InspectNode(ref.Name, ref.Path)
		if err != nil {
			fmt.Printf("Skipped node %s on %s: %v\n", ref.Path, ref.Name, err)
			continue
		}
		extractor.WriteElement(el)
	}

	fmt.Println("Extraction completed. Check atspi_dump.txt")
}
