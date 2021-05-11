package main

import (
	"fmt"
	"reflect"
)

func main() {
	var students [5]string

	students[1] = "Divyansh"
	fmt.Println(students[1])

	// This is slice
	arr := []int{1, 2, 3}
	arrCopy := arr
	arrCopy[2] = 4

	// This is an array
	workList := [...]int{1, 2, 3, 4}

	copyWorklist := workList

	copyWorklist[2] = 4
	fmt.Printf("value %v, unicode %U type is %v\n", arr, arr, reflect.TypeOf(arr))
	fmt.Printf("value %v, unicode %U type is %v\n", arrCopy, arrCopy, reflect.TypeOf(arrCopy))
	fmt.Printf("value %v, unicode %U type is %v\n", workList, workList, reflect.TypeOf(workList))
	fmt.Printf("value %v, unicode %U type is %v", copyWorklist, copyWorklist, reflect.TypeOf(copyWorklist))
}
