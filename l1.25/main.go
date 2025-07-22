package main

import (
	"fmt"
	"time"
)

func sleep(second int) {
	<-time.After(time.Duration(second) * time.Second)
	fmt.Println("Прошло времени: ", second)
}
func main() {
	second := 3
	sleep(second)
}
