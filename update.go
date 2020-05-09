package mongovault

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Update updates the done value in a feedback entry
func (session *Session) Update(query []Filter, updates UpdateValue) error {
	if !session.isConnectionAlive() {
		err := session.loadCredsAndConnect()
		if err != nil {
			return err
		}
	}

	update := bson.M{
		"$set": bson.M(updates),
	}

	filter := make([]primitive.E, len(query))
	for i, entry := range query {
		filter[i] = primitive.E(entry)
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelCtx()

	_, err := session.MongoCollection.UpdateOne(ctx, filter, update)

	return err
}
