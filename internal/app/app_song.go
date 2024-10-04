package app

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"log"
	"log/slog"
	"net/http"
	"song_service/configs"
	"song_service/internal/songservice/repository"
	"song_service/internal/songservice/service"
	"song_service/internal/songservice/transport"
)

type SongServer struct {
	songRouter *mux.Router
	logger     *slog.Logger
	swagRouter *mux.Router
}

func StartService(conf configs.Config) error {
	s := &SongServer{
		songRouter: mux.NewRouter(),
		logger:     slog.Default(),
		swagRouter: mux.NewRouter(),
	}

	db, err := NewDB()
	if err != nil {
		return fmt.Errorf("cannot connect to db on pqx: %v\n ", err)
	}

	//httpserver
	repo := repository.NewRepo(db)
	songService := service.NewSongService(repo)
	songHandler := transport.NewHTTPSongHandle(songService)
	songHandler.RegisterSong(s.songRouter)

	s.logger.Info("Starting MessageService at port: 8080")
	return http.ListenAndServe(":8080", s)
}

func NewDB() (*pgx.Conn, error) {
	c, err := configs.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("Can't load config in restaurant repo: %v\n", err)
	}

	psqlInfo := fmt.Sprint(c.Conn)

	db, err := pgx.Connect(context.Background(), psqlInfo)

	m, err := migrate.New(
		"file://../song_base/migrations",
		"postgres://"+c.Migrate,
		//root:root@localhost:5444/time_tracker?sslmode=disable",
	)
	if err != nil {
		log.Println(err)
		return db, fmt.Errorf("can't automigrate: %v\n", err)
	}
	if err := m.Up(); err != nil {
		log.Println(err)
		fmt.Errorf("%v\n", err)
	}
	return db, err
}

// ServeHTTP
func (h *SongServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.songRouter.ServeHTTP(w, r)

}
