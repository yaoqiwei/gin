package http_error

type ResData map[string]interface{}

type HttpError struct {
	ErrorCode int    `json:"code" example:"400"`
	ErrorMsg  string `json:"msg" example:"status bad request"`
}

var (
	MissingParametersError = HttpError{
		ErrorCode: 5001,
		ErrorMsg:  "缺少参数",
	}

	NoRedisKey = HttpError{
		ErrorCode: 604,
		ErrorMsg:  "没有redis key",
	}

	FrequentOperations = HttpError{
		ErrorCode: 605,
		ErrorMsg:  "操作太频繁，请稍后再试",
	}

	NoAESKey = HttpError{
		ErrorCode: 606,
		ErrorMsg:  "没有aes key",
	}
)
