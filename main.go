package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-martini/martini"
)

type Profile struct {
	Handle string
	Magic  Skill
}

type Skill struct {
	Name      string
	CurrentXP int
	Level     int
}

type APISkill struct {
	data []*Skill `json:"data"`
}

func getProfileHighscore(handle string) []*Skill {
	resp, err := http.Get("http://silabsoft.org/rs-web/highscore/tonnu")

	if err != nil {
		fmt.Println(err)
	}
	var skill APISkill
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	json.Unmarshal(body, &skill)
	var skills []*Skill

	fmt.Println(skill.data)
	if err != nil {
		fmt.Println(err)
	}
	return skills
}

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello, test!"
	})

	m.Get("/profiles/:handle", func(params martini.Params) string {
		skill := getProfileHighscore(params["handle"])
		return "Hello " + skill[0].Name
	})
	m.Run()

}
