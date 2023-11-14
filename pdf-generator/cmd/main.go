package main

import (
	"github.com/FavorDespaches/pdf-generator/pkg/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.HandleUnsupportedMethod()
	case "POST":
		return handlers.SolicitarEtiqueta(req)
	case "PUT":
		return handlers.HandleUnsupportedMethod()
	case "DELETE":
		return handlers.HandleUnsupportedMethod()
	default:
		return handlers.HandleUnsupportedMethod()
	}
}
