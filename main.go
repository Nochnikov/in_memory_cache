package main

import (
	"fmt"
	"time"

	cache "packages/utils"
)

func main() {
	cache := cache.New[string, int](time.Minute)
	defer cache.Close()

	cache.Set("a", 1, 5*time.Second)
	cache.Set("b", 5432, 0)

	v, ok := cache.Get("a")

	if !ok {
		fmt.Println("Some error accured", ok)
	} else {
		fmt.Print(v)
	}
}
