package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/golang/glog"
)

// Generate a UUID
func GenerateUUID() string {
	var symbols = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")
	var uuid string
	for i := 0; i < 12; i++ {
		uuid += string(symbols[rand.Intn(len(symbols)-1)])
	}

	return uuid
}

// Upload a file, save and attribute a hash
func upload(w http.ResponseWriter, r *http.Request) {
	glog.Info("Request recieved")

	var uuid = GenerateUUID()
	path := fmt.Sprintf("./storage/%s/", uuid)

	// Prepare to get the file
	file, header, err := r.FormFile("file")
	if err != nil {
		glog.Errorf("Error retrieving file. Error: %s", err.Error())
		return
	}
	defer func() {
		file.Close()
		glog.Infof(`File "%s" closed.`, header.Filename)
	}()

	// Creates directory with UUID
	_, err = os.Stat(path)
	for !os.IsNotExist(err) {
		uuid = GenerateUUID()
		path := fmt.Sprintf("./storage/%s/", uuid)
		_, err = os.Stat(path)
	}

	if err := os.Mkdir(path, 0777); err != nil {
		glog.Error("Error saving file on server...")
	}

	// Build and Write the file.
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		glog.Errorf("Error retrieving file content.\n Error: %s", err.Error())
	}
	err = os.WriteFile(path+header.Filename, bytes, 0777)
	if err != nil {
		glog.Errorf("Error writing file.\n Error: %s", err.Error())
	}

	// To review all possible status codes
	w.WriteHeader(200)
	fmt.Fprintf(w, "Successfully Uploaded File\n http://localhost:8000/%s\n", uuid)
}

// To implement
func getFile(w http.ResponseWriter, r *http.Request) {

}

// Route handling, logging and application serving
func main() {
	// Random seed creation
	rand.Seed(time.Now().Unix())

	// Flags for the leveled logging
	flag.Usage = func() { fmt.Println("USAGE: To implement") }
	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "WARNING")
	flag.Set("v", "2")
	flag.Parse()

	glog.Flush()

	// Routing
	http.HandleFunc("/", upload)
	http.HandleFunc("/:hash", getFile)
	http.ListenAndServe(":8000", nil)
}
