package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type InTheater struct {
	Count    int
	Title    string
	Total    int
	Subjects []Movie
}

type Movie struct {
	Title  string
	Rating Rating
	Casts  []Cast
}

type Rating struct {
	Max     int
	Average float64
	Starts  string
	Min     int
}

type Cast struct {
	Name     string
	DoubanID string `json:"id"`
}

func GetInTheater(city string) {
	res, err := http.Get("https://api.douban.com/v2/movie/in_theaters?city=" + city)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	var data InTheater
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Title", "Rating", "Casts"})

	for _, movie := range data.Subjects {
		castNames := make([]string, len(movie.Casts))
		for i, cast := range movie.Casts {
			castNames[i] = cast.Name
		}

		table.Append([]string{
			movie.Title,
			strconv.FormatFloat(movie.Rating.Average, 'f', 1, 64),
			strings.Join(castNames, ", "),
		})
	}

	table.Render()
}
