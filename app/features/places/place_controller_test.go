package places

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BeforeEachGetPlaces() *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", "/places/?location=41.584272,-93.6376587", nil)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetPlaces)
	handler.ServeHTTP(recorder, req)

	return recorder
}

func TestShouldReturnPlaces(t *testing.T) {
	assert := assert.New(t)

	setup := BeforeEachGetPlaces()

	actual := setup.Body.String()
	assert.NotNil(actual)
}
