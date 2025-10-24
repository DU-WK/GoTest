package main

import (
	"KGin/KGin"
	"fmt"
	"net/http"
)

func main() {
	//httpserver()

	kgin := KGin.NewEngine()

	kgin.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET:Hello, World!")
	})

	kgin.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET:Hello, %s!", r.URL.Query().Get("name"))
	})

	kgin.POST("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "POST:Hello, World!")
	})
	kgin.PUT("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "PUT:Hello, World!")
	})
	kgin.DELETE("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "DELETE:Hello, World!")
	})

	kgin.Run(":8080")
}
