package minimock

//go:generate minimock -i github.com/jamesjoshuahill/test-doubles-golang/minimock.logger -o .
type logger interface {
	Info(msg string)
}

func Do(logger logger) {
	logger.Info("running do function")
}
