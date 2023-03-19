package criteria

import (
	"strings"
	"time"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/criteria"
)

type InMemoryCriteriaBuilderAdapter struct {
}

func NewInmemoryCriteriaBuilderAdapter() InMemoryCriteriaBuilderAdapter {
	return InMemoryCriteriaBuilderAdapter{}
}

func (a InMemoryCriteriaBuilderAdapter) Between(c criteria.BetweenCriteria, fieldsType map[string]string, entityValue any) bool {

	fieldType := fieldsType[c.Field]

	if entityValue == nil {
		return false
	}

	switch fieldType {
	case "time":
		entityValue := entityValue.(time.Time).Truncate(time.Minute)
		criteriaValue1, _ := time.Parse(time.RFC3339, c.V1.(string))
		criteriaValue1 = criteriaValue1.Truncate(time.Minute)
		criteriaValue2, _ := time.Parse(time.RFC3339, c.V1.(string))
		criteriaValue2 = criteriaValue2.Truncate(time.Minute)

		return (entityValue.After(criteriaValue1) || entityValue.Equal(criteriaValue1)) && (entityValue.Before(criteriaValue2) || entityValue.Equal(criteriaValue2))
	}

	return false
}

func (a InMemoryCriteriaBuilderAdapter) Eq(c criteria.EqCriteria, fieldsType map[string]string, entityValue any) bool {
	fieldType := fieldsType[c.Field]

	if entityValue == nil {
		return false
	}

	switch fieldType {
	case "int":
		return entityValue == c.Value

	case "string":
		return strings.ToLower(entityValue.(string)) == strings.ToLower(c.Value.(string))

	case "time":
		entityValue := entityValue.(time.Time).Truncate(time.Minute)
		criteriaValue, _ := time.Parse(time.RFC3339, c.Value.(string))
		criteriaValue = criteriaValue.Truncate(time.Minute)
		return entityValue.Equal(criteriaValue)
	}

	return false
}

func (a InMemoryCriteriaBuilderAdapter) Gt(c criteria.GtCriteria, fieldsType map[string]string, entityValue any) bool {

	fieldType := fieldsType[c.Field]

	if entityValue == nil {
		return false
	}

	switch fieldType {
	case "int":
		return entityValue.(int64) < c.Value.(int64)

	case "string":
		return entityValue.(string) < c.Value.(string)

	case "time":
		entityValue := entityValue.(time.Time).Truncate(time.Minute)
		criteriaValue, _ := time.Parse(time.RFC3339, c.Value.(string))
		criteriaValue = criteriaValue.Truncate(time.Minute)
		return entityValue.After(criteriaValue)
	}

	return false
}

func (a InMemoryCriteriaBuilderAdapter) Gte(c criteria.GteCriteria, fieldsType map[string]string, entityValue any) bool {

	fieldType := fieldsType[c.Field]

	if entityValue == nil {
		return false
	}

	switch fieldType {
	case "int":
		return entityValue.(int64) <= c.Value.(int64)

	case "string":
		return strings.ToLower(entityValue.(string)) <= strings.ToLower(c.Value.(string))

	case "time":
		entityValue := entityValue.(time.Time).Truncate(time.Minute)
		criteriaValue, _ := time.Parse(time.RFC3339, c.Value.(string))
		criteriaValue = criteriaValue.Truncate(time.Minute)
		return entityValue.After(criteriaValue) || entityValue.Equal(criteriaValue)
	}

	return false
}

func (a InMemoryCriteriaBuilderAdapter) In(c criteria.InCriteria, fieldsType map[string]string, entityValue any) bool {
	fieldType := fieldsType[c.Field]

	for _, criteriaValue := range c.Value.([]any) {
		switch fieldType {
		case "int":
			if entityValue.(int64) == criteriaValue.(int64) {
				return true
			}

		case "string":
			if strings.ToLower(entityValue.(string)) == strings.ToLower(criteriaValue.(string)) {
				return true
			}
		}
	}

	return false
}

func (a InMemoryCriteriaBuilderAdapter) Like(c criteria.LikeCriteria, fieldsType map[string]string, entityValue any) bool {
	fieldType := fieldsType[c.Field]

	if entityValue == nil {
		return false
	}

	switch fieldType {
	case "string":
		if strings.Contains(strings.ToLower(entityValue.(string)), strings.ToLower(c.Value)) {
			return true
		}
	}

	return false
}

func (a InMemoryCriteriaBuilderAdapter) Lt(c criteria.LtCriteria, fieldsType map[string]string, entityValue any) bool {

	fieldType := fieldsType[c.Field]

	if entityValue == nil {
		return false
	}

	switch fieldType {
	case "int":
		return entityValue.(int64) > c.Value.(int64)

	case "string":
		return strings.ToLower(entityValue.(string)) > strings.ToLower(c.Value.(string))

	case "time":
		entityValue := entityValue.(time.Time).Truncate(time.Minute)
		criteriaValue, _ := time.Parse(time.RFC3339, c.Value.(string))
		criteriaValue = criteriaValue.Truncate(time.Minute)
		return entityValue.Before(criteriaValue)
	}

	return false
}

func (a InMemoryCriteriaBuilderAdapter) Lte(c criteria.LteCriteria, fieldsType map[string]string, entityValue any) bool {

	fieldType := fieldsType[c.Field]

	if entityValue == nil {
		return false
	}

	switch fieldType {
	case "int":
		return entityValue.(int64) >= c.Value.(int64)

	case "string":
		return entityValue.(string) >= c.Value.(string)

	case "time":
		entityValue := entityValue.(time.Time).Truncate(time.Minute)
		criteriaValue, _ := time.Parse(time.RFC3339, c.Value.(string))
		criteriaValue = criteriaValue.Truncate(time.Minute)
		return entityValue.Before(criteriaValue) || entityValue.Equal(criteriaValue)
	}

	return false
}

func (a InMemoryCriteriaBuilderAdapter) Sort(c criteria.SorterCriteria, fieldsType map[string]string, v1 any, v2 any) bool {

	fieldType := fieldsType[c.By()]

	switch fieldType {
	case "int":
		if c.Sort() == "asc" {
			return v1.(int64) < v2.(int64)
		}
		return v1.(int64) > v2.(int64)

	case "string":
		if c.Sort() == "asc" {
			return v1.(string) < v2.(string)
		}
		return v1.(string) > v2.(string)

	case "time":
		v1, _ := time.Parse(time.RFC3339, v1.(string))
		v1 = v1.Truncate(time.Minute)
		v2, _ := time.Parse(time.RFC3339, v2.(string))
		v2 = v2.Truncate(time.Minute)

		if c.Sort() == "asc" {
			return v1.Before(v2)
		}

		return v1.After(v2)
	}

	return false
}
