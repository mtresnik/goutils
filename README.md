# goutils
[![build status](https://github.com/mtresnik/goutils/actions/workflows/go.yml/badge.svg)](https://github.com/mtresnik/goutils/actions/workflows/go.yml/)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://github.com/mtresnik/goutils/blob/main/LICENSE)
[![version](https://img.shields.io/badge/version-1.1.11-blue)](https://github.com/mtresnik/goutils/releases/tag/v1.1.11)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-green.svg?style=flat-square)](https://makeapullrequest.com)
<hr>

Goutils (pronounced gout-ils) is a root sdk package for my other go projects.


### Sample Code

In your project run:
```
go mod download github.com/mtresnik/goutils
```

Your `go.mod` file should look like this:
```go 
module mymodule

go 1.23.3

require github.com/mtresnik/goutils v1.1.11
```


Then in your go files you should be able to access sdk:

```go 
package main

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"strings"
)

func main() {
	// nested arrays
	// [[[0]] 1 2 3 [[[4]]] [5 6 7]]
	var tempArray = []any{[]any{[]any{0}}, 1, 2, 3, []any{[]any{[]any{4}}}, []any{5, 6, 7}}
	// [0 1 2 3 4 5 6 7]
	var flattened = goutils.DeepFlatten(tempArray)
	println(flattened)
	var intSlice = goutils.MapToInts(flattened)
	sum := goutils.SumBy(intSlice, func(i int) float64 { return float64(i) })
	println("sum:", sum)
	
	// [10, 20, 50, 90, 20, 10]
	var toZip = []int{10,20,50,90,20,10}
	// [10, 30, 40, -70, -10]
	deltas := goutils.ZipWithNext(toZip, func(i1 int, i2 int) int { return i2 - i1 })
	println("deltas", goutils.ArrayToString(deltas))
	
	expected := []float64{1,2,3,4,3,2,1}
	actual := []float64{1.1, 1.9, 2.5, 3.5, 2, 1.1}
	squaredDistance := goutils.Zip(expected, actual, func(ex float64, ac float64) float64 {
		diff := ac - ex
		return diff * diff
    })
	totalError := goutils.SumOf(squaredDistance...)
	println("totalError:", totalError)
	
	
	// sets
	tempSet := goutils.ToSet([]float64{1.0, 5.0, -10.0})
	// true 
	result := goutils.SetContains(tempSet, 5.0)
	println(result)

	// strings
	testString := "abc123"
	endIndex := strings.Index(testString, "123")
	println(goutils.Substring(testString, 0, endIndex))
}
```