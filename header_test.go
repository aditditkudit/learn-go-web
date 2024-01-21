package learn_go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	content_type := request.Header.Get("content-type")
	fmt.Fprint(writer, content_type)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:9090", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-POWERED-BY", "Kudts.Corp")
	fmt.Fprint(writer, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:9090", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Header.Get("X-POWERED-BY"))
}
