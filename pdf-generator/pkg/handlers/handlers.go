package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FavorDespaches/pdf-generator/pkg/helpers"
	"github.com/FavorDespaches/pdf-generator/pkg/types"
	"github.com/aws/aws-lambda-go/events"
)

func SolicitarEtiquetaLocal(correiosLog types.CorreiosLog) {
	fmt.Println("\n\n========== INICIANDO LAMBDA ==========")
}

func SolicitarEtiqueta(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("\n\n========== INICIANDO LAMBDA ==========")
	var correiosLog types.CorreiosLog
	err := json.Unmarshal([]byte(req.Body), &correiosLog)

	if err != nil {
		errText := fmt.Sprintf("Erro no Parser do JSON: %s", err.Error())

		errorBody := ErrorResponse{
			Message: errText,
		}
		return ApiResponse(http.StatusBadRequest, errorBody)
	}

	base64String, err := helpers.GenerateLabelsPDF(correiosLog)

	if err != nil {
		errText := fmt.Sprintf("Erro GenerateLabelsPDF: %s", err.Error())

		errorBody := ErrorResponse{
			Message: errText,
		}
		ApiResponse(http.StatusBadRequest, errorBody)
	}

	successBody := SuccessResponse{
		StringBase64: base64String,
	}

	return ApiResponse(http.StatusOK, successBody)
}

func HandleUnsupportedMethod() (*events.APIGatewayProxyResponse, error) {
	errorBody := ErrorResponse{
		Message: "Método inválido, utilize somente POST",
	}

	return ApiResponse(http.StatusBadRequest, errorBody)
}
