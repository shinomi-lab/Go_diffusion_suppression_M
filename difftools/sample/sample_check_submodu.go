package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	diff "m/difftools/diffusion"
	opt "m/difftools/optimization"
	"os"
	"strings"
	//"time"
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

	//test part
	var S []int
	var hist [][]float64
	//
	// sample_size = 1000000
	// S, hist = sim_submod(adj, sample_size,pop_list,interest_list,assum_list,SeedSet_F,k_T,prob_map)

	// loop_n := 1000
	// sample_size = 1000
	// list1 := []int{0,6,8}
	// list2 := []int{0,20}
	// opt.FocusLoop(loop_n,list1,list2,SeedSet_F,1,sample_size,adj,prob_map,pop_list,interest_list,assum_list)
	//
	// time.Sleep(time.Second * 2)
	// list1 = []int{20,48,2,8}
	// list2 = []int{8,48,37,2}
	// opt.FocusLoop(loop_n,list1,list2,SeedSet_F,1,sample_size,adj,prob_map,pop_list,interest_list,assum_list)
	//
	// time.Sleep(time.Second * 2)
	// list1 = []int{20,15,6}
	// list2 = []int{48,0,18}
	// opt.FocusLoop(loop_n,list1,list2,SeedSet_F,1,sample_size,adj,prob_map,pop_list,interest_list,assum_list)

	filename := "GreedyAndStrict2.csv"
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(f)

	colmns := []string{"greedy_ans", "strict_ans", "greedy_value", "greedy_value2", "strict_value", "strict_value2", "kinjiritu", "random_seed"}
	w.Write(colmns)

	//start loop
	sample_size = 100
	sample_size2 := 100
	var random_seed int64
	random_seed = 0

	//選ばれうる0 1 2 6 8 15 18 20 37 48

	seedsetfs := []int{0, 1, 2, 6, 8, 15, 18, 20, 37, 48}
	for i := 0; i < 10; i++ {
		SeedSet_Greedy := make([]int, len(adj))
		SeedSet_Greedy[seedsetfs[i]] = 1
		//偽情報の発信源を色々と
		for random_seed = 0; random_seed < 10; random_seed++ {
			greedy_ans, greedy_value, greedy_value2 := opt.Greedy(random_seed, sample_size, adj, SeedSet_Greedy, prob_map, pop_list, interest_list, assum_list, 3, true, sample_size2)

			fmt.Println("greedy_ans")
			fmt.Println(greedy_ans, greedy_value, greedy_value2)

			strict_ans, strict_value, strict_value2 := opt.Strict(random_seed, sample_size, adj, SeedSet_Greedy, prob_map, pop_list, interest_list, assum_list, 3, true, sample_size2)

			fmt.Println("strict_ans")
			fmt.Println(strict_ans, strict_value, strict_value2)

			fmt.Println("近似率")
			fmt.Println(greedy_value2 / strict_value2)

			Sets_string := make([][]string, 2)
			Sets_string[0] = opt.Int_to_String(greedy_ans)
			Sets_string[1] = opt.Int_to_String(strict_ans)

			part0 := []string{strings.Join(Sets_string[0], "-"), strings.Join(Sets_string[1], "-")} //here

			a := []float64{greedy_value, greedy_value2, strict_value, strict_value2, greedy_value2 / strict_value2, float64(random_seed)}

			part1 := opt.Float_to_String(a)

			retu := append(part0, part1...)

			w.Write(retu)
		}
	}

	//loop end

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	//SeedSet_Greedy := make([]int,len(adj))
	//SeedSet_Greedy[15] = 2

	fmt.Println(S, hist)
}

func sim_submod(adj [][]int, sample_size int, pop_list [2]int, interest_list [][]int, assum_list [][]int, SeedSet_F []int, k_T int, prob_map [2][2][2][2]float64) ([]int, [][]float64) {
	var S []int
	var hist [][]float64
	S, hist = opt.Check_submod(1, k_T, sample_size, adj, SeedSet_F, prob_map, pop_list, interest_list, assum_list)

	return S, hist
}

func main() {
	sample1()
}
