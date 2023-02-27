package infrastructure

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

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

func (r *InmemoryUsersRepository) FindByCriteria(ctx *context.Context, f criteria.Criteria, o criteria.SortCriteria, p criteria.PaginatorCriteria) (user.Users, *errors.AppError) {
	var result user.List

	if f != nil {
		for _, u := range r.users {

			ok := filterByCriteria(u, f)

			if ok {
				result = append(result, u)
			}
		}
	} else {
		result = r.users
	}

	result = sortUsers(result, o)
	result = paginateUsers(result, p)

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
		betweenCriteria := c.(inmemoryCriteria.BetweenInmemoryCriteria)

		if betweenCriteria.Field == "created_at" {
			dateValue := getDateValueByField(u, betweenCriteria.Field).Truncate(time.Minute)
			dateToCompare1, err := time.Parse(time.RFC3339, betweenCriteria.V1.(string))
			if err != nil {
				fmt.Println(err)
			}
			dateToCompare1 = dateToCompare1.Truncate(time.Minute)
			dateToCompare2, _ := time.Parse(time.RFC3339, betweenCriteria.V2.(string))
			dateToCompare2 = dateToCompare2.Truncate(time.Minute)

			return (dateValue.After(dateToCompare1) || dateValue.Equal(dateToCompare1)) && (dateValue.Before(dateToCompare2) || dateValue.Equal(dateToCompare2))
		}

		return false

	case "eq":
		eqCriteria := c.(inmemoryCriteria.EqInmemoryCriteria)

		if eqCriteria.Field == "created_at" {
			dateValue := getDateValueByField(u, eqCriteria.Field).Truncate(time.Minute)
			dateToCompare, _ := time.Parse(time.RFC3339, eqCriteria.Value.(string))
			dateToCompare = dateToCompare.Truncate(time.Minute)

			return dateValue.Equal(dateToCompare)
		}

		value := getStringValueByField(u, eqCriteria.Field)
		return value == eqCriteria.Value

	case "gt":
		gtCriteria := c.(inmemoryCriteria.GtInmemoryCriteria)

		if gtCriteria.Field == "created_at" {
			dateValue := getDateValueByField(u, gtCriteria.Field).Truncate(time.Minute)
			dateToCompare, _ := time.Parse(time.RFC3339, gtCriteria.Value.(string))
			dateToCompare = dateToCompare.Truncate(time.Minute)

			return dateValue.After(dateToCompare)
		}

		return false

	case "gte":
		gteCriteria := c.(inmemoryCriteria.GteInmemoryCriteria)

		if gteCriteria.Field == "created_at" {
			dateValue := getDateValueByField(u, gteCriteria.Field).Truncate(time.Minute)
			dateToCompare, _ := time.Parse(time.RFC3339, gteCriteria.Value.(string))
			dateToCompare = dateToCompare.Truncate(time.Minute)

			return dateValue.After(dateToCompare) || dateValue.Equal(dateToCompare)
		}

		return false

	case "in":
		inCriteria := c.(inmemoryCriteria.InInmemoryCriteria)
		value := getStringValueByField(u, inCriteria.Field)
		for _, criteriaValue := range inCriteria.Value.([]any) {
			if criteriaValue == value {
				return true
			}
		}
		return false

	case "like":
		likeCriteria := c.(inmemoryCriteria.LikeInmemoryCriteria)
		value := getStringValueByField(u, likeCriteria.Field)
		if strings.Contains(strings.ToLower(value), strings.ToLower(likeCriteria.Value)) {
			return true
		}
		return false

	case "lt":
		ltCriteria := c.(inmemoryCriteria.LtInmemoryCriteria)

		if ltCriteria.Field == "created_at" {
			dateValue := getDateValueByField(u, ltCriteria.Field).Truncate(time.Minute)
			dateToCompare, _ := time.Parse(time.RFC3339, ltCriteria.Value.(string))
			dateToCompare = dateToCompare.Truncate(time.Minute)

			return dateValue.Before(dateToCompare)
		}

		return false

	case "lte":
		lteCriteria := c.(inmemoryCriteria.LteInmemoryCriteria)

		if lteCriteria.Field == "created_at" {
			var dateValue time.Time = getDateValueByField(u, lteCriteria.Field)
			dateValue = dateValue.Truncate(time.Minute)
			dateToCompare, _ := time.Parse(time.RFC3339, lteCriteria.Value.(string))
			dateToCompare = dateToCompare.Truncate(time.Minute)

			return dateValue.Before(dateToCompare) || dateValue.Equal(dateToCompare)
		}

		return false

	case "or":
		orCriteria := c.(inmemoryCriteria.OrInmemoryCriteria)
		return filterByCriteria(u, orCriteria.C1) || filterByCriteria(u, orCriteria.C2)

	case "not":
		notCriteria := c.(inmemoryCriteria.NotInmemoryCriteria)
		return !filterByCriteria(u, notCriteria.Criteria)
	}

	return true
}

func getStringValueByField(u user.User, field string) string {
	var value string

	switch field {
	case "email":
		value = u.Email().Value

	case "first_name":
		value = u.FirstName().Value

	case "last_name":
		value = u.LastName().Value
	}

	return value
}

func getDateValueByField(u user.User, field string) time.Time {
	var value time.Time

	switch field {
	case "created_at":
		value = u.CreatedAt().Value
	}

	return value
}

func sortUsers(u user.List, o criteria.SortCriteria) user.List {
	sort.Slice(u, func(i, j int) bool {

		switch o.By() {
		case "created_at":
			v1 := getDateValueByField(u[i], o.By())
			v2 := getDateValueByField(u[j], o.By())

			if o.Sort() == "asc" {
				return v1.Before(v2)
			}

			return v1.After(v2)

		default:
			v1 := getStringValueByField(u[i], o.By())
			v2 := getStringValueByField(u[j], o.By())

			if o.Sort() == "asc" {
				return v1 < v2
			}

			return v1 > v2
		}

	})

	return u
}

func paginateUsers(u user.List, c criteria.PaginatorCriteria) user.List {
	return u
}

var _ user.UserRepository = (*InmemoryUsersRepository)(nil)
