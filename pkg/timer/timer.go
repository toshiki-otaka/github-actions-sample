package timer

import "time"

func Now() time.Time {
	return time.Now()
}

func checkTime(t time.Time) error {
	var err error
	if t.Before(time.Now()) {
		return err
	}
	return nil
}
