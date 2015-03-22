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
	url := fmt.Sprintf("http://services.runescape.com/m=hiscore/index_lite.ws?player=%s", handle)
	resp, err := http.Get(url)
	reader := csv.NewReader(resp.Body)
	reader.FieldsPerRecord = -1
	cvsData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	skills := csvToSkill(cvsData)
	fmt.Printf("%s\nlevel:%d\nXP:%d\n", skills[0].Name, skills[0].Level, skills[0].CurrentXP)

}

func csvToSkill(cvsData [][]string) []*Skill {
	var skills []*Skill

	for x, each := range cvsData {
		// index 7 = magic
		if len(each) == 3 && x == 7 {
			if each[1] != "-1" && each[2] != "-1" {
				rank, _ := strconv.Atoi(each[0])
				level, _ := strconv.Atoi(each[1])
				xp, _ := strconv.Atoi(each[2])
				skill := &Skill{
					Name:      "Magic",
					CurrentXP: xp,
					Level:     level,
					Rank:      rank,
				}
				skills = append(skills, skill)
			}
		}
	}

	return skills

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
