package httpserver

import (
	"bitbucket.org/sLn/hlc2018/store"
	"bytes"
	"fmt"
	"log"
	"strings"

	"bitbucket.org/sLn/hlc2018/store/service"
	"github.com/valyala/fasthttp"
)

type Server struct {
	DataStore *service.DataStore
}

func (s *Server) fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	uri := string(ctx.RequestURI())
	parts := strings.Split(uri, "/")

	switch string(ctx.Method()) {
	case "GET":
		switch parts[2] {
		case "filter":
			s.filterHandler(ctx)
			return
		case "group":
			log.Println("group")
		default:
			switch parts[3] {
			case "recommend":
				log.Println("recommend")
				break
			case "suggest":
				log.Println("suggest")
				break
			default:
				ctx.SetStatusCode(fasthttp.StatusNotFound)
				return
			}
		}
		break
	case "POST":
		switch parts[2] {
		case "new":
			log.Println("new")
			break
		case "likes":
			log.Println("likes")
			break
		default:
			log.Println("update")
			break
		}
	}
}

func (ds *Server) ParseFilters(query []byte) (args store.FilterArgs, error) {
	badRequestError := &store.BadRequestError{}
	pairs := bytes.Split(query, []byte("&"))
	args = store.FilterArgs{}

	for _, p := range pairs {
		pair := bytes.Split(p, []byte("="))

		if len(pair) < 2 {
			return args, badRequestError
		}

		if len(pair[1]) == 0 {
			return args, badRequestError
		}

		args = append(args, pair)
	}

	return args, nil
}

func (s *Server) ErrorBadRequest(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusBadRequest)
}

func (s *Server) Run(port int) {
	err := fasthttp.ListenAndServe(fmt.Sprintf(":%v", port), s.fastHTTPHandler)
	log.Printf("http server terminated, %s", err)
}
