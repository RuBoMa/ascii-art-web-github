package main

import (
	"html/template"
	"log"
	"net/http"
	"pkg/ascii"
	"strings"
)

type PageData struct {
	AsciiArt string
}

func main() {
	// Serving static files like css or images
	fileServer := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fileServer))

	// Handling the GET request to the root URL ("/")
	http.HandleFunc("/", handler)

	log.Println("server is running on http://localhost:8080")
	// start the server on port 8080 and listen to incoming requests
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/home.html", "templates/404.html", "templates/500.html", "templates/400.html")
	if err != nil {
		log.Printf("Error loading template: %v", err)
		serverErrorHandler(w, tmpl)
		return
	}
	// Creating dynamic data
	data := PageData{}

	if !(r.URL.Path == "/" || r.URL.Path == "/ascii-art") {
		notFoundHandler(w, tmpl)
		return
	}
	if r.Method == http.MethodPost {
		input := r.FormValue("userText")
		log.Println("Selected input: " + input)

		banner := r.FormValue("style")
		log.Println("Selected banner: " + banner)

		cleanInput, valid := ascii.ValidInput(input)
		bannerCont, err := ascii.AvailableBanner(strings.ToLower(banner))

		if err != nil || !valid {
			log.Printf("Bad request, valid input = %v, banner error: %v", valid, err)
			badRequestHandler(w, tmpl)
			return
		}

		asciiArt := ascii.PrintAsciiArt(cleanInput, bannerCont)

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
	w.WriteHeader(http.StatusInternalServerError)
	tmpl.ExecuteTemplate(w, "500.html", nil)
}

// 404 page not found handler
func notFoundHandler(w http.ResponseWriter, tmpl *template.Template) {
	w.WriteHeader(http.StatusNotFound)
	tmpl.ExecuteTemplate(w, "404.html", nil)
}

// 400 bad request handler
func badRequestHandler(w http.ResponseWriter, tmpl *template.Template) {
	w.WriteHeader(http.StatusBadRequest)
	tmpl.ExecuteTemplate(w, "400.html", nil)
}
