package util

import (
	"database/sql/driver"
	"strconv"
	"time"
)

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(b []byte) error {
	date, err := ParseDate(string(b[1 : len(b)-1]))
	if err != nil {
		return err
	}

	*d = *date

	return nil
}

func (d *Date) String() string {
	return d.Format("2006-01-02")
}

func (d *Date) Scan(src interface{}) error {
	date, err := ParseDate(string(src.([]byte)))
	if err != nil {
		return err
	}

	*d = *date

	return nil
}

func (d *Date) Value() (driver.Value, error) {
	return d.String(), nil
}

func ParseDate(date string) (*Date, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}

	return &Date{t}, nil
}

func IntToTime(time int) string {
	hours := time / 60
	minutes := time % 60

	var result string

	if hours > 9 {
		result = strconv.Itoa(hours) + ":"
	} else {
		result = "0" + strconv.Itoa(hours) + ":"
	}

	if minutes > 9 {
		return result + strconv.Itoa(minutes)
	}

	return result + "0" + strconv.Itoa(minutes)
}
