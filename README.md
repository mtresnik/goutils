# goutils
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://github.com/mtresnik/goutils/blob/main/LICENSE)
[![version](https://img.shields.io/badge/version-1.1.4-blue)](https://github.com/mtresnik/goutils/releases/tag/v1.1.4)
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

require github.com/mtresnik/goutils v1.1.4
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