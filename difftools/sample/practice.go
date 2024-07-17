//このファイルはgo言語の使い方を実験的に調べるために球に使うファイルです

package main

import(
  "fmt"
  "encoding/json"
  "io/ioutil"
  "math/rand"
  "os"
  "reflect"
)

func main(){

  for i:=0;i<10;i++{
    fmt.Println(i)
  }
  for i:=0;i<10;i++{
    fmt.Println(i)
  }
  var intn float64
  intn = 1.1
  var floatn float64
  floatn = 555.5
  fmt.Println(floatn/intn)

  list1n := 3
  list2n := 4

  dp := make([][]float64,list1n)
  for listi:=0;listi<list1n;listi++{
    dp[listi] = make([]float64,list2n)
  }

  dp[2][3] = 2.2

  fmt.Println(dp)

  os.Exit(0)
  n := 5
  max_users := make([]int,n)

	user_num_counter := 0
	max_user_nums := make([]int,n)
  list := []int{22, 5, 2, 1, 15, 11, 18}
  for i:=0;i<len(list);i++{
    user_num_counter = list[i]
    for j:=0;j<n;j++{
      if(max_user_nums[j] < user_num_counter){
        m := n-1
        for k:=j;k<m;k++{
          max_users[m-k+j] = max_users[m-k+j-1]
          max_user_nums[m-k+j] = max_user_nums[m-k+j-1]
          fmt.Println(i,max_user_nums)
        }
        max_users[j] = i
        max_user_nums[j] = user_num_counter
        fmt.Println(i,max_user_nums)

        break
      }
  }
}
  fmt.Println(max_users,max_user_nums)
  os.Exit(0)

  slice := [][]int{{1, 15, 18}, {1, 15, 18}}
  slice2 := [][]int{{1, 15, 18}, {1, 15, 18}}

  fmt.Println(reflect.DeepEqual(slice, slice2))
  os.Exit(0)

  fmt.Println("practice")

  for i:=0;i<20;i++{
    rand.Seed(1)
    MakeRand()
  }
  os.Exit(0)

  node_list_path := "Python_random_nodelists/node_list.txt"
	bytes, err := ioutil.ReadFile(node_list_path)
	if err != nil {
		panic(err)
	}

  var dataJson string = string(bytes)

	arr := make(map[int]map[int]int)
	// var arr []string
	_ = json.Unmarshal([]byte(dataJson), &arr)
	// fmt.Println(arr)

	fmt.Println(arr[2][0])
  fmt.Println(arr)
  // n := len(arr[0])
	var adj [][]int = make([][]int, len(arr))

	//make adj
	for i := 0; i < len(arr); i++ {
		adj[i] = make([]int, len(arr[i]))
		for j := 0; j < len(arr[i]); j++ {
			adj[i][j] = arr[i][j]
      // print(adj[i][j])
		}
    // print("---------------")
	}
  fmt.Println(adj)
}


func MakeRand(){
  for j:=0;j<4;j++{

    fmt.Println(rand.Float64())
  }

}
