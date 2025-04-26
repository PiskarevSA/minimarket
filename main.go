package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/PiskarevSA/minimarket/internal/gen/oapi"
	"github.com/go-chi/chi/v5"
)

var _ oapi.ServerInterface = (*Handlers)(nil)

type Handlers struct{}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) GetApiUserBalance(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func main() {
	rootCtx := context.Background()
	signals := []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL}

	stopCtx, stop := signal.NotifyContext(rootCtx, signals...)
	defer stop()

	handlers := NewHandlers()
	router := chi.NewRouter()

	h := oapi.HandlerFromMux(handlers, router)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: h,
	}

	log.Println("listening server on 127.0.0.1:8080...")

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("failed to listen server")
		}
	}()

	<-stopCtx.Done()
}
