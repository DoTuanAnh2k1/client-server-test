package client

import (
	"crypto/tls"
	"log"
	"os"
)

func NewClient() {
	// caCertBytes, err := os.ReadFile("ca.crt")
	// if err != nil {
	// 	panic(err)
	// }
	// caCertPem, _ := pem.Decode(caCertBytes)
	// caCert, err := x509.ParseCertificate(caCertPem.Bytes)
	// if err != nil {
	// 	panic(err)
	// }
	// caList := x509.NewCertPool()
	// caList.AddCert(caCert)

	keyLogFile, err := os.OpenFile("keylog.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed to open key log file: %v", err)
	}
	defer keyLogFile.Close()

	conf := &tls.Config{
		MaxVersion: tls.VersionTLS12,
		// MaxVersion:         tls.VersionTLS13,
		InsecureSkipVerify: true, // to by pass verify certificate
		// RootCAs:            caList,
		KeyLogWriter: keyLogFile,
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

	log.Println(string(buf[:n]))
}
