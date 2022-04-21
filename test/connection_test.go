package test

import (
	"go-lang-database/db"
	"testing"
)

func TestSetupConnectionDatabase(t *testing.T) {

	_, err := db.SetupDatabase()

	if err != nil {
		t.Error("Connection error:", err.Error())
	}

}
