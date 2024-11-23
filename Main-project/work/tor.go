package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/proxy"
)

const (
	torSocks5Proxy = "socks5://127.0.0.1:9050" // Tor proxy
)

func work() {
	// Setup the Tor SOCKS5 proxy
	dialer, err := proxy.SOCKS5("tcp", torSocks5Proxy, nil, proxy.Direct)
	if err != nil {
		log.Fatalf("Failed to connect to Tor proxy: %v", err)
	}

	// Create an HTTP client that uses the Tor proxy
	client := &http.Client{
		Transport: &http.Transport{
			Dial: dialer.Dial,
		},
		Timeout: 10 * time.Second,
	}

	// Make a request through the Tor network
	url := "http://check.torproject.org" // This URL will check if the request comes from the Tor network
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Output the result
	fmt.Printf("Response status: %s\n", resp.Status)
}
