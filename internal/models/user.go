package models

import (
	"time"
)

type User struct {
	ID        int32     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	DOB       time.Time `json:"dob" db:"dob"`
	Age       int       `json:"age,omitempty" db:"-"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type CreateUserRequest struct {
	Name string    `json:"name" validate:"required,min=1,max=100"`
	DOB  time.Time `json:"dob" validate:"required"`
}

type UpdateUserRequest struct {
	Name string    `json:"name" validate:"required,min=1,max=100"`
	DOB  time.Time `json:"dob" validate:"required"`
}

type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age,omitempty"`
}

func (u *User) CalculateAge() int {
	now := time.Now()
	age := now.Year() - u.DOB.Year()

	if now.YearDay() < u.DOB.YearDay() {
		age--
	}

	return age
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:   u.ID,
		Name: u.Name,
		DOB:  u.DOB.Format("2006-01-02"),
		Age:  u.CalculateAge(),
	}
}
