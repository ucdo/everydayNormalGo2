package model

import (
	"gorm.io/gorm"
	"testing"
)

func TestConn(t *testing.T) {
	t.Helper()
	_, err := Conn()
	if err != nil {
		t.Fatal(err)
	}

	err = db.AutoMigrate(&Task{})
	if err != nil {
		t.Fatal("create table failed")
	}

	tx := db.Create(&Task{
		Name:    "testConn",
		Context: "testConn",
		Status:  0,
		Model:   gorm.Model{},
	})

	if tx.Error != nil {
		t.Fatal(tx.Error.Error())
	}

}
