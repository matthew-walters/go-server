package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()

	server := http.Server{
		Addr:    ":10002",
		Handler: mux,
	}

	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(35 * time.Second)
		w.Write([]byte("ok\n"))
	})

	mux.HandleFunc("/fast", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok\n"))
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok\n"))
	})

	go func() {
		log.Println("starts on :10002, pid:", os.Getpid())
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Println("http.ListenAndServe fails: ", err)
		}
	}()

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM)

	<-stop

	log.Println("shutdown...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("server.Shutdown fails: ", err)
	}

	log.Println("successfully stopped")
}
