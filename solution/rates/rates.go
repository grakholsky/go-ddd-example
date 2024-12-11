package rates

import (
	"context"

	"task/store"
)

var prices = map[string]float64{
	"Bitcoin": 98269.56,
}

type AccountsStore struct {
	AccountsStore store.Store[store.Account]
}

func (s *AccountsStore) FindOne(ctx context.Context, args store.Account) (store.Account, error) {
	account, err := s.AccountsStore.FindOne(ctx, args)
	if err != nil {
		return store.Account{}, err
	}

	if price, ok := prices[account.Asset]; ok {
		// Apply bitwise AND operation to account tags
		if account.Tags&store.AccountTagRecount != 0 {
			account.BalanceUSD = account.Balance * price
		}
	}

	return account, nil
}
