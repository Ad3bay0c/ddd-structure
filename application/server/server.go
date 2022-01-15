package server

import (
	"context"
	"fmt"
	"github.com/Ad3bay0c/ddd-structure/application/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type server struct {
	App	*handler.Application
}

func NewServer(app *handler.Application) *server {
    return &server{
        App: app,
    }
}

func (s *server) Start() {
	router := s.SetupRouter()
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if PORT == ":" {
		PORT = ":8080"
	}
	wait := make(chan os.Signal, 1)
	server := &http.Server{
		Addr: PORT,
		Handler: router,
	}
	go func() {
		log.Printf("Server started on http://localhost%s\n", PORT)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("%v\n", err)
		}
	}()

	// Ctrl + C on the terminal for example sends a terminate signal.
	signal.Notify(wait, os.Interrupt)
	// Blocks until a signal is sent to the channel.
	<-wait

	log.Printf("Shutting down the server...")
	time.Sleep(time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Gracefully shutdown the server when receiving the TERM or KILL signals.
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Shutting down forcefully", err)
	}
	log.Println("Server Shut Down")
}
