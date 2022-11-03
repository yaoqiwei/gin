package http_error

type ResData map[string]interface{}

type HttpError struct {
	ErrorCode int    `json:"code" example:"400"`
	ErrorMsg  string `json:"msg" example:"status bad request"`
}

var MissingParametersError = HttpError{
	ErrorCode: 5001,
	ErrorMsg:  "缺少参数",
}
