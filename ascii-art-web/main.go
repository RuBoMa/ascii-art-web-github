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

	tmpl, err := template.ParseFiles("templates/home.html", "templates/notfound.html", "templates/500.html")
	if err != nil {
		log.Printf("Error loading template: %v", err)
		serverErrorHandler(w, tmpl)
		return
	}
	data := PageData{}

	if r.URL.Path != "/" {
		notFoundHandler(w, tmpl)
		return
	}
	if r.Method == http.MethodPost {
		input := r.FormValue("userText")
		banner := r.FormValue("style")
		log.Println("Selected input: " + input + " and banner: " + banner)

		asciiArt, err := ascii.PrintAsciiArt(input, banner)
		if err != nil {
			serverErrorHandler(w, tmpl)
			return
		}
		data.AsciiArt = asciiArt

	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		serverErrorHandler(w, tmpl)
	}

}

// 500 internal server error handler
func serverErrorHandler(w http.ResponseWriter, tmpl *template.Template) {
	// Set the HTTP status code to 500
	w.WriteHeader(http.StatusInternalServerError)

	// Render the template without additional data
	tmpl.ExecuteTemplate(w, "500.html", nil)
}

// 404 page not found handler
func notFoundHandler(w http.ResponseWriter, tmpl *template.Template) {
	w.WriteHeader(http.StatusNotFound)
	tmpl.ExecuteTemplate(w, "notfound.html", nil)
}
