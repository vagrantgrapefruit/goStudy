package reflectStudy

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func TagInJson() {
	fmt.Println("tagInJson Start")

	movie := Movie{
		Title:  "The Shawshank Redemption",
		Year:   1994,
		Price:  29,
		Actors: []string{"Tim Robbins", "Morgan Freeman", "Bob Gunton"},
	}
	fmt.Printf("movie = %v\n", movie)
	// 序列化
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json.Marshal error", err)
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)
	// 反序列化
	var movie2 Movie
	err = json.Unmarshal(jsonStr, &movie2)
	if err != nil {
		fmt.Println("json.Unmarshal error", err)
	}
	fmt.Printf("movie2 = %v\n", movie2)
	fmt.Println("tagInJson End")
}

func init() {
	fmt.Println("tagInJson init")
}
