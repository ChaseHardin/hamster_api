package hello_world

import (
	"net/http"
	"io"
)

func SetWriterJson(writer http.ResponseWriter) http.ResponseWriter {
	writer.Header().Set("Content-Type", "application/json")
	return writer
}

func GreetingHandler(writer http.ResponseWriter, request *http.Request) {
	writer = SetWriterJson(writer)

	io.WriteString(writer,"Hello World!")
}