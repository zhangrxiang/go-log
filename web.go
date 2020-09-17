package log

import (
	"encoding/json"
	"github.com/kataras/neffos"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var socket *webSocket

type Message struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Response struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type webSocket struct {
	ws      *neffos.Server
	Content chan []byte
	hook    logrus.Hook
}

type WebJSONFormatter struct {
	f logrus.JSONFormatter
}

func (wj *WebJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	delete(entry.Data, "file")
	entry.Data["type"] = "log"
	return wj.f.Format(entry)
}

func LoadWebSocketLog() {
	logger.AddHook(DefaultWebSocketLog())
}

func DefaultWebSocketLog() *lfshook.LfsHook {
	socket = &webSocket{
		Content: make(chan []byte),
	}
	return lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  socket,
		logrus.WarnLevel:  socket,
		logrus.ErrorLevel: socket,
		logrus.FatalLevel: socket,
		logrus.PanicLevel: socket,
	}, &WebJSONFormatter{f: logrus.JSONFormatter{
		TimestampFormat: DefaultLocalTimeDateFormat,
	}})
}

func Socket(t string, args interface{}) {
	data, _ := json.Marshal(Response{
		Data: args,
		Type: t,
	})
	select {
	case socket.Content <- data:
	default:
	}
}

func Read(ws *neffos.Server) {
	if socket == nil {
		return
	}

	for {
		select {
		case content := <-socket.Content:
			ws.Broadcast(nil, neffos.Message{
				Body:     content,
				IsNative: true,
			})
		}
	}
}

func (w *webSocket) Write(p []byte) (n int, err error) {
	select {
	case w.Content <- p:
	default:
	}
	return len(p), nil
}
