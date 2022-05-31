package todo

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

type Todo struct {
	Title     string `json:"text" binding:"required"`
	ID        uint   `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Todo) TableName() string {
	return "todos"
}

type storer interface {
	New(*Todo) error
	Find(*[]Todo) error
	Delete(*Todo, int) error
}

type TodoHandler struct {
	// db *gorm.DB
	store storer
}

type Context interface {
	Bind(interface{}) error
	TransactionID() string
	Audience() string
	Id() string
	JSON(int, interface{})
}

func NewTodoHandler(store storer) *TodoHandler {
	return &TodoHandler{store: store}
}

func (t *TodoHandler) NewTask(c Context) {
	var todo Todo
	// if err := c.ShouldBindJSON(&todo); err != nil {
	if err := c.Bind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	if todo.Title == "sleep" {
		// transactionID := c.Request.Header.Get("TransactionID")
		transactionID := c.TransactionID()
		// aud, _ := c.Get("aud")
		aud := c.Audience()
		log.Println(transactionID, aud, "not allowed")
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": "not allowed",
		// })

		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "not allowed",
		})

		return
	}

	err := t.store.New(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"ID": todo.ID,
	})
}

func (t *TodoHandler) List(c Context) {
	var todos []Todo
	err := t.store.Find(&todos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (t *TodoHandler) Remove(c Context) {
	// idParam := c.Param("id")
	idParam := c.Id()

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	err = t.store.Delete(&Todo{}, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}
