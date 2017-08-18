package places

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// GetPlaces has a comment
func GetPlaces(writer http.ResponseWriter, request *http.Request) {
	urlString, _ := url.Parse(request.URL.String())
	raw := urlString.RawQuery
	valueMap, _ := url.ParseQuery(raw)
	location := valueMap["location"][0]

	fmt.Println(location)

	writer.Header().Set("Content-Type", "application/json")

	resp, _ := GetPlacesService(location)

	responseByteArray, _ := json.Marshal(resp)
	io.WriteString(writer, string(responseByteArray[:]))
}
