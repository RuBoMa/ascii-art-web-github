package main

import (
	"html/template"
	"log"
	"net/http"
	"pkg/ascii"
)

type PageData struct {
	AsciiArt       string
	SelectedBanner string
	Input          string
}

// reads the templates and saves it in tmpl
var tmpl = template.Must(template.ParseFiles("templates/home.html", "templates/400.html", "templates/404.html", "templates/500.html"))

func main() {
	// Serving static files like css or images
	fileServer := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fileServer))

	// Handling the GET request to the root URL ("/")
	http.HandleFunc("/", handler)

	log.Println("server started")
	// start the server on port 8080 and listen to incoming requests
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}

}

// Parsing templates, executing home.html with dynamic data
func handler(w http.ResponseWriter, r *http.Request) {

	//giving error for any other URL paths than root and /ascii-art
	if !(r.URL.Path == "/" || r.URL.Path == "/ascii-art") {
		notFoundHandler(w, tmpl)
		return
	}

	//giving error for any other methods than GET or POST
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		log.Printf("Bad request, method not GET or POST")
		badRequestHandler(w, tmpl)
		return

	}

	// Creating dynamic data
	data := PageData{
		SelectedBanner: "standard",
	}

	//handling POST validating banner and input, creating Ascii-Art
	if r.Method == http.MethodPost {

		banner := r.FormValue("style")
		log.Println("Selected banner: " + banner)
		bannerCont, err := ascii.ReadBanner(banner)
		if err != nil {
			serverErrorHandler(w, tmpl)
			return
		}

		// Keeping the selected banner as default
		data.SelectedBanner = banner

		input := r.FormValue("userText")
		log.Println("Selected input: " + input)
		data.Input = "\n" + input
		output, valid := ascii.ValidInput(input)
		if !valid {
			w.WriteHeader(http.StatusBadRequest)
			data.AsciiArt = output
		} else {
			data.AsciiArt = ascii.PrintAsciiArt(output, bannerCont)
		}
	}

	err := tmpl.ExecuteTemplate(w, "home.html", data)
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
