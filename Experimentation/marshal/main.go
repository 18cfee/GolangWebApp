package main

import (
	"encoding/json"
	"fmt"
)

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
	Rando  int      `json:"rando"`
}

func main() {
	str := `{"page": 1, "fruits": ["apple", "peach", 1], "rando": 2}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
	fmt.Println(str)
}
