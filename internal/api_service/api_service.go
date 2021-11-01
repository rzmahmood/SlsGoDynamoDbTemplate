package api_service

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"net/http"
)

func formatBodyMessage(msg string) string {
	body, err := json.Marshal(map[string]interface{}{
		"message": msg,
	})
	if err != nil {
		log.Fatalf("Failed to Marshal json response in function SuccessResponse")
	}
	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)
	return buf.String()
}

func SuccessResponse(msg string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode:      http.StatusOK,
		IsBase64Encoded: false,
		Body: formatBodyMessage(msg),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

func FailureResponse(statusCode int, msg string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode:      statusCode,
		IsBase64Encoded: false,
		Body: formatBodyMessage(msg),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}