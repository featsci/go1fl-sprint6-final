package main

import (
	"fmt"
	"log"
	"os"

	"github.com/featsci/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	srv := server.ServerGo(logger)
	if err := srv.RunSrv(); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
