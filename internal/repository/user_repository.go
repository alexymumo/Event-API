package repository

import (
	"events/internal/models"
)

//var db *sql.DB

type UserRepository interface {
	Login()
	SignUp()
}

func RegisterUser(user *models.User) (*models.User, error) {

	return user, nil
	/*stmt, err := db.Prepare("INSERT INTO users(name, email,password) VALUES(?,?,?)")
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
	*/
}

func LoginUser() {

}
