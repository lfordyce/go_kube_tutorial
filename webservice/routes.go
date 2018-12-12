package webservice

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"HeroIndex",
		"GET",
		"/heros",
		HeroIndex,
	},
	Route{
		"HeroShow",
		"GET",
		"/heros/{heroId}",
		HeroShow,
	},
	Route{
		"HeroCreate",
		"POST",
		"/heros",
		HeroCreate,
	},
}