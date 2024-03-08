package main

import (
	"flag"
	"fmt"
	"net"

	"mycrypto"
)

func main() {
	address := flag.String("t", "127.0.0.1:4000", "dial IP:Port")
	data := flag.String("c", "Hello world", "data")
	key := flag.String("k", "password12345678", "key")
	flag.Parse()

	fmt.Printf("[info] client start dial %s\n", *address)
	conn, err := net.Dial("tcp", *address)
	if err != nil {
		fmt.Printf("[error] client dial %s failed, err: %+v\n", *address, err)
		return
	}
	defer conn.Close()

	encryptedData := mycrypto.AesEncryptCFB([]byte(*data), []byte(*key))
	_, err = conn.Write(encryptedData)
	if err != nil {
		fmt.Printf("[error] client write %s failed: %+v\n", encryptedData, err)
		return
	}

	fmt.Printf("[info] client write success, data: %s\n", encryptedData)
}
