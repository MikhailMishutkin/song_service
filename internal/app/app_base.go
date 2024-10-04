package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"song_service/configs"
	"song_service/internal/songbase"
)

type BaseServer struct {
	router *mux.Router
}

func StartBase(conf configs.Config) error {
	s := &BaseServer{
		router: mux.NewRouter(),
	}

	handler := songbase.NewBaseHandle()
	handler.RegisterBase(s.router)

	log.Println("Starting SongBase at port: 9000")
	return http.ListenAndServe(":9000", s)
}

// ServeHTTP
func (h *BaseServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
