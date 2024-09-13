package main

import (
	"fmt"
	cache "packages/utils"
)

func main() {
	cache := cache.New()

	

	cache.Set("1", []int{1, 2, 3})
	cache.Set("2", map[string]int{
		"Test1": 15,
		"Test2": 16,
		"Test3": 17,})

	fmt.Println(cache.Get("1"))
	fmt.Println(cache.Get("2"))

	cache.Delete("1")
	cache.Delete("2")
	fmt.Println(cache.Get("1"))
	fmt.Println(cache.Get("2"))


}
