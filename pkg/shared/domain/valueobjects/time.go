package valueobjects

import (
	"time"
)

type CustomTime struct {
	Value time.Time
}

func NewTime() CustomTime {
	return CustomTime{Value: time.Now().UTC()}
}

func (t *CustomTime) ToString() string {
	return t.Value.Format(time.RFC3339)
}
