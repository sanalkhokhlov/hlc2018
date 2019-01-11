package httpserver

import (
	"bytes"
	"testing"
)

func TestParseFilters(t *testing.T) {
	s := Server{}
	query := []byte("sex_eq=f&country_eq=Лол&fname_any=Петр,Сергей&limit=10&query_id=1231")
	args, err := s.ParseFilters(query)
	if err != nil {
		t.Error(err)
		return
	}

	if !bytes.Equal(args.Parts[0][0], []byte("sex")) {
		t.Errorf("first arg must be equal sex")
	}

	if !bytes.Equal(args.Parts[0][1], []byte("eq")) {
		t.Errorf("first arg must be equal eq")
	}

	if !bytes.Equal(args.Parts[0][2], []byte("f")) {
		t.Errorf("value of first arg must be equal f")
	}

	if args.Limit != 10 {
		t.Errorf("limit must be equal 10")
	}

	if len(args.Parts) != 3 {
		t.Errorf("must be 3 pairs")
	}

	query2 := []byte("sex_eq=&limit=10")
	_, err2 := s.ParseFilters(query2)
	if err2 == nil {
		t.Errorf("err must be exists")
	}
}
