package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for i := 0; i < 100; i++ {

		// ارسال متن به متغیر متنی
		robotgo.TypeStr("I Love You")

		// فشردن کلید Enter
		robotgo.KeyTap("enter")

		// انتظار برای تاخیر
		time.Sleep(time.Second)
	}

}
