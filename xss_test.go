package learn_go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  "<p>Ini body coba di post.gohtml <script>alert('Anda di heck ')</script></p>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    ":9090",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateAutoEscapeDisabled(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML("<h1>Ini Adalah Server</h1>"),
	})
}

func TestTemplateAutoEscapeDisabled(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeDisabledServer(t *testing.T) {
	server := http.Server{
		Addr:    ":9090",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateXSS(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML(request.URL.Query().Get("body")),
	})
}

func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090?body=%3Cscript%3Ealert%28%27Samlekom%27%29%3C%2Fscript%3E", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr:    ":9090",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
