package main

import (
	"fmt"
	"log"
	"net/http"
)

var port = ":4000"

func main() {
	http.HandleFunc("/", ServeFiles)
	http.HandleFunc("/script.js", JsFile)
	http.HandleFunc("/style.css", CssFile)
	http.HandleFunc("/media/background.png", backgroundImg)
	fmt.Println("Serving @ ", "http://127.0.0.1" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func backgroundImg(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case "GET":
		path := r.URL.Path
		fmt.Println(path)
		if path == "/" {
			path = "/background.png"
		} else {
			path = "."+path
		}
		http.ServeFile(w, r, path)
	default:
			fmt.Fprintln(w, "Request type other than GET or POST not supported")
	}
}

func CssFile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		path := r.URL.Path

		fmt.Println(path)

		if path == "/style.css"{
			path = "./static/style.css"
		} else {
			path = "."+path
		}

		http.ServeFile(w, r, path)
	default:
		fmt.Fprintln(w, "Request type other than GET or POST not supported")
	}
}

func JsFile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		path := r.URL.Path

		fmt.Println(path)

		if path == "/script.js"{
			path = "./static/script.js"
		} else {
			path = "."+path
		}

		http.ServeFile(w, r, path)
	default:
		fmt.Fprintln(w, "Request type other than GET or POST not supported")
	}
}

func ServeFiles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		path := r.URL.Path

		fmt.Println(path)

		if path == "/"{
			path = "./static/index.html"
		} else {
			path = "."+path
		}

		http.ServeFile(w, r, path)
	default:
		fmt.Fprintln(w, "Request type other than GET or POST not supported")
	}
}
