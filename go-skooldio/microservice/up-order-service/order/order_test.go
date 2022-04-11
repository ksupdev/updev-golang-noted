package order

// go test . -v
// go test -v
// go test

import (
	"errors"
	"net/http"
	"testing"
)

type fakeContext struct {
	channel  string
	code     int
	response map[string]string
}

func (f fakeContext) Order() (Order, error) {
	return Order{
		SalesChannel: f.channel,
	}, nil
}

func (c *fakeContext) JSON(code int, v interface{}) {
	c.code = code
	c.response = v.(map[string]string)
}

// go test -run TestOnlyAcceptOnlin -v
func TestOnlyAcceptOnlinechannel(t *testing.T) {
	handler := &Handler{
		channel: "Online",
	}

	c := &fakeContext{
		channel: "Offline",
	}
	handler.Order(c)

	want := "Offline is not accepted"

	if want != c.response["message"] {
		t.Errorf("%q is expected but got %q\n", want, c.response["message"])
	}

}

// go test -run TestOnlyAcceptOfflineChannel -v
func TestOnlyAcceptOfflineChannel(t *testing.T) {
	handler := &Handler{
		channel: "Offline",
	}

	c := &fakeContext{
		channel: "Online",
	}
	handler.Order(c)

	want := "Online is not accepted"
	if want != c.response["message"] {
		t.Errorf("%q is expected but got %q \n", want, c.response["message"])
	}
}

type fakeContextBadRequest struct {
	code     int
	response map[string]string
}

func (c fakeContextBadRequest) Order() (Order, error) {
	return Order{}, errors.New("went wrong")
}

func (c *fakeContextBadRequest) JSON(code int, v interface{}) {
	c.code = code
	c.response = v.(map[string]string)
}

func TestBadRequestOrderWentWrong(t *testing.T) {
	handler := &Handler{}

	c := &fakeContextBadRequest{}
	handler.Order(c)

	want := http.StatusBadRequest
	if want != c.code {
		t.Errorf("%d status code is expected but got %v \n", want, c.code)
	}

}

type fakeContextBadRequestWithChannel struct {
	channel         string
	jsonCalledCount int
}

func (c fakeContextBadRequestWithChannel) Order() (Order, error) {
	return Order{
		SalesChannel: c.channel,
	}, errors.New("went wrong")
}

func (c *fakeContextBadRequestWithChannel) JSON(code int, v interface{}) {
	c.jsonCalledCount++
}

func TestOnlyCallCalledJSONOneTime(t *testing.T) {
	handler := &Handler{
		channel: "Offline",
	}

	c := &fakeContextBadRequestWithChannel{}
	handler.Order(c)

	want := 1
	if want != c.jsonCalledCount {
		t.Errorf("it should called one time but got %d times", c.jsonCalledCount)
	}

}

type spyStore struct {
	wasCalled bool
}

func (s *spyStore) Save(Order) error {
	s.wasCalled = true
	return nil
}

func TestOrderWasSaved(t *testing.T) {
	spy := &spyStore{}
	handler := &Handler{
		channel: "Online",
		store:   spy,
	}

	c := &fakeContext{
		channel: "Online",
	}

	handler.Order(c)

	want := true
	if want != spy.wasCalled {
		t.Errorf("it should store data ")
	}

}

type failStore struct{}

func (failStore) Save(Order) error {
	return errors.New("")
}

func TestOrderFailAtSave(t *testing.T) {
	store := &failStore{}
	handler := &Handler{
		channel: "Online",
		store:   store,
	}

	c := &fakeContext{
		channel: "Online",
	}

	handler.Order(c)

	want := http.StatusInternalServerError

	if want != c.code {
		t.Errorf("%d is expected but got %d \n", want, c.code)
	}
}

func TestOrderIsOK(t *testing.T) {
	store := &spyStore{}
	handler := &Handler{
		channel: "Online",
		store:   store,
	}

	c := &fakeContext{
		channel: "Online",
	}

	handler.Order(c)

	// v = "value"
	// ok = true, false , true : found data, false : not found data

	if _, ok := c.response["message"]; !ok {
		t.Errorf("message key is expected")
	}

}
