package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(w, err)
			return
		}
		defer file.Close()
		filename := fmt.Sprintf("%s.jpg", handler.Filename)
		out, err := os.Create(filename)
		if err != nil {
			fmt.Println(w, err)
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Println(w, err)
			return
		}
		fmt.Fprintf(w, "OK")
	} else {
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	}
}

func main() {
	http.HandleFunc("/upload", handleUpload)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
