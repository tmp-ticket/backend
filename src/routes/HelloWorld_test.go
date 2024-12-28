package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test case to make sure the /HelloWorld route is working and returning "Hello World!"
func TestGetHelloWorld(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/HelloWorld", nil)
	response := httptest.NewRecorder()

	HelloWorld(response, request)

	if response.Body.String() != "Hello World!" {
		t.Errorf("got: %s, wanted %s", response.Body.String(), "Hello World!")
	}
}
