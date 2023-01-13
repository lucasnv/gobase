package tools

import (
	"[REPO_URL]/pkg/shared/domain/valueobjects"
)

// NAME
type Name struct {
	value string
}

func NewName(value string) (Name, error) {
	return Name{value: value}, nil
}

// LINK
type Link struct {
	value string
}

func NewLink(value string) (Link, error) {
	return Link{value: value}, nil
}

// DESCRIPTION
type Description struct {
	value string
}

func NewDescription(value string) (Description, error) {
	return Description{value: value}, nil
}

// TOOL
type Tool struct {
	id          valueobjects.Id
	name        Name
	link        Link
	description Description
}
