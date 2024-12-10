package web

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"task/store"
)

type AccountsHandler struct {
	// TODO: Add dependencies
}

func (h *AccountsHandler) ServeHTTP(ctx context.Context, accountID uint64) error {
	accounts, err := h.AccountsStore.FindOne(ctx, store.Account{
		ID: accountID,
	})
	if err != nil {
		return err
	}

	bits, err := json.Marshal(accounts)
	if err != nil {
		return fmt.Errorf("failed to marshal account: %w", err)
	}

	// Writing response body to client
	log.Println(string(bits))
	return nil
}
