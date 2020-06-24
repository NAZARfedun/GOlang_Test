package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"

	"github.com/UserDatabaseApi/src/config"
	"github.com/UserDatabaseApi/src/logger"
)

type Player struct {
	Nickname string `json:"Nickname"`
	Score    int    `json:"Score"`
}

type PlayersList struct {
	PlayerList []Player `json:"playerList"`
}

var ListOfPlayers []Player

func main() {
	ListOfPlayers = make([]Player, 0, 10)

	port := ":8080"

	configFilePath := "config.json"
	var (
		Config config.Configuration
		log    logger.Logger
		err    error
	)

	//Create service configuration
	if err, Config = config.Load(configFilePath); err != nil {
		log.Fatal(err)
	}

	//Create service logger
	if err, log = logger.Load(Config.Log); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/PostValues", PostValues)
	http.HandleFunc("/GetValues", GetValues)
	http.HandleFunc("/PostPlayersList", PostPlayersList)

	fmt.Println("Server started at  - ", port)
	http.ListenAndServe(port, nil)
}

func PostValues(w http.ResponseWriter, r *http.Request) {
	var player Player

	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		return
	}

	ListOfPlayers = append(ListOfPlayers, player)
	fmt.Println("**********")
	fmt.Println(player)
	fmt.Fprint(w, player)
}

func PostPlayersList(w http.ResponseWriter, r *http.Request) {
	var playerslist PlayersList

	if err := json.NewDecoder(r.Body).Decode(&playerslist); err != nil {
		return
	}

	fmt.Println("Sorted Players by score")
	sort.Slice(playerslist.PlayerList, func(i, j int) bool { return playerslist.PlayerList[i].Score < playerslist.PlayerList[j].Score })

	fmt.Println("Players' list: ", playerslist)
	fmt.Fprint(w, playerslist)
}

func GetValues(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintln(ListOfPlayers))
}

/*Exersice: Loops and Functions*/
// package main

// import (
// 	"fmt"
// 	"math"
// 	"github.com/Nazar_Test/logger"
// 	"github.com/Nazar_Test/config"
// )

// func Sqrt(x float64) float64 {
// 	z1 := 1.0
// 	z0 := z1

// 	for {

// 		z1 = z1 - ((math.Pow(z1, 2) - x) / (2 * z1))
// 		if (z1 - z0) < 0.001 {
// 			return z1
// 		}
// 		z0 = z1
// 		fmt.Println(z1, "**", z0)

// 	}

// }

// func main() {
// 	logger.Logger()
// 	config.Config()
// 	x := 2.0
// 	fmt.Println(Sqrt(x))
// }
