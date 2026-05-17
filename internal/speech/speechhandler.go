package speech

import (
	"bufio"
	"fmt"
	"net"
	"os/user"
	"strings"
)

type Client struct {
	conn net.Conn
}

func NewClient() (*Client, error) {
	currentUser, err := user.Current()
	uid := "1000"
	username := "unknown"

	if err == nil {
		uid = currentUser.Uid
		username = currentUser.Username
	}

	socketPath := fmt.Sprintf("/run/user/%s/speech-dispatcher/speechd.sock", uid)
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return nil, err
	}

	client := &Client{conn: conn}

	// Inicia a drenagem assíncrona. Isso impede que o socket fique entupido
	// e trave o programa principal.
	go client.drain(bufio.NewReader(conn))

	clientName := fmt.Sprintf("%s:aura:core", username)

	// Envia o handshake sem bloquear
	client.rawCommand(fmt.Sprintf("SET self CLIENT_NAME %s", clientName))

	return client, nil
}

// drain lê e descarta as respostas do servidor (ex: 200 OK, 225 MESSAGE QUEUED)
func (c *Client) drain(reader *bufio.Reader) {
	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			return // Sai da goroutine se a conexão cair
		}
	}
}

func (c *Client) rawCommand(cmd string) {
	if c.conn != nil {
		fmt.Fprintf(c.conn, "%s\r\n", cmd)
	}
}

func (c *Client) Speak(text string) {
	if c.conn == nil || text == "" {
		return
	}
	payload := fmt.Sprintf("SPEAK\r\n%s\r\n.\r\n", text)
	fmt.Fprint(c.conn, payload)
}

func (c *Client) Cancel() {
	c.rawCommand("CANCEL self")
}

func (c *Client) SetPriority(priority string) {
	c.rawCommand(fmt.Sprintf("SET self PRIORITY %s", strings.ToUpper(priority)))
}

func (c *Client) SetRate(rate int) {
	c.rawCommand(fmt.Sprintf("SET self RATE %d", rate))
}

func (c *Client) Close() {
	if c.conn != nil {
		c.rawCommand("QUIT")
		c.conn.Close()
	}
}
