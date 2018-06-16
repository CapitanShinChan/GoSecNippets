package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	/*
		//proxyUrl, err := url.Parse("http://127.0.0.1:8080")

		dialer := &net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}
		//http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
		// or create your own transport, there's an example on godoc.
		http.DefaultTransport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			if addr == "6d6.es:80" {
				addr = "104.31.83.248:80"
				fmt.Println("Direccion cambiada")
			}
			return dialer.DialContext(ctx, network, addr)
		}

		resp, err := http.Get("http://6d6.es/")
		log.Println(resp.Header, err)

	*/
	ipToRedirect := flag.String("ip", "", "IP address to redirect the HTTP request. Usually from a CDN.")

	flag.Parse()
	if *ipToRedirect == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}
	// or create your own transport, there's an example on godoc.
	http.DefaultTransport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		if addr == "6d6.es:80" {
			addr = *ipToRedirect
		}
		return dialer.DialContext(ctx, network, addr)
	}
	resp, err := http.Get("http://6d6.es/")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(resp.Header)
	}
}
