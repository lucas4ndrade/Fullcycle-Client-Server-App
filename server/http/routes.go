package http

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-fullcicle/server/config"
)

func StartRoutes() {
	router := chi.NewRouter()

	router.Get("/cotacao", getBid)

	http.Handle("/", router)

	sv := &http.Server{
		Addr:         ":" + config.GetConfig().Port,
		Handler:      http.TimeoutHandler(http.DefaultServeMux, 10*time.Second, "Timeout"),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Service listening at port %s\n\n", sv.Addr)
	if err := sv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start service, %v", err)
	}
}
