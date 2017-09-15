package esalert

// ConfigError 配置错误
type ConfigError struct {
	message string
}

func (configError ConfigError) Error() string {
	return configError.message
}
