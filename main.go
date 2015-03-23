package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-martini/martini"
	"github.com/melvinmt/firebase"
)

type Profile struct {
	Handle string
	Magic  Skill
}

type Skill struct {
	timeStamp time.Time
	name      string `json:"name"`
	currentXP int    `json:"experience"`
	level     int    `json:"level"`
	rank      int    `json:"rank"`
}

func getProfileHighscore(handle string) *Skill {
	url := fmt.Sprintf("http://services.runescape.com/m=hiscore/index_lite.ws?player=%s", handle)
	resp, err := http.Get(url)
	reader := csv.NewReader(resp.Body)
	reader.FieldsPerRecord = -1
	cvsData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	skills := csvToSkill(cvsData)
	return skills[0]

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
					name:      "Magic",
					currentXP: xp,
					level:     level,
					rank:      rank,
				}
				skills = append(skills, skill)
			}
		}
	}

	return skills

}

func writeToFirebase(skill *Skill) {
	var err error
	skill.timeStamp = time.Now()

	authtoken := "n9h6qpBDPgt6q8XKvwunbNVTBdM4RSYxqfdTjlXx"
	// day := skill.timeStamp.Day()
	// month := skill.timeStamp.Month()
	// year := skill.timeStamp.Year()
	// url := fmt.Sprintf("https://calchmyalch.firebaseio.com/users/tonnu/%d-%d-%d/magic", day, month, year)
	url := "https://calchmyalch.firebaseio.com/users/tonnu/magic"
	fmt.Println(url)
	ref := firebase.NewReference(url).Auth(authtoken)

	if err = ref.Write(skill); err != nil {
		log.Fatal(err)
	}
	fmt.Println("looks like it worked.")

}

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello, test!"
	})

	m.Get("/profiles/:handle", func(params martini.Params) string {
		skill := getProfileHighscore(params["handle"])
		writeToFirebase(skill)
		return fmt.Sprintf("Skill:%s\nLevel:%d\nExperience:%d\nRank:%d\n", skill.name, skill.level, skill.currentXP, skill.rank)
	})
	m.Run()

}
