package repository

import (
	"database/sql"
	"user-service/entity/users"
)

func (ur *UserRepository) CreateUser(user *users.User) error {
	stmt, err := ur.db.Prepare(`INSERT INTO users (email, password_hash, first_name, last_name, birth_date, address, contact_no, role) values (?, ?, ?, ? , ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.BirthDate,
		user.Address,
		user.ContactNo,
		user.Role,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}

func (ur *UserRepository) UpdateUser(user *users.User) error {
	stmt, err := ur.db.Prepare(`UPDATE users SET email = ?, password_hash = ?, first_name = ?, last_name = ?, birth_date = ?, address = ?, contact_no = ?, verified = ? WHERE user_id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.BirthDate,
		user.Address,
		user.ContactNo,
		user.Verified,
		user.ID,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
