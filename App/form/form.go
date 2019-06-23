package form

import (
	"fmt"
	"net/http"
)

var forms = map[string]map[string]string{}

func Put(w http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		// need some error handling here
		fmt.Println("the method was ", req.Method)
	}
	fmt.Println("the put was called")

}

func Get(w http.ResponseWriter, req *http.Request) {
	fmt.Println("the get was called")
}
