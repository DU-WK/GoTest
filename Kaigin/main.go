package main

import (
	"KGin/KGin"
	"fmt"
)

func main() {
	//httpserver()

	kgin := KGin.NewEngine()

	kgin.GET("/", func(context *KGin.Context) {
		fmt.Fprintf(context.Writer, "GET:Hello, World!")
	})

	kgin.GET("/hello", func(context *KGin.Context) {
		fmt.Fprintf(context.Writer, "GET:Hello, %s!", context.Request.URL.Query().Get("name"))
	})

	kgin.POST("/", func(context *KGin.Context) {
		fmt.Fprintf(context.Writer, "POST:Hello, World!")
	})
	kgin.PUT("/", func(context *KGin.Context) {
		fmt.Fprintf(context.Writer, "PUT:Hello, World!")
	})
	kgin.DELETE("/", func(context *KGin.Context) {
		fmt.Fprintf(context.Writer, "DELETE:Hello, World!")
	})

	kgin.Run(":8080")
}
