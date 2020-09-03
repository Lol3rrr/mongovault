package mongovault

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (session *Session) isConnectionAlive() bool {
	if session.MongoClient == nil {
		return false
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelCtx()

	err := session.MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return false
	}

	return true
}
