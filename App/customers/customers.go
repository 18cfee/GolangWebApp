package customers

import (
	"GolangWebApp/App/dao"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex

func Create(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		fmt.Println("the method was ", req.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	mu.Lock()
	highID, err := dao.GetHighestCustID()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	newID := highID + 1
	err = dao.CreateNewCust(newID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	mu.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(newID)))
}

func UpdateCustomer(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		fmt.Println("the method was ", req.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	data, err := ioutil.ReadAll(req.Body)
	var cust Customer
	json.Unmarshal(data, &cust)
	dao.UpdateCustomer(cust)
	w.WriteHeader(http.StatusOK)
}

func GetById(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		fmt.Println("the method was ", req.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id, ok := req.URL.Query()["custId"]

	var custID int

	if !ok {
		fmt.Println("there was no id")
	} else {
		var err error
		custID, err = strconv.Atoi(id[0])
		if err != nil {
			fmt.Println(err)
		}
	}

	limitA, ok2 := req.URL.Query()["limit"]

	var limit int

	if !ok2 {
		fmt.Println("there was no id")
	} else {
		var err error
		limit, err = strconv.Atoi(limitA[0])
		if err != nil {
			fmt.Println(err)
		}
	}

	customers, err := dao.RetrieveCustomers(custID, int64(limit))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)

	custs, _ := json.Marshal(customers)
	w.Write([]byte(custs))
}
