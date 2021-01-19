package models

type PlayersResponseBody struct {
	Players []Player `json:"players"`
}

type PlayerStatus struct {
}

type Player struct {
	ID     string       `json:"id"`
	Name   string       `json:"name"`
	Live   bool         `json:"live"`
	Status PlayerStatus `json:"status"`
}
