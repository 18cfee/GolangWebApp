package form

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// map[formId]json
var forms = map[string]string{}

func Put(w http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		// need some error handling here
		fmt.Println("the method was ", req.Method)
	}
	// get the form from the body
	formId := req. .Query().Get("formId")
	req.Body.Close()
	json, _ := ioutil.ReadAll(req.Body)
	forms[formId] = string(json)
}

func Get(w http.ResponseWriter, req *http.Request) {
	
}
