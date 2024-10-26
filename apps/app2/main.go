package main

import (
	"net/http"

	"genaigateway/packages/logger"
)

func main() {
	router := http.NewServeMux()

	logger, err := logger.InitLogger("debug", "console", "development", "")
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("App 2"))
	})
	
	logger.Info().Msg("Server App 2 is running on port 4242")
	http.ListenAndServe(":4242", router)

}