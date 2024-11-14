package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"studentapi/internal/config"
	"syscall"
	"time"
)

func main() {
	//load the config
	cfg := config.MustLoad()

	//setup routes
	router := http.NewServeMux()

	router.HandleFunc("GET /",func(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World"))	
	})


	//start the server
	server := http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}

	slog.Info("starting server", slog.String("address", cfg.Addr))
	fmt.Println("server started on", cfg.Addr)


	done := make(chan os.Signal,1)


	signal.Notify(done,os.Interrupt,syscall.SIGINT,syscall.SIGTERM)

	go func(){
	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("failed to start server: %v", err.Error)
		
	}
	} ()

	<-done

	slog.Info("shutting down server")

	ctx,cancel := context.WithTimeout(context.Background(),5 * time.Second)

	defer cancel()

	err :=server.Shutdown(ctx)

	if err != nil {
		slog.Error("failed to shutdown server:", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown")

	
}
