package engine

import (
	"sort"
	"sync"

	"github.com/sanalkhokhlov/hlc2018/store"
)

var (
	isNull    = "is_null"
	isNotNull = "is_not_null"
)

type memoryEngine struct {
	sync.RWMutex
	accounts map[uint32]store.Account

	indexes Indexes
}

func NewMemoryEngine() Engine {
	engine := &memoryEngine{
		accounts: map[uint32]store.Account{},
		indexes:  Indexes{},
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
		m.indexes.Add("sex", a.Sex, a.ID, false)
		m.indexes.Add("country", a.Country, a.ID, true)
		m.indexes.Add("city", a.City, a.ID, true)
		m.indexes.Add("fname", a.Status, a.ID, true)
		m.indexes.Add("sname", a.Status, a.ID, true)
		m.indexes.Add("status", a.Status, a.ID, false)
	}

	m.indexes.Sort()

	return nil
}

func (m *memoryEngine) Filter(filters store.FilterArgs) (accounts []uint32, err error) {
	accountIndexes := [][]uint32{}
	for _, parts := range filters.Parts {
		if a, ok := m.indexes.Get(string(parts[0]), string(parts[1]), string(parts[2])); !ok || len(a) == 0 {
			return []uint32{}, nil
		} else {
			accountIndexes = append(accountIndexes, a)
		}
	}

	if len(accountIndexes) == 0 {
		return []uint32{}, nil
	}

	if len(accountIndexes) == 1 {
		if filters.Limit > 0 && len(accountIndexes[0]) > filters.Limit {
			return accountIndexes[0][:filters.Limit], nil
		}

		return accountIndexes[0], nil
	}

	sort.SliceStable(accountIndexes, func(i, j int) bool {
		return len(accountIndexes[i]) < len(accountIndexes[j])
	})

	accounts = []uint32{}
	for _, id := range accountIndexes[0] {
		for ai := 1; ai < len(accountIndexes); ai++ {
			i := sort.Search(len(accountIndexes[ai]), func(i int) bool { return accountIndexes[ai][i] <= id })
			if i < len(accountIndexes[ai]) && accountIndexes[ai][i] == id {
				accounts = append(accounts, id)
			}
		}

		if len(accounts) == filters.Limit {
			return accounts, err
		}
	}

	return accounts, err
}

func (m *memoryEngine) MakeAccountsResponse(ids []uint32, filters store.FilterArgs) []store.Account {
	m.RLock()
	accounts := []store.Account{}

	for _, id := range ids {
		acc := store.Account{
			ID:    m.accounts[id].ID,
			Email: m.accounts[id].Email,
		}

		for _, f := range filters.Parts {
			switch string(f[0]) {
			case "sex":
				acc.Sex = m.accounts[id].Sex
				break
			case "country":
				acc.Country = m.accounts[id].Country
				break
			case "city":
				acc.City = m.accounts[id].City
				break
			case "status":
				acc.Status = m.accounts[id].Status
				break
			case "fname":
				acc.Fname = m.accounts[id].Fname
				break
			case "sname":
				acc.Sname = m.accounts[id].Sname
				break
			}

		}
		accounts = append(accounts, acc)
	}

	m.RUnlock()
	return accounts
}

type Indexes map[string]Index

func (idx Indexes) Add(field string, key string, value uint32, nullable bool) {
	if _, ok := idx[field]; !ok {
		idx[field] = Index{}
	}

	idx[field].Add(key, value, nullable)
}

func (idx Indexes) Get(field string, predicate string, key string) ([]uint32, bool) {
	if _, ok := idx[field]; !ok {
		return []uint32{}, false
	}

	return idx[field].Get(predicate, key)
}

func (idx Indexes) Sort() {
	for _, i := range idx {
		i.Sort()
	}
}

type Index map[string][]uint32

func (idx Index) Add(key string, value uint32, nullable bool) {
	if _, ok := idx[key]; !ok {
		idx[key] = []uint32{value}
		return
	}

	idx[key] = append(idx[key], value)

	if nullable {
		if key == "" {
			idx.Add(isNull, value, false)
		} else {
			idx.Add(isNotNull, value, false)
		}
	}
}

func (idx Index) Get(predicate string, key string) ([]uint32, bool) {
	if predicate == "null" {
		if key == "1" {
			key = isNull
		} else {
			key = isNotNull
		}
	}

	if _, ok := idx[key]; !ok {
		return []uint32{}, false
	}

	return idx[key], true
}

func (idx Index) Sort() {
	for _, val := range idx {
		sort.SliceStable(val, func(i, j int) bool {
			return val[i] > val[j]
		})
	}
}
