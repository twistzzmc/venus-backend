package logger

import "fmt"

type ConsoleLogger struct{}

func (ConsoleLogger) LogPanic(err any) {
	fmt.Printf("Panic! %+v\n", err)
}

func (ConsoleLogger) LogResponse(status int, resp interface{}) {
	if status < 400 {
		return
	}

	fmt.Printf("[%d] %s\n]", status, resp)
}
