package game

type PlayerProcedures int

func (p *PlayerProcedures) Start(playerName string, result *bool) error {
	*result = true
	return nil
}
