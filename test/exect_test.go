package test

import (
	"context"
	"go-lang-database/db"
	"testing"
)

//exect =  tidak mengembalian nilai
//query  = mengembalikan nilai
func TestExecContextInsertDatabase(t *testing.T) {

	db, _ := db.SetupDatabase()

	ctx := context.Background()
	query := "INSERT INTO users VALUES(0,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		t.Error("Failed to prepare data: ", err.Error())
	}

	user := map[string]string{
		"name":  "Sandrian",
		"email": "sandrian@gmail.com",
	}

	_, err = stmt.ExecContext(ctx, user["name"], user["email"])
	if err != nil {
		t.Error("Failed to insert data: ", err.Error())
	}

}
