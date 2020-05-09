package mongovault

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connect takes the given Session and connects to the Database
func (session *Session) Connect() error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelCtx()

	credentials := options.Credential{
		Username:      session.Username,
		Password:      session.Password,
		AuthMechanism: "SCRAM-SHA-1",
	}

	options := options.Client()
	options = options.ApplyURI("mongodb://" + session.URL + ":" + session.Port)
	options = options.SetAuth(credentials)
	options = options.SetAppName(session.ApplicationName)
	options = options.SetMaxConnIdleTime(10 * time.Second)

	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return err
	}

	pingCtx, cancelPingCtx := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelPingCtx()

	err = client.Ping(pingCtx, readpref.Primary())
	if err != nil {
		return err
	}

	session.MongoClient = client
	session.MongoCollection = client.Database(session.Database).Collection(session.Collection)

	return nil
}
