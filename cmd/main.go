package main

import (
	"context"
	// "fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/VarunNSUT/Students-api/internal/config"
	"github.com/VarunNSUT/Students-api/internal/http/handlers/student"
)

func main() {
	//steps to set up the entry point 
	// 1. load config 
	// 2. database setup
	// 3. setup router 
	// 4. set up server 


	//LOAD CONFIG
	cfg := config.MustLoad()

	// SETUP ROUTER 
	// go has a powerful net http package to make server or router instead of a third party package 
	router := http.NewServeMux()

	router.HandleFunc("/api/student" , student.New())
	router.HandleFunc("/api/student/" , student.New())

	

	// SETUP SERVER
	server := http.Server {
		Addr: cfg.HTTPServer.Addr , 
		Handler: router,
	}

	slog.Info("server started" , slog.String("address" , cfg.HTTPServer.Addr)) // this is the format of writing in slog 

	done := make(chan os.Signal , 1) // we take signal as the channel input here from the os 

	signal.Notify(done , os.Interrupt , syscall.SIGINT , syscall.SIGTERM)

	go func (){ // this will run concurrently so when the execution arrives here then our code will end , so we resolve this issue by channels / synchronising
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed{
			log.Fatal("failed to start the server" , err)
		}
	}()
	
	// the statement below is blocking
	<- done //due to this our channel is blocked and main program will run unless complete process is over 

	slog.Info("shutting down the server") // slog is reocording what your program is doing 

	// report the shutdown by context pacakge 
	ctx , cancel :=context.WithTimeout(context.Background() , 5*time.Second) // pass an empty starting point that is the background function in this packae itsellf 
	defer cancel() 
	
	if err := server.Shutdown(ctx) ; err != nil {
		slog.Error("failed to shutdhown server" , slog.String("error", err.Error())) 
	}
	// this method gracefully shuts down the server 
	

	slog.Info("server shutdown successfully")
}
