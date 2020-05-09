package mongovault

import (
	"context"
	"time"
)

// Get loads the Feedback with the given ID
func (session *Session) Get(query []Filter, result interface{}) error {
	if !session.isConnectionAlive() {
		err := session.reconnect()
		if err != nil {
			return err
		}
	}

	filter := convertToPrimary(query)

	ctx, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelCtx()

	err := session.MongoCollection.FindOne(ctx, filter).Decode(result)

	return err
}
