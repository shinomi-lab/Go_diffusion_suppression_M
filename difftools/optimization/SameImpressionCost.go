package optimization

import (
	"os"
	"fmt"
	diff "m/difftools/diffusion"
  "bufio"
)

func SameImpressionCost(seed int64, sample_size int, adj [][]int, non_use_list []int, prob_map [2][2][2][2]float64, pop [2]int, interest_list [][]int, assum_list [][]int, OnlyInfler bool){
	//sample_size2はグリーディで求めた解をより詳しくやる
	var n int = len(adj)
	var result float64


	S := make([]int, len(adj))
	S_test := make([]int, len(adj))


	var info_num int


	info_num = 2

  file, err := os.Create("SameImporessionCost.csv")
  if err != nil {
      fmt.Println("Error creating file:", err)
      return
  }
  // 関数の終了時にファイルを閉じる
  defer file.Close()
  writer := bufio.NewWriter(file)

	for j := 0; j < n; j++ {

		_ = copy(S_test, S)//初期化

		if S_test[j] != 0 { //すでに発信源のユーザだったら
      fmt.Println("error in SameImpressionCost.go")
      os.Exit(0)
		}
    if(contains(non_use_list,j)){
      continue
    }
		if(OnlyInfler){
			if(FolowerSize(adj,j)==10000000000000){
				continue
			}
		}

		S_test[j] = info_num

		dist := Infl_prop_exp(seed, sample_size, adj, S_test, prob_map, pop, interest_list, assum_list)

		result = dist[diff.InfoType_T]



    // ファイルを作成または開く


    // ファイルに書き込む
    _, err = writer.WriteString(fmt.Sprintf("%d",j)+","+fmt.Sprintf("%d",FolowerSize(adj,j))+","+fmt.Sprintf("%f",result)+"\n")
    // _, err = file.WriteString("aaaaa\n")
    // fmt.Println("aa")
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    // fmt.Println("File written successfully")
	}
  err = writer.Flush()
    if err != nil {
        fmt.Println("Error flushing buffer:", err)
        return
    }
}


func contains(slice []int, target int) bool {
    for _, element := range slice {
        if element == target {
            return true
        }
    }
    return false
}
