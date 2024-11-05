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
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Printf("Error loading template: %v", err)
		renderErrorPage(w)
		return
	}
	data := PageData{}
	if r.Method == http.MethodPost {
		input := r.FormValue("userText")
		banner := r.FormValue("style")
		log.Println("Selected input: " + input + " and banner: " + banner)

		asciiArt, err := ascii.PrintAsciiArt(input, banner)
		if err != nil {
			renderErrorPage(w)
			return
		}
		data.AsciiArt = asciiArt

	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		renderErrorPage(w)
	}

}

func renderErrorPage(w http.ResponseWriter) {
	// Set the HTTP status code to 500
	w.WriteHeader(http.StatusInternalServerError)

	// Parse and execute the 500 error template
	tmpl, err := template.ParseFiles("templates/500.html")
	if err != nil {
		// If template loading fails, return a simple fallback message
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error loading 500 error template: %v", err)
		return
	}

	// Render the template without additional data
	tmpl.Execute(w, nil)
}
