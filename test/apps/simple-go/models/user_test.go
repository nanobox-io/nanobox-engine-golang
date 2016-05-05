package models

import "testing"

func TestCreateUser(t *testing.T) {
	usr, err := CreateUser()
	if usr == nil || err != nil {
		t.Error("User should be created %s", err.Error())
	}
}

func TestListUsers(t *testing.T) {
	users, err := ListUsers()
	if len(*users) < 1 || err != nil {
		t.Error("User list should atleast have an admin user")
	}
}

func TestDeleteUser(t *testing.T) {
	users, _ := ListUsers()
	for _, usr := range *users {
		if !usr.Admin {
			DeleteUser(usr.ID)
		}
	}

	users, _ = ListUsers()
	if len(*users) != 1 {
		t.Error("I Just deleted all of users.. why are you still here?")
	}

}
