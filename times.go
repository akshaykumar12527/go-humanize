package humanize

import (
	"fmt"
	"math"
	"sort"
	"time"
)

// Seconds-based time units
const (
	Minute   = 60
	Hour     = 60 * Minute
	Day      = 24 * Hour
	Week     = 7 * Day
	Month    = 30 * Day
	Year     = 12 * Month
	LongTime = 37 * Year
)

// Time formats a time into a relative string.
//
// Time(someT) -> "3 weeks ago"
func Time(then time.Time) string {
	return RelTime(then, time.Now(), "ago", "from now")
}

var magnitudes = []struct {
	d      int64
	format string
	divby  int64
}{
	{1, "now", 1},
	{2, "1 sec %s", 1},
	{Minute, "%d sec %s", 1},
	{2 * Minute, "1 min %s", 1},
	{Hour, "%d mins %s", Minute},
	{2 * Hour, "1 hr %s", 1},
	{Day, "%d hrs %s", Hour},
	{2 * Day, "1 day %s", 1},
	// {8 * Day, "4 days %s", 1},
	{4 * Day, "%d days %s", Day},
	// {2 * Week, "1 week %s", 1},
	// {Month, "%d weeks %s", Week},
	// {2 * Month, "1 month %s", 1},
	// {Year, "%d months %s", Month},
	// {18 * Month, "1 year %s", 1},
	// {2 * Year, "2 years %s", 1},
	// {LongTime, "%d years %s", Year},
	{math.MaxInt64, "few days %s", 1},
}

// RelTime formats a time into a relative string.
//
// It takes two times and two labels.  In addition to the generic time
// delta string (e.g. 5 minutes), the labels are used applied so that
// the label corresponding to the smaller time is applied.
//
// RelTime(timeInPast, timeInFuture, "earlier", "later") -> "3 weeks earlier"
func RelTime(a, b time.Time, albl, blbl string) string {
	lbl := albl
	diff := b.Unix() - a.Unix()

	after := a.After(b)
	if after {
		lbl = blbl
		diff = a.Unix() - b.Unix()
	}
	fmt.Println(len(magnitudes))
	n := sort.Search(len(magnitudes), func(i int) bool {
		fmt.Println(i, " ", magnitudes[i].d, " ", diff)
		return magnitudes[i].d > diff
	})

	mag := magnitudes[n]
	args := []interface{}{}
	escaped := false
	for _, ch := range mag.format {
		fmt.Println(string(ch))
		if escaped {
			switch ch {
			case '%':
			case 's':
				args = append(args, lbl)
			case 'd':
				args = append(args, diff/mag.divby)
			}
			escaped = false
		} else {
			escaped = ch == '%'
		}
	}
	return fmt.Sprintf(mag.format, args...)
}
