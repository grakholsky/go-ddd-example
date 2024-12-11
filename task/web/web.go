package web

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"task/store"
)

type AccountHandler struct {
	// TODO: Add dependencies
}

func (h *AccountHandler) ServeHTTP(ctx context.Context, accountID uint64) error {
	// In the real-world application we will call ctx.UserValue(key)
	// to retrieve URL parameter :account_id from the *fasthttp.RequestCtx

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
