package places

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetPlaces has a comment
func GetPlaces(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	resp, _ := GetPlacesService()

	responseByteArray, _ := json.Marshal(resp)
	io.WriteString(writer, string(responseByteArray[:]))
}
