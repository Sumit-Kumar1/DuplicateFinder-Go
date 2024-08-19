package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"dupfinder/internal/handler"
	"dupfinder/internal/service"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stderr, nil))

	s := service.New(log)
	h := handler.New(log, s)

	http.HandleFunc("/", h.RenderPage)

	http.HandleFunc("/info", h.SystemInfo)
	http.HandleFunc("/current", h.CurrentUsage)

	server := &http.Server{
		Addr:         "localhost:9000",
		Handler:      nil,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	log.Info("server started: ", "host:port->", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Error("error in starting up server: ", "error", err)
	}
}
