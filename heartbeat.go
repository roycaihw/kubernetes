package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

var httpVersion = flag.Int("version", 2, "HTTP version")

func main() {
	flag.Parse()
	go startListenAndServe()
	heartbeatToServer()
}

func startListenAndServe() {
	handle := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got connection: %s", r.Proto)
		w.Write([]byte("Hello"))
	}
	srv := &http.Server{Addr: ":9999", Handler: http.HandlerFunc(handle)}
	log.Printf("Serving on https://0.0.0.0:9999")
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}

func heartbeatToServer() {
	log.Printf("start heartbeat with apiserver\n")
	t := time.NewTicker(time.Second)
	defer t.Stop()

	c := initHttpClient()
	for range t.C {
		str, err := requestByClient(c)
		if err != nil {
			log.Printf("occur error: %v", err)
			continue
		}
		log.Println(str)
	}
}

func initHttpClient() *http.Client {
	caCert, err := ioutil.ReadFile("server.crt")
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS configuration with the certificate of the server
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	if *httpVersion == 2 {
		if err := http2.ConfigureTransport(transport); err != nil {
			log.Printf("Transport failed http2 configuration: %v", err)
		}
	}

	client.Transport = transport

	return client
}

func requestByClient(c *http.Client) (string, error) {
	resp, err := c.Get("https://localhost:9999")
	if err != nil {
		return "", fmt.Errorf("http request error: %v\n", err)
	}

	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ioutils read response failed: %v", err)
	}
	return string(bs), nil
}
