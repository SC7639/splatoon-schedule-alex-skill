package request

import (
	"encoding/json"

	"github.com/SC7639/splatoon-schedule-alex-skill/api/types"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const url = "https://splatoon3.ink/data/schedules.json"

func getSchedule() (types.SplatoonSchedule, error) {
	var schedule types.SplatoonSchedule
	client := resty.New()

	resp, err := client.R().
		Get(url)
	if err != nil {
		return schedule, errors.Wrap(err, "Failed to get schedule")
	}

	b := resp.Body()
	err = json.Unmarshal(b, &schedule)
	if err != nil {
		return schedule, errors.Wrap(err, "Failed to unmashal schedule")
	}

	return schedule, nil
}
