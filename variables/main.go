package main

import (
	"fmt"
	"learning-go/basic_syntax"
	"strconv"
)

/*
1) Visibility of variables with naming conventino
	lower case first letter for package level scope
	upper case first letter for exporting (available across multiple packages)

2) Type conversions
	destinationType(variable)
	strconv for strings
*/

// Package level variables
var packageLevelVariable int = 10

// Grouping vairables

var (
	userName string = "divyanshbhowmick"
	age      int    = 21
	fullName string = "Divyansh Bhowmick"
	role     string = "developer"
)

var (
	readPermission   bool = true
	writePermission  bool = true
	deletePermission bool = false
)

var shadowedVariable int = 10

func main() {
	packageLevelVariable = 100

	var shadowedVariable int = 100
	// shadowedVariable := 100

	fmt.Println("shadowedVariable " + strconv.Itoa(shadowedVariable))

	// First Type
	var i int
	i = 7
	fmt.Println(i)
	i = 77
	// Second Type
	var j int = 14
	// Third Type
	k := 21
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(packageLevelVariable)

	message := basic_syntax.Hello("Divyansh!")
	fmt.Println(message)
}
