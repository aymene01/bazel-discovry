package main

import (
	"net/http"

	"genaigateway/packages/logger"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("App 1"))
	})

	logger, err := logger.InitLogger("debug", "console", "development", "")
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	logger.Info().Msg("Server App 1 is running on port 8080")
	http.ListenAndServe(":8080", router)

}