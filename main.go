package main

import (
    "encoding/json"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/time", onTimeRequest)
    http.ListenAndServe(":8795", nil)
}

func onTimeRequest(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now()
    formattedTime := currentTime.Format(time.RFC3339)

    response := map[string]string{"time": formattedTime}

    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}