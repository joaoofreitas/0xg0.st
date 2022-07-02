package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/golang/glog"
)

// Dead simple router that just does the **perform** the job
func router(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		upload(w, r)
		break
	default:
		getFile(w, r)
		break
	}
}

// Route handling, logging and application serving
func main() {
	// Random seed creation
	rand.Seed(time.Now().Unix())

	// Flags for the leveled logging
	flag.Usage = func() { fmt.Println("USAGE: To implement") }
	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "INFO")
	flag.Set("v", "2")
	flag.Parse()

	glog.Flush()

	// Routing
	http.HandleFunc("/", router)
	http.ListenAndServe(":8000", nil)
}
