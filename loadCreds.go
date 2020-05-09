package mongovault

import (
	"errors"
)

func (s *Session) loadCreds() error {
	data, err := s.VaultSession.ReadData(s.CredsEndpoint)
	if err != nil {
		return errors.New("Could not load Credentials from Vault")
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
