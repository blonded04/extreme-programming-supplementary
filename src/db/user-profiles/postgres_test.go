package db_user_profiles

import "testing"

func TestDbCreation(t *testing.T) {
	err := InitializeUserDb()
	if err != nil {
		t.Errorf("InitializeUserDb returned an error: %s", err.Error())
	}

	defer DropUserDb()
}
