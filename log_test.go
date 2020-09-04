package log_test

import (
	"github.com/zhangrxiang/go-log"
	"testing"
)

func TestName(t *testing.T) {
	log.Default()
	log.LoadFileLog()
	log.Info("Info")
}
