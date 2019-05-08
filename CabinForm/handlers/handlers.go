package handlers

import (
	"fmt"
	"io"
	"net/http"
)

type H map[string]interface{}

func GetPolls(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "hello world from Carl\n")
	fmt.Println("the get polls function is being called in go")
	io.WriteString(w, "{sample:json}")
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
