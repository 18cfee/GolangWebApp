package form

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// map[formId]json
var forms = map[int]string{}

func Put(w http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		fmt.Println("the method was ", req.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	// get the form from the body
	formId, err := strconv.Atoi(req.URL.Query().Get("formId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(formId)
	//req.Body.Close()
	json, _ := ioutil.ReadAll(req.Body)
	forms[formId] = string(json)
	fmt.Println("the body was set to: ", forms[formId])
	w.WriteHeader(http.StatusOK)
}

func Get(w http.ResponseWriter, req *http.Request) {
	fmt.Println("the get method is being called")
	formId, err := strconv.Atoi(req.URL.Query().Get("formId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	json, exists := forms[formId]
	if !exists {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(json))
}
