package main

import (
	"GolangWebApp/CabinForm/handlers"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func startHTTPServer() *http.Server {

	db := initDB("storage.db")
	migrate(db)

	srv := &http.Server{Addr: ":8081"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world\n")
	})

	http.HandleFunc("/otherEndpoint", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header)
		io.WriteString(w, "hello world other option\n")
	})

	http.HandleFunc("/polls", handlers.GetPolls)

	http.HandleFunc("/poll/:index", handlers.UpdatePoll)

	http.HandleFunc("/testJson", returnJSON)

	http.HandleFunc("/serveVue", serveVue)

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

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS polls(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL,
			topic VARCHAR NOT NULL,
			src VARCHAR NOT NULL,
			upvotes INTEGER NOT NULL,
			downvotes INTEGER NOT NULL,
			UNIQUE(name)
	);

	INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Angular','Awesome Angular', 'https://cdn.colorlib.com/wp/wp-content/uploads/sites/2/angular-logo.png', 1, 0);

	INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Vue', 'Voguish Vue','https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Vue.js_Logo.svg/400px-Vue.js_Logo.svg.png', 1, 0);

	INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('React','Remarkable React','https://upload.wikimedia.org/wikipedia/commons/thumb/a/a7/React-icon.svg/1200px-React-icon.svg.png', 1, 0);

	INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Ember','Excellent Ember','https://cdn-images-1.medium.com/max/741/1*9oD6P0dEfPYp3Vkk2UTzCg.png', 1, 0);

	INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Knockout','Knightly Knockout','https://images.g2crowd.com/uploads/product/image/social_landscape/social_landscape_1489710848/knockout-js.png', 1, 0);
`
	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}
