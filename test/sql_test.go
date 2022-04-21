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

type User struct {
	ID        int
	Name      string
	Email     string
	Age       int
	IsStudent bool
	CreatedAt time.Time
}

func TestQueryContextSelectCondition(t *testing.T) {
	//setup database
	db, _ := db.SetupDatabase()

	//create context
	ctx := context.Background()

	//create query and create condition
	query := "SELECT * FROM users WHERE id = ? LIMIT 1"
	stmt, err := db.Prepare(query)
	if err != nil {
		t.Error("Failed to prepare data: ", err.Error())
	}
	//process query
	UserID := 4
	rows, err := stmt.QueryContext(ctx, UserID)
	if err != nil {
		t.Error("Failed to insert data: ", err.Error())
	}

	user := User{}
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.IsStudent, &user.CreatedAt)
		if err != nil {
			t.Error("Faild scan data:", err.Error())
		}
	}

	log.Println(user)

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
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.IsStudent, &user.CreatedAt)
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
