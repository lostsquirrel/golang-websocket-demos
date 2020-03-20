package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

type MyEvent struct {
	Path string `json:"path"`
	Body string `json:"body"`
}

func HandleRequest(ctx context.Context, name MyEvent) (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body: name.Body,
	}
	return resp, nil
}

func main() {
	lambda.Start(HandleRequest)
}
