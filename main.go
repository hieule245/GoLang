package main

import (
	"errors"
	"fmt"
	"net/http"
)

const postNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of bytes written: %d\n", n)
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(5, 6)
	_, err := fmt.Fprintf(w, "Hello %d", sum)
	if err != nil {
		fmt.Println(err)
	}
}

func addValues(a, b int) int {
	return a + b
}

func Devide(w http.ResponseWriter, r *http.Request) {
	value, err := devideNumber(5, 0)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("The result of the devide is: %.1f", value))
}

func devideNumber(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("Can't devide")
	}
	result = x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/devide", Devide)
	fmt.Printf(fmt.Sprintf("Starting application on post %s", postNumber))
	_ = http.ListenAndServe(postNumber, nil)
}
