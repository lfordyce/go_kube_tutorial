package webservice

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func HeroIndex(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(heros); err != nil {
		panic(err)
	}
}

func HeroShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var heroId int
	var err error

	if heroId, err = strconv.Atoi(vars["heroId"]); err != nil {
		panic(err)
	}

	hero := RepoFindHero(heroId)
	if hero.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset-UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(hero); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

/*
Test with this curl command:
curl -H "Content-Type: application/json" -d '{"name":"New Hero"}' http://localhost:8080/heros
*/
func HeroCreate(w http.ResponseWriter, r *http.Request) {
	var hero Hero

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &hero); err != nil {
		panic(err)
	}

	h := RepoCreateHero(hero)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(h); err != nil {
		panic(err)
	}
}