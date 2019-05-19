package main

import (
	"fmt"
	"net/http"
	"runtime"
	"text/template"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	un           string
	lastActivity time.Time
}

var tpl *template.Template
var Users = map[string]user{}       // user ID, user
var Sessions = map[string]session{} // session ID, session

const sessionLength int = 10

func initUsers() {
	u1 := user{
		UserName: "carlfee",
		Password: []byte("testing"),
		First:    "Carl",
		Last:     "Fee",
		Role:     "Developer",
	}
	Users[u1.UserName] = u1
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u user
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, ok := Users[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
		Sessions[c.Value] = session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := Sessions[c.Value]
	if ok {
		s.lastActivity = time.Now()
		Sessions[c.Value] = s
	}
	_, ok = Users[s.un]
	// refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
}

func timeCleans() {
	for {
		runtime.Gosched()
		time.Sleep(1 * time.Second)
		cleanSessions()
	}
}

// todo add manual cleaning and logging for cleaning
func cleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	showSessions()              // for demonstration purposes
	for k, v := range Sessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(Sessions, k)
		}
	}
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	showSessions()             // for demonstration purposes
}

// for demonstration purposes
func showSessions() {
	fmt.Println("********")
	for k, v := range Sessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}