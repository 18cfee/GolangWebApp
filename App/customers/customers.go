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
	newId := dao.GetHighestCustId()
	mu.Unlock()

	// if error != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([byte](error))
	// 	return
	// }

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(newId)))
}
