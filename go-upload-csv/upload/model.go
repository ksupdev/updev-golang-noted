package upload

type Resp struct {
	Result       string `json:"result"`
	IsError      bool   `json:"is_error"`
	ErrorMessage string `json:"error_message"`
}
