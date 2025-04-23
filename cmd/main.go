package main

import (
	"fmt"
	"log"
	"os"

	"github.com/featsci/go1fl-sprint6-final/internal/server"
	"github.com/featsci/go1fl-sprint6-final/internal/service"
)

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Printf("Server is start")

	// t := service.ServiceMorse("Привет")

	// fmt.Println(t)

	// функция для вывода каждой строки

	// читаем файл построчно
	service.ServiceMorse("test12")

	// запускаем сервер
	if err := server.ServerGo(logger); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}

	// mux := http.NewServeMux()
	// // mux.HandleFunc("/time", handleTime)
	// // mux.HandleFunc("/", handleMain)

	// err := http.ListenAndServe(":8080", mux)
	// if err != nil {
	// 	panic(err)
	// }
}
