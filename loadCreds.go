package mongovault

import (
	"errors"
	"fmt"
)

func (s *Session) loadCreds() error {
	data, err := s.readCreds()
	if err != nil {
		return fmt.Errorf("Could not load Credentials: %s", err)
	}

	creds := data.Data
	username, found := creds["username"]
	if !found {
		return errors.New("Vault response did not include a username")
	}
	password, found := creds["password"]
	if !found {
		return errors.New("Vault response did not include a password")
	}

	s.Username = username.(string)
	s.Password = password.(string)

	return nil
}
