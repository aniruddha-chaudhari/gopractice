package main

import (
	"fmt"
	"log"
	"net/http"
	"studentapi/internal/config"
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

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("failed to start server: %v", err.Error)
		
	}

	fmt.Println("server started on", cfg.Addr)
}
