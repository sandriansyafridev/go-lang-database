package test

import (
	"context"
	"go-lang-database/db"
	"log"
	"testing"
	"time"
)

//exect =  tidak mengembalian nilai
//query  = mengembalikan nilai
func TestExecContextInsertDatabase(t *testing.T) {

	db, _ := db.SetupDatabase()
	tx, err := db.Begin()
	if err != nil {
		t.Fatal("Error to Begin Transaction")
	}

	ctx := context.Background()
	query := "INSERT INTO users VALUES(0,?,?,?,?,?)"
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		t.Error("Failed to prepare data: ", err.Error())
	}

	defer stmt.Close()

	user := map[string]interface{}{
		"name":       "random",
		"email":      "random@gmail.com",
		"age":        25,
		"is_student": false,
		"created_at": time.Now(),
	}

	_, err = stmt.ExecContext(ctx, user["name"], user["email"], user["age"], user["is_student"], user["created_at"])
	if err != nil {
		t.Error("Failed to insert data: ", err.Error())
	}

	log.Println("CREATED")
	tx.Rollback()

}
