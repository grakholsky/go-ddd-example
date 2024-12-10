package main

import (
	"context"
	"log"

	"task/earning"
	"task/inmem"
	"task/rates"
	"task/web"
)

func main() {
	ctx := context.Background()

	handler := &web.AccountHandler{
		AccountsStore: &rates.AccountsStore{
			AccountsStore: &earning.AccountsStore{
				AccountsStore: &inmem.AccountsStore{},
			},
		},
	}

	// -------------
	// SPOT
	// -------------
	if err := handler.ServeHTTP(ctx, 1); err != nil {
		log.Fatal(err)
	}

	// -------------
	// EARN
	// -------------
	if err := handler.ServeHTTP(ctx, 2); err != nil {
		log.Fatal(err)
	}
}
