package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	localhost = "127.0.0.1"
)

var (
	root = flag.String("d", "", "Root path to serve files from (default: current working directory)")
	port = flag.Int("p", 8080, "Port to serve files from")
)

func main() {
	flag.Parse()

	if *root == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("error getting working directory: %v", err)
		}
		*root = wd
	}

	cleanPath := filepath.Clean(*root)
	info, err := os.Stat(cleanPath)
	if err != nil || !info.IsDir() {
		flag.Usage()
		log.Fatalf("Invalid path: %s", cleanPath)
	}

	if *port < 0 || *port > 65535 {
		flag.Usage()
		log.Fatalf("invalid port: %v", *port)
	}

	address := strings.Join([]string{localhost, fmt.Sprintf("%v", *port)}, ":")

	log.Fatal(serve(address, cleanPath))
}

func serve(address, path string) error {
	log.Printf("Serving files from %s on http://%s\n", path, address)
	return http.ListenAndServe(
		address,
		http.FileServer(http.Dir(path)))
}
