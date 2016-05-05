package api

import (
	"fmt"
	"../backends/backendtest"
	"../models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateObject(t *testing.T) {
	br := backendtest.NewBackendRecorder()
	models.Initialize("dbname=test sslmode=disable", br)
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Userid", adminUser().ID)
	req.Header.Add("Key", adminUser().Key)
	req.Header.Add("Bucketname", "admin")
	req.Header.Add("Objectalias", "test")
	req.Body = ioutil.NopCloser(strings.NewReader("body"))

	w := httptest.NewRecorder()
	createObject(w, req)
	if w.Code != 201 {
		t.Error("File not created:", w)
	}
	if string(br.Written) != "body" {
		t.Error("the body didnt get written to the backend")
	}
}

func TestListObject(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Userid", adminUser().ID)
	req.Header.Add("Key", adminUser().Key)
	req.Header.Add("Bucketname", "admin")

	w := httptest.NewRecorder()
	listObjects(w, req)
	if w.Code != 200 {
		t.Error("Cannot list objects:", w)
	}
}

func TestGetObject(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/objects/test", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Userid", adminUser().ID)
	req.Header.Add("Key", adminUser().Key)
	req.Header.Add("Bucketname", "admin")

	w := httptest.NewRecorder()
	getObject(w, req)
	if w.Code != 200 {
		t.Error("I couldnt get bucket:", w)
	}
	if w.Body.String() != "object data" {
		t.Error("I couldnt read the body")
	}

}

func TestDeleteObject(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://example.com/objects/test", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Userid", adminUser().ID)
	req.Header.Add("Key", adminUser().Key)
	req.Header.Add("Bucketname", "admin")

	w := httptest.NewRecorder()
	deleteObject(w, req)
	if w.Code != 202 {
		t.Error("object not deleted:", w)
	}
}
