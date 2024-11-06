package main

import "fmt"

func main() {
	myNum := 0
	fmt.Println("Enter a number")
	fmt.Scanf("%d", &myNum)
	if myNum <= 2 {
		fmt.Println("number must be greater than 2")
	} else {
		myArr := make([]int, myNum)
		myArr[0] = 1
		myArr[1] = 1
		myArr[2] = myArr[0] + myArr[1]
		for i := 3; i < myNum; i++ {
			prevArrayItem := myArr[i-1]
			prevArrayItem2 := myArr[i-2]
			myArr[i] = prevArrayItem + prevArrayItem2
		}
		fmt.Println(myArr)
	}

}
