package main

import (
	"fmt"
	"time"
)

func Time() {
	n := time.Now()
	fmt.Println(n.String())

	d, err := time.ParseDuration("1h")
	if err != nil {
		panic(err)
	}
	n = n.Add(d)
	fmt.Println(n.String())

	fmt.Println(n.Year(), n.Month(), n.Day(), n.Hour(), n.Minute(), n.Second(), n.Nanosecond(), n.Weekday(), n.YearDay())

	fmt.Println(n.UTC(), n.Unix(), n.UnixNano())

	y, m, day := n.Date()
	h, min, s := n.Clock()
	fmt.Println(y, m, day, h, min, s)

	fmt.Println(n.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))

	fmt.Println(time.Monday.String())
}

func Duration() {
	d, err := time.ParseDuration("1h")
	if err != nil {
		panic(err)
	}
	fmt.Println(d.Hours(), d.Microseconds(), d.Milliseconds(), d.Minutes(), d.Nanoseconds(), d.Seconds(), d.String())

	t, err := time.Parse("20060102", "20201231")
	if err != nil {
		panic(err)
	}

	nd := time.Until(t)
	fmt.Printf("%+v\n", nd)
}

func Ticker() {
	tk := time.NewTicker(1 * time.Second)
	for {
		_ = <-tk.C
		fmt.Println(time.Now().Unix())
	}
}

func main() {
	Time()
	Duration()
	Ticker()
}
