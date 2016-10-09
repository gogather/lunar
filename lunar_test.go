package lunarsolar

import (
	"github.com/gogather/com/log"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func Test_Com(t *testing.T) {
	Convey("Lunar Test", t, func() {
		n := time.Now()

		s := TimeToSolar(n)
		printSolar(*s)

		l := SolarToLunar(*s)
		printLunar(*l)
	})
}

func printSolar(s Solar) {
	log.Greenf("阳历 %d-%d-%d\n", s.SolarYear, s.SolarMonth, s.SolarDay)
}

func printLunar(s Lunar) {
	isLeap := ""
	if s.IsLeap {
		isLeap = " (闰年)"
	}
	log.Pinkf("农历 %d-%d-%d%s\n", s.LunarYear, s.LunarMonth, s.LunarDay, isLeap)
}
