package repository

import (
	"database/sql"
	"echo-go-api/models"
)

type UserRepository interface {
	GetUser(id int) (*models.User, error)
	CreateUser(user *models.User) error
	GetUsers() ([]*models.User, error)
	UpdateUser(id string, user *models.User) error
	DeleteUser(id string) error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) GetUser(id int) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query("SELECT r.id, r.name FROM roles r JOIN user_role ur ON r.id = ur.role_id WHERE ur.user_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []models.Role
	for rows.Next() {
		var role models.Role
		if err := rows.Scan(&role.ID, &role.Name); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	user.Roles = roles

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) CreateUser(user *models.User) error {
	result, err := r.db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	for _, role := range user.Roles {
		_, err := r.db.Exec("INSERT INTO user_role (user_id, role_id) VALUES (?, ?)", userID, role.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *userRepo) GetUsers() ([]*models.User, error) {
	var users []*models.User

	rows, err := r.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepo) UpdateUser(id string, user *models.User) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, id)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM user_role WHERE user_id = ?", id)
	if err != nil {
		return err
	}

	for _, role := range user.Roles {
		_, err := tx.Exec("INSERT INTO user_role (user_id, role_id) VALUES (?, ?)", id, role.ID)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *userRepo) DeleteUser(id string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM user_role WHERE user_id = ?", id)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
