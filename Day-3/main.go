package main

import (
	"html/template"
	"net/http"
)

// Template structure
type PageData struct {
	Title   string
	Content string
}

func main() {
	// Set up HTTP routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	// Start the server
	port := ":8080"
	println("Server is running on http://localhost" + port)
	http.ListenAndServe(port, nil)
}

// Handler for "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Home Page",
		Content: "Welcome to the Home Page!",
	}
	renderTemplate(w, "home", data)
}

// Handler for "/about"
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "About Page",
		Content: "This is the About Page. Learn more about us!",
	}
	renderTemplate(w, "home", data)
}

// Template rendering helper
func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	tmplPath := tmpl + ".html"
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}
