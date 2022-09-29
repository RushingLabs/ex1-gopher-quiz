# Exercise 1 - Quiz Program via CSV

Ref: [https://courses.calhoun.io/lessons/les_goph_01](https://courses.calhoun.io/lessons/les_goph_01)

- CSV file should default to `problems.csv`


## Getting it Running from Scratch

- `go mod init exercises/ex1`
- `go run .`



## Questions, Research

How to find length of array

- Ref: [https://reactgo.com/go-array-length/](https://reactgo.com/go-array-length/)

```go
records = [...data here]
length = len(records)
```

How to write a for-loop

- Ref: [https://yourbasic.org/golang/for-loop/](https://yourbasic.org/golang/for-loop/)

```go
for i := 0; i < 9; i++ {
    fmt.Println(i) // print the counter
}
```


### CSV Resources

- csv package: [https://pkg.go.dev/encoding/csv](https://pkg.go.dev/encoding/csv)


### Print int in string

Ref: [https://golang.cafe/blog/golang-int-to-string-conversion-example.html](https://golang.cafe/blog/golang-int-to-string-conversion-example.html)

```go
input = "test"
fmt.Printf("length: %v", len(input))
```