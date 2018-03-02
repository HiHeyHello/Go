package myTime

import (
	"fmt"
	"time"
)

func TimeTest(s string) {
	fmt.Println(s)

	wipsd, err := time.Parse("2006/01/02 15:04:05", s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wipsd.Unix())

	t := time.Now()

	for i := 0; i < 2; i++ {
		t = t.AddDate(0, 0, -1)
		fmt.Println(t.Format("2006/01/02 15:04:05"))
		// fmt.Println(t.Format("2006/01/02"))
		fmt.Println(t.Unix(), "\n")
		fmt.Println(t.UTC(), "\n")
	}

	// fmt.Println(t.Format("2006/01/02 15:04:05"))
}
