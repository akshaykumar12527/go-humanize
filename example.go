package main

import (
	"fmt"
	"github.com/akshaykumar12527/go-humanize"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(humanize.Time(time.Now()))
}
