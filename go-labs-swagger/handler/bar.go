package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	barDomain "puza.dev/go-labs-swagger/bar"
)

type bar struct{}

func NewBar() *bar {
	return &bar{}
}

func (b *bar) GetById(c *gin.Context) {
	// filterStatus, hasFilter := c.GetQuery("status")
	//if hasFilter {

	barId := c.Param("id")
	if len(barId) == 0 || barId == "" {
		c.JSON(http.StatusBadRequest, fail(fmt.Errorf("please fill id").Error()))
	}

	c.JSON(http.StatusOK, success("200", barDomain.MBar{BarId: barId, BarName: " test "}))

}

func (b *bar) Create(c *gin.Context) {

}

func (b *bar) Update(c *gin.Context) {

}

func (b *bar) Delete(c *gin.Context) {

}
