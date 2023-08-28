package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

const port = ":8080"

var (
	cert = flag.String("tlsCertFile", "/etc/certs/cert.pem", "The certificate file.")
	key  = flag.String("tlsKeyFile", "/etc/certs/key.pem", "The key file")
)

func main() {
	flag.Parse()

	r := router.New()
	r.GET("/validate", Validate)

	server := fasthttp.Server{
		Handler: r.Handler,
		Name:    "security-policy-admissions",
	}

	go func() {
		err := server.ListenAndServeTLS(port, *cert, *key)
		if err != nil {
			panic(err)
		}
	}()

	// Log status message to stdout
	fmt.Printf("Server running listening on port: %v\n", port)

	// Listen for shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	// Emit a shutdown log
	fmt.Println("Got shutdown signal, shutting down webhook server gracefully...")
	server.Shutdown()
}
