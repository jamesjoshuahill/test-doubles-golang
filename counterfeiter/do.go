package counterfeiter

//go:generate counterfeiter . logger
type logger interface {
	Info(msg string)
}

func Do(logger logger) {
	logger.Info("running do function")
}
