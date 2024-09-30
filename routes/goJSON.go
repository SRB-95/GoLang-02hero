package routes

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int
	Color  bool
	Actors []string
}

var movies = []Movie{
	{Title: "Shole", Year: 1980, Color: true, Actors: []string{"Dharamendra", "Amitab Bachan", "Hemamalini"}},
	{Title: "Kuch Kuch Hota Hai", Year: 2000, Color: true, Actors: []string{"Shahrukh Khan", "Kajol", "Rani Mukherji"}},
	{Title: "Bajrani Bhaijan", Year: 2010, Color: true, Actors: []string{"Salman Khan", "Kareena Kapoor", "Nawajudin"}},
	{Title: "Mother India", Year: 1960, Color: false, Actors: []string{"Nargis", "Sunil Dutt"}},
}

func playWithMoviesMarshallingAndUnmarshalling() {
	fmt.Println("\nMovies Array Data...")

	for index, movie := range movies {
		fmt.Println(index, movie)
	}

	jsonData, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON Marshalling Failed: %s", err)
	}

	fmt.Println("\nMarshalled JSON Data...")
	fmt.Printf("%s \n", jsonData)

	var moviesData []Movie
	if err := json.Unmarshal(jsonData, &moviesData); err != nil {
		log.Fatalf("JSON UnMarshalling Failed: %s", err)
	}

	fmt.Println("\nUnmarshalled JSON Data...")
	for index, movie := range moviesData {
		fmt.Println(index, movie)
	}

	// Named Structure
	// type MovieTitle struct {
	// 		Title string
	// }
	// var moviesTitles []MovieTitle

	// Anonymous Structure
	var moviesTitles []struct{ Title string }
	if err := json.Unmarshal(jsonData, &moviesTitles); err != nil {
		log.Fatalf("JSON Marshalling Failed: %s", err)
	}

	fmt.Println("\nUnmarshalled JSON Data... Movies Titles")
	for index, movieTitle := range moviesTitles {
		fmt.Println(index, movieTitle)
	}
}

//________________________________________________________________

type MovieAgain struct {
	Title  string
	Year   int `json:"ReleasedYear"`
	Color  bool
	Actors []string
}

var moviesAgain = []MovieAgain{
	{Title: "Shole", Year: 1980, Color: true, Actors: []string{"Dharamendra", "Amitab Bachan", "Hemamalini"}},
	{Title: "Kuch Kuch Hota Hai", Year: 2000, Color: true, Actors: []string{"Shahrukh Khan", "Kajol", "Rani Mukherji"}},
	{Title: "Bajrani Bhaijan", Year: 2010, Color: true, Actors: []string{"Salman Khan", "Kareena Kapoor", "Nawajudin"}},
	{Title: "Mother India", Year: 1960, Color: false, Actors: []string{"Nargis", "Sunil Dutt"}},
}

func playWithMoviesMarshallingAndUnmarshallingAgain() {
	fmt.Println("\nMovies Array Data...")

	for index, movie := range moviesAgain {
		fmt.Println(index, movie)
	}

	{ // Local Scope
		jsonData, err := json.Marshal(moviesAgain)
		if err != nil {
			log.Fatalf("JSON Marshalling Failed: %s", err)
		}

		fmt.Println("\nMarshalled JSON Data...")
		fmt.Printf("%s \n", jsonData)
	}

	{ // Local Scope
		jsonData, err := json.MarshalIndent(moviesAgain, "", "    ")
		if err != nil {
			log.Fatalf("JSON Marshalling Failed: %s", err)
		}

		fmt.Println("\nMarshalled JSON Data Indented...")
		fmt.Printf("%s \n", jsonData)
	}
}

func playWithJSON() {
	fmt.Println("Function: playWithMoviesMarshallingAndUnmarshalling")
	playWithMoviesMarshallingAndUnmarshalling()
	fmt.Println("--------------------------------------->")

	fmt.Println("Function: playWithMoviesMarshallingAndUnmarshallingAgain")
	playWithMoviesMarshallingAndUnmarshallingAgain()
	fmt.Println("--------------------------------------->")
}
