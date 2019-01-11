package engine

import "github.com/sanalkhokhlov/hlc2018/store"

type Engine interface {
	BulkCreate(accounts []store.Account) error
	MakeIndexes() error
	Filter(filters store.FilterArgs) (accounts []uint32, err error)
	MakeAccountsResponse(ids []uint32, filters store.FilterArgs) []store.Account
}
