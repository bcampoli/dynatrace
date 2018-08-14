package notification

// Listen TODO: documentation
func Listen(config *Config, handler Handler) {
	newListener(config, handler).listen()
}
