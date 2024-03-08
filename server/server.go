package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"

	"mycrypto"
)

func process(conn net.Conn, key []byte) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	buf := make([]byte, 128)
	n, err := reader.Read(buf)
	if err != nil {
		fmt.Printf("[error] server read failed, err: %+v\n", err)
		return
	}

	decryptedData := mycrypto.AesDecryptCFB(buf[:n], key)
	fmt.Printf("[info] server receive data: %s\n", decryptedData)
}

func main() {
	address := flag.String("l", "127.0.0.1:5000", "listen IP:Port")
	key := flag.String("k", "password87654321", "key")
	flag.Parse()

	fmt.Printf("[info] server start listen %s\n", *address)
	listen, err := net.Listen("tcp", *address)
	if err != nil {
		fmt.Printf("[error] server listen %s failed, err: %+v\n", *address, err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("[error] server accept failed, err: %+v\n", err)
			continue
		}

		go process(conn, []byte(*key))
	}
}
