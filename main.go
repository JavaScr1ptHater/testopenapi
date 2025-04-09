package main

import (
    "context"
    "fmt"
    openapi "github.com/JavaScr1ptHater/testopenapi/go-client"
)

func main() {
    // Configurazione del client
    cfg := openapi.NewConfiguration()
    cfg.BasePath = "http://localhost:8080" // Cambia l'URL se necessario
    client := openapi.NewAPIClient(cfg)

    // Chiamata all'API
    response, _, err := client.DefaultAPI.HelloGet(context.Background())
    if err != nil {
        fmt.Println("Errore chiamata API:", err)
        return
    }

    fmt.Println("Risposta dell'API:", response.Message)
}
