package inmem

import "task/store"

const (
	// 0xb = 0x1 | 0x2 | 0x8
	spotAccountTags = store.AccountTagAvailable | store.AccountTagSpot | store.AccountTagRecount

	// 0xd = 0x1 | 0x4 | 0x8
	earnAccountTags = store.AccountTagAvailable | store.AccountTagEarn | store.AccountTagRecount
)

var AccountsIndex = map[uint64]store.Account{
	1: {
		ID:      1,
		Tags:    spotAccountTags,
		Asset:   "Bitcoin",
		Balance: 0.001,
	},
	2: {
		ID:      2,
		Tags:    earnAccountTags,
		Asset:   "Bitcoin",
		Balance: 0.00001,
	},
}
