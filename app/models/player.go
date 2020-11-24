package models

type PlayersResponseBody struct {
	Players []Player `json:"players"`
}

type Player struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
