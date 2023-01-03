# JSlice

<img src="https://i.ibb.co/9s1s4cv/Frame-2.png" alt="Swagno" align="right" width="300"/>

JSlice is a Golang package that provides a set of methods to manipulate arrays which are [built-in functions in JavaScript](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array).

## About the Project

I implemented this package as hobby project. I personally do not use external package for array manipulation. I prefer to implement them with for loop etc. when necessary. But if you are looking for a package that provides array manipulation methods, this package may help you.

## Getting started

1. Get jslice package in your project

```sh
go get github.com/anilsenay/jslice
```

2. Import jslice

```go
import (
  . "github.com/anilsenay/jslice"
)
```

3. Use jslice methods

```go
slice := jslice.Of(1, 2, 3, 4, 5)
filteredSlice = slice.Filter(func(i int) bool {return i > 3}) // [4,5]
filteredSlice = filteredSlice.Take(1) // [4]
```

Chaning methods is also possible.

```go
intSlice := []int{1, 2, 3, 4, 5}
slice = jslice.From(intSlice).Filter(func(i int) bool {
  return i > 3
}).Map(func(i int) int {
  return i * 2
})

fmt.Println(slice) // [8,10]
```

## Methods

| Function         | JavaScript | JSlice |
| ---------------- | ---------- | ------ |
| at()             | ✅         | ✅     |
| concat()         | ✅         | ✅     |
| copyWithin()     | ✅         | ❌     |
| entries()        | ✅         | ❌     |
| every()          | ✅         | ✅     |
| fill()           | ✅         | ❌     |
| filter()         | ✅         | ✅     |
| find()           | ✅         | ✅     |
| findIndex()      | ✅         | ✅     |
| findLast()       | ✅         | ✅     |
| findLastIndex()  | ✅         | ✅     |
| flat()           | ✅         | ❌     |
| flatMap()        | ✅         | ❌     |
| forEach()        | ✅         | ✅     |
| group()          | ✅         | ❌     |
| groupToMap()     | ✅         | ❌     |
| includes()       | ✅         | ✅     |
| indexOf()        | ✅         | ✅     |
| join()           | ✅         | ✅     |
| keys()           | ✅         | ✅     |
| lastIndexOf()    | ✅         | ✅     |
| map()            | ✅         | ✅     |
| pop()            | ✅         | ✅     |
| push()           | ✅         | ✅     |
| reduce()         | ✅         | ❌     |
| reduceRight()    | ✅         | ❌     |
| reverse()        | ✅         | ✅     |
| shift()          | ✅         | ✅     |
| slice()          | ✅         | ✅     |
| some()           | ✅         | ✅     |
| sort()           | ✅         | ✅     |
| splice()         | ✅         | ❌     |
| toLocaleString() | ✅         | ❌     |
| toString()       | ✅         | ✅     |
| unshift()        | ✅         | ✅     |
| values()         | ✅         | ✅     |

# Contribution

I am welcome to any contribution. Missing methods can be implemented. And also, you can improve the existing methods. Please feel free to open an issue or pull request.
