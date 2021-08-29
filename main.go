package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:2701")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		echo(listener)
	}
}

func echo(listener net.Listener) {
	conn, err := listener.Accept()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	go func() {
		defer conn.Close()
		for {
			buf := make([]byte, 1500)
			n, err := conn.Read(buf)
			if err != nil {
				if err != io.EOF {
					fmt.Fprintln(os.Stderr, err.Error())
				}
				break
			}
			fmt.Println("---")
			fmt.Printf("received from %v\n", conn.RemoteAddr().String())
			fmt.Printf("%v", string(buf[:n]))

			_, err = conn.Write(buf[:n])
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				break
			}
		}
	}()
}
