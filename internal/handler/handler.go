package handler

import (
	"calc_service/internal/calculator"
	"encoding/json"
	"net/http"
	"strings"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Проверяем метод
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(Response{Error: "Method not allowed"})
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Response{Error: "Expression is not valid"})
		return
	}

	// Проверяем на пустое выражение
	if req.Expression == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Response{Error: "Expression is not valid"})
		return
	}

	result, err := calculator.Evaluate(req.Expression)
	if err != nil {
		// Определяем тип ошибки по сообщению
		errMsg := err.Error()
		if strings.Contains(errMsg, "invalid character") ||
		   strings.Contains(errMsg, "invalid number") ||
		   strings.Contains(errMsg, "empty expression") {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(Response{Error: "Expression is not valid"})
			return
		}
		
		// Для всех остальных ошибок возвращаем 500
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Error: "Internal server error"})
		return
	}

	// Успешный ответ
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Result: result})
}
