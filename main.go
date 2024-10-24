package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

//nolint:gochecknoglobals // need these global to set during build
var (
	version   string
	commitSHA string
	buildDate string
	timeout   = 3 * time.Second
)

func helloHandler(w http.ResponseWriter, _ *http.Request) {
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
	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: timeout,
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Error("Could not start server", "error", err)
		panic(err)
	}
}
