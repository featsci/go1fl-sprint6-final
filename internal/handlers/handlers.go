package handlers

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

/*
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>Document</title>
  </head>
  <body>
    <form
      enctype="multipart/form-data"
      action="http://localhost:8080/upload"
      method="post"
    >
      <input type="file" name="myFile" />
      <input type="submit" value="upload" />
    </form>
  </body>
</html>
*/

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
	fmt.Fprintf(w, string(data), "")
}

func UploadHandle(w http.ResponseWriter, req *http.Request) {
	// получаем файл из формы
	file, handler, err := req.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusBadRequest)
		return
	}
	// закрываем файл
	defer file.Close()

	dst, err := os.Create(handler.Filename)
	if err != nil {
		http.Error(w, "ошибка при создании файла", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// копируем содержимое загруженного файла в новый файл
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "ошибка при записи файла", http.StatusInternalServerError)
		return
	}

}
