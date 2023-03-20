package criteria

import (
	"<MODULE_URL_REPLACE>/pkg/shared/domain/criteria"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoCriteriaBuilderAdapter struct {
}

func NewMongoCriteriaBuilderAdapter() MongoCriteriaBuilderAdapter {
	return MongoCriteriaBuilderAdapter{}
}

func (a MongoCriteriaBuilderAdapter) Build(c criteria.Criteria) bson.M {
	if c == nil {
		return bson.M{}
	}

	switch c.Type() {
	case "and":
		andCriteria := c.(criteria.AndCriteria)
		return bson.M{
			"$and": []bson.M{
				a.Build(andCriteria.C1),
				a.Build(andCriteria.C2),
			},
		}

	case "between":
		betweenCriteria := c.(criteria.BetweenCriteria)
		return a.Between(betweenCriteria.Field, betweenCriteria.V1, betweenCriteria.V2)

	case "eq":
		eqCriteria := c.(criteria.EqCriteria)
		return a.Eq(eqCriteria.Field, eqCriteria.Value)

	case "gt":
		gtCriteria := c.(criteria.GtCriteria)
		return a.Gt(gtCriteria.Field, gtCriteria.Value)

	case "gte":
		gteCriteria := c.(criteria.GteCriteria)
		return a.Gt(gteCriteria.Field, gteCriteria.Value)

	case "in":
		inCriteria := c.(criteria.InCriteria)
		return a.In(inCriteria.Field, inCriteria.Value)

	case "like":
		likeCriteria := c.(criteria.LikeCriteria)
		return a.Like(likeCriteria.Field, likeCriteria.Value)

	case "lt":
		ltCriteria := c.(criteria.LtCriteria)
		return a.Lt(ltCriteria.Field, ltCriteria.Value)

	case "lte":
		lteCriteria := c.(criteria.LteCriteria)
		return a.Lte(lteCriteria.Field, lteCriteria.Value)

	case "or":
		orCriteria := c.(criteria.OrCriteria)

		return bson.M{
			"$or": []bson.M{
				a.Build(orCriteria.C1),
				a.Build(orCriteria.C2),
			},
		}

	case "not":
		notCriteria := c.(criteria.NotCriteria)
		return bson.M{"$not": a.Build(notCriteria.Criteria)}
	}

	return bson.M{}
}

func (MongoCriteriaBuilderAdapter) Between(field string, v1 any, v2 any) bson.M {
	return bson.M{
		field: bson.M{
			"$gte": v1,
			"$lte": v2,
		},
	}
}

func (MongoCriteriaBuilderAdapter) Eq(field string, value any) bson.M {
	return bson.M{field: value}
}

func (MongoCriteriaBuilderAdapter) Gt(field string, value any) bson.M {
	return bson.M{
		field: bson.M{
			"$gt": value,
		},
	}
}

func (MongoCriteriaBuilderAdapter) Gte(field string, value any) bson.M {
	return bson.M{
		field: bson.M{
			"$gte": value,
		},
	}
}

func (MongoCriteriaBuilderAdapter) In(field string, values any) bson.M {
	return bson.M{
		field: bson.M{
			"$in": values,
		},
	}
}

func (MongoCriteriaBuilderAdapter) Like(field string, value string) bson.M {
	return bson.M{
		field: bson.M{
			"$regex":   ".*" + value + ".*",
			"$options": "i",
		},
	}
}

func (MongoCriteriaBuilderAdapter) Lt(field string, value any) bson.M {
	return bson.M{
		field: bson.M{
			"$lt": value,
		},
	}
}

func (MongoCriteriaBuilderAdapter) Lte(field string, value any) bson.M {
	return bson.M{
		field: bson.M{
			"$lte": value,
		},
	}
}

func (MongoCriteriaBuilderAdapter) Not(field string, value any) bson.M {
	return bson.M{
		field: bson.M{"$not": bson.M{"$eq": value}},
	}
}

func (MongoCriteriaBuilderAdapter) Sort(o criteria.SorterCriteria) bson.D {
	var sorting map[string]int8 = map[string]int8{"asc": 1, "desc": -1}

	if len(o.By()) == 0 || len(o.Sort()) == 0 {
		return bson.D{}
	}

	return bson.D{primitive.E{Key: o.By(), Value: sorting[o.Sort()]}}
}
