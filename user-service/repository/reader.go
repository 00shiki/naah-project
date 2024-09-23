package repository

import (
	"database/sql"
	"user-service/entity/users"
)

func (ur *UserRepository) GetUserByID(userID int64) (*users.User, error) {
	user := &users.User{
		ID: userID,
	}
	stmt, err := ur.db.Prepare(`SELECT email, password_hash, first_name, last_name, birth_date, address, contact_no, role, verified FROM users WHERE user_id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(userID)
	if row == nil {
		return nil, sql.ErrNoRows
	}
	err = row.Scan(
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.BirthDate,
		&user.Address,
		&user.ContactNo,
		&user.Role,
		&user.Verified,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*users.User, error) {
	user := &users.User{
		Email: email,
	}
	stmt, err := ur.db.Prepare(`SELECT user_id, password_hash, first_name, last_name, birth_date, address, contact_no, role, verified FROM users WHERE email = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(email)
	if row == nil {
		return nil, sql.ErrNoRows
	}
	err = row.Scan(
		&user.ID,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.BirthDate,
		&user.Address,
		&user.ContactNo,
		&user.Role,
		&user.Verified,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
