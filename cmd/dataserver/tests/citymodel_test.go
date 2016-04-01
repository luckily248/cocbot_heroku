package tests

import (
	. "dataserver/models"
	"testing"
)

func TestCity(t *testing.T) {
	cityname := "广州"
	cityid, err := GetOneCityid(cityname)
	if err != nil {
		t.Fatalf("getcityid err:%s\n", err.Error())
	}
	t.Logf("%s cityid is %s\n", cityname, cityid)

	content, err := GetCityContent()
	if err != nil {
		t.Fatalf("getcitycontent err:%s\n", err.Error())
	}
	t.Logf("getcontent success [0]:%s\n", content.Areas[0].Area)
}
