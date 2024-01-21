package learn_go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Field Name is Empty")
	} else {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponseCodeBadRequest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}

func TestResponseCodeOkRequest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090?name=Adhitya", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}

func ResponceCodeWithFormPost(writer http.ResponseWriter, request *http.Request) {
	lastName := request.PostFormValue("last_name")
	if lastName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Field Name is Empty")
	} else {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Hello %s", lastName)
	}
}

func TestResponseCodeOkFormPost(t *testing.T) {
	requestBody := strings.NewReader("last_name=Jihan")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:9090", requestBody)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	ResponceCodeWithFormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
