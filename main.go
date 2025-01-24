package main

import (
	"awesomeProject/internal/db"
	"awesomeProject/internal/todo"
	"awesomeProject/internal/transport"
	"log"
)

func main() {
	dbHandler, err := db.New("postgres", "example", "localhost", "postgres", 5432)
	if err != nil {
		log.Fatal(err)
	}
	svc := todo.NewService(dbHandler)

	srv := transport.NewServer(svc)
	err = srv.Serve()
	if err != nil {
		log.Println("Server Shutting Down ", err.Error())
	}

}
