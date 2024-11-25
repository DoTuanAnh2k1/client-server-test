package main

import (
	"crypto/tls"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	conf := &tls.Config{
		MaxVersion:         tls.VersionTLS12,
		InsecureSkipVerify: true,
		// CipherSuites: []uint16{
		// 	tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		// 	tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		// 	tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		// 	tls.TLS_RSA_WITH_AES_128_CBC_SHA,
		// },
	}

	conn, err := tls.Dial("tcp", "127.0.0.1:8443", conf)
	if err != nil {
		log.Println(err)
		return
	}

	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}
	err = conn.Close()
	if err != nil {
		log.Println(err)
	}

	println(string(buf[:n]))
}
