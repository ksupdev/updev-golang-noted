package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Title string `json:"text"`
}

func (Todo) Tablename() string {
	return "todos "
}

type TodoHandler struct {
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{}
}

func (t *TodoHandler) NewTask(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create with gorm

}
