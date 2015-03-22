package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
}

func getProfileHighscore(handle string) {
	resp, err := http.Get("http://services.runescape.com/m=hiscore/index_lite.ws?player=tonnu")
	reader := csv.NewReader(resp.Body)
	reader.FieldsPerRecord = -1
	cvsData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	csvToSkill(cvsData)

}

func csvToSkill(cvsData [][]string) {

	var skills []*Skill

	for x, each := range cvsData {
		if len(each) == 3 {
			if each[1] != "-1" && each[2] != "-1" {
				if x == 7 {
					rank, _ := strconv.Atoi(each[0])
					level, _ := strconv.Atoi(each[1])
					xp, _ := strconv.Atoi(each[2])
					skill := &Skill{
						Name:      "magic",
						CurrentXP: xp,
						Level:     level,
						Rank:      rank,
					}
					skills = append(skills, skill)
					fmt.Println(each[1], each[2])
				}
			}
		}
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
