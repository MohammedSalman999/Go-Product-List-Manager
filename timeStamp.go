package main

import (
	"strconv"
	"time"
)

type timestamp struct {
	time.Time
}

func (ts timestamp) String() string {
	if ts.IsZero() {
		return "Unknown"
	}
	const layout = "2006/01"
	return ts.Format(layout)
}

func isTimeStamp(v interface{}) (ts timestamp) {
	var t int
	switch v := v.(type) {
	case int:
		t = v
	case string:
		t , _ = strconv.Atoi(v)
	}

	ts.Time = time.Unix(int64(t),0)

	return ts
}