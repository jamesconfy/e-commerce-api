package service

import "time"

const layout = "2006-01-02 15:04:05"

type TimeService interface {
	// Parse(ti string) time.Time, error
	CurrentTime() string
}

type timeSrv struct{}

func (t timeSrv) Parse(ti string) (time.Time, error) {
	return time.ParseInLocation(layout, ti, time.Local)
}

func (t timeSrv) CurrentTime() string {
	return time.Now().Local().Format(layout)
}

func NewTimeService() TimeService {
	return &timeSrv{}
}
