package main

import (
	"context"
	"log"

	"task/inmem"
	"task/rates"
	"task/web"
)

func main() {
	// TODO: initialize accounts handler

	err := handler.ServeHTTP(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
