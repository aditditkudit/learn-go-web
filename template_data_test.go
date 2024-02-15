package learn_go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./template/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Percobaan Template Data Map",
		"Name":  "Adhitya",
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Page struct {
	Title string
	Name  string
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./template/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Halo Halo",
		Name:  "Jancok",
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Address struct {
	Street string
}

type People struct {
	Name    string
	Job     string
	Address Address
}

func TemplateDataStructNested(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./template/nested.gohtml"))
	t.ExecuteTemplate(writer, "nested.gohtml", People{
		Job:  "Raja Ngoding",
		Name: "Jancok",
		Address: Address{
			Street: "Jalan Silicon Valley",
		},
	})
}

func TestTemplateDataStructNested(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStructNested(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateDataMapNested(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./template/nested.gohtml"))
	t.ExecuteTemplate(writer, "nested.gohtml", map[string]interface{}{
		"Job":  "Raja Gacor",
		"Name": "Dewa Dewa",
		"Address": map[string]interface{}{
			"Street": "Di Internet",
		},
	})
}

func TestTemplateDataMapNested(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMapNested(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
