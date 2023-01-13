package registeruser

import (
	"<MODULE_URL_REPLACE>/internal/commandbus"
)

const COMMMAND_TYPE commandbus.Type = "command.registering.user"

type Command struct {
	firstName string
	lastName  string
	email     string
	password  string
}

func NewCommand(firstName, lastName, email, password string) Command {
	return Command{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		password:  password,
	}
}

func (c Command) Type() commandbus.Type {
	return COMMMAND_TYPE
}
