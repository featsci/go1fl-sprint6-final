package handlers

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/featsci/go1fl-sprint6-final/internal/service"
)

func MainHandle(w http.ResponseWriter, req *http.Request) {
	path := "index.html"
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Файл %s не существует\n", path)
			return
		}
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, string(data),
			fmt.Sprintf("Сервер не поддерживает %s запросы", req.Method))
		return
	}
	fmt.Fprintf(w, "%s", string(data))
}

func UploadHandle(w http.ResponseWriter, req *http.Request) {
	// получаем файл из формы
	file, _, err := req.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusBadRequest)
		return
	}
	// закрываем файл
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Файл %s не существует\n", err)
			return
		}
		log.Fatal(err)
	}

	service.ServiceMorse(string(data))

}
