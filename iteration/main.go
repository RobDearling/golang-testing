package iteration

import "time"

func Repeat(character string, amountOfTimes int) string {
	var repeated string
	for i := 0; i < amountOfTimes; i++ {
		repeated += character
	}

	return repeated
}

func RepeatWithSleep(character string, amountOfTimes int) string {
	var repeated string
	for i := 0; i < amountOfTimes; i++ {
		time.Sleep(10 * time.Millisecond)
		repeated += character
	}

	return repeated
}
