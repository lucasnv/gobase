package infrastructure

import (
	"context"
	"sort"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/collection"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/criteria"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	vo "<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
	criteriaAdapter "<MODULE_URL_REPLACE>/pkg/shared/infrastructure/criteria"
	user "<MODULE_URL_REPLACE>/pkg/users/domain"
)

type InmemoryUsersRepository struct {
	users           user.List
	criteriaBuilder criteriaAdapter.InMemoryCriteriaBuilderAdapter
	fieldsTypes     map[string]string
}

func NewInmemoryUsersRepository() *InmemoryUsersRepository {
	return &InmemoryUsersRepository{
		criteriaBuilder: criteriaAdapter.NewInmemoryCriteriaBuilderAdapter(),
		fieldsTypes: map[string]string{
			"first_name": "string",
			"last_name":  "string",
			"email":      "string",
			"created_at": "time",
		},
	}
}

func (r *InmemoryUsersRepository) Save(ctx *context.Context, newUser user.User) *errors.AppError {

	for i, u := range r.users {
		id := u.GetId()

		if id.ToString() == newUser.GetId().Value.String() {
			r.users[i] = newUser
			return nil
		}
	}

	r.users = append(r.users, newUser)

	return nil
}

func (r *InmemoryUsersRepository) Find(ctx *context.Context, searchedId vo.Id) (user.User, *errors.AppError) {

	for _, u := range r.users {
		id := u.GetId()

		if id.ToString() == searchedId.ToString() {
			return u, nil
		}
	}

	return user.User{}, user.NewUserError(user.NOT_FOUND_ERROR)
}

func (r *InmemoryUsersRepository) FindByCriteria(ctx *context.Context, f criteria.Criteria, o criteria.SorterCriteria, p criteria.PaginatorCriteria) (collection.Collection, *errors.AppError) {
	var result user.List

	if f != nil {
		for _, u := range r.users {

			ok := r.filterByCriteria(u, f)

			if ok {
				result = append(result, u)
			}
		}
	} else {
		result = r.users
	}

	totalUsers := uint32(len(result) & 0xffffffff)
	result = r.sortUsers(result, o)
	result = r.paginateUsers(result, p)

	return collection.NewCollection(result, p.Page(), p.PageSize(), totalUsers), nil
}

func (r *InmemoryUsersRepository) Delete(ctx *context.Context, searchedId vo.Id) *errors.AppError {

	for i, u := range r.users {
		id := u.GetId()

		if id.ToString() == searchedId.ToString() {
			r.users = append(r.users[:i], r.users[i+1:]...)
		}
	}

	return nil
}

var _ user.UserRepository = (*InmemoryUsersRepository)(nil)

func (r *InmemoryUsersRepository) filterByCriteria(u user.User, c criteria.Criteria) bool {

	switch c.Type() {
	case "and":
		andCriteria := c.(criteria.AndCriteria)
		return r.filterByCriteria(u, andCriteria.C1) && r.filterByCriteria(u, andCriteria.C2)

	case "between":
		betweenCriteria := c.(criteria.BetweenCriteria)
		entityValue := getValueByField(u, betweenCriteria.Field)
		return r.criteriaBuilder.Between(betweenCriteria, r.fieldsTypes, entityValue)

	case "eq":
		eqCriteria := c.(criteria.EqCriteria)
		entityValue := getValueByField(u, eqCriteria.Field)
		return r.criteriaBuilder.Eq(eqCriteria, r.fieldsTypes, entityValue)

	case "gt":
		gtCriteria := c.(criteria.GtCriteria)
		entityValue := getValueByField(u, gtCriteria.Field)
		return r.criteriaBuilder.Gt(gtCriteria, r.fieldsTypes, entityValue)

	case "gte":
		gteCriteria := c.(criteria.GteCriteria)
		entityValue := getValueByField(u, gteCriteria.Field)
		return r.criteriaBuilder.Gte(gteCriteria, r.fieldsTypes, entityValue)

	case "in":
		inCriteria := c.(criteria.InCriteria)
		entityValue := getValueByField(u, inCriteria.Field)
		return r.criteriaBuilder.In(inCriteria, r.fieldsTypes, entityValue)

	case "like":
		likeCriteria := c.(criteria.LikeCriteria)
		entityValue := getValueByField(u, likeCriteria.Field)
		return r.criteriaBuilder.Like(likeCriteria, r.fieldsTypes, entityValue)

	case "lt":
		ltCriteria := c.(criteria.LtCriteria)
		entityValue := getValueByField(u, ltCriteria.Field)
		return r.criteriaBuilder.Lt(ltCriteria, r.fieldsTypes, entityValue)

	case "lte":
		lteCriteria := c.(criteria.LteCriteria)
		entityValue := getValueByField(u, lteCriteria.Field)
		return r.criteriaBuilder.Lte(lteCriteria, r.fieldsTypes, entityValue)

	case "or":
		orCriteria := c.(criteria.OrCriteria)
		return r.filterByCriteria(u, orCriteria.C1) || r.filterByCriteria(u, orCriteria.C2)

	case "not":
		notCriteria := c.(criteria.NotCriteria)
		return !r.filterByCriteria(u, notCriteria.Criteria)
	}

	return false
}

func (r *InmemoryUsersRepository) sortUsers(u user.List, o criteria.SorterCriteria) user.List {
	sort.Slice(u, func(i, j int) bool {
		v1 := getValueByField(u[i], o.By())
		v2 := getValueByField(u[j], o.By())

		return r.criteriaBuilder.Sort(o, r.fieldsTypes, v1, v2)
	})

	return u
}

func (r *InmemoryUsersRepository) paginateUsers(u user.List, c criteria.PaginatorCriteria) user.List {
	start := c.Offset()
	end := c.Offset() + c.Limit()

	if start >= uint32(len(u)&0xffffffff) {
		return user.List{}
	}

	if end > uint32(len(u)&0xffffffff) {
		end = uint32(len(u) & 0xffffffff)
	}

	return u[start:end]
}

func getValueByField(u user.User, field string) any {
	switch field {
	case "created_at":
		return u.GetCreatedAt().Value
	case "first_name":
		return u.GetFirstName().Value
	case "last_name":
		return u.GetLastName().Value
	case "email":
		return u.GetEmail().Value
	}

	return nil
}
