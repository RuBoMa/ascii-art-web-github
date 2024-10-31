package main

import (
	"fmt"
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
	http.HandleFunc("/", homePage)
	

	fmt.Println("server is running on http://localhost:8080")
	// start the server on port 8080 and listen to incoming requests
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func homePage(w http.ResponseWriter, r *http.Request) {
	//loads home.html from templates folder
	tmpl, _ := template.ParseFiles("templates/home.html")
	data := PageData{}
	if r.Method == http.MethodPost {
		input := r.FormValue("text")
		banner := r.FormValue("style")

		asciiArt := ascii.PrintAsciiArt(input, banner)

		data.AsciiArt = asciiArt

	}
	tmpl.Execute(w, data)

}
