package handlers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type H map[string]interface{}

func GetPolls(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "hello world from Carl\n")
	test := `{"items":[{"id":2,"name":"Vue","topic":"Voguish Vue","src":"https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Vue.js_Logo.svg/400px-Vue.js_Logo.svg.png","upvotes":3,"downvotes":1},{"id":3,"name":"React","topic":"Remarkable React","src":"https://upload.wikimedia.org/wikipedia/commons/thumb/a/a7/React-icon.svg/1200px-React-icon.svg.png","upvotes":1,"downvotes":1},{"id":4,"name":"Ember","topic":"Excellent Ember","src":"https://cdn-images-1.medium.com/max/741/1*9oD6P0dEfPYp3Vkk2UTzCg.png","upvotes":1,"downvotes":1},{"id":5,"name":"Knockout","topic":"Knightly Knockout","src":"https://images.g2crowd.com/uploads/product/image/social_landscape/social_landscape_1489710848/knockout-js.png","upvotes":7,"downvotes":0}]}`
	fmt.Println("the get polls function is being called in go")
	io.WriteString(w, test)
}

func UpdatePoll(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "hello world from Carl\n")
	fmt.Println("update is being called")
	//fmt.Println(req.Method, req.Header, req.Host, req.Context)
	//fmt.Println("the context is", req.Context())
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		io.WriteString(w, "there was a problem")
	}

	fmt.Println("the content is", string(body))
	fmt.Printf("the type of the body is %T\n", req.Body)
	fmt.Println("the params passed in are", len(req.PostForm))
	fmt.Println("the host is:", req.Host)
	fmt.Println("the url is:", req.URL)
	fmt.Println("the whatUp is:", req.URL.Query().Get("whatUp"))
	io.WriteString(w, "{someOther:goingDown}")
}

// func GetPolls(db *sql.DB) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, models.GetPolls(db))
// 	}
// }

// func UpdatePoll(db *sql.DB) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var poll models.Poll

// 		c.Bind(&poll)

// 		index, _ := strconv.Atoi(c.Param("index"))

// 		id, err := models.UpdatePoll(db, index, poll.Name, poll.Upvotes, poll.Downvotes)

// 		if err == nil {
// 			return c.JSON(http.StatusCreated, H{
// 				"affected": id,
// 			})
// 		}

// 		return err
// 	}
// }
