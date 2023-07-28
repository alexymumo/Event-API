package repository

import (
	"events/models"
)

//var db *sql.DB

func RegisterUser(user *models.User) (*models.User, error) {
	stmt, err := db.Prepare("INSERT INTO users(name, email,password) VALUES(?,?,?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := db.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}
	userid, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = int(userid)
	return user, nil
}

func DeleteUser(id int) (int64, error) {
	user, err := db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		return 0, err
	}
	result, err := user.RowsAffected()
	if err != nil {
		return 0, err
	}

	return result, err

}
