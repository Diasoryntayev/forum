package main

import (
	"fmt"
	"forum/internal/handler"
	"forum/internal/server"
	"log"
)

const port = ":8888"

func main() {
	handler := handler.NewHandler(services)

	server := new(server.Server)
	fmt.Printf("Starting server at port %s\nhttp://localhost%s/\n", port, port)
	if err := server.Run(port, handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
