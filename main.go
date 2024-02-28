package main

import (
    "encoding/json"
    "net/http"
    "time"
)

func main() {
	port := "8080"

  http.HandleFunc("/time", onTimeRequest)
  http.ListenAndServe(":" + port, nil)
}

func onTimeRequest(w http.ResponseWriter, r *http.Request) {
    time := time.Now().Format(time.RFC3339)

    response := map[string]string{"time": time}

    jsonResponse, err := json.Marshal(response)
		
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}