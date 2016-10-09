package lunarsolar

import (
	"time"
)

func TimeToSolar(t time.Time) *Solar {
	return &Solar{SolarDay: int(t.Day()), SolarMonth: int(t.Month()), SolarYear: int(t.Year())}
}
