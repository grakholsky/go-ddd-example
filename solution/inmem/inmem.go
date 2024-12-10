package inmem

import (
	"context"

	"task/store"
)

type AccountsStore struct{}

func (s *AccountsStore) FindOne(ctx context.Context, args store.Account) (store.Account, error) {
	return AccountsIndex[args.ID], nil
}
