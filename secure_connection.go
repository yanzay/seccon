package seccon

import (
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

// Implements net.Conn
type secureConnection struct {
	conn    net.Conn
	channel ssh.Channel
}

func (sc *secureConnection) Read(b []byte) (n int, err error) {
	return sc.channel.Read(b)
}

func (sc *secureConnection) Write(b []byte) (n int, err error) {
	return sc.channel.Write(b)
}

func (sc *secureConnection) Close() error {
	err := sc.channel.Close()
	if err != nil {
		return err
	}
	return sc.conn.Close()
}

func (sc *secureConnection) LocalAddr() net.Addr {
	return sc.conn.LocalAddr()
}

func (sc *secureConnection) RemoteAddr() net.Addr {
	return sc.conn.RemoteAddr()
}

func (sc *secureConnection) SetDeadline(t time.Time) error {
	return sc.conn.SetDeadline(t)
}

func (sc *secureConnection) SetReadDeadline(t time.Time) error {
	return sc.conn.SetReadDeadline(t)
}

func (sc *secureConnection) SetWriteDeadline(t time.Time) error {
	return sc.conn.SetWriteDeadline(t)
}
