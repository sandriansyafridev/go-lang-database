package test

import (
	"context"
	"go-lang-database/db"
	"go-lang-database/entity"
	"go-lang-database/repository"
	"log"
	"testing"
	"time"
)

func TestUserRepositoryFindByID(t *testing.T) {
	db, err := db.SetupDatabase()
	ctx := context.Background()
	if err != nil {
		t.Fatal("Connection error:", err.Error())
	}

	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.FindByID(ctx, 4)
	if err != nil {
		t.Fatal("Not Found:", err.Error())
	}

	log.Println(user)

}

func TestUserRepositoryFindByAll(t *testing.T) {
	db, err := db.SetupDatabase()
	ctx := context.Background()
	if err != nil {
		t.Fatal("Connection error:", err.Error())
	}

	userRepository := repository.NewUserRepository(db)
	users, err := userRepository.FindAll(ctx)
	if err != nil {
		t.Fatal("Not Found:", err.Error())
	}

	for _, user := range users {
		log.Println(user)
	}

}

func TestUserRepositoryInser(t *testing.T) {
	db, err := db.SetupDatabase()
	ctx := context.Background()
	if err != nil {
		t.Fatal("Connection error:", err.Error())
	}

	user := entity.User{
		Name:      "test",
		Email:     "test@gmail.com",
		Age:       28,
		IsStudent: true,
		CreatedAt: time.Now(),
	}
	userRepository := repository.NewUserRepository(db)
	userCreated, err := userRepository.Insert(ctx, user)
	if err != nil {
		t.Fatal("Not Found:", err.Error())
	}

	log.Println("CREATED")
	log.Println(userCreated)
	log.Println("CREATED")

}
