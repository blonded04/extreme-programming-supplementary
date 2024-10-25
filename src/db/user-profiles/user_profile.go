package db_user_profiles

import (
	"crypto/sha256"
	"errors"
)

const (
	maxLoginSize = 500
)

func GetUidByLogin(login string) (uid, error) {
	u := User{login: login}
	err := pgdb.Model(&u).Where("login = ?", login).Select()
	if err != nil {
		return 0, err
	}

	return u.userId, nil
}

func CheckUidLoginMatch(login string, userId uid) (bool, error) {
	uid, err := GetUidByLogin(login)
	if err != nil {
		return false, err
	}

	return uid == userId, nil
}

func insertNewUserToDb(login string, passwordHash hash) (uid, error) {
	u := User{login: login, passwordHash: passwordHash}
	_, err := pgdb.Model(&u).Returning("id").OnConflict("DO NOTHING").Insert()
	if err != nil {
		return 0, err
	}

	if u.userId == 0 {
		return 0, errors.New("login already exists")
	}
	return u.userId, nil
}

func CreateNewUser(login string, password string) (User, error) {
	if len(login) > 500 {
		return User{}, errors.New("login must be at most 500 characters")
	}
	passwordHash := sha256.Sum256([]byte(password))
	uid, err := insertNewUserToDb(login, passwordHash)
	if err != nil {
		return User{}, err
	}

	return User{login, passwordHash, uid}, nil
}
