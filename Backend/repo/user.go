package repo

import (
	"database/sql"
	"first-program/domain"
	"first-program/user"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db, // assign the sqlx database connection to the user repository(userRepo) struct
	}
}

func (r *userRepo) Create(user domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (
			first_name,
			last_name,
			email,
			password,
			is_shop_owner
		)
		VALUES (
			:first_name,
			:last_name,
			:email,
			:password,
			:is_shop_owner
		)
		RETURNING id
	`
	// Exectute named query

	var userID int

	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		rows.Scan(&userID)
	}
	user.ID = userID

	return &user, nil
}

func (r *userRepo) Find(email, pass string) (*domain.User, error) {
	var user domain.User
	query := `
		SELECT id, first_name, last_name, email, password, is_shop_owner
		FROM users
		WHERE email = $1 and password = $2
		LIMIT 1
	`
	err := r.db.Get(&user, query, email, pass) // db is *sqlx.DB
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no matching user
		}
		return nil, err
	}
	return &user, nil
}
