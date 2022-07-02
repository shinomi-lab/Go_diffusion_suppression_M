package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	diff "m/difftools/diffusion"
	"os"
	"reflect"
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

	file, err := os.Open("adjdf.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(file)

	fmt.Println(r)

	rows, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range rows {
		fmt.Println(reflect.TypeOf(v))
	}

	print(reflect.TypeOf(rows))

	// fp, err := os.Open("adj_json.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer fp.Close()

	// buf := make([]byte, 64)
	// for {
	// 	n, err := fp.Read(buf)
	// 	if n == 0 {
	// 		break
	// 	}
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// }
	// fmt.Println(string(buf))

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

	fmt.Println(arr[0][1])

	var adj [50][50]int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			adj[i][j] = arr[i][j]
		}
	}

	// fmt.Println(adj)

	var SeedSet_F []int = diff.Make_seedSet_F(n, K_F, seed)

	var interest_list [][]int = diff.Make_interest_list(n, seed)

	var assum_list [][]int = diff.Make_assum_list(n, seed)

	var seq [16]float64 = diff.Make_probability()

	var prob_map [2][2][2][2]float64 = diff.Map_probagbility(seq)

	fmt.Println(SeedSet_F, interest_list, assum_list)

	fmt.Println((seq))

	fmt.Println(prob_map)
}

func main() {
	sample1()
}
