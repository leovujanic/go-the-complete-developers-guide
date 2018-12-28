package main

import (
	"fmt"
	"strings"
	"time"
)

const PirChar string = "*"
const Fill = "."
const SleepTime = 80 * time.Millisecond

func pyramid(lines int) {
	for i := 1; i <= lines; i++ {

		dots := strings.Repeat(PirChar, i)
		space := strings.Repeat(Fill, (lines-i)*2)
		fmt.Printf("%s%s%s\n", dots, space, dots)
		time.Sleep(SleepTime)
	}
}

func reversePyramid(lines int) {
	for i := lines; i > 0; i-- {
		dots := strings.Repeat(PirChar, i)
		space := strings.Repeat(Fill, (lines-i)*2)
		fmt.Printf("%s%s%s\n", dots, space, dots)
		time.Sleep(SleepTime)
	}
}

func main() {

	cnt := 20
	mara := "MARA"
	maraSeparator := strings.Repeat(Fill, cnt-len(mara))
	maraString := fmt.Sprintf("*%s %s %s*", maraSeparator, mara, maraSeparator)
	for i := 0; i < 10; i++ {
		reversePyramid(cnt)
		fmt.Println(maraString)
		pyramid(cnt)
	}

}
