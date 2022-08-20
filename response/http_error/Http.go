package http_error

type HttpError struct {
	ErrorCode int    `json:"code"`
	ErrorMsg  string `json:"msg"`
}
