package todo

import "github.com/gin-gonic/gin"

type MyContext struct {
	// Context *gin.Context
	// Compositions
	*gin.Context
}

func NewMyContext(c *gin.Context) *MyContext {
	return &MyContext{c}
}

func (c *MyContext) Bind(v interface{}) error {
	return c.Context.ShouldBindJSON(v)
}
func (c *MyContext) TransactionID() string {
	return c.Request.Header.Get("TransactionID")
}
func (c *MyContext) Audience() string {
	if aud, ok := c.Get("aud"); ok {
		if s, ok := aud.(string); ok {
			return s
		}
	}
	return ""
}
func (c *MyContext) Id() string {
	return c.Param("id")
}
func (c *MyContext) JSON(code int, v interface{}) {
	c.Context.JSON(code, v)
}

func NewGinHandler(handler func(Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewMyContext(c))
	}
}
