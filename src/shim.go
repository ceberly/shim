package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	listenPort  int
	connectHost string
	verbose     bool
)

func init() {
	const (
		defaultListenPort = 8080
	)

	flag.IntVar(&listenPort, "l", defaultListenPort, "local port to listen on")
	flag.StringVar(&connectHost, "h", "http://localhost/", "remote url to connect to")
	flag.BoolVar(&verbose, "v", false, "print debugging information")
}

func main() {
	flag.Parse()

	l := fmt.Sprintf(":%d", listenPort)

	u, err := url.Parse(connectHost)
	if err != nil {
		log.Fatal(err)
	}

	shim := httputil.NewSingleHostReverseProxy(u)
	origDirector := shim.Director
	shim.Director = func(r *http.Request) {
		r.Host = u.Host

		if verbose {
			log.Printf("%+v", r)
		}

		origDirector(r)
	}

	http.Handle("/", shim)

	log.Printf("Ok, we are listening on port %d and connecting to %s ...", listenPort, connectHost)
	log.Fatal(http.ListenAndServe(l, nil))
}
