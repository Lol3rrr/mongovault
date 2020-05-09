package mongovault

import (
	"context"
	"time"
)

// Delete deletes the Entry with the matching filter from the Database
func (session *Session) Delete(query []Filter) error {
	if !session.isConnectionAlive() {
		err := session.reconnect()
		if err != nil {
			return err
		}
	}

	filter := convertToPrimary(query)

	ctx, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelCtx()

	_, err := session.MongoCollection.DeleteOne(ctx, filter)

	return err
}
