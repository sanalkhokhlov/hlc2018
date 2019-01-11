package httpserver

import (
	"encoding/json"

	"github.com/sanalkhokhlov/hlc2018/store"
	"github.com/valyala/fasthttp"
)

func (s *Server) filterHandler(ctx *fasthttp.RequestCtx) {
	filters, err := s.ParseFilters(ctx.URI().QueryString())
	if err != nil {
		s.ErrorBadRequest(ctx)
		return
	}

	accounts, err := s.DataStore.Filter(filters)
	if err != nil {
		s.ErrorBadRequest(ctx)
		return
	}

	result := make(map[string][]store.Account)
	result["accounts"] = s.DataStore.MakeAccountsResponse(accounts, filters)
	err = json.NewEncoder(ctx.Response.BodyWriter()).Encode(result)
	if err != nil {
		s.ErrorBadRequest(ctx)
	} else {
		ctx.SetStatusCode(fasthttp.StatusOK)
	}
}
