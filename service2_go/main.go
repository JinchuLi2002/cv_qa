package main

import (
    "encoding/json"
    "log"
    "net/http"
    "sync"
	"fmt"
)

type Message struct {
    Data int `json:"data"`
}

var (
    // lastProcessedNumber holds the last processed number.
    // Using a mutex to ensure safe access from multiple goroutines.
    lastProcessedNumber int
    mutex               sync.Mutex
)

func main() {
    http.HandleFunc("/api/send", sendHandler)
    log.Println("Service 2 (Go) started on :5001")
    err := http.ListenAndServe(":5001", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*") // CORS header
    w.Header().Set("Content-Type", "application/json")

    switch r.Method {
    case http.MethodGet:
        // Handle GET request
        mutex.Lock()
        responseMessage := Message{Data: lastProcessedNumber}
        mutex.Unlock()
        json.NewEncoder(w).Encode(responseMessage)
		log.Println("Number requested by UI, sending:", lastProcessedNumber)
    case http.MethodPost:
        // Handle POST request
        var requestMessage Message
        decoder := json.NewDecoder(r.Body)
        if err := decoder.Decode(&requestMessage); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        processedNumber := requestMessage.Data * 2
		// print: {requestMessage.Data} * 2 = {processedNumber} -> c# server
		fmt.Printf("%d * 2 = %d -> c# server\n", requestMessage.Data, processedNumber)
        mutex.Lock()
        lastProcessedNumber = processedNumber
        mutex.Unlock()
        responseMessage := Message{Data: processedNumber}
        json.NewEncoder(w).Encode(responseMessage)
    default:
        http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
    }
}
