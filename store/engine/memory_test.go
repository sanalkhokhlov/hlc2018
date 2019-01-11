package engine

import (
	"testing"

	"github.com/sanalkhokhlov/hlc2018/store"
)

func TestFilter(t *testing.T) {
	ds := NewMemoryEngine()
	ds.BulkCreate([]store.Account{
		store.Account{ID: 1, Sex: "f", Country: "Lol"},
		store.Account{ID: 2, Sex: "m", Country: "Lol"},
		store.Account{ID: 3, Sex: "f", Country: "Lol"},
	})
	ds.MakeIndexes()

	filters := store.FilterArgs{
		Limit: 1,
		Parts: [][][]byte{
			[][]byte{
				[]byte("sex"),
				[]byte("eq"),
				[]byte("f"),
			},
			[][]byte{
				[]byte("country"),
				[]byte("eq"),
				[]byte("Lol"),
			},
		},
	}

	accounts, err := ds.Filter(filters)
	if err != nil {
		t.Error(err)
		return
	}

	if len(accounts) != 1 {
		t.Errorf("accounts must be equal [3]")
	}
}
