package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func startHTTPServer() *http.Server {
	srv := &http.Server{Addr: ":8081"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world\n")
	})

	http.HandleFunc("/otherEndpoint", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header)
		io.WriteString(w, "hello world other option\n")
	})

	http.HandleFunc("/serveFile", serveFile)

	go func() {
		// returns ErrServerClosed on graceful close
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// NOTE: there is a chance that next line won't have time to run,
			// as main() doesn't wait for this goroutine to stop. don't use
			// code with race conditions like these for production. see post
			// comments below on more discussion on how to handle this.
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	// returning reference so caller can call Shutdown()
	return srv
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world from Carl\n")
}

func main() {
	log.Printf("main: starting HTTP server")

	srv := startHTTPServer()

	var t time.Duration

	t = 2

	fmt.Printf("%T\n", t)

	log.Printf("main: serving for %v seconds", int(t))

	time.Sleep(time.Second * t)

	var anything string
	fmt.Scanf("%v", &anything)

	log.Printf("main: stopping HTTP server")

	// now close the server gracefully ("shutdown")
	// timeout could be given with a proper context
	// (in real world you shouldn't use TODO()).
	if err := srv.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}

	log.Printf("main: done. exiting")
}
