package service

import (
	"bitbucket.org/sLn/hlc2018/store/engine"
)

type DataStore struct {
	engine.Engine
}

var limitBytes = []byte("limit")

func (ds *DataStore) Filter(query []byte) error {

	return nil
}

