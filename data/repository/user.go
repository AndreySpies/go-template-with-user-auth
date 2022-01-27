package repository

import (
	"database/sql"

	"github.com/AndreySpies/doccer/domain/constants"
	"github.com/AndreySpies/doccer/domain/contract"
	"github.com/AndreySpies/doccer/domain/entity"
	"github.com/AndreySpies/doccer/server/request"
	"github.com/go-sql-driver/mysql"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) contract.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user *request.CreateUser, encryptedPassword []byte) (err error) {
	queryInsert := `INSERT INTO doccer.user (
		first_name,
		last_name,
		email,
		password,
		birthday,
		location,
		created_at
	) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err = r.db.Exec(queryInsert,
		user.FirstName,
		user.LastName,
		user.Email,
		encryptedPassword,
		user.Birthday,
		user.Location,
		user.CreatedAt,
	)

	if err != nil && err.(*mysql.MySQLError).Number == constants.MySQLDuplicateEntryErrorCode {
		err = constants.CustomError(constants.EmailAlreadyUsedErrorCode)
	}

	return err
}

func (r *userRepo) FindOneByEmail(email string) (user entity.User, err error) {
	query := `SELECT * FROM user WHERE email = ?`
	rows, err := r.db.Query(query, email)
	if err != nil {
		return user, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.Birthday,
			&user.Location,
			&user.TotalBalance,
			&user.CreatedAt,
		)
	}

	return user, err
}

func (r *userRepo) FindOneByID(id int) (user entity.User, err error) {
	query := `SELECT * FROM user WHERE id = ?`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return user, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.Birthday,
			&user.Location,
			&user.TotalBalance,
			&user.CreatedAt,
		)
	}

	return user, err
}
