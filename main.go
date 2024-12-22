package main

import (
	"calc_service/internal/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Регистрируем обработчик
	http.HandleFunc("/api/v1/calculate", handler.CalculateHandler)

	// Запускаем сервер
	port := ":8080"
	fmt.Printf("Server is starting on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
