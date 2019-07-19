package customers

import (
	"GolangWebApp/App/dao"
	"fmt"
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
	highID, err := dao.GetHighestCustId()
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
