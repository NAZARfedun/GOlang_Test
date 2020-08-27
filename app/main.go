package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"

	config "github.com/Nazar_Test/app/config"
	"github.com/Nazar_Test/app/logger"
)

type Player struct {
	Nickname string `json:"Nickname"`
	Score    int    `json:"Score"`
}

type PlayersList struct {
	Player []Player `json:"playerList"`
}

type Handlers struct {
	Logger logger.Logger
}

var ListOfPlayers []Player

func main() {
	ListOfPlayers = make([]Player, 0, 10)

	configFilePath := "./config/config.json"
	var (
		Config config.Configuration
		log    logger.Logger
		err    error
	)

	//Create service configuration
	if err, Config = config.Load(configFilePath); err != nil {
		log.Fatal(err)
	}
	fmt.Println(Config)

	//Create service logger
	if err, log = logger.Load(Config.Log); err != nil {
		log.Fatal(err)
	}

	handlers := Handlers{log}

	http.HandleFunc("/PostValue", handlers.PostValue)
	http.HandleFunc("/GetValues", handlers.GetValues)
	http.HandleFunc("/PostPlayersList", handlers.PostPlayersList)

	fmt.Println("Server started at  - ", Config.ListenPort)
	http.ListenAndServe(Config.ListenPort, nil)
}

func (h Handlers) PostValue(w http.ResponseWriter, r *http.Request) {
	var player Player

	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		h.Logger.Errorln(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	ListOfPlayers = append(ListOfPlayers, player)
	//fmt.Println("**********")
	fmt.Println(player)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, player)
}

func (h Handlers) PostPlayersList(w http.ResponseWriter, r *http.Request) {
	var playerslist PlayersList

	if err := json.NewDecoder(r.Body).Decode(&playerslist); err != nil {
		h.Logger.Errorln(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println("Sorted Players by score")
	sort.Slice(playerslist.Player, func(i, j int) bool { return playerslist.Player[i].Score < playerslist.Player[j].Score })

	w.WriteHeader(http.StatusCreated)
	fmt.Println("Players' list: ", playerslist)
	fmt.Fprint(w, playerslist)
}

func (h Handlers) GetValues(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintln(ListOfPlayers))
}
