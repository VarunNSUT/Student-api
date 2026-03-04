package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/VarunNSUT/Students-api/internal/config"
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

	//http.ResponseWriter is an interface provided by Go's net/http package that lets your server write the HTTP response.
	router.HandleFunc("/" , func(w http.ResponseWriter , r *http.Request){
		w.Write([]byte("welcome to students api"))
	})

	// SETUP SERVER
	server := http.Server {
		Addr: cfg.HTTPServer.Addr , 
		Handler: router,
	}

	fmt.Println("server started")

	done := make(chan os.Signal , 1) // we take signal as the channel input here from the os 

	signal.Notify(done , os.Interrupt , syscall.SIGINT , syscall.SIGTERM)

	go func (){ // this will run concurrently so when the execution arrives here then our code will end , so we resolve this issue by channels / synchronising
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start the server")
		}
	}()
	
	// the statement below is blocking
	<- done //due to this our channel is blocked and main program will run unless complete process is over 

	slog.Info("shutting down the server") // slog is reocording what your program is doing 

	// report the shutdown by context pacakge 
	ctx , cancel :=context.WithTimeout(context.Background() , 5*time.Second) // pass an empty starting point that is the background function in this packae itsellf 
	defer cancel() 
	
	err := server.Shutdown(ctx) // this method gracefully shuts down the server 
	if err != nil {
		slog.Error("failed to shutdown the server" , slog.String("error" , err.Error()))
	}

	slog.Info("server shutdown successfully")
}
