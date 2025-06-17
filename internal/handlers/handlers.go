package handlers

import (
    "io"
    "net/http"
    "os"
    "path/filepath"
    "time"
    
    "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// IndexHandler обрабатывает запрос главной страницы
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

// UploadHandler обрабатывает загрузку файла
func UploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
        return
    }

    file, header, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Ошибка получения файла", http.StatusBadRequest)
        return
    }
    defer file.Close()

    content, err := io.ReadAll(file)
    if err != nil {
        http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
        return
    }

    result, err := service.Convert(string(content))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    filename := time.Now().UTC().Format("20060102150405") + filepath.Ext(header.Filename)
    if err := os.WriteFile(filename, []byte(result), 0644); err != nil {
        http.Error(w, "Ошибка сохранения файла", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte(result))
}

