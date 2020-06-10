package main

import (
	"fmt"
)

func main() {
	var counter int32
	var subcount int32
	counter = 999

	for counter > 2 {
		subcount = counter - 1
		var flag = true
		for subcount > 1 {
			if counter%subcount != 0 {
				subcount--
				continue
			} else {
				flag = false
				break
			}
		}
		if flag == true {
			fmt.Printf("%v is prime\n", counter)
		}
		counter = counter - 2
	}
}
