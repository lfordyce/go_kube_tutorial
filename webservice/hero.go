package webservice

type Hero struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	RealName string `json:"realName"`
}

type Heros []Hero