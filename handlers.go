package main

import (
    "encoding/json"
    "net/http"
)

func createMessageHandler(w http.ResponseWriter, r *http.Request) {
    var msg Message
    if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := saveMessageToDB(msg); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := sendMessageToKafka(msg); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func getStatsHandler(w http.ResponseWriter, r *http.Request) {
    stats, err := getProcessedMessagesStats()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(stats)
}
