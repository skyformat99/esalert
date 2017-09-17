package esalert

type baseError struct {
	Message string
}

func (baseError baseError) Error() string {
	return baseError.Message
}

func (baseError baseError) SetError(error string) {
	 baseError.Message = error
}

// ConfigError 配置错误
type ConfigError struct {
	Message string
	baseError
}

// RequestError 请求错误
type RequestError struct {
	Message string
	baseError
}

type NotFoundError struct {
	Message string
	baseError
}
