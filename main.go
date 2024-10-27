package main

import (
	_ "embed"
	"gitlab.com/arkadooti.sarkar/relay-raspberry-pi/apis"
	"gitlab.com/arkadooti.sarkar/relay-raspberry-pi/toggle"
	"log"
	"net/http"
)

//go:embed static/index.html
var indexHtml []byte

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	// Serve the embedded index.html file
	w.Header().Set("Content-Type", "text/html")
	w.Write(indexHtml)
}

func main() {
	go toggle.DayTimeScheduler()
	mux := http.NewServeMux()

	// Serve embedded index.html at the "/" route
	mux.HandleFunc("/", serveHome)

	mux.HandleFunc("/toggle", apis.ToggleLightHandler)
	mux.HandleFunc("/status", apis.GetStatusHandler)

	handler := corsMiddleware(mux)
	log.Fatal(http.ListenAndServe(":8082", handler))
}
