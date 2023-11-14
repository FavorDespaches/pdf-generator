package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/FavorDespaches/pdf-generator/pkg/helpers"
	"github.com/FavorDespaches/pdf-generator/pkg/types"
	"github.com/aws/aws-lambda-go/events"
)

func SolicitarEtiqueta(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var correiosLog types.CorreiosLog

	err := json.Unmarshal([]byte(req.Body), &correiosLog)
	if err != nil {
		errorBody := ErrorResponse{
			Message: err.Error(),
		}
		return apiResponse(http.StatusBadRequest, errorBody)
	}

	base64String, etiquetas, err := helpers.GenerateLabelsPDF(correiosLog)

	if err != nil {
		errorBody := ErrorResponse{
			Message: err.Error(),
		}
		apiResponse(http.StatusBadRequest, errorBody)
	}

	successBody := SuccessResponse{
		IdPlp:                    correiosLog.Plp.IdPlp,
		Etiquetas:                etiquetas,
		LabelBase64:              base64String,
		ContentDeclarationBase64: "",
	}

	return apiResponse(http.StatusOK, successBody)
}

func HandleUnsupportedMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusBadRequest, nil)
}
