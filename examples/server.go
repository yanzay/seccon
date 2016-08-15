package main

import (
	"io"
	"log"
	"os"

	"github.com/yanzay/seccon"
)

func main() {
	listener, err := seccon.Listen(":2022", "")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go func() {
			_, err := io.Copy(os.Stdout, conn)
			if err != nil {
				log.Println(err)
				return
			}
		}()
	}
}
