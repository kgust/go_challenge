package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
	"path/filepath"
	"os"
	"encoding/json"
)

func main() {

	var cwd, _ = os.Getwd()
	path := filepath.Join(cwd, "challenge_02/templates/*.html")
	tmpl := template.Must(template.ParseGlob(path))

	// template handler
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		params := struct {}{}
		err := tmpl.ExecuteTemplate(w, "hi.html", params)
		if err != nil {
			fmt.Print(w, err)
		}
	})

	// POST handler
	http.HandleFunc("/save", func(w http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		var t struct{}
		err := decoder.Decode(&t)
		if err != nil {
			fmt.Print(w, err)
		}

		defer req.Body.Close()
		log.Println(t)
	})

	// file handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(cwd, "./challenge_02/templates/", r.URL.Path[1:])
		if r.URL.Path == "/" {
			path += "/index.html"
		}
		fmt.Println(path)

		http.ServeFile(w, r, path)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
