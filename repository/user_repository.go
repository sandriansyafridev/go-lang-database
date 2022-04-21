package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-lang-database/entity"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]entity.User, error)
	FindByID(ctx context.Context, UserID int) (entity.User, error)
	Insert(ctx context.Context, user entity.User) (entity.User, error)
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (userRepository *userRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	querySQL := "SELECT * FROM users"
	users := []entity.User{}

	stmt, err := userRepository.DB.PrepareContext(ctx, querySQL)
	if err != nil {
		return users, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return users, err
	}

	defer rows.Close()
	for rows.Next() {

		user := entity.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.IsStudent, &user.CreatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)

	}

	return users, nil

}

func (userRepository *userRepository) FindByID(ctx context.Context, UserID int) (entity.User, error) {

	querySQL := "SELECT * FROM users WHERE id = ? LIMIT 1"
	user := entity.User{}

	stmt, err := userRepository.DB.PrepareContext(ctx, querySQL)
	if err != nil {
		return user, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, UserID)
	if err != nil {
		return user, err
	}

	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.IsStudent, &user.CreatedAt)
		if err != nil {
			return user, err
		}
	} else {
		return user, errors.New("user not found")
	}

	return user, nil

}

func (userRepository *userRepository) Insert(ctx context.Context, user entity.User) (entity.User, error) {

	querySQL := "INSERT INTO users(name, email, age, is_student, created_at) VALUES (?,?,?,?,?)"
	stmt, err := userRepository.DB.PrepareContext(ctx, querySQL)
	if err != nil {
		return user, err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, user.Name, user.Email, user.Age, user.IsStudent, user.CreatedAt)
	if err != nil {
		return user, err
	}

	UserID, _ := result.LastInsertId()

	user.ID = int(UserID)

	return user, nil

}
