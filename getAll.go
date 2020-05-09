package mongovault

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// GetAll loads all Feedback entrys from the Database
func (session *Session) GetAll(results interface{}) error {
	if !session.isConnectionAlive() {
		err := session.loadCredsAndConnect()
		if err != nil {
			return err
		}
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelCtx()

	cur, err := session.MongoCollection.Find(ctx, bson.D{{}})
	if err != nil {
		return err
	}

	if err := cur.All(context.TODO(), results); err != nil {
		return err
	}

	return nil
}
