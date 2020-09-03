package mongovault

import "go.mongodb.org/mongo-driver/bson/primitive"

func convertToPrimary(query []Filter) []primitive.E {
	filter := make([]primitive.E, len(query))
	for i, entry := range query {
		filter[i] = primitive.E(entry)
	}

	return filter
}
