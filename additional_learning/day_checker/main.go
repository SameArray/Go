package main

import (
	"fmt"
	"time"
)

func main() {

	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()
	switch time.Now().Weekday() {

	case time.Monday:
		fmt.Println("Сегодня понедельник.")
		fmt.Printf("Сейчас на часах %v часов, %v минут и %v секунд.", hour, minute, second)
		fmt.Printf("\n%v:%v", hour, minute)
	case time.Tuesday:
		fmt.Println("Сегодня вторник.")
		fmt.Printf("Сейчас на часах %v часов, %v минут и %v секунд.", hour, minute, second)
		fmt.Printf("\n%v:%v", hour, minute)
	case time.Wednesday:
		fmt.Println("Сегодня среда.")
		fmt.Printf("Сейчас на часах %v часов, %v минут и %v секунд.", hour, minute, second)
		fmt.Printf("\n%v:%v", hour, minute)
	case time.Thursday:
		fmt.Println("Сегодня четверг.")
		fmt.Printf("Сейчас на часах %v часов, %v минут и %v секунд.", hour, minute, second)
		fmt.Printf("\n%v:%v", hour, minute)
	case time.Friday:
		fmt.Println("Сегодня пятница.")
		fmt.Printf("Сейчас на часах %v часов, %v минут и %v секунд.", hour, minute, second)
		fmt.Printf("\n%v:%v", hour, minute)
	case time.Saturday:
		fmt.Println("Сегодня суббота.")
		fmt.Printf("Сейчас на часах %v часов, %v минут и %v секунд.", hour, minute, second)
		fmt.Printf("\n%v:%v", hour, minute)
	case time.Sunday:
		fmt.Println("Сегодня воскресенье.")
		fmt.Printf("Сейчас на часах %v часов, %v минут и %v секунд.", hour, minute, second)
		fmt.Printf("\n%v:%v", hour, minute)
	}
}
