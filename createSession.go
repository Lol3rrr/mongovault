package mongovault

// CreateSession simply returns a valid DB-Session struct
func CreateSession(dbSettings DBOptions, vSettings VaultSettings) (DB, error) {
	result := &Session{
		URL:             dbSettings.URL,
		Port:            dbSettings.Port,
		Database:        dbSettings.Database,
		Collection:      dbSettings.Collection,
		ApplicationName: dbSettings.ApplicationName,

		VaultSession:  vSettings.Session,
		CredsEndpoint: vSettings.CredsName,
	}

	err := result.loadCreds()
	if err != nil {
		return result, err
	}

	err = result.Connect()
	if err != nil {
		return result, err
	}

	return result, nil
}
