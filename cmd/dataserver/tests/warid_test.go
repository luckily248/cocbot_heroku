package tests

import (
	. "heroku-dataserver/cmd/dataserver/models"
	"testing"
)

func TestWarid(t *testing.T) {
	id, err := GetNextWarId()
	if err != nil {
		t.Fatalf("warid err:%s\n", err.Error())
	}
	t.Logf("warid is %d\n", id)
}
