package engine

import "bitbucket.org/sLn/hlc2018/store"

type Engine interface {
	BulkCreate(accounts []store.Account) error
	MakeIndexes() error
}
