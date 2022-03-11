package server

import (
	native "net/http"

	"github.com/evanjennings/progprac/App/handler"
	"github.com/kevinanthony/gorps/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type HTTPServer struct {
	mux *chi.Mux
}

func NewServer(reqh http.RequestHandler, calc handler.Calc) HTTPServer {
	server := HTTPServer{
		mux: chi.NewRouter(),
	}
	server.mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	return server
}

func (s HTTPServer) Run() {
	svr := &native.Server{
		Addr:    ":8080",
		Handler: s.mux,
	}

	svr.SetKeepAlivesEnabled(false)

	if err := svr.ListenAndServe(); err != nil {
		panic(err)
	}
}
