package models

import (
	"errors"
	"time"

	"test.com/event-api/db"
	"test.com/event-api/utils"
)

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email" binding:"required"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email" binding:"required"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) Save() error {

	query := `
	INSERT INTO users (email, first_name, last_name, password, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	// HASH PASSWORD AND GET TIME
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	timeCreated := time.Now()

	data, err := statement.Exec(user.Email, user.FirstName, user.LastName, hashedPassword, timeCreated, timeCreated)
	if err != nil {
		return err
	}
	id, err := data.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = id
	user.Password = hashedPassword
	user.CreatedAt = timeCreated
	return nil
}

func GetUserByEmail(email string) (*UserResponse, error) {
	var query = `SELECT * FROM users WHERE email = ?`
	data := db.DB.QueryRow(query, email)
	user := &UserResponse{}
	err := data.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (login *Login) ValidateUserCredentials() (*UserResponse, error) {
	query := `
 SELECT * FROM users WHERE email = ?
 `
	row := db.DB.QueryRow(query, login.Email)
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, errors.New("User not found")
	}
	var userRes = UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	err = utils.ComparePassword(user.Password, login.Password)
	if err != nil {
		return nil, errors.New("wrong password")
	}

	return &userRes, nil
}

func (u *User) Update() (*User, error) {

	return nil, nil
}
