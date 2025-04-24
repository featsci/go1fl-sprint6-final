package handlers

import (
	"fmt"
	"net/http"
	"os"
)

// const pattern = `<!DOCTYPE html>
//   <html lang="ru"><head>
//   <meta charset="utf-8" />
//   <title>Тестовый сервер</title>
//   </head>
// <body>%s</body></html>`

func MainHandle(w http.ResponseWriter, req *http.Request) {
	data, err := os.ReadFile("index.html")
	if err != nil {
		// if errors.Is(err, os.ErrNotExist) {
		// 	return "", nil
		// }
		// return "", err
		// fmt.Println(err)
		return
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
	file, handler, err := r.FormFile("attach")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusBadRequest)
		return
	}
	// закрываем файл
	defer file.Close()
}
