package main

import (
	"fmt"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
)

func call() {
	accountSid := "AC33c65894f3bd633a6cf0d76645687de0"
	authToken := os.Getenv("CallAuthToken")

	//urlStr := "http://demo.twilio.com/docs/voice.xml"

	from := "+15407128161"
	to := "+919654674248"

	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &openapi.CreateCallParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetUrl("http://twimlets.com/holdmusic?Bucket=com.twilio.music.ambient")

	resp, err := client.ApiV2010.CreateCall(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Call Status: " + *resp.Status)
		fmt.Println("Call Sid: " + *resp.Sid)
		fmt.Println("Call Direction: " + *resp.Direction)
	}
}
