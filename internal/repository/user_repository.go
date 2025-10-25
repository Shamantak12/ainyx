package repository

import (
	"context"
	"database/sql"
	"time"

	"user-api/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, name string, dob time.Time) (*models.User, error)
	GetByID(ctx context.Context, id int32) (*models.User, error)
	List(ctx context.Context) ([]*models.User, error)
	Update(ctx context.Context, id int32, name string, dob time.Time) (*models.User, error)
	Delete(ctx context.Context, id int32) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, name string, dob time.Time) (*models.User, error) {
	query := `
		INSERT INTO users (name, dob)
		VALUES ($1, $2)
		RETURNING id, name, dob, created_at, updated_at
	`

	var user models.User
	err := r.db.QueryRowContext(ctx, query, name, dob).Scan(
		&user.ID,
		&user.Name,
		&user.DOB,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int32) (*models.User, error) {
	query := `
		SELECT id, name, dob, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user models.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.DOB,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) List(ctx context.Context) ([]*models.User, error) {
	query := `
		SELECT id, name, dob, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.DOB,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) Update(ctx context.Context, id int32, name string, dob time.Time) (*models.User, error) {
	query := `
		UPDATE users
		SET name = $2, dob = $3, updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
		RETURNING id, name, dob, created_at, updated_at
	`

	var user models.User
	err := r.db.QueryRowContext(ctx, query, id, name, dob).Scan(
		&user.ID,
		&user.Name,
		&user.DOB,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Delete(ctx context.Context, id int32) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return models.ErrUserNotFound
	}

	return nil
}
