package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/golang/glog"
)

// Upload a file, save and attribute a hash
func upload(w http.ResponseWriter, r *http.Request) {
	glog.Info("Request recieved")

	file, header, err := r.FormFile("file")
	if err != nil {
		glog.Errorf("Error retrieving file. Error: %s", err.Error())
		return
	}
	defer func() {
		file.Close()
		glog.Infof(`File "%s" closed.`, header.Filename)
	}()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		glog.Errorf("Error retrieving file content.\n Error: %s", err.Error())
	}

	err = os.WriteFile(fmt.Sprintf("./storage/%s", header.Filename), bytes, 0644)
	if err != nil {
		glog.Errorf("Error writing file.\n Error: %s", err.Error())
	}

	//// Building the file and reading all bytes from array
	//tempFile, err := ioutil.TempDir("storage", header.Filename)
	//bytes, err := ioutil.ReadAll(file)
	//if err != nil {
	//	glog.Errorf("Error retrieving file. Error: %s", err.Error())
	//}
	//defer tempFile.Close()
	//tempFile.Write(bytes)

	fmt.Fprintf(w, "Successfully Uploaded File\n")

}

// Get Alder-32 Hash.
func hash(w http.ResponseWriter, r *http.Request) {

}

// Route handling and application serving
func main() {
	flag.Usage = func() { fmt.Println("USAGE: To implement") }
	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "WARNING")
	flag.Set("v", "2")

	flag.Parse()

	glog.Flush()
	http.HandleFunc("/", upload)
	http.HandleFunc("/:hash", hash)
	http.ListenAndServe(":8000", nil)
}
