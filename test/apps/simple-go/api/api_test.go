package api

import (
	"fmt"
	"../backends/backendtest"
	"../models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var aUser *models.User

func TestMain(m *testing.M) {
	models.Initialize("dbname=test sslmode=disable", backendtest.NewBackendRecorder())
	os.Exit(m.Run())
}

func TestAdminAccess(t *testing.T) {
	handler := adminAccess(listUsers)
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		fmt.Println(err)
	}

	w := httptest.NewRecorder()
	handler(w, req)
	if w.Code != 404 {
		t.Error("the admin access let someone through without access")
	}

	req.Header.Add("Userid", adminUser().ID)
	req.Header.Add("Key", adminUser().Key)
	w = httptest.NewRecorder()
	handler(w, req)
	if w.Code != 200 {
		t.Error("the admin access didnt let someone through with access")
	}

}

func adminUser() *models.User {
	if aUser == nil {
		users, _ := models.ListUsers()
		for _, usr := range *users {
			if usr.Admin {
				aUser = &usr
				return aUser
			}
		}
	}
	return aUser
}
