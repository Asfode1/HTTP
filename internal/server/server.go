package server

import (
    "log"
    "net/http"
    "time"
    
    "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
    logger     *log.Logger
    httpServer *http.Server
}

// New создает новый экземпляр сервера
func New(logger *log.Logger) *Server {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handlers.IndexHandler)
    mux.HandleFunc("/upload", handlers.UploadHandler)

    return &Server{
        logger: logger,
        httpServer: &http.Server{
            Addr:         ":8080",
            Handler:      mux,
            ErrorLog:     logger,
            ReadTimeout:  5 * time.Second,
            WriteTimeout: 10 * time.Second,
            IdleTimeout:  15 * time.Second,
        },
    }
}

// Start запускает HTTP-сервер
func (s *Server) Start() error {
    return s.httpServer.ListenAndServe()
}


