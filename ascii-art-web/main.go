package main

import (
	"html/template"
	"log"
	"net/http"
	"pkg/ascii"
)

type PageData struct {
	AsciiArt string
	SelectedBanner string
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
	SelectedBanner := "standard"
	//reads the templates and saves it in tmpl
	tmpl, err := template.ParseFiles("templates/home.html", "templates/404.html", "templates/500.html", "templates/400.html")
	if err != nil {
		log.Printf("Error loading template: %v", err)
		serverErrorHandler(w, tmpl)
		return
	}
	// Creating dynamic data
	data := PageData{
		SelectedBanner: SelectedBanner,
	}

	if !(r.URL.Path == "/" || r.URL.Path == "/ascii-art") {
		notFoundHandler(w, tmpl)
		return
	}
	if r.Method == http.MethodPost {
		input := r.FormValue("userText")
		log.Println("Selected input: " + input)
		cleanInput, valid := ascii.ValidInput(input)
		if !valid {
			badRequestHandler(w, tmpl)
			return
		}

		banner := r.FormValue("style")
		log.Println("Selected banner: " + banner)
		bannerCont, err := ascii.ReadBanner(banner)
		if err != nil {
			serverErrorHandler(w, tmpl)
			return
		}

		data.AsciiArt = ascii.PrintAsciiArt(cleanInput, bannerCont)

		data.SelectedBanner = banner

	} else if r.Method != http.MethodGet {
		log.Printf("Bad request, not GET or POST")
		badRequestHandler(w, tmpl)
		return

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
