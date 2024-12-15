package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const (
	HTTPVersion1 = 1
	HTTPVersion2 = 2
)

func NewHTTPServer(httpServerVersion uint8) {
	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
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

	var s *http.Server
	switch httpServerVersion {
	case HTTPVersion1:
		s = newServerVersion1()
	case HTTPVersion2:
		s = newServerVersion2()
	default:
		panic("wrong version")
	}
	s.TLSConfig = config
	err = s.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		panic(err)
	}
}

func newServerVersion1() *http.Server {
	mux := newRouter()
	return &http.Server{
		Addr:    ":8457",
		Handler: mux,
	}
}

func newServerVersion2() *http.Server {
	h2s := &http2.Server{}
	mux := newRouter()
	return &http.Server{
		Addr:    ":8457",
		Handler: h2c.NewHandler(mux, h2s),
	}
}

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", reqHandler)
	mux.HandleFunc("/post", postHandler)
	return mux
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler post request")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func NewClientVer2() {
	caCertBytes, err := os.ReadFile("ca.crt")
	if err != nil {
		panic(err)
	}
	caCertPem, _ := pem.Decode(caCertBytes)
	caCert, err := x509.ParseCertificate(caCertPem.Bytes)
	if err != nil {
		panic(err)
	}
	caList := x509.NewCertPool()
	caList.AddCert(caCert)

	keyLogFile, err := os.OpenFile("keylog.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed to open key log file: %v", err)
	}
	defer keyLogFile.Close()

	conf := &tls.Config{
		MaxVersion:         tls.VersionTLS12,
		InsecureSkipVerify: true,
		RootCAs:            caList,
		KeyLogWriter:       keyLogFile,
	}

	c := &http.Client{
		Transport: &http2.Transport{
			// AllowHTTP: true,
			// DisableCompression: true,
			TLSClientConfig: conf,
		},
		Timeout: 2 * time.Second,
	}
	bodyMess := "Abc-ASDJIOEF"
	url := "https://localhost:8457/post"
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(bodyMess)))
	req.Header.Set("Abc-ASDJIOEF-----b", "value")
	// resp, err := c.Get("http://localhost:8457/")
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Body)
}

func NewClientVer1() {
	caCertBytes, err := os.ReadFile("ca.crt")
	if err != nil {
		panic(err)
	}
	caCertPem, _ := pem.Decode(caCertBytes)
	caCert, err := x509.ParseCertificate(caCertPem.Bytes)
	if err != nil {
		panic(err)
	}
	caList := x509.NewCertPool()
	caList.AddCert(caCert)

	keyLogFile, err := os.OpenFile("keylog.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed to open key log file: %v", err)
	}
	defer keyLogFile.Close()

	conf := &tls.Config{
		MaxVersion:         tls.VersionTLS12,
		InsecureSkipVerify: true,
		RootCAs:            caList,
		KeyLogWriter:       keyLogFile,
	}

	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: conf,
		},
	}
	resp, err := c.Get("https://localhost:8457/")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Body)
}

func main() {
	go NewHTTPServer(HTTPVersion2)
	time.Sleep(2 * time.Second)
	// // go NewHTTPServer(HTTPVersion1)
	NewClientVer2()
	// NewClientVer1()
}
