package valueobjects

import (
	"time"
)

type DateTime struct {
	Value time.Time
}

func NewDateTime(value time.Time) DateTime {
	return DateTime{Value: value}
}

func NewDateTimeNow() DateTime {
	return DateTime{Value: time.Now().UTC()}
}

func (t *DateTime) ToString() string {
	return t.Value.Format(time.RFC3339)
}
