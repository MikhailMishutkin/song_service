package transport

import (
	"context"
	"github.com/gorilla/mux"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"song_service/internal/models"
	"song_service/pkg/middleware"
)

type HTTPSongHandle struct {
	sh HTTPSongManager
}

func NewHTTPSongHandle(sh HTTPSongManager) *HTTPSongHandle {
	return &HTTPSongHandle{
		sh: sh,
	}
}

type HTTPSongManager interface {
	CreateSong(context.Context, *models.Song) error
	GetSongText(context.Context, *models.FiltAndPagin) (*models.Song, error)
	UpdateSong(context.Context, *models.Song) error
	DeleteSong(context.Context, *models.Song) error
	GetAllSongs(context.Context, *models.FiltAndPagin) ([]*models.Song, error)
}

func (h *HTTPSongHandle) RegisterSong(router *mux.Router) {
	router.Use(middleware.Logging)
	router.HandleFunc("/create", h.CreateSong).Methods("POST")
	router.HandleFunc("/edit/{id}", h.UpdateSong).Methods("PUT")
	router.HandleFunc("/delete/{id}", h.DeleteSong).Methods("DELETE")
	router.HandleFunc("/gettext/{id}", h.GetSongText).Methods("GET")
	router.HandleFunc("/getall", h.GetAllSongs).Methods(
		"Get",
	).Queries("id", "{id}", "group", "{group}", "song", "{song}", "release_date", "{release_date}", "text", "{text}", "link", "{link}")

}
