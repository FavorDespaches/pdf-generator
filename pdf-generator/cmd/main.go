package main

import (
	"strings"

	"github.com/FavorDespaches/pdf-generator/pkg/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch strings.ToUpper(req.HTTPMethod) {
	case "GET":
		return handlers.HandleUnsupportedMethod()
	case "POST":
		//handlers.SolicitarEtiqueta(req) //handlers.SolicitarEtiqueta(req)

		/*
			dirPath := "/opt/bin/images" // Replace with the path of your directory
			contents := printAndReturnContents(dirPath)
			successBody := handlers.SuccessResponse{
				IdPlp:       0,
				Etiquetas:   []string{""},
				LabelBase64: contents,
			}
		*/

		return handlers.SolicitarEtiqueta(req)
	case "PUT":
		return handlers.HandleUnsupportedMethod()
	case "DELETE":
		return handlers.HandleUnsupportedMethod()
	default:
		return handlers.HandleUnsupportedMethod()
	}
}
