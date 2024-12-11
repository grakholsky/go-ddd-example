package earning

import (
	"context"

	"task/store"
)

var rewards = map[string]float64{
	"Bitcoin": 0.000002,
}

type AccountsStore struct {
	AccountsStore store.Store[store.Account]
}

func (s *AccountsStore) FindOne(ctx context.Context, args store.Account) (store.Account, error) {
	account, err := s.AccountsStore.FindOne(ctx, args)
	if err != nil {
		return store.Account{}, err
	}

	if reward, ok := rewards[account.Asset]; ok {
		// Apply bitwise AND operation to account tags
		if account.Tags&store.AccountTagEarn != 0 {
			account.Balance = account.Balance + reward
		}
	}

	return account, nil
}
