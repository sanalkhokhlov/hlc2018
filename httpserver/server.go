package httpserver

import (
	"fmt"
	"log"
	"strings"

	"github.com/sanalkhokhlov/hlc2018/store/service"
	"github.com/valyala/fasthttp"
)

type Server struct {
	DataStore *service.DataStore
}

func (s *Server) fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	uri := string(ctx.RequestURI())
	parts := strings.Split(uri, "/")
	ctx.SetContentType("application/json")

	switch string(ctx.Method()) {
	case "GET":
		switch parts[2] {
		case "filter":
			s.filterHandler(ctx)
			return
		case "group":
			// log.Println("group")
			ctx.Write([]byte(`{"groups": []}`))
			ctx.SetStatusCode(fasthttp.StatusOK)
			return
		default:
			switch parts[3] {
			case "recommend":
				// log.Println("recommend")
				ctx.Write([]byte(`{"accounts": []}`))
				ctx.SetStatusCode(fasthttp.StatusOK)
				return
			case "suggest":
				// log.Println("suggest")
				ctx.Write([]byte(`{"accounts": []}`))
				ctx.SetStatusCode(fasthttp.StatusOK)
				return
			default:
				ctx.SetStatusCode(fasthttp.StatusNotFound)
				return
			}
		}
	case "POST":
		switch parts[2] {
		case "new":
			// log.Println("new")
			ctx.SetStatusCode(fasthttp.StatusCreated)
			ctx.Write([]byte(`{}`))
			return
		case "likes":
			// log.Println("likes")
			ctx.SetStatusCode(fasthttp.StatusAccepted)
			ctx.Write([]byte(`{}`))
			return
		default:
			// log.Println("update")
			ctx.SetStatusCode(fasthttp.StatusAccepted)
			ctx.Write([]byte(`{}`))
			return
		}
	}
}

func (s *Server) ErrorBadRequest(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusBadRequest)
}

func (s *Server) Run(port int) {
	err := fasthttp.ListenAndServe(fmt.Sprintf(":%v", port), s.fastHTTPHandler)
	log.Printf("http server terminated, %s", err)
}
