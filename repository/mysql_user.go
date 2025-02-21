package repository

import (
	"database/sql"
	"echo-go-api/models"
)

// UserRepository defines the methods available for user repository operations.
type UserRepository interface {
	GetUser(id int) (*models.User, error)
	CreateUser(user *models.User) error
	GetUsers() ([]*models.User, error) // New method
	UpdateUser(id string, user *models.User) error
	DeleteUser(id string) error
}

// userRepo is a struct that implements UserRepository.
type userRepo struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of userRepo.
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

// GetUserByID retrieves a user by their ID along with their associated roles.
func (r *userRepo) GetUser(id int) (*models.User, error) {
	var user models.User
	// Query user data
	err := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	// Now, get the roles associated with the user
	rows, err := r.db.Query("SELECT r.id, r.name FROM roles r JOIN user_role ur ON r.id = ur.role_id WHERE ur.user_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Append roles to the user object
	var roles []models.Role
	for rows.Next() {
		var role models.Role
		if err := rows.Scan(&role.ID, &role.Name); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	// Assign the roles to the user
	user.Roles = roles

	// Check for errors from iterating over the rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser creates a new user in the database.
func (r *userRepo) CreateUser(user *models.User) error {
	// Insert user data
	result, err := r.db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		return err
	}

	// Get the inserted user's ID
	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Insert roles for the user
	for _, role := range user.Roles {
		_, err := r.db.Exec("INSERT INTO user_role (user_id, role_id) VALUES (?, ?)", userID, role.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetAllUsers retrieves all users from the database.
func (r *userRepo) GetUsers() ([]*models.User, error) {
	var users []*models.User

	// Query all users
	rows, err := r.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Loop through the result set and populate the users slice
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	// Check for errors from iterating over the rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// UpdateUserById updates a user's name, email, and roles by their ID.
func (r *userRepo) UpdateUser(id string, user *models.User) error {
	// Begin transaction to ensure data consistency
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Update the user's basic information (name, email)
	_, err = tx.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, id)
	if err != nil {
		return err
	}

	// Remove existing roles for the user
	_, err = tx.Exec("DELETE FROM user_role WHERE user_id = ?", id)
	if err != nil {
		return err
	}

	// Insert new roles for the user
	for _, role := range user.Roles {
		_, err := tx.Exec("INSERT INTO user_role (user_id, role_id) VALUES (?, ?)", id, role.ID)
		if err != nil {
			return err
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// DeleteUserById deletes a user by their ID along with their associated roles.
func (r *userRepo) DeleteUser(id string) error {
	// Begin transaction to ensure data consistency
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Remove user roles first
	_, err = tx.Exec("DELETE FROM user_role WHERE user_id = ?", id)
	if err != nil {
		return err
	}

	// Then remove the user from the users table
	_, err = tx.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
