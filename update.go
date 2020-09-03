package mongovault

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Update updates all the entys that match the query
func (session *Session) Update(query []Filter, updates UpdateValue, opts ...*options.UpdateOptions) error {
	if !session.isConnectionAlive() {
		err := session.reconnect()
		if err != nil {
			return err
		}
	}

	update := bson.M{
		"$set": bson.M(updates),
	}

	filter := convertToPrimary(query)

	ctx, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelCtx()

	_, err := session.MongoCollection.UpdateMany(ctx, filter, update, opts)

	return err
}
