package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"

	"github.com/go-martini/martini"
)

type Profile struct {
	Handle string
	Magic  Skill
}

type Skill struct {
	Name      string `json:"name"`
	CurrentXP int    `json:"experience"`
	Level     int    `json:"level"`
	Rank      int    `json:"rank"`
	ID        int    `json:"id"`
}

type APISkill struct {
	data []Skill `json:"data"`
}

func getProfileHighscore(handle string) {
	resp, err := http.Get("http://silabsoft.org/rs-web/highscore/tonnu")
	reader := csv.NewReader(resp.Body)
	reader.FieldsPerRecord = -1

	cvsData, err := reader.ReadAll()
	fmt.Printf("%+v\n", cvsData)
	if err != nil {
		log.Fatal(err)
	}

	// var skill Skill
	// var skills []Skill
	for _, each := range cvsData {
		fmt.Println(each[1])
	}

}

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello, test!"
	})

	m.Get("/profiles/:handle", func(params martini.Params) string {
		getProfileHighscore(params["handle"])
		return "Hello "
	})
	m.Run()

}
