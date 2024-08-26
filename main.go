package main

import (
	"fmt"
	cache "packages/utils"
)

func main() {
	cache := cache.New() 

	cache.Set("userId", 42)
	
	userId := cache.Get("userId")

	fmt.Printf("User info:%+v ",userId)
	
	cache.Delete("userId")

	userId = cache.Get("userId")

	fmt.Printf("User info:%+v", userId)
}
