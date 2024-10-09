package aocutils

import (
	"fmt"
	"time"
)

func Timer[T interface{}, R interface{}](name string, f func(T) R, arg T) R {
	start := time.Now()
	result := f(arg)
	fmt.Println("Answer:", result)
	fmt.Println()
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
	return result
}
