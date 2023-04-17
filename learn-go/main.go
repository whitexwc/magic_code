package main

import "fmt"

func main() {
	//  array := []string{"randy", "pausch", "json", "xiao", "wc"}

	// array2 := [...]string{"hello", "new", "array"}

	//  modify(array)
	// fmt.Println(array)
	// fmt.Println(len(array2))

	//slice(起始指针，长度，容量)
	// slice := []string{"wu", "xiao"}
	// fmt.Println(slice)

	// //make会初始0值    预留了100个值的空间，但是目前仅赋值两个
	// newSlice := make([]int, 2, 100)
	// fmt.Println(newSlice)

	// //测试扩容, 预留空间不够时自动扩成2倍 扩容会把原内容移到新地址，会有开销
	// mySlice := make([]int, 2)
	// fmt.Println(mySlice, len(mySlice), cap(mySlice))
	// mySlice = append(mySlice, 10)
	// fmt.Println(mySlice, len(mySlice), cap(mySlice))

	// modify(mySlice)
	// fmt.Println(mySlice)

	//subslice
	slice := []string{"hello", "ustc", "today", "is", "a", "good", "day"}
	subSlice := slice[:3]
	fmt.Println(subSlice)

	fmt.Println(slice)
	modify(subSlice)
	fmt.Println(slice)

	//映射map
	scores := map[string]int{
		"ming" :10,
		"jon" : 66,
	}
	fmt.Println(scores["ming"])
	fmt.Println(scores["jon"])
}

func modify(slice []string) {
	slice[0] = "gtmd"
}
