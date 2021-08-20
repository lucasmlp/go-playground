package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	fmt.Println("Hora UTC:", time.Now())
	fmt.Println("Hora UTC-3:", time.Now().In(loc))
	time.Sleep(time.Minute * 10)
}
