package main

import (
	"encoding/json"
	gr "github.com/awesome-fc/golang-runtime"
)

func initialize(ctx *gr.FCContext) error {
	fcLogger := gr.GetLogger().WithField("requestId", ctx.RequestID)
	fcLogger.Infoln("init golang!")
	return nil
}

func handler(ctx *gr.FCContext, event []byte) ([]byte, error) {
	fcLogger := gr.GetLogger().WithField("requestId", ctx.RequestID)
	b, err := json.Marshal(ctx)
	if err != nil {
		fcLogger.Error("error:", err)
	}
	fcLogger.Infof("hello golang! \ncontext = %s", string(b))
	return event, nil
}

func main() {
	gr.Start(handler, initialize)
}
