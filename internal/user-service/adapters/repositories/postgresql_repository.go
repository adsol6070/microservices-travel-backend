package repositories

import (
	"database/sql"
	"errors"
	"microservices-travel-backend/internal/user-service/domain/models"
	"microservices-travel-backend/internal/user-service/domain/ports"
)

type PostgreSQLUserRepository struct {
	db *sql.DB
}

func NewPostgreSQLUserRepository(db *sql.DB) ports.UserRepositoryPort {
	return &PostgreSQLUserRepository{db: db}
}

// Create creates a new user in the database
func (repo *PostgreSQLUserRepository) Create(user models.User) (*models.User, error) {
	query := `INSERT INTO users (id, email, name, password, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := repo.db.QueryRow(query, user.ID, user.Email, user.Name, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByID retrieves a user by their ID
func (repo *PostgreSQLUserRepository) GetByID(id string) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, name, password FROM users WHERE id = $1`
	err := repo.db.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Name, &user.Password)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail retrieves a user by their email
func (repo *PostgreSQLUserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, name, password FROM users WHERE email = $1`
	err := repo.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Name, &user.Password)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAll retrieves all users
func (repo *PostgreSQLUserRepository) GetAll() ([]models.User, error) {
	rows, err := repo.db.Query(`SELECT id, email, name, password FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

// Update updates a user's information
func (repo *PostgreSQLUserRepository) Update(id string, user models.User) (*models.User, error) {
	query := `UPDATE users SET email = $1, name = $2, password = $3, updated_at = $4 WHERE id = $5 RETURNING id`
	err := repo.db.QueryRow(query, user.Email, user.Name, user.Password, user.UpdatedAt, id).Scan(&user.ID)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

// Delete removes a user by their ID
func (repo *PostgreSQLUserRepository) Delete(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
