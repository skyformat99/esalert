package esalert

// ConfigError 配置错误
type ConfigError struct {
	message string
}

func (configError ConfigError) Error() string {
	return configError.message
}

// RequestError 请求错误
type RequestError struct {
	message string
}

func (requestError RequestError) Error() string {
	return requestError.message
}
