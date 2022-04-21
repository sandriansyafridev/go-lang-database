package test

import (
	"context"
	"go-lang-database/db"
	"log"
	"testing"
)

//exect =  tidak mengembalian nilai
//query  = mengembalikan nilai

type User struct {
	ID    int
	Name  string
	Email string
}

func TestQueryContextSelect(t *testing.T) {

	//setup database
	db, _ := db.SetupDatabase()

	//create context
	ctx := context.Background()

	//create query
	query := "SELECT * FROM users"
	stmt, err := db.Prepare(query)
	if err != nil {
		t.Error("Failed to prepare data: ", err.Error())
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		t.Error("Failed to insert data: ", err.Error())
	}

	//process scanning data
	users := []User{}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			t.Error("Faild scan data:", err.Error())
		}
		users = append(users, user)
	}

	// log result.
	for _, user := range users {
		log.Println(user)
	}

}
