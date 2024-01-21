package learn_go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello Broks")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParameterWeb(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090/hello?name=adhitya", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("firstName")
	lastName := request.URL.Query().Get("lastName")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParameterWeb(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090/hello?firstName=adhitya&lastName=Kudit", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func MultipleParameterValue(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]

	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultipleParameterValue(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090/hello?name=adhitya&name=Keren", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValue(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
