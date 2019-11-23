package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tmcentee/lets-go/api/models"
)

func main() {
	models.InitDB("")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/players", playerIndex)
	router.HandleFunc("/player/{id}", playerSingle)
	log.Fatal(http.ListenAndServe(":8888", router))
}

func playerIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	players, err := models.AllPlayers()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Println(err)
		return
	}
	for _, players := range players {
		fmt.Fprintf(w, "%d, %s, %s, %d\n", players.ID, players.FirstName, players.LastName, players.Team)
	}
}

func playerSingle(w http.ResponseWriter, r *http.Request) {
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

	player, err := models.SinglePlayer(playerID)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Println(err)
		return
	}

	fmt.Fprintf(w, "%d, %s, %s, %d\n", player.ID, player.FirstName, player.LastName, player.Team)
}
