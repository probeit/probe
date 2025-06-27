package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	message = flag.String("message", "Hello from probe!", "Message to echo")
	port    = flag.String("port", "8943", "Port to listen on")
)

func main() {
	flag.Parse()

	// If no message provided via flag, check environment variable
	if *message == "Hello from probe!" {
		if envMsg := os.Getenv("ECHO_MESSAGE"); envMsg != "" {
			*message = envMsg
		}
	}

	// If still default and args provided, use first arg as message
	if *message == "Hello from probe!" && len(os.Args) > 1 {
		// Check if the argument is not a flag
		if !flag.Parsed() || len(flag.Args()) > 0 {
			*message = flag.Args()[0]
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := fmt.Sprintf(`{
  "message": "%s",
  "method": "%s",
  "path": "%s",
  "headers": %v,
  "host": "%s",
  "remote_addr": "%s"
}`, *message, r.Method, r.URL.Path, r.Header, r.Host, r.RemoteAddr)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, response)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status": "healthy"}`)
	})

	log.Printf("Starting server on port %s with message: %s", *port, *message)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
