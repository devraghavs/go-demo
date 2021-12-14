package main

import ("fmt"
"sort")

func try(){
	arr:= []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(arr)
	// array append
	arr=append(arr,11)
	fmt.Println(arr[1:5])
	fmt.Println(arr)
	// for e:=0; e<len(arr); e++{
	// 	if arr[e]%2==0{
	// 		fmt.Println(arr[e])
	// 	}
	
	// }
	// for _,v:=range arr{
	// 	if v%2!=0{
	// 		fmt.Println(v)
	// 	}
	// }
	// fmt.Println(arr[1:5]==nil)

	// multidimensional array
	// s1:= [][] int{
	// 	{1,2},
	// 	{3,4},
	// }
	// fmt.Println(s1[0][1])

	//sorting the array

	s2:=[]int{10,2,4,40,20,4,1}
	sort.Ints(s2) //need to use selector for sorting
	fmt.Println(s2)
}

func main() {
	fmt.Println("Hello, World!")
	try()
}