package service

import (
	"context"
	"errors"
	"fmt"
	"song_service/internal/models"
	"strings"
)

// first check existence names of group and/or song, then create uniq record with details
func (s *SongService) CreateSong(ctx context.Context, song *models.Song) error {
	idGr, err := s.sr.CheckExistGroup(ctx, song)
	if err != nil && idGr == 0 {
		return fmt.Errorf("something wrong with check group existence: %v", err)
	} else if idGr == 0 {
		idGr, err = s.sr.CreateGroup(ctx, song)
		if err != nil {
			return err
		}
	}

	idS, err := s.sr.CheckExistSong(ctx, song)
	if err != nil && idS == 0 {
		return fmt.Errorf("something wrong with check songname existence: %v", err)
	} else if idS == 0 {
		idS, err = s.sr.CreateSong(ctx, song)
		if err != nil {
			return err
		}
	}

	idU, err := s.sr.CheckExistSongUniq(ctx, idGr, idS)
	if err != nil && idU == 0 {
		return fmt.Errorf("something wrong with check unique song existence: %v", err)
	} else if idU == 0 {
		idU, err = s.sr.CreateSongUniqRec(ctx, idGr, idS)
		if err != nil {
			return err
		}
	}
	song.Id = idU

	err = s.sr.AddDetails(ctx, song)

	return err

}

// ...
func (s *SongService) UpdateSong(ctx context.Context, song *models.Song) (err error) {

	if song.GroupName != "" && song.Song != "" && song.Link != "" && song.Text != "" && song.ReleaseDate != "" {

		err = s.sr.UpdateSong(ctx, song)
	} else {
		return errors.New("empty fields in update data")
	}

	return err
}

// ...
func (s *SongService) DeleteSong(ctx context.Context, song *models.Song) error {
	err := s.sr.DeleteSong(ctx, song)
	return err
}

// ...
func (s *SongService) GetAllSongs(
	ctx context.Context,
	fp *models.FiltAndPagin,
) (
	songs []*models.Song,
	err error,
) {
	var values []interface{}
	var where []string
	for k, v := range fp.FilterMap {
		values = append(values, v)
		where = append(where, fmt.Sprintf("%s = ?", k))
	}

	fp.Values = values
	fp.Where = where
	songs, err = s.sr.GetAllSongs(ctx, fp)
	return songs, err
}

// ...
func (s *SongService) GetSongText(ctx context.Context, fp *models.FiltAndPagin) (song *models.Song, err error) {
	songText, err := s.sr.GetSongText(ctx, fp)
	slice := strings.Fields(songText.Text)
	var slice1 []string
	var text string
	chorus := 0
	for _, v := range slice {

		if v == "chorus" {
			chorus++
		}
		if chorus == fp.Limit {
			for _, v := range slice1 {
				text = text + v
			}
			song.Text = text
			return song, err
		}
		slice1 = append(slice1, v)
	}

	return songText, err
}
