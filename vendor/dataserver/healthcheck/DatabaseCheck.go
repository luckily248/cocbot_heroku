package healthcheck

import (
	"cocbot_heroku/cmd/dataserver/models"
)

type DatabaseCheck struct {
}

func (dbc *DatabaseCheck) Check() error {
	database := &models.BaseDBmodel{}
	return database.Check()
}
