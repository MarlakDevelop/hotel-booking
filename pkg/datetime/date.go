package datetime

import (
	"fmt"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

func (date *Date) UnmarshalJSON(content []byte) error {
	layout := "2006-01-02"

	s := strings.Trim(string(content), "\"")

	if s == "null" {
		return nil
	}

	var err error

	date.Time, err = time.Parse(layout, s)

	return fmt.Errorf("time.Parse: %s", err.Error())
}
