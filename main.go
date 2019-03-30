package main

import "fmt"
import "./space-age"

func main() {
	for _, tc := range testCases {
		fmt.Println("Your age on planet %s is %d", tc.planet, Age(1230746400, tc.planet))
	}
}
