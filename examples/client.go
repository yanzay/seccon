package main

import (
	"io"
	"os"

	"github.com/yanzay/seccon"
)

func main() {
	client := seccon.NewClient("yanzay")
	conn, err := client.Dial("localhost:2022")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(conn, os.Stdin)
	if err != nil {
		panic(err)
	}
}
