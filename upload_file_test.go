package learn_go_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "upload.form.gohtml", nil)
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	file, fileHeader, err := request.FormFile("uploadFile")
	if err != nil {
		panic(err)
	}
	fileDestionation, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestionation, file)
	if err != nil {
		panic(err)
	}
	name := request.FormValue("name")
	myTemplates.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    ":9090",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed "resources/Murphy Radio.jpg"
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	writer.WriteField("name", "Adhitya")
	file, _ := writer.CreateFormFile("uploadFile", "CONTOH_UPLOAD.jpg")
	file.Write(uploadFileTest) //menggunakan bantuan embed
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:9090/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyRespon, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyRespon))
}
