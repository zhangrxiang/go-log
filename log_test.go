package log_test

import (
	"github.com/zing-dev/go-log"
	"testing"
)

func TestDefault(t *testing.T) {
	log.Info("default")
}

func TestLogLoad(t *testing.T) {
	log.LoadLog(log.ConsoleLog)
	log.Info("hello world")
}

func TestLogLoadMore(t *testing.T) {
	log.LoadLog(log.ConsoleLog, log.SocketLog, log.FileLog)
	log.Info("hello")
}

func TestLoad(t *testing.T) {
	log.LoadConsoleLog()
	log.LoadFileLog()
	log.LoadWebSocketLog()
	log.Info("Info")
	log.Socket("test", "test")
}
