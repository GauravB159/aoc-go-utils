package aocutils

import (
	"fmt"
	"time"
)

func TiemIt[T interface{}, R interface{}](name string, f func(T) R, arg T) R {
	start := time.Now()
	result := f(arg)
	fmt.Println("Answer:", result)
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
	return result
}
