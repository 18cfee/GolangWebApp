package handlers

import (
	"fmt"
	"io"
	"net/http"
)

type H map[string]interface{}

func GetPolls(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "hello world from Carl\n")
	test := `{"items":[{"id":1,"name":"Angular","topic":"Awesome Angular","src":"https://cdn.colorlib.com/wp/wp-content/uploads/sites/2/angular-logo.png","upvotes":3,"downvotes":3},{"id":2,"name":"Vue","topic":"Voguish Vue","src":"https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Vue.js_Logo.svg/400px-Vue.js_Logo.svg.png","upvotes":3,"downvotes":1},{"id":3,"name":"React","topic":"Remarkable React","src":"https://upload.wikimedia.org/wikipedia/commons/thumb/a/a7/React-icon.svg/1200px-React-icon.svg.png","upvotes":1,"downvotes":1},{"id":4,"name":"Ember","topic":"Excellent Ember","src":"https://cdn-images-1.medium.com/max/741/1*9oD6P0dEfPYp3Vkk2UTzCg.png","upvotes":1,"downvotes":1},{"id":5,"name":"Knockout","topic":"Knightly Knockout","src":"https://images.g2crowd.com/uploads/product/image/social_landscape/social_landscape_1489710848/knockout-js.png","upvotes":7,"downvotes":0}]}`
	fmt.Println("the get polls function is being called in go")
	io.WriteString(w, test)
}

func UpdatePoll(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "hello world from Carl\n")
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
