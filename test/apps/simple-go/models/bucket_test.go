package models

import "testing"

var aUser *User

func TestCreatebucket(t *testing.T) {
	buck, err := CreateBucket(adminUser().ID, adminUser().Key, "test")
	if buck == nil || err != nil {
		t.Error("bucket should be created %s", err.Error())
	}
}

func TestListbuckets(t *testing.T) {
	bucks, err := ListBuckets(adminUser().ID, adminUser().Key)
	if len(*bucks) < 1 || err != nil {
		t.Error("Bucket should exist")
	}
}

func TestDeletebucket(t *testing.T) {
	bucks, _ := ListBuckets(adminUser().ID, adminUser().Key)
	for _, usr := range *bucks {
		DeleteBucket(adminUser().ID, adminUser().Key, usr.ID)
	}

	bucks, _ = ListBuckets(adminUser().ID, adminUser().Key)
	if len(*bucks) > 0 {
		t.Error("Bucket shouldn't exist")
	}

}

func adminUser() *User {
	if aUser == nil {
		users, _ := ListUsers()
		for _, usr := range *users {
			if usr.Admin {
				aUser = &usr
				return aUser
			}
		}
	}
	return aUser
}
