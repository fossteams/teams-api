package api

import (
	"encoding/json"
	"fmt"
	"time"
)

type RFC3339Time time.Time

var _ json.Unmarshaler = &RFC3339Time{}
var _ json.Marshaler = &RFC3339Time{}

func (t *RFC3339Time) MarshalJSON() ([]byte, error){
	return []byte(fmt.Sprintf("\"%s\"", time.Time(*t).Format(time.RFC3339Nano))), nil
}

func (t *RFC3339Time) UnmarshalJSON(bytes []byte) error {
	var str string
	err := json.Unmarshal(bytes, &str)
	if err != nil {
		return err
	}

	if str == "" {
		return nil
	}

	parsed, err := time.Parse(time.RFC3339Nano, str)
	if err != nil {
		return err
	}
	*t = RFC3339Time(parsed)
	return nil
}