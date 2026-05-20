//go:build dev

package main

import (
	"log/slog"
	"net/http"
	"os"
)

func public() http.Handler {
	slog.Info("serving static files from disk")
	return http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public")))
}
