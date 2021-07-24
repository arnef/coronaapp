package covpass

import "time"

func ParseDay(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

func Today() time.Time {
	today, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	return today
}
