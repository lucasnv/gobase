package createtool

import (
	"<MODULE_URL_REPLACE>/internal/commandbus"
)

const COMMMAND_TYPE commandbus.Type = "command.creating.tool"

type Command struct {
	id          string
	name        string
	link        string
	description string
}

func NewCommand(id, name, link, description string) Command {
	return Command{
		id:          id,
		name:        name,
		link:        link,
		description: description,
	}
}

func (c Command) Type() commandbus.Type {
	return COMMMAND_TYPE
}
