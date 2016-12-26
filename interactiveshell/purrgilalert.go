package ishell

import (
	"fmt"
	"github.com/fatih/color"
	"strconv"
	"time"
)

func PurrgilAlert(message string) {
	magenta := color.New(color.FgMagenta).SprintFunc()
	t := time.Now()
	hour, minute, second := normalizeTime(t.Hour(), t.Minute(), t.Second())

	fmt.Printf("[%s:%s:%s] %s || %s \n", hour, minute, second, magenta("PURRGIL"), message)
}

func normalizeTime(hour int, minute int, second int) (string, string, string) {
	hourString := strconv.Itoa(hour)
	minuteString := strconv.Itoa(minute)
	secondString := strconv.Itoa(second)

	paddingTime(hour, &hourString)
	paddingTime(minute, &minuteString)
	paddingTime(second, &secondString)

	return hourString, minuteString, secondString
}

func paddingTime(time int, space *string) {
	if time < 10 {
		*space = "0" + *space
	}
}
