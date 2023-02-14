package domain

import vo "<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"

type UserRepository interface {
	Save(u User) error
	Find(id vo.Id) User
}
