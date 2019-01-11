package httpserver

import (
	"bytes"
	"net/url"
	"strconv"

	"github.com/sanalkhokhlov/hlc2018/store"
)

var (
	paramQueryID = []byte("query_id")
	paramLimit   = []byte("limit")
)

func (s *Server) ParseFilters(query []byte) (args store.FilterArgs, err error) {
	badRequestError := &store.BadRequestError{}
	pairs := bytes.Split(query, []byte("&"))
	args = store.FilterArgs{}

	for _, p := range pairs {
		pair := bytes.Split(p, []byte("="))

		if len(pair) < 2 {
			return args, badRequestError
		}

		if bytes.Equal(pair[0], paramQueryID) {
			continue
		}

		if len(pair[1]) == 0 {
			return args, badRequestError
		}

		if bytes.Equal(pair[0], paramLimit) {
			limit, err := strconv.Atoi(string(pair[1]))
			if err != nil {
				return args, err
			}

			args.Limit = limit
			continue
		}

		parts := bytes.Split(pair[0], []byte("_"))
		if len(parts) < 2 {
			return args, badRequestError
		}

		if bytes.Equal([]byte("country"), parts[0]) ||
			bytes.Equal([]byte("city"), parts[0]) ||
			bytes.Equal([]byte("sname"), parts[0]) ||
			bytes.Equal([]byte("fname"), parts[0]) ||
			bytes.Equal([]byte("status"), parts[0]) {
			// val, _ := strconv.Unquote("\"" + string(pair[1]) + "\"")
			// parts = append(parts, []byte(val))
			// } else if bytes.Equal([]byte("status"), parts[0]) {
			val, _ := url.QueryUnescape(string(pair[1]))
			parts = append(parts, []byte(val))
		} else {
			parts = append(parts, pair[1])
		}

		args.Parts = append(args.Parts, parts)
	}

	return args, nil
}
