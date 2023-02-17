package domain

type UserList []User

type Metadata struct {
	page       int
	pageSize   int
	totalPages int
	total      int
}

type Users struct {
	metadata Metadata
	data     UserList
}

func (u *Users) Data() UserList {
	return u.data
}

func newUsers(u UserList) *Users {
	return &Users{
		data: u,
	}
}
