package mongovault

import (
	"context"
	"time"
)

// GetAll loads all Feedback entrys from the Database
func (session *Session) GetAll(query []Filter, results interface{}) error {
	if !session.isConnectionAlive() {
		err := session.reconnect()
		if err != nil {
			return err
		}
	}

	filter := convertToPrimary(query)

	ctx, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelCtx()

	cur, err := session.MongoCollection.Find(ctx, filter)
	if err != nil {
		return err
	}

	if err := cur.All(context.TODO(), results); err != nil {
		return err
	}

	return nil
}
