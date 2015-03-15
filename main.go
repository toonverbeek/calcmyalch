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
	var data map[string][]json.RawMessage
	var skills []Skill
	jsonStr, err := ioutil.ReadAll(resp.Body)
	str := string(jsonStr[26 : len(jsonStr)-1])
	// str = str[26:len(str)]
	fmt.Println(str)
	json.Unmarshal([]byte(str), &skills)
	// jsonStr := []byte(`
	//  {
	//  	"data":[
	//   {
	//     "id":0,
	//     "isSkill":true,
	//     "name":"Overall",
	//     "rank":576774,
	//     "level":1302,
	//     "experience":12477942
	//   },
	//   {
	//     "id":1,
	//     "isSkill":true,
	//     "name":"Attack",
	//     "rank":626083,
	//     "level":70,
	//     "experience":805403
	//   }
	//  ]}`)
	// if err != nil {
	// 	log.Fatal("Error reading All ", err)
	// }
	err = json.Unmarshal(jsonStr, &data)
	if err != nil {
		log.Fatal("Error unmarshalling ", err)
	}

	for _, each := range data["data"] {
		skill := &Skill{}
		if err = json.Unmarshal(each, &skill); err != nil {
			log.Println(err)
		} else {
			if skill != new(Skill) {
				skills = append(skills, *skill)
				fmt.Printf("%+v", skill)
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
