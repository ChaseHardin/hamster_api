package helloWorld

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func BeforeEach()  *httptest.ResponseRecorder  {
	req, _ := http.NewRequest("GET", "/greeting", nil)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GreetingHandler)
	handler.ServeHTTP(recorder, req)

	return recorder
}

func TestShouldReturnHelloWorld(t *testing.T) {
	setup := BeforeEach()

	expected := `Hello World!`
	actual := setup.Body.String()

	if actual != expected {
		t.Errorf("Invalid message was returned. Expected: %v Actual: %v", expected, actual)
	}
}

func TestShouldReturnOkayResponse(t *testing.T) {
	setup := BeforeEach()

	expected := http.StatusOK
	actual := setup.Code

	if expected != actual {
		t.Errorf("handler returned wrong status code: Expected: %v Actual: %v", expected, actual)
	}
}

func TestShouldReturnJsonContentType(t *testing.T) {
	setup := BeforeEach()

	expected := "application/json"
	actual := setup.Header().Get("Content-Type")

	if expected != actual {
		t.Errorf("handler retuned wrong content type: Expected: %v Actual: %v", expected, actual)
	}
}