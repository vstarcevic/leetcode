// Given multiple times, like "10:00am", "11:45pm", "5:00am", "12:01am"
// Find minimum difference in minutes for any of two
package medium

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func StringTimeMinDiff(strArr []string) string {

	minutesOnly := make([]int, len(strArr))
	minDiff := math.MaxInt

	for idx, element := range strArr {
		time := strings.Split(element, ":")
		hour, _ := strconv.Atoi(time[0])
		rest := time[1]
		minutes, _ := strconv.Atoi(rest[0:2])
		amPm := rest[len(rest)-2:]
		if amPm == "am" {
			if hour == 12 {
				hour = 0
			}
		}
		if amPm == "pm" {
			hour += 12
		}
		minutesOnly[idx] = hour*60 + minutes
	}
	slices.Sort(minutesOnly)

	minDiff = minutesOnly[1] - minutesOnly[0]
	for idx, _ := range minutesOnly {
		if idx < 2 {
			continue
		}
		minDiff = int(math.Min(float64(minutesOnly[idx]-minutesOnly[idx-1]), float64(minDiff)))
	}

	specialCase := minutesOnly[0] + 1440 - minutesOnly[len(minutesOnly)-1]
	if specialCase < minDiff {
		minDiff = specialCase
	}

	return strconv.Itoa(minDiff)
}
