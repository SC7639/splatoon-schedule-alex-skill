package types

type AlexaRequest struct {
	Version string `json:"version"`
	Request struct {
		Type   string `json:"type"`
		Time   string `json:"timestamp"`
		Intent struct {
			Name               string `json:"name"`
			ConfirmationStatus string `json:"confirmationstatus"`
		}
	}
}

type AlexaResponse struct {
	Version  string   `json:"version"`
	Response Response `json:"response"`
}

type Response struct {
	OutputSpeech OutputSpeech `json:"outputSpeech"`
}

type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
