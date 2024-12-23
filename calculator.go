package calculator

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := Response{Error: "Invalid request"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	result, err := EvaluateExpression(req.Expression)
	if err != nil {
		response := Response{Error: "Expression is not valid"}
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := Response{Result: strconv.FormatFloat(result, 'f', -1, 64)}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func EvaluateExpression(expression string) (float64, error) {
	// Пример обработки только чисел (упрощённая логика)
	expression = strings.ReplaceAll(expression, " ", "")
	result, err := strconv.ParseFloat(expression, 64)
	return result, err
}
