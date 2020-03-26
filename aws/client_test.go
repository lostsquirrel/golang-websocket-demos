package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"testing"
)

//func TestJson(t *testing.T) {
//	c := MyMessage{
//		Path: "json",
//		Body: "content",
//	}
//	bob, err := json.Marshal(c)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Print(string(bob))
//}

func TestProxyResponse(t *testing.T) {
	resp := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body: "string(bb)",
	}
	fmt.Print(resp)
}