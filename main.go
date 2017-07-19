package main

import (
	"flag"
)

func main() {
	var city string
	flag.StringVar(&city, "city", "上海", "输入城市名称，如：上海")

	flag.Parse()

	GetInTheater(city)
}
