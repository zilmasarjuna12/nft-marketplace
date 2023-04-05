package graph

import "nft-marketplace/entity"

type itemQuery struct {
	entity.ItemQuery
}

func NewItemQuery(filter *Filter) itemQuery {
	i := itemQuery{}

	if filter != nil {
		i.ItemQuery = entity.ItemQuery{
			Rating:          filter.Rating,
			ReputationBadge: filter.ReputationBadge,
			Category:        filter.Category,
		}

		if filter.Availability != nil {
			i.ItemQuery.Availability = &entity.RangeInput{
				Gte: filter.Availability.Gte,
				Lte: filter.Availability.Lte,
			}
		}
	}

	return i
}
