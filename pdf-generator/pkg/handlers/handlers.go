package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FavorDespaches/pdf-generator/pkg/helpers"
	"github.com/FavorDespaches/pdf-generator/pkg/types"
	"github.com/aws/aws-lambda-go/events"
)

func SolicitarEtiqueta(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("\n\n========== INICIANDO LAMBDA ==========")
	var correiosLog types.CorreiosLog
	err := json.Unmarshal([]byte(req.Body), &correiosLog)

	if err != nil {
		errText := fmt.Sprintf("Erro no Parser do JSON: %s", err.Error())

		errorBody := ErrorResponse{
			Message: errText,
		}
		return apiResponse(http.StatusBadRequest, errorBody)
	}

	base64String, etiquetas, err := helpers.GenerateLabelsPDF(correiosLog)

	if err != nil {
		errText := fmt.Sprintf("Erro GenerateLabelsPDF: %s", err.Error())

		errorBody := ErrorResponse{
			Message: errText,
		}
		apiResponse(http.StatusBadRequest, errorBody)
	}

	successBody := SuccessResponse{
		IdPlp:       correiosLog.Plp.IdPlp,
		Etiquetas:   etiquetas,
		LabelBase64: base64String,
	}

	return apiResponse(http.StatusOK, successBody)
}

func HandleUnsupportedMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusBadRequest, nil)
}
