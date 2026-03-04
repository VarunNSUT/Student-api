package main

import (
	"fmt"
	"log"
	"net/http"

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

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start the server")
	}

	fmt.Println("server started")
}
