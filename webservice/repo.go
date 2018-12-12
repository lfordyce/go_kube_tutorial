package webservice

import "fmt"

var currentId int

var heros Heros

func init() {
	RepoCreateHero(Hero{Name: "Batman", RealName: "Bruce Wayne", Id: 0})
	RepoCreateHero(Hero{Name: "Iron Man", RealName: "Tony Stark", Id: 1})
}

func RepoFindHero(id int) Hero {
	for _, h := range heros {
		if h.Id == id {
			return h
		}
	}
	return Hero{}
}

func RepoCreateHero(h Hero) Hero {
	currentId += 1
	h.Id = currentId
	heros = append(heros, h)
	return h
}

func ReposDestroyHero(id int) error {
	for i, h := range heros {
		if h.Id == id {
			heros = append(heros[:i], heros[i + 1:]...)
			return nil
		}
 	}
	return fmt.Errorf("could not find Hero with id of %d to delete", id)
}
