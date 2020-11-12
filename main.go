package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Game struct {
	Id     string
	Board  string
	Status string
}

// Global array for storing several games
var games []Game


// Placeholder for homepage
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}


// Get random key from map

func getGameID(r *http.Request) string {
	params := mux.Vars(r)
	var key string
	for _, v := range params {
		key = v
	}
	return key
}

// Gets all games from global array, and lists them

func getGames(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	for i := 0; i < len(games); i++ {
		json.NewEncoder(w).Encode(games[i])
	}
	fmt.Fprintf(w, "getGames")
	fmt.Println("getGames")
}

// Finds a game from global array by UUID, and lists its information

func getGame(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	key := getGameID(r)
	for i := 0; i < len(games); i++ {
		if games[i].Id == key {
			json.NewEncoder(w).Encode(games[i])
		}
		fmt.Println("game_id:" + games[i].Id)
	}

	fmt.Fprintf(w, "getGame")
	fmt.Println("getGame")
}

// Unfinished code to create a new game and store it

func startGame(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var msg Game
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	gameID := uuid.New().String()
	fmt.Println(msg.Board)
	json.NewEncoder(w).Encode(gameID)
	games = append(games, Game{Id: gameID, Board: "----X----", Status: "RUNNING"})
	fmt.Fprintf(w, "startGame")
	fmt.Println("startGame")
}

// Unfinished code to take a move and update the correspoding game

func newMove(w http.ResponseWriter, r *http.Request) {


	w.Header().Set("Content-Type", "application/json")

	key := getGameID(r)

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var msg Game
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	for i := 0; i < len(games); i++ {
		if games[i].Id == key {
			json.NewEncoder(w).Encode(games[i])
			games[i].Board = msg.Board
		}
		fmt.Println("game_id:" + games[i].Id)
	}
	fmt.Println(msg.Board)
	fmt.Fprintf(w, "newMove")
	fmt.Println("newMove")
}

/*
**	https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
 */

func remove(s []Game, i int) []Game {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// Finds a game from global array by UUID, and deletes it

func deleteGame(w http.ResponseWriter, r *http.Request) {

	key := getGameID(r)

	for i := 0; i < len(games); i++ {
		if games[i].Id == key {
			json.NewEncoder(w).Encode(games[i])
			games = remove(games, i)
			break
		}
		fmt.Println("game_id:" + games[i].Id)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "deleteGames")
	fmt.Println("deleteGames")
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/games", getGames).Methods("GET")
	router.HandleFunc("/api/v1/games/{game_id}", getGame).Methods("GET")
	router.HandleFunc("/api/v1/games", startGame).Methods("POST")
	router.HandleFunc("/api/v1/games/{game_id}", newMove).Methods("PUT")
	router.HandleFunc("/api/v1/games/{game_id}", deleteGame).Methods("DELETE")
	router.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", router))
}
