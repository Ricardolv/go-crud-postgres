package repository

import (
	"context"
	"fmt"
	"github.com/Ricardolv/go-crud-postgres/models"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	conn *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{conn: conn}
}

// Create
func (r *UserRepository) Create(user *models.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at`
	err := r.conn.QueryRow(context.Background(), query, user.Name, user.Email).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return fmt.Errorf("falha ao criar usuário: %v", err)
	}
	return nil
}

// Read (by ID)
func (r *UserRepository) GetByID(id int) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1`
	err := r.conn.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("falha ao buscar usuário: %v", err)
	}
	return user, nil
}

// Update
func (r *UserRepository) Update(user *models.User) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
	_, err := r.conn.Exec(context.Background(), query, user.Name, user.Email, user.ID)
	if err != nil {
		return fmt.Errorf("falha ao atualizar usuário: %v", err)
	}
	return nil
}

// Delete
func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.conn.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("falha ao deletar usuário: %v", err)
	}
	return nil
}

// List All
func (r *UserRepository) List() ([]models.User, error) {
	query := `SELECT id, name, email, created_at FROM users`
	rows, err := r.conn.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("falha ao listar usuários: %v", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			return nil, fmt.Errorf("falha ao escanear usuário: %v", err)
		}
		users = append(users, user)
	}
	return users, nil
}
