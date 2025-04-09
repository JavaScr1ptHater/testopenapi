package main

import (
    "context"
    "fmt"
    openapi "github.com/JavaScr1ptHater/testopenapi"
)

func main() {
    // Configurazione del client
    cfg := openapi.NewConfiguration()
    cfg.Servers = openapi.ServerConfigurations{
        {
            URL: "http://localhost:8080", // URL del tuo server API
        },
    }
    client := openapi.NewAPIClient(cfg)

    // Chiamata all'API
    req := client.DefaultAPI.HelloGet(context.Background())
    response, _, err := req.Execute()
    if err != nil {
        fmt.Println("Errore chiamata API:", err)
        return
    }   
    fmt.Println("Risposta dell'API:", response.GetMessage())
}
