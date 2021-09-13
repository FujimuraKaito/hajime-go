package main

import (
	f "fmt"
	"gosample"
)

func main() {
	f.Println(gosample.Message)

	message := "temp message"
	f.Println(message)

	// constant
	const Hello string = "hello world"
	f.Println(Hello)

	// ゼロ値（型によって初期化される値）
	var i int
	f.Println(i) // 0

	a, b := 10, 100
	if a > b {
		f.Println("a is larger than b")
	} else if a < b {
		f.Println("a is smaller than b")
	} else {
		f.Println("a equals b")
	}

	for i := 0; i < 3; i++ {
		f.Println(i)
	}

	// how to use while
	n := 0
	for n < 5 {
		f.Printf("n = %d\n", n)
		n++
	}

	n = 0
	for {
		n++
		if n > 10 {
			break
		}
		if n%2 == 0 {
			// 偶数なら処理を飛ばす
			continue
		}
		f.Println(n)
	}

}
