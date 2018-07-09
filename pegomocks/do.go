package pegomocks

//go:generate pegomock generate --use-experimental-model-gen logger
type logger interface {
	Info(msg string)
}

func Do(logger logger) {
	logger.Info("running do function")
}
