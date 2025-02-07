package main

import (
	"fmt"
	"github-actoins-sample/pkg/timer"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {
	c := cache.New(5*time.Minute, 10*time.Minute)
	c.Set("foo", "bar", cache.DefaultExpiration)
	now := timer.Now()
	fmt.Println("Hello, playground", now)
}
