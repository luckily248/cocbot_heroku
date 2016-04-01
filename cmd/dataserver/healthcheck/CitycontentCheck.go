package healthcheck

import (
	"dataserver/models"
)

type CitycontentCheck struct {
}

func (cc *CitycontentCheck) Check() error {
	return models.CheckCitycontentInstance()
}
