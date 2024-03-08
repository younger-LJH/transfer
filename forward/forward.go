package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"

	"mycrypto"
)

func process(conn net.Conn, key []byte) []byte {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	buf := make([]byte, 128)
	n, err := reader.Read(buf)
	if err != nil {
		fmt.Printf("[error] forward read failed, err: %+v\n", err)
		return nil
	}

	decryptedData := mycrypto.AesDecryptCFB(buf[:n], key)
	return decryptedData
}

func forward(address string, data []byte, key []byte) {
	fmt.Printf("[info] forward start dial %s\n", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("[error] forward dial %s failed, err: %+v\n", address, err)
		return
	}
	defer conn.Close()

	encryptedData := mycrypto.AesEncryptCFB(data, key)
	_, err = conn.Write(encryptedData)
	if err != nil {
		fmt.Printf("[error] client write %s failed: %+v\n", encryptedData, err)
		return
	}

	fmt.Printf("[info] forward write success, data: %s\n", encryptedData)
}

func main() {
	listenAddress := flag.String("l", "127.0.0.1:4000", "listen IP:Port")
	forwardAddress := flag.String("t", "127.0.0.1:5000", "forward IP:Port")
	key1 := flag.String("k1", "password12345678", "key1")
	key2 := flag.String("k2", "password87654321", "key2")
	flag.Parse()

	fmt.Printf("[info] forward start listen %s\n", *listenAddress)
	listen, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		fmt.Printf("[error] forward listen %s failed, err: %+v\n", *listenAddress, err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("[error] forward accept failed, err: %+v\n", err)
			continue
		}

		data := process(conn, []byte(*key1))
		forward(*forwardAddress, data, []byte(*key2))
	}
}
