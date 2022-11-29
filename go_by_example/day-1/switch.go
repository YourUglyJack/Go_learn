package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("un")
	case 2:
		fmt.Println("deux")
	case 3:
		fmt.Println("trois")
	}

	fmt.Println("Now:", time.Now())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I'm bool")
		case int:
			fmt.Println("I'm int")
		default:
			fmt.Println("Je ne sais pas")
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI(3.14)
}
