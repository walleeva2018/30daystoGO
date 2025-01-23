### make a web server 

#### Let's learn about 2 packages first 

net/http
This package is used to create HTTP servers and clients in Go. It’s the backbone of any web server in Go.

Key Features:

- HTTP Server: It provides tools to listen for HTTP requests, handle routes, and respond to requests.
- Request Handling: The http.HandleFunc function maps URL paths (like /about) to handler functions that define how to respond to those requests.
- Response Writing: The http.ResponseWriter interface is used to write data (like HTML, JSON, or plain text) back to the client.

Common Functions:

- http.HandleFunc: Maps routes (like /) to handler functions.
- http.ListenAndServe: Starts the HTTP server on a specified port.
- http.Get: Makes HTTP GET requests as a client.
- http.Post: Makes HTTP POST requests as a client.

[Learn more](https://pkg.go.dev/net/http)

----------------------------------------------------------------------


 html/template
This package is used to work with HTML templates in Go. It allows you to create dynamic HTML pages by combining static HTML with data (provided by your Go code).

Key Features:

- Templating Syntax: It supports placeholders like {{.Title}} or {{.Content}}, where the dot (.) refers to the data passed to the template.
- Data Injection: You can pass structs, maps, or other data types to populate your templates dynamically.
- XSS Protection: It automatically escapes user-supplied data to prevent Cross-Site Scripting (XSS) attacks.

Common Functions:

- template.ParseFiles: Reads and parses one or more template files.
- template.Execute: Executes the template and writes the output to a ResponseWriter or another writer.


```
package main

import (
	"html/template"
	"net/http"
)


func main() {
	// Set up HTTP routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	// Start the server
	port := ":8080"
	println("Server is running on http://localhost" + port)
	http.ListenAndServe(port, nil)
}
```

Here we took those packages and  in the main we define which route should use which function and finally we started the server it's that simple 

Now lets write those function 

```
func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Home Page",
		Content: "Welcome to the Home Page!",
	}
	renderTemplate(w, "home", data)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "About Page",
		Content: "This is the About Page. Learn more about us!",
	}
	renderTemplate(w, "home", data)
}
```

Here for type you must define PageData type 
```
type PageData struct {
	Title   string
	Content string
}

```
Then simply tell renderTemplate func to w (write) home (filename of your html) and pass data to it 

Now let's write the renderTemplate func

```
func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	tmplPath := tmpl + ".html"
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}
```

what is does is simply take the name of the html (which must be in this case in the same directory) make a html file path and renders it using package and passing the data With error handling 

See the html file how we have shown the data using go templating 