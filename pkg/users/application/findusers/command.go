package findusers

import (
	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
)

const COMMAND_TYPE commandbus.Type = "command.finding.users"

type Command struct {
	filter    string
	orderBy   string
	orderSort string
	page      int
	perPage   int
}

func NewCommand(f string, ob string, os string, p int, pp int) commandbus.Command {
	return Command{
		filter:    f,
		orderBy:   ob,
		orderSort: os,
		page:      p,
		perPage:   pp,
	}
}

func (c Command) Type() commandbus.Type {
	return COMMAND_TYPE
}

var _ commandbus.Command = (*Command)(nil)
