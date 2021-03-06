package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve MacQueen", "Jacqueline Bisset"}},
}

func main() {
	//data, err := json.Marshal(movies)

	// a more "people friendly" version
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		fmt.Printf("JSON marshalling failed: %s\n", err)
	}

	fmt.Printf("%s\n", data)
}
