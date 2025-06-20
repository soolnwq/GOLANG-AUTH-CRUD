package repository

import (
	"go-crud/entities"

	"github.com/jmoiron/sqlx"
)

type authUserRepositoryDB struct {
	db *sqlx.DB
}

func NewAuthUserRepository(db *sqlx.DB) AuthUserRepository {
	return &authUserRepositoryDB{db: db}
}

func (r *authUserRepositoryDB) Insert(user *entities.User) (*entities.User, error) {
	query := "insert into users (username, email, password) values (?,?,?)"
	result, err := r.db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	lastUserID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = int(lastUserID)
	return user, nil
}

func (r *authUserRepositoryDB) FindByUsernameOrEmail(username string, email string) (*entities.User, error) {
	user := entities.User{}
	query := "select * from users where username = ? or email = ?"
	if err := r.db.Get(&user, query, username, email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authUserRepositoryDB) FindByUsername(username string) (*entities.User, error) {
	user := entities.User{}
	query := "select * from users where username = ?"
	if err := r.db.Get(&user, query, username); err != nil {
		return nil, err
	}

	return &user, nil
}
