package models

import (
	"../backends"
	"os"
	"regexp"
	"testing"
)

func TestMain(m *testing.M) {
	Initialize("dbname=test sslmode=disable", backends.NewLocalStorage("./data"))
	os.Exit(m.Run())

}

func TestGenerateID(t *testing.T) {
	key := generateID()
	matched, err := regexp.MatchString(`^\w\w\w\w\w\w\w\w-\w\w\w\w-\w\w\w\w-\w\w\w\w-\w\w\w\w\w\w\w\w\w\w\w\w$`, key)

	if err != nil || !matched {
		t.Error("the id was not a uuid")
	}
}

func TestInitialize(t *testing.T) {
	err := Initialize("badstring", backends.NewLocalStorage("./data"))
	if err == nil {
		t.Error("Initialize with a bad string should return an error")
	}

	err = Initialize("dbname=test sslmode=disable", backends.NewLocalStorage("./data"))
	if err != nil {
		t.Error("Initialize should not return an error")
	}
}
