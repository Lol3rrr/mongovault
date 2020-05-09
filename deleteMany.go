package mongovault

import (
	"context"
	"time"
)

// DeleteMany is used to delete multiple items from the database
func (session *Session) DeleteMany(query []Filter) error {
	if !session.isConnectionAlive() {
		err := session.reconnect()
		if err != nil {
			return err
		}
	}

	filter := convertToPrimary(query)

	ctx, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelCtx()

	_, err := session.MongoCollection.DeleteMany(ctx, filter)

	return err
}
