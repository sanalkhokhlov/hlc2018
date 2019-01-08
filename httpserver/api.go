package httpserver

import (
	"log"

	"github.com/valyala/fasthttp"
)

func (s *Server) filterHandler(ctx *fasthttp.RequestCtx) {
	log.Println("filter")
	err := s.DataStore.Filter(ctx.URI().QueryString())
	if err != nil {
		s.ErrorBadRequest(ctx)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}
