package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"kankan", "xyan"}}

	//编码的过程 结构体-->json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	//解码的过程 json-->结构体
	jsonVar := `{"title":"喜剧之王","year":2000,"price":10,"actors":["kankan","xyan"]}`
	myMovie := Movie{}
	err = json.Unmarshal([]byte(jsonVar), &myMovie)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", myMovie)

}
