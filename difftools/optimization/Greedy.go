package optimization

import (
	// "os"
	"fmt"
	diff "m/difftools/diffusion"
)

func Greedy(seed int64, sample_size int, adj [][]int, Seed_set []int, prob_map [2][2][2][2]float64, pop [2]int, interest_list [][]int, assum_list [][]int, ans_len int, Count_true bool, sample_size2 int) ([]int, float64, []float64) {
	//sample_size2はグリーディで求めた解をより詳しくやる
	var n int = len(adj)
	var max float64 = 0
	var result float64
	var index int
	var ans []int
	var ans_v []float64

	ans = make([]int, 0, ans_len)
	S := make([]int, len(Seed_set))
	_ = copy(S, Seed_set)
	S_test := make([]int, len(Seed_set))
	_ = copy(S_test, Seed_set)

	var info_num int

	if Count_true {
		info_num = 2
	} else {
		info_num = 1
	}

	for i := 0; i < ans_len; i++ {
		fmt.Println(i)
		max = 0
		for j := 0; j < n; j++ {
			if (j+1)%100 == 0 {
				fmt.Println(i, "-", (j+1)/100)
			}
			_ = copy(S_test, S)
			if S_test[j] != 0 { //すでに発信源のユーザだったら
				continue
			}
			S_test[j] = info_num

			dist := Infl_prop_exp(seed, sample_size, adj, S_test, prob_map, pop, interest_list, assum_list)
			if Count_true {
				result = dist[diff.InfoType_T]
			} else {
				result = dist[diff.InfoType_F]
			}

			if result > max {
				max = result
				index = j
			}
		} //subloop end

		ans = append(ans, index)
		ans_v = append(ans_v, max)
		S[index] = info_num

	} //mainloop end

	// var max_2 float64
	// dist2 := Infl_prop_exp(seed, sample_size2, adj, S, prob_map, pop, interest_list, assum_list)
	// if Count_true {
	// 	max_2 = dist2[diff.InfoType_T]
	// } else {
	// 	max_2 = dist2[diff.InfoType_F]
	// }
	return ans, max, ans_v
}

func Greedy_exp(seed int64, sample_size int, adj [][]int, Seed_set []int, prob_map [2][2][2][2]float64, pop [2]int, interest_list [][]int, assum_list [][]int, ans_len int, Count_true bool, capacity float64, max_user int, OnlyInfler bool, user_weight float64)([]int, []float64) {
	//sample_size2はグリーディで求めた解をより詳しくやる
	var n int = len(adj)
	var max float64 = 0
	var result float64
	var index int
	var ans []int
	var ans_v []float64

	ans = make([]int, 0, ans_len)
	S := make([]int, len(Seed_set))
	_ = copy(S, Seed_set)
	S_test := make([]int, len(Seed_set))
	_ = copy(S_test, Seed_set)
	cap_use := capacity

	var info_num int

	if Count_true {
		info_num = 2
	} else {
		info_num = 1
	}

	for{
		max = -1
		for j := 0; j < n; j++ {

			_ = copy(S_test, S)//初期化
			for i:=0; i< len(ans); i++{//初期設定
				S_test[ans[i]] = info_num
			}
			if S_test[j] != 0 { //すでに発信源のユーザだったら
				continue
			}
			if(OnlyInfler){
				if(FolowerSize(adj,j)==0){
					continue
				}
			}
			if Cal_cost(user_weight, 1-user_weight, adj, j, max_user)>cap_use{//コストが大きすぎるユーザなら
				continue
			}
			S_test[j] = info_num

			dist := Infl_prop_exp(seed, sample_size, adj, S_test, prob_map, pop, interest_list, assum_list)
			if Count_true {
				result = dist[diff.InfoType_T]/Cal_cost(user_weight,1-user_weight,adj,j,max_user)
			} else {
				result = dist[diff.InfoType_F]/Cal_cost(user_weight,1-user_weight,adj,j,max_user)
			}//must change 重要 修正

			if result > max {
				max = result
				index = j
			}
		} //subloop end

		if max == -1{
			break;
		}
		ans = append(ans, index)
		cap_use -= Cal_cost(user_weight,1-user_weight,adj,index,max_user)

		S[index] = info_num
	}
	return ans, ans_v
}

func Cal_cost(u_weight float64, f_wight float64,adj [][]int,node int, max_user int)float64{
	f := FolowerSize(adj,node)
	u_max := len(adj)
	f_max := 0
	for i:=0; i<u_max; i++{
		if(adj[max_user][i] == 1){
			f_max ++
		}
	}
	// fmt.Println("folowersize",f)
	// fmt.Println("f:",f,"f_wight:",f_wight,"f_max:",f_max,"u_weight:",u_weight)
	return float64(f)*f_wight/float64(f_max) + u_weight/2
}
