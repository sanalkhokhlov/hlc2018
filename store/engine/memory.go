package engine

import (
	"sync"

	"bitbucket.org/sLn/hlc2018/store"
)


type memoryEngine struct {
	sync.RWMutex
	accounts map[uint32]store.Account

	sexMIndex []uint32
	sexFIndex []uint32
}

func NewMemoryEngine() Engine {
	engine := &memoryEngine{
		accounts: map[uint32]store.Account{},
		sexMIndex: []uint32{},
		sexFIndex: []uint32{},
	}

	return engine
}

func (m *memoryEngine) BulkCreate(accounts []store.Account) error {
	m.Lock()
	defer m.Unlock()

	for _, a := range accounts {
		m.accounts[a.ID] = a
	}

	return nil
}

func (m *memoryEngine) MakeIndexes() error {
	for _, a := range m.accounts {
		if a.Sex == "m" {
			m.sexMIndex = append(m.sexMIndex, a.ID)
		} else {
			m.sexFIndex = append(m.sexFIndex, a.ID)
		}
	}

	return nil
}

func (m *memoryEngine) filter(query string) (err error) {
	//badRequestError := &store.BadRequestError{}
	//pairs := strings.Split(query, "&")
	//for _, p := range pairs {
	//	pair := strings.Split(p, "=")
	//	if _, ok := filters[pair[0]]; ok {
	//		filters[pair[0]]++
	//	} else {
	//		filters[pair[0]] = 1
	//	}
	//	//if len(pair) < 2 {
	//	//	return badRequestError
	//	//}
	//
	//	//if pair[1] == "" {
	//	//	return badRequestError
	//	//}
	//	//
	//	//switch pair[0] {
	//	//case "sex_eq":
	//	//	break
	//	//}
	//}

	return err
}
