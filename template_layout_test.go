package learn_go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./template/header.gohtml",
		"./template/footer.gohtml",
		"./template/layout.gohtml",
	))
	t.ExecuteTemplate(writer, "layout.gohtml", map[string]interface{}{
		"Name":  "Adhitya",
		"Title": "Template layout",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateLayoutDefine(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./template/header_define.gohtml",
		"./template/footer_define.gohtml",
		"./template/layout_define.gohtml",
	))
	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Name":  "Adhitya",
		"Title": "Template layout Define",
	})
}

func TestTemplateLayoutDefine(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090", nil)
	recorder := httptest.NewRecorder()

	TemplateLayoutDefine(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
