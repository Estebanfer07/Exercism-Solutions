package clock

import (
	"fmt"
	"strconv"
)

// Define the Clock type here.
type Clock struct {
	hour    int
	minutes int
}

func TransformMinutes(min int) (h int, m int) {
	hours := min / 60
	remainingMinutes := min % 60

	if remainingMinutes < 0 {
		remainingMinutes += 60
		hours--
	}

	return hours, remainingMinutes
}

func FormatHours(h int) int {

	if h >= 0 && h < 24 {
		return h
	}
	if h >= 24 {
		h -= 24
	}
	if h < 0 {
		h += 24
	}

	return FormatHours(h)
}

func New(h, m int) Clock {

	extraHours, minutes := TransformMinutes(m)

	return Clock{
		hour:    FormatHours(h + extraHours),
		minutes: minutes,
	}
}

func (c Clock) Add(m int) Clock {

	extraHours, minutes := TransformMinutes(c.minutes + m)

	return Clock{
		hour:    FormatHours(c.hour + extraHours),
		minutes: minutes,
	}

}

func (c Clock) Subtract(m int) Clock {
	extraHours, minutes := TransformMinutes(c.minutes - m)

	return Clock{
		hour:    FormatHours(c.hour + extraHours),
		minutes: minutes,
	}
}

func (c Clock) String() string {

	strHour := strconv.Itoa(c.hour)
	strMin := strconv.Itoa(c.minutes)
	if c.hour < 10 {
		strHour = "0" + strHour
	}
	if c.minutes < 10 {
		strMin = "0" + strMin
	}

	return fmt.Sprintf("%s:%s", strHour, strMin)
}
