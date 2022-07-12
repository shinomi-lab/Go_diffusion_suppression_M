package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	diff "m/difftools/diffusion"
	opt "m/difftools/optimization"
)

func sample1() {
	var n int = 50
	var seed int64 = 1
	var K_F int = 5
	var k_T int = 10
	var sample_size int = 100
	var pop_list [2]int
	pop_list[0] = diff.Pop_high
	pop_list[1] = diff.Pop_high

	fmt.Println(n, seed, k_T, K_F, diff.InfoType_F, sample_size, pop_list)



	bytes, err := ioutil.ReadFile("adj_json.txt")
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(bytes))

	var dataJson string = string(bytes)

	arr := make(map[int]map[int]int)
	// var arr []string
	_ = json.Unmarshal([]byte(dataJson), &arr)
	// fmt.Println(arr)

	// fmt.Println(arr[0][1])

	var adj [][]int = make([][]int, n)

	for i := 0; i < n; i++ {
		adj[i] = make([]int, n)
		for j := 0; j < n; j++ {
			adj[i][j] = arr[j][i]
		}
	}

	// fmt.Println(adj)

	var SeedSet_F []int = diff.Make_seedSet_F(n, 1, seed, adj)

	var interest_list [][]int = diff.Make_interest_list(n, seed)

	var assum_list [][]int = diff.Make_assum_list(n, seed)

	var seq [16]float64 = diff.Make_probability()

	var prob_map [2][2][2][2]float64 = diff.Map_probagbility(seq)

	fmt.Println("Seedsetf")
	fmt.Println(SeedSet_F)

	// fmt.Println((seq))
	//
	// fmt.Println(prob_map)

	var S []int
	var hist [][]float64

	S, hist = sim_submod(adj, sample_size,pop_list,interest_list,assum_list,SeedSet_F,k_T,prob_map)

	fmt.Println(S,hist)
}

func sim_submod(adj [][]int, sample_size int, pop_list [2]int, interest_list [][]int, assum_list [][]int, SeedSet_F []int, k_T int, prob_map [2][2][2][2]float64)([]int, [][]float64){
	var S []int
	var hist [][]float64
	S,hist = opt.Check_submod(1, k_T, sample_size, adj, SeedSet_F, prob_map, pop_list, interest_list, assum_list)

	return S,hist
}

func main() {
	sample1()
}
