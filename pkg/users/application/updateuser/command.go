package updateuser

import (
	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
)

const COMMAND_TYPE commandbus.Type = "command.updating.user"

type Command struct {
	id        string
	firstName string
	lastName  string
	email     string
}

func NewCommand(id string, firstName string, lastName string, email string) commandbus.Command {
	return Command{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		email:     email,
	}
}

func (c Command) Type() commandbus.Type {
	return COMMAND_TYPE
}

var _ commandbus.Command = (*Command)(nil)
