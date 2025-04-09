package main

import (
    "context"
    "fmt"
    openapi "github.com/JavaScr1ptHater/testopenapi.git"
)

func main() {
    // Crea una configurazione API
    cfg := openapi.NewConfiguration()
    client := openapi.NewAPIClient(cfg)

    // Chiama l'endpoint "/hello"
    response, _, err := client.DefaultAPI.HelloGet(context.Background())
    if err != nil {
        fmt.Println("Errore:", err)
        return
    }

    fmt.Println("Risposta dell'API:", response.Message)
}
