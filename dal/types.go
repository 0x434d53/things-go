package dal

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// Time is stored as REAL, Days since 24 Nov 4714 BC.
type NullRealTime struct {
	Time  time.Time
	Valid bool
}

// Support either int or float for date
func (d NullRealTime) Value() (driver.Value, error) {
	if d.Valid {
		return d.Time.Unix(), nil
	}
	return nil, nil
}

func (d *NullRealTime) Scan(value interface{}) error {
	var s float64
	s, ok := value.(float64)
	if ok {
		d.Time = time.Unix(int64(s), 0)
		d.Valid = true
		return nil
	}
	d.Valid = false
	return nil
}

func (d *NullRealTime) String() string {
	if d.Valid {
		return fmt.Sprint(d.Time)
	}
	return ""
}

// Time is stored as int. Unix Time Seconds
type NullIntTime struct {
	Time  time.Time
	Valid bool
}

// Support either int or float for date
func (d NullIntTime) Value() (driver.Value, error) {
	if d.Valid {
		return d.Time.Unix(), nil
	}
	return nil, nil
}

func (d *NullIntTime) Scan(value interface{}) error {
	var s int64
	s, ok := value.(int64)
	if ok {
		d.Time = time.Unix(s, 0)
		d.Valid = true
		return nil
	}
	d.Valid = false
	return nil
}

func (d *NullIntTime) String() string {
	if d.Valid {
		return fmt.Sprint(d.Time)
	}
	return ""
}
