package handlers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// Define the SuccessResponse structure
type SuccessResponse struct {
	StringBase64 string `json:"stringBase64"`
}

// Define the ErrorResponse structure
type ErrorResponse struct {
	Message string `json:"message"`
}

func ApiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	stringBody, err := json.Marshal(body)

	if err != nil {
		return nil, err // Return the error if marshaling fails
	}

	resp := events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: status,
		Body:       string(stringBody),
	}

	return &resp, nil
}
