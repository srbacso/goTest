package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})

	http.HandleFunc("/visit", Visit)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting server on port: %v", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}

// Divide divides one value into another and returns message with result
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 12.0)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Cannot divide by 0")
		return
	}
	_, _ = fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 12.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func addValues(x, y int) (int, error) {
	return x + y, nil
}

func Visit(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "visit.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
	}
}
