package main

import (
	"fmt"
	"net/http"
)

func main() {
	//http.HandleFunc("/", indexHandler)
	//http.HandleFunc("/hello", helloHandler)
	//log.Fatal(http.ListenAndServe(":9999", nil))


	//2021-02-24:leecode  maxSatisfied
	//cus := []int{1,0,1,2,1,1,7,5}
	//gru := []int{0,1,0,1,0,1,0,1}
	//x := 3
	//var ret int
	//ret = maxSatisfied(cus,gru,x)
	//fmt.Println(ret)
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
//2021-02-23:leecode  maxSatisfied
func maxSatisfied(customers []int, grumpy []int, X int) int {
	length := len(customers)
	satisfyNum := 0
	for i, num:= range customers {
		if grumpy[i] == 0 {
			satisfyNum += num
		}
	}

	maxUnSatisfyNum,unSatisfyNum := 0,0

	for i:=0; i < length; i++{
		unSatisfyNum = 0
		for j,k:=i,X+i; j<k && j<length; j++ {
			if grumpy[j] == 1 {
				unSatisfyNum += customers[j]
			}
			//fmt.Printf("i is %d, unSatisfyNum is %d\n",i, unSatisfyNum )
		}

		if maxUnSatisfyNum < unSatisfyNum {
			maxUnSatisfyNum = unSatisfyNum
		}
	}
	//fmt.Printf("satisfy num is %d\n", satisfyNum )
	//fmt.Printf("maxUnSatisfyNum num is %d\n", maxUnSatisfyNum )
	return maxUnSatisfyNum+satisfyNum
}

//2021-02-24:leecode  maxSatisfied
func flipAndInvertImage(A [][]int ) [][]int {
	for _, row :=range A  {
		left, right := 0, len(row)-1
		for left < right {
			if row[left] == row[right] {
				row[left] ^= 1
				row[right] ^= 1
			}
			left++
			right--
		}
		if left == right {
			row[left] ^= 1
		}
	}

	return A
}


