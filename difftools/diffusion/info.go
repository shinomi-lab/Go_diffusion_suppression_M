package diffusion

import (
	"fmt"
	"math/rand"
)

var InfoType_F int = 0
var InfoType_T int = 1
var InfoTypes_n int = 2
var Pop_low int = 0
var Pop_high int = 1
var Pops_n int = 2

// func make_Info(pop int) {
// 	var a [InfoTypes_n][pops_n]int
// }

// func make_InfoTypes() [2]int{
// 	a [InfoTypes_n]int := [InfoType_F,InfoType_T]
// 	return a
// }

func Make_seedSet_F(n int, k int, seed int64) []int {
	rand.Seed(seed)
	Fs := make([]int, n)
	for i := 0; i < k; {
		n := rand.Intn(n)
		if Fs[n] == 0 {
			Fs[n] = 1
			i++
		}
	}
	fmt.Println(rand.Intn(n))

	return Fs
}
