package response

import (
	"fmt"
	"time"

	"github.com/SC7639/splatoon-schedule-alex-skill/api/types"
	"github.com/arienmalec/alexa-go"
)

func NewWhatsonResponse(gameTypeSettings []types.GameTypeSettings) alexa.Response {
	text := parseDataToSpeech(gameTypeSettings)
	return alexa.Response{
		Version: "1.0",
		Body: alexa.ResBody{
			OutputSpeech: &alexa.Payload{
				Type: "PlainText",
				Text: text,
			},
		},
	}
}

func parseDataToSpeech(gameTypeSettings []types.GameTypeSettings) string {
	var text string
	for _, gameTypeSetting := range gameTypeSettings {
		startTime := formatTime(gameTypeSetting.Node.StartTime)
		endTime := formatTime(gameTypeSetting.Node.EndTime)
		text += fmt.Sprintf("%s is at %s to %s\n", gameTypeSetting.Setting.VsRule.Name, startTime, endTime)
		break
	}
	return text
}

func formatTime(timeStr string) string {
	// Try parsing as RFC3339 (ISO 8601) format
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		// If parsing fails, return the original string
		return timeStr
	}
	// Format as 12-hour time with AM/PM (e.g., "3:00 PM")
	return t.Format("3:04 PM")
}
