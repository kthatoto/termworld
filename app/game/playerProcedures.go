package game

import (
	"fmt"
)

type PlayerProcedures int
type PlayerProcedureArgs struct {
	PlayerName string
	Options []string
}

func (p *PlayerProcedures) Start(args PlayerProcedureArgs, result *bool) error {
	fmt.Printf("%+v", args)
	*result = true
	return nil
}
