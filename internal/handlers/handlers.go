package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandlerOne(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func HandlerTwo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "parse failed", http.StatusBadRequest)
		return
	}

	file, head, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "form failed", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	src, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "read failed", http.StatusInternalServerError)
		return
	}

	res, err := service.Parse(string(src))
	if err != nil {
		http.Error(w, "parse failed", http.StatusInternalServerError)
		return
	}

	t := time.Now().UTC().Format("02-01-06_15-04-05")
	extFile := filepath.Ext(filepath.Base(head.Filename))

	dst, err := os.Create(t + extFile)
	if err != nil {
		http.Error(w, "create failed", http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	_, err = dst.Write([]byte(res))
	if err != nil {
		http.Error(w, "write error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(res))
}
