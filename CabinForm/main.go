package main

import (
	"GolangWebApp/CabinForm/handlers"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	//_ "github.com/mattn/go-sqlite3"
)

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

	// db := initDB("storage.db")
	// migrate(db)

	handlers.InitMap()

	port := ":" + portFromArgs()

	srv := &http.Server{Addr: port}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world\n")
	})

	http.HandleFunc("/otherEndpoint", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header)
		io.WriteString(w, "hello world other option\n")
	})

	http.HandleFunc("/polls", handlers.GetPolls)

	http.HandleFunc("/poll/", handlers.UpdatePoll)

	http.HandleFunc("/testJson", returnJSON)

	http.HandleFunc("/serveVue", serveVue)

	http.HandleFunc("/serveFile", serveFile)

	http.HandleFunc("/testCookie", generateCookie)

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

func generateCookie(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "my-cookie",
		Value:  "some value",
		MaxAge: 5,
		Path:   "/",
	})
}

func serveFile(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "hello world from Carl\n")
	http.ServeFile(w, req, "index.html")
}

func serveVue(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "vue.html")
}

type Person struct {
	Name string
	Age  int
}

func returnJSON(w http.ResponseWriter, req *http.Request) {
	carl := Person{"carl", 22}
	c, err := json.Marshal(&carl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c, carl)
	io.WriteString(w, string(c))
}

func main() {
	log.Printf("main: starting HTTP server")

	srv := startHTTPServer()

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
