package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"strconv"
)

type MyEvent struct {
	Path string `json:"path"`
	Body string `json:"body"`
	Counter int `json:"counter"`
}

func HandleRequest(ctx context.Context, e MyEvent) (MyEvent, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	if lc.ClientContext.Custom == nil {
		lc.ClientContext.Custom = make(map[string]string)
	}
	custom := lc.ClientContext.Custom

	val, ok := custom["foo"]
	if ok {
		//do something here
		custom["counter"] = "1"
	}
	old, _ := strconv.Atoi(val)
	newVal := old + 1
	custom["counter"] = string(newVal)
	e.Counter = newVal
	return e, nil
}

func main() {
	lambda.Start(HandleRequest)
}
