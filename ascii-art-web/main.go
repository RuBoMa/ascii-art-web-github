package main

import (
	"html/template"
	"log"
	"net/http"
	"pkg/ascii"
)

type PageData struct {
	AsciiArt string
}

func main() {
	// server static files like css or images
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	//route handler for home page
	http.HandleFunc("/", handler)

	log.Println("server is running on http://localhost:8080")
	// start the server on port 8080 and listen to incoming requests
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	//loads home.html from templates folder
	tmpl, _ := template.ParseFiles("templates/home.html")
	data := PageData{}
	if r.Method == http.MethodPost {
		input := r.FormValue("userText")
		banner := r.FormValue("style")

		data.AsciiArt = ascii.PrintAsciiArt(input, banner)

	}
	tmpl.Execute(w, data)

}
