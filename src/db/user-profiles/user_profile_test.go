package db_user_profiles

import "testing"

func TestUserInsertion(t *testing.T) {
	err := InitializeUserDb()
	if err != nil {
		t.Errorf("InitializeUserDb returned an error: %s", err.Error())
	}
	defer DropUserDb()

	user, err := CreateNewUser("valera", "cool2004")
	if err != nil {
		t.Errorf("CreateNewUser returned an error: %s", err.Error())
	}

	_, err = CreateNewUser("valera", "another valera attempts to create another account")
	if err == nil {
		t.Error("CreateNewUser just allowed to create 1 more user with same login: valera")
	}

	uid_got, err := GetUidByLogin("valera")
	if err != nil {
		t.Errorf("GetUidByLogin returned an error: %s", err.Error())
	}
	if user.userId != uid_got {
		t.Errorf("GetUidByLogin returned an error: %s", err.Error())
	}

	match, err := CheckUidLoginMatch("valera", user.userId)
	if err != nil || !match {
		t.Error("CheckUidLoginMatch failed")
	}
}
