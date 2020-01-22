package main

import (
	"fmt"
)

type TimeFrame struct {
	Name string
	StartRealTimeIn24HourClock int
	EndRealTimeIn24HourClock int
}

const hoursInDay int = 24
const minutesInHour int = 60
const secondsInMinute int = 60
const fiveMinutesBeforeEnd int = 5 * 60

func Name(tf TimeFrame) string {
	return tf.Name
}

func DurationInSeconds(tf TimeFrame) int {
	durationInHours := hoursInDay - tf.StartRealTimeIn24HourClock + tf.EndRealTimeIn24HourClock
	return secondsInMinute * minutesInHour * durationInHours
}

func Dawn(seconds int) bool {
	if seconds == 0 {
		return false
	}
	return fiveMinutesBeforeEnd >= seconds
}

func DisplayTime(s int) string {
	if s % 60 == 0 { // On the minute
		return fmt.Sprintf("\n%d:00", s / 60)
	} else if s < 30 {
		return "!"
	}
	return "."
}

func main() {
	tf := TimeFrame{
		"Night",
		20,
		6,
	}
	fmt.Printf("%+v\n", tf)

	tfd := DurationInSeconds(tf)

	fmt.Println("Starting timer...")
	for i := tfd; i >= 0; i-- {
		if Dawn(i) {
			fmt.Print(DisplayTime(i))
		}
	}
	fmt.Println(" => DONE!!!")
}