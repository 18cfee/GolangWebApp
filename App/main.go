package main

import (
	"GolangWebApp/App/form"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func init() {
}

func portFromArgs() string {
	port := "8081"
	if len(os.Args) == 2 {
		port = os.Args[1]
	}
	fmt.Println("The port is:", port)
	portInt, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		fmt.Println("There was an error parsing the port, err:", err)
		os.Exit(3)
	}
	if portInt < 1 || portInt > 9999 {
		fmt.Println("Error the port number is not in the valid range")
		os.Exit(3)
	}
	return port
}

func startHTTPServer() *http.Server {
	port := ":" + portFromArgs()

	srv := &http.Server{Addr: port}

	fs := http.FileServer(http.Dir("FrontEnd"))
	http.Handle("/", fs)
	http.HandleFunc("/updateForm", form.Put)
	http.HandleFunc("/getFormInfo", form.Get)

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
	fmt.Printf("Starting server on Port%v\n", port)
	return srv
}

func main() {
	log.Printf("main: starting HTTP server")

	srv := startHTTPServer()

	// stop the server by typing anything and hitting enter
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
