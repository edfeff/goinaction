package gotype

import (
	"fmt"
	"strings"
	"time"
)

//
func InnerTypeShow() {
	s := " string "
	fmt.Println(s)
	ts := strings.Trim(s, " ")
	fmt.Println(s)
	fmt.Println(ts)
	osTest()

}

func osTest() {
	t := time.Time{}
	fmt.Println(t)

}
