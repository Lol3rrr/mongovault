package mongovault

import (
	"context"
	"errors"
	"fmt"
)

func (session *Session) reconnect() error {
	if session.MongoClient != nil {
		session.MongoClient.Disconnect(context.TODO())
	}

	err := session.Connect()
	if errors.Is(err, errors.New("test")) {
		session.loadCreds()

		return session.Connect()
	}

	fmt.Printf("Error: '%s' \n", err)

	return err
}
