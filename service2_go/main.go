package main

import (
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // Connected clients
var broadcast = make(chan int)               // Broadcast channel

var upgrader = websocket.Upgrader{
    // Allow connections from any origin
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func main() {
    http.HandleFunc("/ws", handleConnections)
    go handleMessages()
    log.Println("Service 2 (Go) started on :5001")
    err := http.ListenAndServe(":5001", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer ws.Close()

    clients[ws] = true

    for {
        var msg int
        err := ws.ReadJSON(&msg)
        if err != nil {
            log.Printf("error: %v", err)
            delete(clients, ws)
            break
        }
        msg *= 2 // Double the number
		log.Printf("Received and doubled: %d", msg)

        broadcast <- msg
    }
}

func handleMessages() {
    for {
        msg := <-broadcast
		log.Printf("Broadcasting: %d", msg)
        for client := range clients {
            err := client.WriteJSON(msg)
            if err != nil {
                log.Printf("error: %v", err)
                client.Close()
                delete(clients, client) // Corrected from 'ws' to 'client'
            }
        }
    }
}
