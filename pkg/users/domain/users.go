package domain

type List []User

type Metadata struct {
	page       int
	pageSize   int
	totalPages int
	total      int
}

type Users struct {
	metadata Metadata
	data     List
}

func (u *Users) Data() List {
	return u.data
}

func NewUsers(u List) Users {
	return Users{
		data: u,
	}
}
