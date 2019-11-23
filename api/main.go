package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tmcentee/lets-go/api/models"
)

type Env struct {
	db models.Datastore
}

func main() {
	db, err := models.NewDB("postgres://")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/players", env.playerIndex)
	router.HandleFunc("/player/{id}", env.playerSingle)
	router.HandleFunc("/teams", env.teamsIndex)
	router.HandleFunc("/team/{id}", env.teamSingle)

	log.Fatal(http.ListenAndServe(":8888", router))
}

func (env *Env) playerIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	players, err := env.db.AllPlayers()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Println(err)
		return
	}
	for _, player := range players {
		fmt.Fprintf(w, "%d, %s, %s, %d\n", player.ID, player.FirstName, player.LastName, player.Team)
	}
}

func (env *Env) teamsIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	teams, err := env.db.AllTeams()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Println(err)
		return
	}
	for _, team := range teams {
		fmt.Fprintf(w, "%d, %s, %s\n", team.ID, team.Name, team.City)
	}
}

func (env *Env) playerSingle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	vars := mux.Vars(r)
	key := vars["id"]

	playerID, err := strconv.Atoi(key)
	if err != nil {
		fmt.Fprintf(w, "Invalid player ID %s", key)
	}

	player, err := env.db.SinglePlayer(playerID)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Println(err)
		return
	}

	fmt.Fprintf(w, "%d, %s, %s, %d\n", player.ID, player.FirstName, player.LastName, player.Team)
}

func (env *Env) teamSingle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	vars := mux.Vars(r)
	key := vars["id"]

	teamID, err := strconv.Atoi(key)
	if err != nil {
		fmt.Fprintf(w, "Invalid team ID %s", key)
	}

	team, err := env.db.SingleTeam(teamID)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Println(err)
		return
	}

	fmt.Fprintf(w, "%d, %s, %s\n", team.ID, team.Name, team.City)
}
