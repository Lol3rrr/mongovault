package mongovault

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get loads the Feedback with the given ID
func (session *Session) Get(query []Filter, result interface{}) error {
	if !session.isConnectionAlive() {
		err := session.reconnect()
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

	err := session.MongoCollection.FindOne(ctx, filter).Decode(result)

	return err
}
