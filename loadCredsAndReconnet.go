package mongovault

import (
	"context"
)

func (s *Session) loadCredsAndConnect() error {
	err := s.loadCreds()
	if err != nil {
		return err
	}

	if s.MongoClient != nil {
		s.MongoClient.Disconnect(context.TODO())
	}

	return s.Connect()
}
