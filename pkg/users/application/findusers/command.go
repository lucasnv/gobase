package findusers

import (
	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
)

const COMMAND_TYPE commandbus.Type = "command.finding.users"

type Command struct {
	filter    string
	orderBy   string
	orderSort string
	page      uint32
	perPage   uint32
}

func NewCommand(f string, ob string, os string, p uint32, pp uint32) commandbus.Command {
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
