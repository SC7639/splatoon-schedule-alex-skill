package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSchedule(t *testing.T) {
	data, err := getSchedule()
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, data)
}
