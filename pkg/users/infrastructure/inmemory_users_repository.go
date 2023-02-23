package infrastructure

import (
	"context"
	"strings"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/criteria"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	vo "<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
	inmemoryCriteria "<MODULE_URL_REPLACE>/pkg/shared/infrastructure/criteria"
	user "<MODULE_URL_REPLACE>/pkg/users/domain"
)

type InmemoryUsersRepository struct {
	users user.List
}

func NewInmemoryUsersRepository() *InmemoryUsersRepository {
	return &InmemoryUsersRepository{}
}

func (r *InmemoryUsersRepository) Save(ctx *context.Context, u user.User) *errors.AppError {
	r.users = append(r.users, u)

	return nil
}

func (r *InmemoryUsersRepository) Find(ctx *context.Context, searchedId vo.Id) (user.User, *errors.AppError) {

	for _, u := range r.users {
		id := u.Id()

		if id.ToString() == searchedId.ToString() {
			return u, nil
		}
	}

	return user.User{}, user.NewUserError(user.NOT_FOUND_ERROR)
}

func (r *InmemoryUsersRepository) FindByCriteria(ctx *context.Context, c criteria.Criteria) (user.Users, *errors.AppError) {
	var result user.List

	if c == nil {
		return user.NewUsers(r.users), nil
	}

	for _, u := range r.users {

		ok := filterByCriteria(u, c)

		if ok {
			result = append(result, u)
		}
	}

	return user.NewUsers(result), nil
}

func (r *InmemoryUsersRepository) Delete(ctx *context.Context, searchedId vo.Id) *errors.AppError {

	for i, u := range r.users {
		id := u.Id()

		if id.ToString() == searchedId.ToString() {
			r.users = append(r.users[:i], r.users[i+1:]...)
		}
	}

	return nil
}

func filterByCriteria(u user.User, c criteria.Criteria) bool {
	filter := c.Filter()
	switch filter.(string) {

	case "and":
		andCriteria := c.(inmemoryCriteria.AndInmemoryCriteria)
		return filterByCriteria(u, andCriteria.C1) && filterByCriteria(u, andCriteria.C2)

	case "between":
		between := c.(inmemoryCriteria.BetweenInmemoryCriteria)
		value := getValueByField(u, between.Field)
		return value.(int) <= between.V1.(int) && between.V2.(int) <= value.(int)

	case "eq":
		eqCriteria := c.(inmemoryCriteria.EqInmemoryCriteria)
		value := getValueByField(u, eqCriteria.Field)
		return value == eqCriteria.Value

	case "gt":
		gtCriteria := c.(inmemoryCriteria.GtInmemoryCriteria)
		value := getValueByField(u, gtCriteria.Field)
		return value.(int) < gtCriteria.Value.(int)

	case "gte":
		gteCriteria := c.(inmemoryCriteria.GteInmemoryCriteria)
		value := getValueByField(u, gteCriteria.Field)
		return value.(int) <= gteCriteria.Value.(int)

	case "in":
		inCriteria := c.(inmemoryCriteria.InInmemoryCriteria)
		value := getValueByField(u, inCriteria.Field)
		for _, criteriaValue := range inCriteria.Value.([]any) {
			if criteriaValue == value {
				return true
			}
		}
		return false

	case "like":
		likeCriteria := c.(inmemoryCriteria.LikeInmemoryCriteria)
		value := getValueByField(u, likeCriteria.Field)
		if strings.Contains(strings.ToLower(value.(string)), strings.ToLower(likeCriteria.Value)) {
			return true
		}
		return false

	case "lt":
		ltCriteria := c.(inmemoryCriteria.LtInmemoryCriteria)
		value := getValueByField(u, ltCriteria.Field)
		return value.(int) > ltCriteria.Value.(int)

	case "lte":
		lteCriteria := c.(inmemoryCriteria.LteInmemoryCriteria)
		value := getValueByField(u, lteCriteria.Field)
		return value.(int) > lteCriteria.Value.(int)

	case "or":
		orCriteria := c.(inmemoryCriteria.OrInmemoryCriteria)
		return filterByCriteria(u, orCriteria.C1) || filterByCriteria(u, orCriteria.C2)

	case "not":
		notCriteria := c.(inmemoryCriteria.NotInmemoryCriteria)
		return !filterByCriteria(u, notCriteria.Criteria)
	}

	return true
}

func getValueByField(u user.User, field string) any {

	var value string

	switch field {
	case "email":
		value = u.Email().Value

	case "first_name":
		value = u.FirstName().Value

	case "last_name":
		value = u.LastName().Value

	case "created_at":
		createdAt := u.CreatedAt()
		value = createdAt.ToString()
	}

	return value
}

var _ user.UserRepository = (*InmemoryUsersRepository)(nil)
