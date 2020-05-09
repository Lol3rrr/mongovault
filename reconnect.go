package mongovault

import (
	"context"
	"errors"
)

const authErrorMessage string = "connection() : auth error: sasl conversation error: unable to authenticate using mechanism \"SCRAM-SHA-1\": (AuthenticationFailed) Authentication failed."

func (session *Session) reconnect() error {
	if session.MongoClient != nil {
		session.MongoClient.Disconnect(context.TODO())
	}

	err := session.Connect()
	if errors.Is(err, errors.New(authErrorMessage)) {
		session.loadCreds()

		return session.Connect()
	}

	return err
}
