package main

import (
	_ "beegoPro/routers"

	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

func test(a int, sb chan int) {
	c := a + 8
	sb <- c
}

func swap(x, y int) (int, int) {
	return x, y
}

func main() {
	sb := make(chan int)
	go test(7, sb)
	x, y := swap(6, 9)
	fmt.Printf("%v\n", <-sb)
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)

	arr := []int{1, 2, 4, 56, 7}
	var num = 0
	for i := 0; i < len(arr); i++ {
		num += arr[i]
	}

	beego.Run()
}
