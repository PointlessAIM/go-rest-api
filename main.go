package main

import (
	"GoRestApi/Server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)


func main() {

	ctx := context.Background()
	serverDoneChannel := make(chan os.Signal, 1)

	signal.Notify(serverDoneChannel, os.Interrupt, syscall.SIGTERM)
	srv:= Server.New(":8080")

	go func (){
		err:= srv.ListenAndServe()
		if err!=nil{
			panic(err)
		}
	}()
	log.Println("Server started")


	<- serverDoneChannel

	srv.Shutdown(ctx)
	log.Println("Server stopped")

	


}