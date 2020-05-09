package mongovault

import (
	"context"
	"time"
)

// Insert inserts the given Feedback into the Database
func (session *Session) Insert(content interface{}) error {
	if !session.isConnectionAlive() {
		err := session.loadCredsAndConnect()
		if err != nil {
			return err
		}
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelCtx()

	_, err := session.MongoCollection.InsertOne(ctx, content)

	return err
}
