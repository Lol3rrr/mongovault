package mongovault

import (
	"context"
)

const authErrorMessage string = "connection() : auth error: sasl conversation error: unable to authenticate using mechanism \"SCRAM-SHA-1\": (AuthenticationFailed) Authentication failed."

func (session *Session) reconnect() error {
	if session.MongoClient != nil {
		session.MongoClient.Disconnect(context.TODO())
	}

	err := session.connect()
	if err.Error() == authErrorMessage {
		session.loadCreds()

		return session.connect()
	}

	return err
}
