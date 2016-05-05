package models

import "testing"

var aBuck *Bucket

func TestCreateObject(t *testing.T) {
	f, err := CreateObject(adminUser().ID, adminUser().Key, adminBucket().ID, "test")
	if f == nil || err != nil {
		t.Error("Object should be created %s", err.Error())
	}
}

func TestListObjects(t *testing.T) {
	f, err := ListObjects(adminUser().ID, adminUser().Key, adminBucket().ID)
	if len(*f) < 1 || err != nil {
		t.Error("Objects should exist")
	}
}

func TestDeleteObject(t *testing.T) {
	fs, _ := ListObjects(adminUser().ID, adminUser().Key, adminBucket().ID)
	for _, obj := range *fs {
		DeleteObject(adminUser().ID, adminUser().Key, adminBucket().ID, obj.ID)
	}

	fs, _ = ListObjects(adminUser().ID, adminUser().Key, adminBucket().ID)
	if len(*fs) > 0 {
		t.Error("Bucket shouldn't exist")
	}

}

func adminBucket() *Bucket {
	if aBuck == nil {
		bucks, _ := ListBuckets(adminUser().ID, adminUser().Key)
		for _, buck := range *bucks {
			if buck.Name == "admin" {
				aBuck = &buck
				return aBuck
			}
		}
		aBuck, _ = CreateBucket(adminUser().ID, adminUser().Key, "admin")
	}
	return aBuck
}
