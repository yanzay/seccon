package seccon

import (
	"io/ioutil"
	"net"
	"os"
	"path"

	"golang.org/x/crypto/ssh"
)

// SecureDialer interface
type SecureDialer interface {
	Dial(string) (net.Conn, error)
	SetPrivateKeyPath(string)
}

// Dial connects to the seccon server
func (c *client) Dial(addr string) (net.Conn, error) {
	return c.dial(addr)
}

// Set custom private key, by default its `~/.ssh/id_rsa`
func (c *client) SetPrivateKeyPath(path string) {
	c.privateKey = path
}

func defaultKeyPath() string {
	home := os.Getenv("HOME")
	if len(home) > 0 {
		return path.Join(home, ".ssh/id_rsa")
	}
	return ""
}

// NewClient creates new seccon client
func NewClient(user string) SecureDialer {
	return &client{
		privateKey: defaultKeyPath(),
		user:       user,
	}
}

type client struct {
	privateKey string
	user       string
}

func (c *client) dial(addr string) (net.Conn, error) {
	signer, err := signerFromFile(c.privateKey)
	if err != nil {
		return nil, err
	}
	config := &ssh.ClientConfig{
		User: c.user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	}

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	clientConn, _, _, err := ssh.NewClientConn(conn, addr, config)
	if err != nil {
		return nil, err
	}

	ch, _, err := clientConn.OpenChannel("data", []byte{})
	if err != nil {
		return nil, err
	}
	return &secureConnection{conn: conn, channel: ch}, nil
}

func signerFromFile(path string) (ssh.Signer, error) {
	key, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}
	return signer, nil
}
