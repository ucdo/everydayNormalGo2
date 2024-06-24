package handler

import (
	"Gonic/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTodoTable(c *gin.Context) {
	db, err := model.Conn()
	if err != nil {
		errStr := fmt.Sprintf("Connection error: %v", err)
		c.JSON(200, gin.H{
			"message": errStr,
		})
		return
	}

	ts := model.Task{
		Name:    "Add first",
		Context: "test create table",
		Status:  0,
		Model:   gorm.Model{},
	}

	tx := db.Create(&ts)

	if tx.Error != nil {
		errStr := fmt.Sprintf("CreateTable with error: %v", tx.Error)
		c.JSON(200, gin.H{
			"message": errStr,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
	})

	return
}

func TodoList(c *gin.Context) {
	db, err := model.Conn()

	if err != nil {
		errStr := fmt.Sprintf("Connection error: %v", err)
		c.JSON(200, gin.H{
			"message": errStr,
		})
		return
	}

	var res []model.Task
	tx := db.First(&res)
	if tx.Error != nil {
		c.JSON(200, gin.H{
			"message": tx.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": tx.Error,
		"data":    res,
	})

}
