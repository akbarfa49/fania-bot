package http

import (
	"context"
	"fania-bot/core"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/lesismal/nbio/nbhttp"
)

type Server struct {
	CoreService *core.Core
	Router      chi.Router
	HostPort    []string
	stopFunc    []func(ctx context.Context)
}

func (srv *Server) WatchStreamer(w http.ResponseWriter, r *http.Request) {
	in := core.WatchStreamer_In{}
	if err := render.DecodeJSON(r.Body, &in); err != nil {
		w.WriteHeader(400)
		return
	}
	out := srv.CoreService.WatchStreamer(r.Context(), in)
	render.JSON(w, r, out)

}

func (srv *Server) RegisterRoute() {
	router := chi.NewRouter()

	router.Post("/v1/stream/watch-streamer", srv.WatchStreamer)

	srv.Router = router
}

func (srv *Server) ServeNBHTTP() {
	srv.RegisterRoute()
	server := nbhttp.NewServer(nbhttp.Config{
		Network: "tcp",
		Handler: srv.Router,
		Addrs:   srv.HostPort,
	}) // pool.Go
	server.Start()
	srv.stopFunc = append(srv.stopFunc, func(ctx context.Context) {
		server.Shutdown(ctx)
	})
}

func (srv *Server) Shutdown(ctx context.Context) {
	for _, v := range srv.stopFunc {
		v(ctx)
	}
}
