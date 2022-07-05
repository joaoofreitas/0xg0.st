package main

import (
	"flag"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang/glog"
)

var port *uint64
var tmpl *template.Template
var ip *string

// Dead simple router that just does the **perform** the job
func router(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.Contains(r.Header.Get("Content-type"), "multipart/form-data"):
		upload(w, r)
		break
	case uuidMatch.MatchString(r.URL.Path):
		getFile(w, r)
		break
	default:
		home(w, r)
		break
	}
}

// Route handling, logging and application serving
func main() {
	// Random seed creation
	rand.Seed(time.Now().Unix())

	// Home template initalization
	tmpl = template.Must(template.ParseFiles("./templates/index.html"))
	// Flags for the leveled logging

	ip = flag.String("i", "0.0.0.0", "ip")
	port = flag.Uint64("p", 8000, "port")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "USAGE: ./0xg0.st -p=8080 -stderrthreshold=[INFO|WARNING|FATAL] -log_dir=[string]\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()
	glog.Flush()

	// Routing
	http.HandleFunc("/", router)
	http.ListenAndServe(fmt.Sprintf("%s:%d",*ip,*port), nil)
}
