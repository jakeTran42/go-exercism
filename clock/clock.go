// clock.go

package clock

import (
	"strconv"
	"time"
)

// Clock is a simple 24 hour clock that rolls over. Does not keep track of the date,
// only the time of day.
type Clock struct {
	hour   int
	minute int
}

// New takes in two ints, an hour and a minute, and returns a Clock struct.
// Corrects for hours and minutes that are negative, or roll over.
func New(hour int, minute int) Clock {
	newTime := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	difference := time.Duration((minute + (hour * 60)) * 60000000000)
	newTime = newTime.Add(difference)
	clock := Clock{hour: newTime.Hour(), minute: newTime.Minute()}
	return clock
}

// String returns a string representation of the clock.
func (clock Clock) String() string {
	hourStr := "0"
	minuteStr := "0"

	hourConv := strconv.Itoa(clock.hour)
	minuteConv := strconv.Itoa(clock.minute)

	if clock.hour < 10 {
		hourStr += hourConv
	} else {
		hourStr = hourConv
	}

	if clock.minute < 10 {
		minuteStr += minuteConv
	} else {
		minuteStr = minuteConv
	}

	representation := hourStr + ":" + minuteStr
	return representation
}

// Add takes in a number of minutes n of type int and returns a Clock struct representing a time
// n minutes after the Clock's current time.
func (clock Clock) Add(minutes int) Clock {
	begin := time.Date(2000, time.January, 1, clock.hour, clock.minute, 0, 0, time.UTC)
	difference := time.Duration((minutes) * 60000000000)
	end := begin.Add(difference)
	final := Clock{hour: end.Hour(), minute: end.Minute()}
	return final
}

// Subtract takes in a number of minutes n of type int and returns a Clock struct representing a time
// n minutes before the Clock's current time.
func (clock Clock) Subtract(minutes int) Clock {
	begin := time.Date(2000, time.January, 1, clock.hour, clock.minute, 0, 0, time.UTC)
	difference := time.Duration((minutes * -1) * 60000000000)
	end := begin.Add(difference)
	final := Clock{hour: end.Hour(), minute: end.Minute()}
	return final
}
