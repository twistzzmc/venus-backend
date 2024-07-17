package logger

var Logger ILogger

type ILogger interface {
	LogPanic(err any)
	LogResponse(status int, resp interface{})
}
