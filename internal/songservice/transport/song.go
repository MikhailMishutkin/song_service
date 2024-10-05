package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"song_service/internal/models"
	"strconv"
)

// CreateUser
//
//	@Summary		CreateSong
//	@Description	create a song
//	@Tags			create
//	@Accept			json
//	@Produce		json
//	@Param			input	body		models.SongInput true "song info"
//	@Success		200		{integer} integer 1
//	@Failure		400,404	{integer}	integer 1
//	@Failure		500		{integer}	integer 1
//	@Router			/create [post]
func (h *HTTPSongHandle) CreateSong(w http.ResponseWriter, r *http.Request) {

	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	var song *models.SongInput
	err = json.Unmarshal(content, &song)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("corrupt json data" + err.Error()))
	}

	songEnrich, err := h.GetSongInfo(song)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("error with user's info request" + err.Error()))
	}

	err = h.sh.CreateSong(context.Background(), songEnrich)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
	}
}

// ...
func (h *HTTPSongHandle) UpdateSong(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)["id"]     //
	id, err := strconv.Atoi(v) //
	FailOnErrorsHttp(w, err, "can't convert UUID", http.StatusBadRequest)

	song := &models.Song{Id: id}

	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	err = json.Unmarshal(content, &song)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("corrupt json data" + err.Error()))
	}

	err = h.sh.UpdateSong(context.Background(), song)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

	}
}

func (h *HTTPSongHandle) DeleteSong(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)["uuid"]
	id, err := strconv.Atoi(v)
	FailOnErrors(err, "can't convert UUID")

	song := &models.Song{Id: id}

	err = h.sh.DeleteSong(context.Background(), song)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

	}
}

// ...
func (h *HTTPSongHandle) GetSongText(w http.ResponseWriter, r *http.Request) {

	// TODO Errors!!
	strLimit := r.URL.Query().Get("limit")
	limit := 0
	if strLimit != "" {
		limit, err := strconv.Atoi(strLimit)
		if err != nil || limit < 0 {
			http.Error(w, "limit query parameter is no valid number", http.StatusBadRequest)
			return
		}
	}
	strOffset := r.URL.Query().Get("offset")
	offset := 0
	if strOffset != "" {
		offset, err := strconv.Atoi(strOffset)
		if err != nil || offset < 0 {
			http.Error(w, "offset query parameter is no valid number", http.StatusBadRequest)
			return
		}
	}

	response, err := h.sh.GetSongText(
		context.Background(),
		&models.FiltAndPagin{Limit: limit, Offset: offset},
	)

	//respJson := &models.Song{: response}

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Connection:", "close")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			fmt.Println(err)
			http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		}

	}
}

// Get All Songs info
//
//	@Summary		GetAllSongs
//	@Description	get song info
//	@Tags			get info
//	@Accept			json
//	@Produce		json
//	@Success		200		{object} models.Song
//	@Failure		400,404	{integer}	integer 1
//	@Failure		500		{integer}	integer 1
//	@Router			/getall [get]
func (h *HTTPSongHandle) GetAllSongs(w http.ResponseWriter, r *http.Request) {

	var err error

	strLimit := r.URL.Query().Get("limit")
	limit := 0
	if strLimit != "" {
		limit, err = strconv.Atoi(strLimit)
		if err != nil || limit < 0 {
			http.Error(w, "limit query parameter is no valid number", http.StatusBadRequest)
			return
		}
	}
	strOffset := r.URL.Query().Get("offset")
	offset := 0
	if strOffset != "" {
		offset, err = strconv.Atoi(strOffset)
		if err != nil || offset < 0 {
			http.Error(w, "offset query parameter is no valid number", http.StatusBadRequest)
			return
		}
	}

	filter := r.URL.Query().Get("filter")
	filterMap := map[string]string{}
	if filter != "" {
		filterMap, err = validateAndReturnFilterMap(filter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	response, err := h.sh.GetAllSongs(context.Background(), &models.FiltAndPagin{FilterMap: filterMap, Limit: limit, Offset: offset})
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Connection:", "close")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			fmt.Println(err)
			http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		}
	}
}
