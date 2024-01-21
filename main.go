package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"

    _ "github.com/lib/pq"
)

const (
    dbConnectionString = "user=postgres dbname=postgres sslmode=disable" // Modify with your DB credentials
    listenAddr         = ":3040"
)

var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("postgres", dbConnectionString)
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }

    db.SetMaxOpenConns(20)

    if err = db.Ping(); err != nil {
        log.Fatalf("Error pinging database: %v", err)
    }
}

func main() {
    http.HandleFunc("/", handleRequest)
    log.Printf("Listening on port %s", listenAddr)
    log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    if _, err := db.Exec("SELECT pg_sleep(1)"); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := map[string]string{"foo": "bar"}
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}
