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

	ctx := context.Background()
	query := "INSERT INTO users VALUES(0,?,?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		t.Error("Failed to prepare data: ", err.Error())
	}

	user := map[string]interface{}{
		"name":       "Yogi",
		"email":      "yogi@gmail.com",
		"age":        27,
		"is_student": false,
		"created_at": time.Now(),
	}

	result, err := stmt.ExecContext(ctx, user["name"], user["email"], user["age"], user["is_student"], user["created_at"])
	if err != nil {
		t.Error("Failed to insert data: ", err.Error())
	}

	UserID, _ := result.LastInsertId()
	log.Println("Last Insert UserID :", UserID)

}
