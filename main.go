package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

var (
	version   string
	commitSHA string
	buildDate string
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	http.HandleFunc("/", helloHandler)

	logger.Info("Starting server on :8080")
	if len(version) == 0 {
		version = "local"
		commitSHA = "none"
		buildDate = time.Now().Format(time.RFC3339)
	}
	logger.Info(fmt.Sprintf("Version: %s, commitSHA: %s, buildDate: %s", version, commitSHA, buildDate))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Could not start server", "error", err)
	}
}
