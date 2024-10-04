package service

import (
	"context"
	"log/slog"

	"song_service/internal/models"
)

type SongService struct {
	sr  SongRepositorier
	log *slog.Logger
}

func NewSongService(sr SongRepositorier) *SongService {
	return &SongService{sr: sr}
}

type SongRepositorier interface {
	CreateGroup(context.Context, *models.Song) (groupId int, err error)
	CreateSong(context.Context, *models.Song) (songId int, err error)
	CreateSongUniqRec(context.Context, int, int) (uniqRecId int, err error)
	AddDetails(context.Context, *models.Song) error
	UpdateSong(context.Context, *models.Song) error
	DeleteSong(context.Context, *models.Song) error
	GetAllSongs(context.Context, *models.FiltAndPagin) ([]*models.Song, error)
	GetSongText(context.Context, *models.FiltAndPagin) (*models.Song, error)
	CheckExistGroup(context.Context, *models.Song) (groupId int, err error)
	CheckExistSong(context.Context, *models.Song) (songId int, err error)
	CheckExistSongUniq(context.Context, int, int) (uniqRecId int, err error)
}
