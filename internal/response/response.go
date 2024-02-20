package response

import (
	"github.com/SC7639/splatoon-schedule-alex-skill/api/types"
)

type Response types.AlexaResponse

func CreateResponse() *Response {
	return &Response{
		Version: "1.0",
		Response: types.Response{
			OutputSpeech: types.OutputSpeech{
				Type: "test",
				Text: "Hello. Please override this default output.",
			},
		},
	}
}

func (resp *Response) Say(text string) {
	resp.Response.OutputSpeech.Text = text
}
