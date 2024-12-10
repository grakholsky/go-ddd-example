package store

import "context"

// AccountTag a unique type to use bitwise values.
type AccountTag uint8

const (
	// AccountTagAvailable indicates that the account is available.
	AccountTagAvailable AccountTag = 1 << iota

	// AccountTagSpot indicates that the account is a spot account.
	AccountTagSpot

	// AccountTagEarn indicates that the account is an earn account for rewards and cashback or bonuses.
	AccountTagEarn

	// AccountTagRecount indicates that the account should be recounted in a stable coin equivalent.
	AccountTagRecount
)

type Account struct {
	ID         uint64     `json:"id"`
	Tags       AccountTag `json:"tags"`
	Asset      string     `json:"asset"`
	Balance    float64    `json:"balance"`
	BalanceUSD float64    `json:"balance_usd"`
}

type Store[T any] interface {
	FindOne(context.Context, T) (T, error)
}
