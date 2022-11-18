package http_error

var PasswordCountError = HttpError{
	ErrorCode: 1008,
	ErrorMsg:  "密码6-12位数字与字母",
}
