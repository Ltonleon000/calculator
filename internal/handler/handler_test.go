package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		name         string
		request      string
		wantStatus   int
		wantResponse map[string]interface{}
	}{
		{
			name:       "valid expression",
			request:    `{"expression": "2+2"}`,
			wantStatus: http.StatusOK,
			wantResponse: map[string]interface{}{
				"result": float64(4),
			},
		},
		{
			name:       "invalid expression",
			request:    `{"expression": "2+a"}`,
			wantStatus: http.StatusUnprocessableEntity,
			wantResponse: map[string]interface{}{
				"error": "Expression is not valid",
			},
		},
		{
			name:       "empty expression",
			request:    `{"expression": ""}`,
			wantStatus: http.StatusUnprocessableEntity,
			wantResponse: map[string]interface{}{
				"error": "Expression is not valid",
			},
		},
		{
			name:       "invalid json",
			request:    `{"expression": 2+2}`,
			wantStatus: http.StatusUnprocessableEntity,
			wantResponse: map[string]interface{}{
				"error": "Expression is not valid",
			},
		},
		{
			name:       "missing expression field",
			request:    `{}`,
			wantStatus: http.StatusUnprocessableEntity,
			wantResponse: map[string]interface{}{
				"error": "Expression is not valid",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем тестовый HTTP запрос
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBufferString(tt.request))
			req.Header.Set("Content-Type", "application/json")

			// Создаем ResponseRecorder для записи ответа
			rr := httptest.NewRecorder()

			// Вызываем тестируемый обработчик
			CalculateHandler(rr, req)

			// Проверяем статус код
			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantStatus)
			}

			// Проверяем тело ответа
			var got map[string]interface{}
			if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
				t.Errorf("Failed to decode response body: %v", err)
				return
			}

			// Сравниваем ответ
			if tt.wantResponse["result"] != nil {
				if got["result"] != tt.wantResponse["result"] {
					t.Errorf("handler returned unexpected body: got %v want %v",
						got["result"], tt.wantResponse["result"])
				}
			} else if tt.wantResponse["error"] != nil {
				if got["error"] != tt.wantResponse["error"] {
					t.Errorf("handler returned unexpected error: got %v want %v",
						got["error"], tt.wantResponse["error"])
				}
			}
		})
	}
}

func TestCalculateHandlerMethodNotAllowed(t *testing.T) {
	// Тестируем GET запрос (должен быть отклонен)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/calculate", nil)
	rr := httptest.NewRecorder()

	CalculateHandler(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}

	var got map[string]interface{}
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Errorf("Failed to decode response body: %v", err)
		return
	}

	expected := "Method not allowed"
	if got["error"] != expected {
		t.Errorf("handler returned unexpected error: got %v want %v",
			got["error"], expected)
	}
}
