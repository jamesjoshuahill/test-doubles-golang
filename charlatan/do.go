package charlatan

type logger interface {
	Info(msg string)
}

func Do(logger logger) {
	logger.Info("running do function")
}
