package main

import (
	"fmt"
	"profiling-demo/utils"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1024), profile.ProfilePath(".")).Stop()
	primes := GetPrimes(1, 100000)
	fmt.Println(primes)
}

func GetPrimes(start, end int) []int {
	var primes [10000]int
	var count = 0
	for i := start; i <= end; i++ {
		if utils.IsPrime(i) {
			//primes = append(primes, i)
			primes[count] = i
			count++
		}
	}
	//return primes
	return primes[:count]
}
