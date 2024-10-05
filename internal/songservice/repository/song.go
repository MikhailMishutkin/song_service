package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"song_service/internal/models"
	"strings"
)

// ...
func (r *Repo) CreateGroup(ctx context.Context, song *models.Song) (id int, err error) {
	const query = `
INSERT INTO groups (group_name) 
	VALUES ($1) RETURNING id
`
	err = r.DB.QueryRow(
		ctx,
		query,
		song.GroupName,
	).Scan(&id)
	// TODO
	r.log.Debug("creategroup repo err: %v", "error", err, "id", id)

	return id, err
}

// ...
func (r *Repo) CreateSong(ctx context.Context, song *models.Song) (id int, err error) {
	const query = `
INSERT INTO songs (song) 
	VALUES ($1) RETURNING id
`
	err = r.DB.QueryRow(
		ctx,
		query,
		song.Song,
	).Scan(&id)
	// TODO
	r.log.Debug("createSong repo err: %v", "error", err, "id", id)

	return id, err
}

// ...
func (r *Repo) CreateSongUniqRec(ctx context.Context, grId, sId int) (uId int, err error) {
	const query = `
INSERT INTO song_unique (group_id, song_id) 
	VALUES ($1, $2) RETURNING id
`
	err = r.DB.QueryRow(
		ctx,
		query,
		grId,
		sId,
	).Scan(&uId)
	// TODO
	r.log.Debug("createSongUniq repo err: %v", "error", err, "id", uId)

	return uId, err
}

// ...
func (r *Repo) AddDetails(ctx context.Context, song *models.Song) error {
	const query = `
INSERT INTO details (uniq_id, release_date, text, link) 
	VALUES ($1, $2, $3, $4)
`
	tag, err := r.DB.Exec(
		ctx,
		query,
		song.Id,
		song.ReleaseDate,
		song.Text,
		song.Link,
	)
	r.log.Debug("do add details", "tag", tag, "error", err)

	return err
}

// TODO
func (r *Repo) UpdateSong(ctx context.Context, song *models.Song) error {
	log.Println("UpateSong ivoked", song)
	const query = `
UPDATE groups 
SET group = $2, song = $3, release_date = $4, text = $5, link = $6 
WHERE id = $1
`
	_, err := r.DB.Exec(
		ctx,
		query,
		song.Id,
		song.GroupName,
		song.Song,
		song.ReleaseDate,
		song.Text,
		song.Link,
	)

	return err
}

// ...
func (r *Repo) DeleteSong(ctx context.Context, song *models.Song) error {
	_, err := r.DB.Exec(ctx, "DELETE FROM song_unique WHERE id = $1", song.Id)
	if err != nil {
		return fmt.Errorf("can't delete record: %v", err)
	}
	return err
}

// ...
func (r *Repo) GetSongText(ctx context.Context, fp *models.FiltAndPagin) (songText *models.Song, err error) {
	log.Println("GetSongText invoked")

	var text string
	query := `SELECT text FROM details 
            WHERE uniq_id = $1
			`
	err = r.DB.QueryRow(ctx, query, fp.FilterMap["id"]).Scan(&text)

	if err != nil {
		return nil, fmt.Errorf("something wrong with get users info: ", err)
	}

	songText.Text = text

	return songText, err

}

// ...
func (r *Repo) GetAllSongs(ctx context.Context, fp *models.FiltAndPagin) (songs []*models.Song, err error) {

	var rows pgx.Rows
	if len(fp.FilterMap) == 0 {
		rows, err = r.DB.Query(ctx, "SELECT song_unique.id, groups.group_name, songs.song, details.release_date, details.text, details.link FROM song_unique "+
			"INNER JOIN groups ON song_unique.group_id = groups.id "+
			"INNER JOIN songs ON song_unique.song_id = songs.id "+
			"INNER JOIN details ON song_unique.id = details.uniq_id ",
		)
	} else {
		rows, err = r.DB.Query(ctx, "SELECT song_unique.id, groups.group_name, songs.song, details.release_date, details.text, details.link FROM song_unique "+
			"INNER JOIN groups ON song_unique.group_id = groups.id "+
			"INNER JOIN songs ON song_unique.song_id = songs.id "+
			"INNER JOIN details ON song_unique.id = details.uniq_id "+
			" WHERE "+strings.Join(fp.Where, " AND "), fp.Values...)
	}

	r.log.Debug("row", "row", rows)
	// TODO
	if err != nil {
		return nil, fmt.Errorf("something wrong with get songs info: ", err)
	}

	for rows.Next() {
		song := &models.Song{}
		if err = rows.Scan(
			&song.Id,
			&song.GroupName,
			&song.Song,
			&song.ReleaseDate,
			&song.Text,
			&song.Link); err != nil {
			return nil, fmt.Errorf("trouble with rows.Next then get users with filter: %s", err)
		}

		songs = append(songs, song)
	}

	return songs, err
}

func (r *Repo) CheckExistGroup(ctx context.Context, song *models.Song) (id int, err error) {

	const query = `SELECT id FROM groups WHERE group_name = $1`

	err = r.DB.QueryRow(ctx, query, song.GroupName).Scan(&id)
	r.log.Debug("check group row existence", "error", err)
	if err != nil {
		r.log.Debug("no such group_name in db")
		return 0, nil
	}
	return id, err

}

func (r *Repo) CheckExistSong(ctx context.Context, song *models.Song) (id int, err error) {

	const query = `SELECT id FROM songs where song = $1`

	err = r.DB.QueryRow(ctx, query, song.Song).Scan(&id)
	r.log.Debug("check song row existence", "error", err)
	if err != nil {
		r.log.Debug("no such song in db")
		return 0, nil
	}

	return id, err

}

func (r *Repo) CheckExistSongUniq(ctx context.Context, gr, s int) (id int, err error) {

	const query = `SELECT id FROM song_unique where group_id = $1 and song_id = $2`

	err = r.DB.QueryRow(ctx, query, gr, s).Scan(&id)
	r.log.Debug("check unique song row existence", "error", err, "id", id)
	if err != nil {
		r.log.Debug("no such unique song in db")
		return 0, nil
	}

	return id, err

}
