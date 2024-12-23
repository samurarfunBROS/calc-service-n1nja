package calculator

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		name           string
		payload        string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Expression",
			payload:        `{"expression":"2+2"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"result":"4"}`,
		},
		{
			name:           "Invalid Expression",
			payload:        `{"expression":"2+a"}`,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   `{"error":"Expression is not valid"}`,
		},
		{
			name:           "Invalid JSON",
			payload:        `{"expr":}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid request"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", strings.NewReader(tc.payload))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			CalculateHandler(rr, req)

			if rr.Code != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, rr.Code)
			}

			if strings.TrimSpace(rr.Body.String()) != tc.expectedBody {
				t.Errorf("expected body %s, got %s", tc.expectedBody, rr.Body.String())
			}
		})
	}
}
