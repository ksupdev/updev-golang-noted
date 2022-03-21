package order

import "net/http"

type Handler struct {
	channel string
}

type Context interface {
	Order() (Order, error)
	JSON(int, interface{})
}

func (h *Handler) Order(c Context) {
	order, err := c.Order()
	if err != nil {

	}

	if order.SalesChannel != h.channel {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Offline is not accepted",
		})
	}
}
