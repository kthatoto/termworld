package models

type PlayersResponseBody struct {
	Players []Player `json:"players"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PlayerStatus struct {
	MaxHP    int      `json:"maxHP"`
	HP       int      `json:"HP"`
	Position Position `json:"position"`
}

type Player struct {
	ID     string       `json:"id"`
	Name   string       `json:"name"`
	Live   bool         `json:"live"`
	Status PlayerStatus `json:"status"`
}
