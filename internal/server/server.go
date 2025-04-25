package server

import (
	"log"
	"net/http"
	"time"

	"github.com/featsci/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger    *log.Logger
	HttpStand *http.Server
}

func (s *Server) RunSrv() error {
	s.Logger.Printf("Server is start")
	return s.HttpStand.ListenAndServe()

}

func ServerGo(logger *log.Logger) *Server {
	// logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	// logger.Printf("Server is start")

	router := http.NewServeMux()
	router.HandleFunc("/", handlers.MainHandle)
	router.HandleFunc("/upload", handlers.UploadHandle)

	// Создайте экземпляр структуры http.Server
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	// Верните ссылку на ваш сервер.
	return &Server{
		HttpStand: srv,
		Logger:    logger,
	}
}
