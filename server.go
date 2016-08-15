package seccon

import (
	"io/ioutil"
	"net"

	"golang.org/x/crypto/ssh"
)

type server struct {
	listener net.Listener
	config   *ssh.ServerConfig
}

// Listen start the TCP server on addr and listening for connections
func Listen(addr string, privateKeyPath string) (net.Listener, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	config, err := serverConfig(privateKeyPath)
	if err != nil {
		return nil, err
	}
	return &server{listener: listener, config: config}, nil
}

// Accept accepts new connection and makes it secure
func (s *server) Accept() (net.Conn, error) {
	conn, err := s.listener.Accept()
	if err != nil {
		return nil, err
	}
	sshChannel, err := s.channelFromConn(conn)
	if err != nil {
		return nil, err
	}
	return &secureConnection{conn: conn, channel: sshChannel}, nil
}

// Close closes secure connection
func (s *server) Close() error {
	return s.listener.Close()
}

// Addr returns resolved local address of listener
func (s *server) Addr() net.Addr {
	return s.listener.Addr()
}

func serverConfig(privateKeyPath string) (*ssh.ServerConfig, error) {
	config := &ssh.ServerConfig{
		PublicKeyCallback: func(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			return &ssh.Permissions{}, nil
		},
	}
	var path string
	if len(privateKeyPath) == 0 {
		path = defaultKeyPath()
	} else {
		path = privateKeyPath
	}
	privateBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	private, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		return nil, err
	}
	config.AddHostKey(private)
	return config, nil
}

func (s *server) channelFromConn(conn net.Conn) (ssh.Channel, error) {
	_, ch, _, err := ssh.NewServerConn(conn, s.config)
	if err != nil {
		return nil, err
	}
	c := <-ch
	sshChannel, _, err := c.Accept()
	if err != nil {
		return nil, err
	}
	return sshChannel, nil
}
