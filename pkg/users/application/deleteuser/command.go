package deleteuser

import (
	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
)

const COMMAND_TYPE commandbus.Type = "command.delete.user"

type Command struct {
	id string
}

func NewCommand(id string) commandbus.Command {
	return Command{
		id: id,
	}
}

func (c Command) Type() commandbus.Type {
	return COMMAND_TYPE
}

var _ commandbus.Command = (*Command)(nil)
