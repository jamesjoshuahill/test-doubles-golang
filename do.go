package doubles

//go:generate charlatan -output charlatan/fake_logger.go -package charlatan logger
//go:generate counterfeiter -o counterfeiter/fake_logger.go . logger
//go:generate minimock -i github.com/jamesjoshuahill/test-doubles-golang.logger -o minimocks -s _mock.go
//go:generate pegomock generate --use-experimental-model-gen -o pegomock/mock_logger.go --package pegomock logger
type logger interface {
	Info(msg string)
}

func Do(logger logger) {
	logger.Info("running do function")
}
