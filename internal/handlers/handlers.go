package handlers

import (
    "io"
    "net/http"
    "os"
    "path/filepath"
    "time"

    "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// IndexHandler обрабатывает запрос главной страницы и отдаёт файл index.html
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

// UploadHandler обрабатывает загрузку файла из формы,
// конвертирует содержимое из текста в Морзе или наоборот,
// сохраняет результат в локальный файл и возвращает результат клиенту.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
        return
    }

    // Ограничение размера загружаемого файла (например, 10 МБ)
    r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

    file, header, err := r.FormFile("myFile")
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
        http.Error(w, "Ошибка конвертации: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Генерируем уникальное имя файла с расширением исходного файла
    filename := time.Now().UTC().Format("20060102150405") + filepath.Ext(header.Filename)
    if err := os.WriteFile(filename, []byte(result), 0644); err != nil {
        http.Error(w, "Ошибка сохранения файла", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.Write([]byte(result))
}

