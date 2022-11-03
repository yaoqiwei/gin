package http_error

var JwtError = HttpError{
	ErrorCode: 700,
	ErrorMsg:  "用户登录状态已失效，请重新登录！",
}
