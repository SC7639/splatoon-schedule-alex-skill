package request

import (
	"context"
	"fmt"
	"log"

	"github.com/SC7639/splatoon-schedule-alex-skill/internal/response"
	"github.com/arienmalec/alexa-go"
	"github.com/yassinebenaid/godump"
)

func HandleRequest(ctx context.Context, req alexa.Request) (alexa.Response, error) {
	// Use spew to output the request for debugging purposes:
	// fmt.Println("---- Dumping Input Map: ----")
	// spew.Dump(req)
	// fmt.Println("---- Done ----")

	log.Printf("Request type is %v", req.Body.Intent.Name)
	log.Printf("Request slots is %v", req.Body.Intent.Slots)

	fmt.Println("---- Schedule data ----")
	schedule, err := getSchedule()
	if err != nil {
		log.Panicf("Failed to get schdule %v", err)
	}
	// godump.Dump(schedule.Data.Anarchy)
	fmt.Println("---- Done ----")

	// Create response object
	var resp alexa.Response
	switch req.Body.Intent.Name {
	case "whatson":
		gameType := req.Body.Intent.Slots["gameType"].Value
		gameTypeSettings := getNextGameType(schedule, gameType)
		godump.Dump(gameTypeSettings)
		resp = response.NewWhatsonResponse(gameTypeSettings)

	}

	return resp, nil
}
