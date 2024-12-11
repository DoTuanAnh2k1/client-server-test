package server

import (
	"crypto/tls"
	"log"
)

func NewServer() {
	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{
		CipherSuites: []uint16{
			// for the purpose of education we avoid AEAD cipher suites
			0xc013, // ECDHE-RSA-AES128-SHA
			0xc009, // ECDHE-ECDSA-AES128-SHA
			0xc014, // ECDHE-RSA-AES256-SHA
			0xc00a, // ECDHE-ECDSA-AES256-SHA
			0x002f, // RSA-AES128-SHA
			0x0035, // RSA-AES256-SHA
			0xc012, // ECDHE-RSA-3DES-EDE-SHA
			0x000a, // RSA-3DES-EDE-SHA
		},
		Certificates: []tls.Certificate{cer},
	}

	ln, err := tls.Listen("tcp", ":8443", config)
	if err != nil {
		panic(err)
	}

	log.Println("tls server running at :8443")
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	rdata := make([]byte, 1024)
	n, err := conn.Read(rdata)
	log.Println("Message received:", string(rdata[:n]))
	if err != nil {
		panic(err)
	}

	_, err = conn.Write([]byte("world\n"))
	if err != nil {
		panic(err)
	}
}
