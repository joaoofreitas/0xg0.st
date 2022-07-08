package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUpload(t *testing.T) {
	var buf bytes.Buffer

	w := multipart.NewWriter(&buf)
	part, _ := w.CreateFormFile("file", "test.txt")
	part.Write([]byte("this is the file content"))
	w.Close()

	req, _ := http.NewRequest("POST", "/", &buf)
	req.Header.Set("Content-Type", w.FormDataContentType())
	res := httptest.NewRecorder()

	upload(res, req)
	if res.Result().StatusCode != 200 {
		t.Fatal("expected 200 got", res.Result().StatusCode)
	}

	b, _ := io.ReadAll(res.Body)

	lines := strings.Split(string(b), "\n")
	downloadPathComponents := strings.Split(lines[1], "/")

	hash := downloadPathComponents[len(downloadPathComponents)-1]

	os.RemoveAll(filepath.Join("storage", hash))
}

func TestDownload(t *testing.T) {

	os.MkdirAll("storage/sillyhash", 0777)

	os.WriteFile("storage/sillyhash/test.txt", []byte("this is the file content"), 0777)

	req := httptest.NewRequest("GET", "/sillyhash", nil)
	res := httptest.NewRecorder()

	getFile(res, req)
	if res.Result().StatusCode != 200 {
		t.Fatal("expected 200 got", res.Result().StatusCode)
	}

	os.RemoveAll(filepath.Join("storage", "sillyhash"))

}
