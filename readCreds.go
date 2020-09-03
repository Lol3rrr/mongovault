package mongovault

import (
	"errors"

	"github.com/hashicorp/vault/api"
)

func (session *Session) readCreds() (*api.Secret, error) {
	resp, err := session.VaultSession.Logical().Read(session.CredsEndpoint)
	if err != nil {
		return nil, err
	}

	if resp == nil || resp.Data == nil {
		return nil, errors.New("No Data was returned")
	}

	return resp, nil
}
