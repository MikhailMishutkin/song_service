package songbase

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"song_service/internal/models"
)

const (
	detailsBase = "./songdetails.txt"
)

type BaseHandle struct {
}

func NewBaseHandle() *BaseHandle {
	return &BaseHandle{}
}

func (h *BaseHandle) GetSongInfo(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	song := &models.SongInput{
		Group: params["group"],
		Song:  params["song"],
	}

	stringSong := song.Group + " " + song.Song

	file, err := os.ReadFile(detailsBase)
	if err != nil {
		fmt.Errorf("errror: %v", err)
	}
	repo := make(map[string]*models.SongDetails)
	json.Unmarshal(file, &repo)

	details := repo[stringSong]

	response := &models.Song{
		GroupName:   song.Group,
		Song:        song.Song,
		ReleaseDate: details.ReleaseDate,
		Text:        details.Text,
		Link:        details.Link,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Connection:", "close")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}

func (h *BaseHandle) RegisterBase(router *mux.Router) {
	router.HandleFunc("/info", h.GetSongInfo).Methods("GET").Queries("group", "{group}", "song", "{song}")

}

//TODO
//func h(name string) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "%s: Вы вызвали %s методом %s\n", name, r.URL.String(), r.Method)
//	}
//}
//
//func main() {
//	m := http.NewServeMux()
//	m.Handle("GET /posts/latest", h("latest"))
//	m.Handle("GET /posts/{id}", h("id"))
//	m.Handle("GET /posts", h("posts"))
//	http.ListenAndServe(":7777", m)
//}
