package main

import (
    "log"
    "os"
    "text/template"
)

// Define a struct to hold template data
type PageData struct {
    Title string
    Body  string
}

func main() {
    // Parse the template file
    tmpl, err := template.ParseFiles("templates/example.tmpl")
    if err != nil {
        log.Fatalf("Error parsing template: %v", err)
    }

    // Define the data to pass to the template
    data := PageData{
        Title: "My Title",
        Body:  "This is the body of the page.",
    }

    // Execute the template and write the output to stdout
    err = tmpl.ExecuteTemplate(os.Stdout, "example", data)
    if err != nil {
        log.Fatalf("Error executing template: %v", err)
    }
}
