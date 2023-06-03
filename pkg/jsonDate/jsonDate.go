package jsonDate

import (
	"encoding/json"
	"strings"
	"time"
)

type JsonDate time.Time

func parseDateString(dateString string) (time.Time, error) {
	dateAsTime, err := time.Parse("2006-01-02", dateString)

	if err != nil {
		return time.Now(), err
	}

	return dateAsTime, nil
}

func New(dateString string) JsonDate {

	dateAsTime, err := parseDateString(dateString)

	if err != nil {
		return JsonDate(dateAsTime)
	}

	date := JsonDate(dateAsTime)

	return date
}

// Implement Marshaler and Unmarshaler interface
func (j *JsonDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")

	dateAsTime, err := parseDateString(s)

	if err != nil {
		return err
	}

	*j = JsonDate(dateAsTime)

	return nil
}

func (j JsonDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

func (j JsonDate) ToTime() time.Time {
	return time.Time(j)
}
