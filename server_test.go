package seccon

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestConnection(t *testing.T) {
	resp := make(chan string)
	listener, _ := Listen(":9876", "")
	go func() {
		conn, _ := listener.Accept()
		b, _ := ioutil.ReadAll(conn)
		resp <- string(b)
	}()
	client := NewClient("test")
	clientConn, err := client.Dial("localhost:9876")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	fmt.Fprint(clientConn, "hi")
	err = clientConn.Close()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if <-resp != "hi" {
		t.Fail()
	}
}
