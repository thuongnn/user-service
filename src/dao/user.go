package dao

import (
	"time"
	"user-service/src/common/utils"
	"user-service/src/models"
)

const (
	passwordVersion = utils.SHA256
)

// Register is used for user to register, the password is encrypted before the record is inserted into database.
func Register(user *models.User) (int64, error) {
	o := GetOrmer()

	salt := utils.GenerateRandomString()
	now := time.Now()
	sql := `insert into "user"
				(username, password, salt, creation_time, update_time)
				 values (?, ?, ?, ?, ?) RETURNING user_id`
	var userID int64
	err := o.Raw(sql, user.Username, utils.Encrypt(user.Password, salt, utils.SHA256), salt, now, now).QueryRow(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil

}

// UserExists returns whether a user exists according username.
func UserExists(user *models.User) (bool, error) {
	o := GetOrmer()
	isExist := o.QueryTable("user").Filter("username", user.Username).Exist()
	return isExist, nil
}

// LoginByDb is used for user to login with database auth mode.
func LoginByDb(auth models.AuthModel) (*models.User, error) {
	var user models.User
	o := GetOrmer()

	err := o.QueryTable("user").Filter("username", auth.Username).Filter("deleted", false).One(&user)
	if err != nil {
		return nil, err
	}

	if !matchPassword(&user, auth.Password) {
		return nil, nil
	}
	user.Password = "" // do not return the password
	return &user, nil
}

func GetUser(userQuery models.User) (*models.User, error) {
	o := GetOrmer()

	var u models.User
	err := o.QueryTable("user").Filter("user_id", userQuery.UserID).Filter("deleted", userQuery.Deleted).One(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// MatchPassword returns true is password matched
func matchPassword(u *models.User, password string) bool {
	return utils.Encrypt(password, u.Salt, passwordVersion) == u.Password
}
