package mongovault

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete deletes the Entry with the matching filter from the Database
func (session *Session) Delete(query []Filter) error {
	if !session.isConnectionAlive() {
		err := session.loadCredsAndConnect()
		if err != nil {
			return err
		}
	}

	filter := make([]primitive.E, len(query))
	for i, entry := range query {
		filter[i] = primitive.E(entry)
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelCtx()

	_, err := session.MongoCollection.DeleteOne(ctx, filter)

	return err
}
